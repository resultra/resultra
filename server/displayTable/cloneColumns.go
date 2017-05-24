package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/textInput"
	"resultra/datasheet/server/generic/uniqueID"
)

func cloneTableCols(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	if err := numberInput.CloneNumberInputs(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	if err := textInput.CloneTextInputs(remappedIDs, parentTableID); err != nil {
		return fmt.Errorf("cloneTableCols: %v", err)
	}

	return nil

}
