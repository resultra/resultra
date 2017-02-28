package form

import (
	"fmt"

	"resultra/datasheet/server/form/components/caption"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/comment"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/formButton"
	"resultra/datasheet/server/form/components/header"
	"resultra/datasheet/server/form/components/htmlEditor"
	"resultra/datasheet/server/form/components/image"
	"resultra/datasheet/server/form/components/progress"
	"resultra/datasheet/server/form/components/rating"
	"resultra/datasheet/server/form/components/selection"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/form/components/userSelection"

	"resultra/datasheet/server/generic/uniqueID"
)

func cloneFormComponents(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	if err := textBox.CloneTextBoxes(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := datePicker.CloneDatePickers(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := progress.CloneProgressIndicators(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := checkBox.CloneCheckBoxes(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := header.CloneHeaders(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := formButton.CloneButtons(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := htmlEditor.CloneHTMLEditors(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := image.CloneImages(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := rating.CloneRatings(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := selection.CloneSelections(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := userSelection.CloneUserSelections(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := comment.CloneComments(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := caption.CloneCaptions(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	return nil

}
