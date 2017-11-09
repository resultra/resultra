package form

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type DeleteComponentParams struct {
	ParentFormID string `json:"parentFormID"`
	ComponentID  string `json:"componentID"`
}

func deleteComponent(trackerDBHandle *sql.DB, params DeleteComponentParams) error {

	if deleteErr := common.DeleteFormComponent(trackerDBHandle, params.ParentFormID, params.ComponentID); deleteErr != nil {
		return fmt.Errorf("deleteComponent: %v", deleteErr)
	}

	return nil
}
