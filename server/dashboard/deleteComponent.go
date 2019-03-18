package dashboard

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/dashboard/components/common"
)

type DeleteComponentParams struct {
	ParentDashboardID string `json:"parentDashboardID"`
	ComponentID       string `json:"componentID"`
}

func deleteComponent(trackerDBHandle *sql.DB, params DeleteComponentParams) error {

	if deleteErr := common.DeleteDashboardComponent(trackerDBHandle, params.ParentDashboardID, params.ComponentID); deleteErr != nil {
		return fmt.Errorf("deleteComponent: %v", deleteErr)
	}

	return nil
}
