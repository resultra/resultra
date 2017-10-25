package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/attachment"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/comment"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/emailAddr"
	"resultra/datasheet/server/displayTable/columns/file"
	"resultra/datasheet/server/displayTable/columns/formButton"
	"resultra/datasheet/server/displayTable/columns/image"
	"resultra/datasheet/server/displayTable/columns/note"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/progress"
	"resultra/datasheet/server/displayTable/columns/rating"
	"resultra/datasheet/server/displayTable/columns/socialButton"
	"resultra/datasheet/server/displayTable/columns/tag"
	"resultra/datasheet/server/displayTable/columns/textInput"
	"resultra/datasheet/server/displayTable/columns/toggle"
	"resultra/datasheet/server/displayTable/columns/urlLink"
	"resultra/datasheet/server/displayTable/columns/userSelection"
	"resultra/datasheet/server/displayTable/columns/userTag"
	"resultra/datasheet/server/trackerDatabase"
)

func cloneTableCols(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	if err := numberInput.CloneNumberInputs(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := textInput.CloneTextInputs(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := datePicker.CloneDatePickers(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := checkBox.CloneCheckBoxes(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := rating.CloneRatings(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := toggle.CloneToggles(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := userSelection.CloneUserSelections(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := userTag.CloneUserTags(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := note.CloneNotes(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := comment.CloneComments(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := attachment.CloneAttachments(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := formButton.CloneButtons(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := progress.CloneProgressIndicators(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := socialButton.CloneSocialButtons(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := tag.CloneTags(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := emailAddr.CloneEmailAddrs(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := urlLink.CloneUrlLinks(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := file.CloneFiles(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := image.CloneImages(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	return nil

}
