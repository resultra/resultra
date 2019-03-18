package calcField

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic/uniqueID"
)

// Parameters for creating a new calculated field. FieldEqn needs to be converted to
// JSON before being saved. TBD - Should the parameters instead be an equation in
// end-user format? If so, this code will need an update once equation parsing is
// done.
type NewCalcFieldParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	RefName          string `json:"refName"`
	FormulaText      string `json:"formulaText"`
}

func newCalcField(trackerDBHandle *sql.DB, calcFieldParams NewCalcFieldParams) (*field.Field, error) {

	compileParams := formulaCompileParams{
		trackerDBHandle:    trackerDBHandle,
		formulaText:        calcFieldParams.FormulaText,
		databaseID:         calcFieldParams.ParentDatabaseID,
		expectedResultType: calcFieldParams.Type,
		// resultFieldID is intentionally left empty. resultFieldID is used
		// to check for cycles in the formula. Since this is a new field, there
		// by definition can't be any existing formulas with references to this field,
		// so cycle detection can be disabled.
		resultFieldID: ""}

	compileResult, err := compileAndEncodeFormula(compileParams)
	if err != nil {
		return nil, fmt.Errorf("Error creating new calculated field %v, can't compile formula: %v",
			calcFieldParams.Name, err)
	}

	// Create the actual field. All the parameters are the same as calcFieldParams, except
	// the equation which is encoded in JSON.
	newField := field.Field{
		ParentDatabaseID:        calcFieldParams.ParentDatabaseID,
		FieldID:                 uniqueID.GenerateUniqueID(),
		Name:                    calcFieldParams.Name,
		Type:                    calcFieldParams.Type,
		RefName:                 calcFieldParams.RefName,
		CalcFieldEqn:            compileResult.jsonEncodedEqn,
		PreprocessedFormulaText: compileResult.preprocessedFormula,
		IsCalcField:             true}

	return field.CreateNewFieldFromRawInputs(trackerDBHandle, newField)
}
