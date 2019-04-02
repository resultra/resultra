// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package displayTable

import (
	"fmt"
	"github.com/resultra/resultra/server/displayTable/columns/attachment"
	"github.com/resultra/resultra/server/displayTable/columns/checkBox"
	"github.com/resultra/resultra/server/displayTable/columns/comment"
	"github.com/resultra/resultra/server/displayTable/columns/datePicker"
	"github.com/resultra/resultra/server/displayTable/columns/emailAddr"
	"github.com/resultra/resultra/server/displayTable/columns/file"
	"github.com/resultra/resultra/server/displayTable/columns/formButton"
	"github.com/resultra/resultra/server/displayTable/columns/image"
	"github.com/resultra/resultra/server/displayTable/columns/note"
	"github.com/resultra/resultra/server/displayTable/columns/numberInput"
	"github.com/resultra/resultra/server/displayTable/columns/progress"
	"github.com/resultra/resultra/server/displayTable/columns/rating"
	"github.com/resultra/resultra/server/displayTable/columns/socialButton"
	"github.com/resultra/resultra/server/displayTable/columns/tag"
	"github.com/resultra/resultra/server/displayTable/columns/textInput"
	"github.com/resultra/resultra/server/displayTable/columns/textSelection"
	"github.com/resultra/resultra/server/displayTable/columns/toggle"
	"github.com/resultra/resultra/server/displayTable/columns/urlLink"
	"github.com/resultra/resultra/server/displayTable/columns/userSelection"
	"github.com/resultra/resultra/server/displayTable/columns/userTag"
	"github.com/resultra/resultra/server/trackerDatabase"
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
