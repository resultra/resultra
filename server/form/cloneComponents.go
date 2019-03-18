// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package form

import (
	"fmt"

	"resultra/tracker/server/form/components/attachment"
	"resultra/tracker/server/form/components/caption"
	"resultra/tracker/server/form/components/checkBox"
	"resultra/tracker/server/form/components/comment"
	"resultra/tracker/server/form/components/datePicker"
	"resultra/tracker/server/form/components/emailAddr"
	"resultra/tracker/server/form/components/file"
	"resultra/tracker/server/form/components/formButton"
	"resultra/tracker/server/form/components/gauge"
	"resultra/tracker/server/form/components/header"
	"resultra/tracker/server/form/components/htmlEditor"
	"resultra/tracker/server/form/components/image"
	"resultra/tracker/server/form/components/label"
	"resultra/tracker/server/form/components/numberInput"
	"resultra/tracker/server/form/components/progress"
	"resultra/tracker/server/form/components/rating"
	"resultra/tracker/server/form/components/selection"
	"resultra/tracker/server/form/components/socialButton"
	"resultra/tracker/server/form/components/textBox"
	"resultra/tracker/server/form/components/toggle"
	"resultra/tracker/server/form/components/urlLink"
	"resultra/tracker/server/form/components/userSelection"
	"resultra/tracker/server/form/components/userTag"

	"resultra/tracker/server/trackerDatabase"
)

func cloneFormComponents(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	if err := textBox.CloneTextBoxes(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := numberInput.CloneNumberInputs(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := datePicker.CloneDatePickers(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := progress.CloneProgressIndicators(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := gauge.CloneGauges(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := checkBox.CloneCheckBoxes(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := toggle.CloneToggles(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := header.CloneHeaders(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := formButton.CloneButtons(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := htmlEditor.CloneHTMLEditors(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := attachment.CloneImages(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := rating.CloneRatings(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := selection.CloneSelections(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := userSelection.CloneUserSelections(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := userTag.CloneUserTags(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := comment.CloneComments(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := caption.CloneCaptions(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := socialButton.CloneSocialButtons(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := label.CloneLabels(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := emailAddr.CloneEmailAddrs(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := urlLink.CloneUrlLinks(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := file.CloneFiles(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := image.CloneImages(cloneParams, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	return nil

}
