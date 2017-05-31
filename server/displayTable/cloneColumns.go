package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/rating"
	"resultra/datasheet/server/displayTable/columns/textInput"
	"resultra/datasheet/server/displayTable/columns/toggle"
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

	return nil

}
