package form

import (
	"fmt"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/datePicker"
	//	"resultra/datasheet/server/form/components/header"
	//	"resultra/datasheet/server/form/components/htmlEditor"
	//	"resultra/datasheet/server/form/components/image"
	//	"resultra/datasheet/server/form/components/rating"
	//	"resultra/datasheet/server/form/components/selection"
	"resultra/datasheet/server/form/components/textBox"
	//	"resultra/datasheet/server/form/components/userSelection"

	"resultra/datasheet/server/generic/uniqueID"
)

func cloneFormComponents(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	if err := textBox.CloneTextBoxes(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := datePicker.CloneDatePickers(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	if err := checkBox.CloneCheckBoxes(remappedIDs, parentFormID); err != nil {
		return fmt.Errorf("cloneFormComponents: %v", err)
	}

	return nil

}
