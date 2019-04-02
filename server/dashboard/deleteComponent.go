// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboard

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/dashboard/components/common"
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
