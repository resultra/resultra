package calcField

import (
	"fmt"
	"log"
	"resultra/datasheet/server/field"
)

type SetFormulaParams struct {
	field.FieldIDHeader
	FormulaText string `json:"formulaText"`
}

func (setFormulaParams SetFormulaParams) UpdateProps(fieldForUpdate *field.Field) error {

	if !fieldForUpdate.IsCalcField {
		return fmt.Errorf("SetFormula: Can't set formula on non-calculated field: %v", fieldForUpdate.Name)
	}

	compileParams, paramErr := assembleCalcFieldCompileParams(setFormulaParams.GetFieldID())
	if paramErr != nil {
		return fmt.Errorf("Error saving formula, can't setup parameters for formula compilation: %v", paramErr)
	}
	compileParams.formulaText = setFormulaParams.FormulaText

	compileResult, compileErr := compileAndEncodeFormula(*compileParams)
	if compileErr != nil {
		return fmt.Errorf("Error saving formula, can't compile formula: %v", compileErr)
	}

	log.Printf("Formula compilation succeeded: %+v", compileResult)

	// The formula source/text is always stored side-by-side with the compile equation.
	// This compile equation is used for equation evaluation.
	fieldForUpdate.CalcFieldEqn = compileResult.jsonEncodedEqn
	fieldForUpdate.PreprocessedFormulaText = compileResult.preprocessedFormula

	// TODO(IMPORTANT) - Saving the equation doesn't yet result in an update to records' values for the field.
	// All the record values need to be updated for the new equation. Or, formula evaluation needs to be done
	// on the fly when records are loaded. Otherwise, there will be stale results in the records.

	return nil
}
