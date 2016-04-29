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

	fieldID := setFormulaParams.GetUniqueRootID().ObjectID

	valParms := ValidateFormulaParams{FieldID: fieldID, FormulaText: setFormulaParams.FormulaText}
	validationResp := validateFormulaText(appEngContext, valParms)
	if !validationResp.IsValidFormula {
		return fmt.Errorf("Can't set calculated field formula: invalid formula: %v", validationResp.ErrorMsg)
	}

	fieldForUpdate.CalcFieldFormulaText = setFormulaParams.FormulaText

	return nil
}
