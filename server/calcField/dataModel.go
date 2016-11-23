package calcField

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/table"
)

// Parameters for creating a new calculated field. FieldEqn needs to be converted to
// JSON before being saved. TBD - Should the parameters instead be an equation in
// end-user format? If so, this code will need an update once equation parsing is
// done.
type NewCalcFieldParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	RefName       string `json:"refName"`
	FormulaText   string `json:"formulaText"`
}

func newCalcField(calcFieldParams NewCalcFieldParams) (*field.Field, error) {

	databaseID, getDatabaseIDErr := table.GetTableDatabaseID(calcFieldParams.ParentTableID)
	if getDatabaseIDErr != nil {
		return nil, fmt.Errorf("assembleCalcFieldCompileParams: Unable to get database ID for table: table id =%v, error=%v ",
			calcFieldParams.ParentTableID, getDatabaseIDErr)
	}

	compileParams := formulaCompileParams{
		formulaText:        calcFieldParams.FormulaText,
		parentTableID:      calcFieldParams.ParentTableID,
		databaseID:         databaseID,
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
		ParentTableID:           calcFieldParams.ParentTableID,
		FieldID:                 uniqueID.GenerateSnowflakeID(),
		Name:                    calcFieldParams.Name,
		Type:                    calcFieldParams.Type,
		RefName:                 calcFieldParams.RefName,
		CalcFieldEqn:            compileResult.jsonEncodedEqn,
		PreprocessedFormulaText: compileResult.preprocessedFormula,
		IsCalcField:             true}

	return field.CreateNewFieldFromRawInputs(newField)
}
