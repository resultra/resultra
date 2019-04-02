// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package inputProps

import (
	"fmt"
	"github.com/resultra/resultra/server/generic/stringValidation"
	"github.com/resultra/resultra/server/generic/uniqueID"
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
