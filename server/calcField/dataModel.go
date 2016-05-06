package calcField

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
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

func newCalcField(appEngContext appengine.Context, calcFieldParams NewCalcFieldParams) (string, error) {

	compileParams := formulaCompileParams{
		appEngContext:      appEngContext,
		formulaText:        calcFieldParams.FormulaText,
		parentTableID:      calcFieldParams.ParentTableID,
		expectedResultType: calcFieldParams.Type}

	compileResult, err := compileAndEncodeFormula(compileParams)
	if err != nil {
		return "", fmt.Errorf("Error creating new calculated field %v, can't compile formula: %v",
			calcFieldParams.Name, err)
	}

	// Create the actual field. All the parameters are the same as calcFieldParams, except
	// the equation which is encoded in JSON.
	newField := field.Field{
		Name:                    calcFieldParams.Name,
		Type:                    calcFieldParams.Type,
		RefName:                 calcFieldParams.RefName,
		CalcFieldEqn:            compileResult.jsonEncodedEqn,
		PreprocessedFormulaText: compileResult.preprocessedFormula,
		IsCalcField:             true}

	return field.CreateNewFieldFromRawInputs(appEngContext, calcFieldParams.ParentTableID, newField)
}
