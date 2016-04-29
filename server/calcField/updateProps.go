package calcField

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

type SetFormulaParams struct {
	datastoreWrapper.UniqueRootIDHeader
	FormulaText string `json:"formulaText"`
}

func (setFormulaParams SetFormulaParams) UpdateProps(appEngContext appengine.Context, fieldForUpdate *field.Field) error {

	jsonEncodedEqn, compileErr := compileAndEncodeFormula(setFormulaParams.FormulaText)
	if compileErr != nil {
		fmt.Errorf("Error saving formula, can't compile formula: %v", compileErr)
	}

	// The formula source/text is always stored side-by-side with the compile equation.
	// This compile equation is used for equation evaluation.
	fieldForUpdate.CalcFieldEqn = jsonEncodedEqn
	fieldForUpdate.CalcFieldFormulaText = setFormulaParams.FormulaText

	// TODO(IMPORTANT) - Saving the equation doesn't yet result in an update to records' values for the field.
	// All the record values need to be updated for the new equation. Or, formula evaluation needs to be done
	// on the fly when records are loaded. Otherwise, there will be stale results in the records.

	return nil
}
