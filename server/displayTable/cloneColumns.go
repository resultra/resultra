package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/attachment"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/comment"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/emailAddr"
	"resultra/datasheet/server/displayTable/columns/formButton"
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
	"resultra/datasheet/server/generic/uniqueID"
)

func cloneTableCols(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	if err := numberInput.CloneNumberInputs(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := textInput.CloneTextInputs(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := datePicker.CloneDatePickers(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := checkBox.CloneCheckBoxes(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := rating.CloneRatings(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := toggle.CloneToggles(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := userSelection.CloneUserSelections(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := note.CloneNotes(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := comment.CloneComments(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := attachment.CloneAttachments(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := formButton.CloneButtons(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := progress.CloneProgressIndicators(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := socialButton.CloneSocialButtons(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := tag.CloneTags(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := emailAddr.CloneEmailAddrs(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := urlLink.CloneUrlLinks(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	return nil

}
