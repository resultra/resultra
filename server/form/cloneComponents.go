// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package form

import (
	"fmt"

	"github.com/resultra/resultra/server/form/components/attachment"
	"github.com/resultra/resultra/server/form/components/caption"
	"github.com/resultra/resultra/server/form/components/checkBox"
	"github.com/resultra/resultra/server/form/components/comment"
	"github.com/resultra/resultra/server/form/components/datePicker"
	"github.com/resultra/resultra/server/form/components/emailAddr"
	"github.com/resultra/resultra/server/form/components/file"
	"github.com/resultra/resultra/server/form/components/formButton"
	"github.com/resultra/resultra/server/form/components/gauge"
	"github.com/resultra/resultra/server/form/components/header"
	"github.com/resultra/resultra/server/form/components/htmlEditor"
	"github.com/resultra/resultra/server/form/components/image"
	"github.com/resultra/resultra/server/form/components/label"
	"github.com/resultra/resultra/server/form/components/numberInput"
	"github.com/resultra/resultra/server/form/components/progress"
	"github.com/resultra/resultra/server/form/components/rating"
	"github.com/resultra/resultra/server/form/components/selection"
	"github.com/resultra/resultra/server/form/components/socialButton"
	"github.com/resultra/resultra/server/form/components/textBox"
	"github.com/resultra/resultra/server/form/components/toggle"
	"github.com/resultra/resultra/server/form/components/urlLink"
	"github.com/resultra/resultra/server/form/components/userSelection"
	"github.com/resultra/resultra/server/form/components/userTag"

	"github.com/resultra/resultra/server/trackerDatabase"
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
