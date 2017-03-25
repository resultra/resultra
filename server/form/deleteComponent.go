package form

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type DeleteComponentParams struct {
	ParentFormID string `json:"parentFormID"`
	ComponentID  string `json:"componentID"`
}

func deleteComponent(params DeleteComponentParams) error {

	if deleteErr := common.DeleteFormComponent(params.ParentFormID, params.ComponentID); deleteErr != nil {
		return fmt.Errorf("deleteComponent: %v", deleteErr)
	}

	return nil
}
