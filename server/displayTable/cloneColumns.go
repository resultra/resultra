package displayTable

import (
	"fmt"
	"resultra/tracker/server/displayTable/columns/attachment"
	"resultra/tracker/server/displayTable/columns/checkBox"
	"resultra/tracker/server/displayTable/columns/comment"
	"resultra/tracker/server/displayTable/columns/datePicker"
	"resultra/tracker/server/displayTable/columns/emailAddr"
	"resultra/tracker/server/displayTable/columns/file"
	"resultra/tracker/server/displayTable/columns/formButton"
	"resultra/tracker/server/displayTable/columns/image"
	"resultra/tracker/server/displayTable/columns/note"
	"resultra/tracker/server/displayTable/columns/numberInput"
	"resultra/tracker/server/displayTable/columns/progress"
	"resultra/tracker/server/displayTable/columns/rating"
	"resultra/tracker/server/displayTable/columns/socialButton"
	"resultra/tracker/server/displayTable/columns/tag"
	"resultra/tracker/server/displayTable/columns/textInput"
	"resultra/tracker/server/displayTable/columns/textSelection"
	"resultra/tracker/server/displayTable/columns/toggle"
	"resultra/tracker/server/displayTable/columns/urlLink"
	"resultra/tracker/server/displayTable/columns/userSelection"
	"resultra/tracker/server/displayTable/columns/userTag"
	"resultra/tracker/server/trackerDatabase"
)

func cloneTableCols(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	if err := numberInput.CloneNumberInputs(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := textInput.CloneTextInputs(cloneParams, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}
	if err := textSelection.CloneTextSelections(cloneParams, parentTableID); err != nil {
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
