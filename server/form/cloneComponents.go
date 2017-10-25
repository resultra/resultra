package form

import (
	"fmt"

	"resultra/datasheet/server/form/components/attachment"
	"resultra/datasheet/server/form/components/caption"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/comment"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/emailAddr"
	"resultra/datasheet/server/form/components/file"
	"resultra/datasheet/server/form/components/formButton"
	"resultra/datasheet/server/form/components/gauge"
	"resultra/datasheet/server/form/components/header"
	"resultra/datasheet/server/form/components/htmlEditor"
	"resultra/datasheet/server/form/components/image"
	"resultra/datasheet/server/form/components/label"
	"resultra/datasheet/server/form/components/numberInput"
	"resultra/datasheet/server/form/components/progress"
	"resultra/datasheet/server/form/components/rating"
	"resultra/datasheet/server/form/components/selection"
	"resultra/datasheet/server/form/components/socialButton"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/form/components/toggle"
	"resultra/datasheet/server/form/components/urlLink"
	"resultra/datasheet/server/form/components/userSelection"
	"resultra/datasheet/server/form/components/userTag"

	"resultra/datasheet/server/trackerDatabase"
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
