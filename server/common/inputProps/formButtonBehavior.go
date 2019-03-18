package inputProps

import (
	"fmt"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/generic/uniqueID"
)

const popupBehaviorModeless string = "modeless"
const popupBehaviorModal string = "modal"

const showFormDestPopup string = "popup"
const showFormDestPage string = "page"
const showFormDestNewPage string = "newPage"

type ButtonPopupBehavior struct {
	PopupMode            string `json:"popupMode"`
	CustomLabelModalSave string `json:"customLabelModalSave"`
	WhereShowForm        string `json:"whereShowForm"`
}

func (srcProps ButtonPopupBehavior) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonPopupBehavior, error) {
	destProps := srcProps
	return &destProps, nil
}

func NewDefaultPopupBehavior() ButtonPopupBehavior {
	defaultPopupBehavior := ButtonPopupBehavior{
		PopupMode:            popupBehaviorModeless,
		CustomLabelModalSave: "",
		WhereShowForm:        showFormDestNewPage}
	return defaultPopupBehavior
}

func (buttonPopupBehavior ButtonPopupBehavior) ValidateWellFormed() error {

	if !(buttonPopupBehavior.PopupMode == popupBehaviorModeless ||
		buttonPopupBehavior.PopupMode == popupBehaviorModal) {
		return fmt.Errorf("Invalid form popup mode: %v", buttonPopupBehavior.PopupMode)
	}

	if validLabelErr := stringValidation.ValidateOptionalItemLabel(buttonPopupBehavior.CustomLabelModalSave); validLabelErr != nil {
		return validLabelErr
	}

	return nil
}
