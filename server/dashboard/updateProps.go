// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboard

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
)

type DashboardIDInterface interface {
	getDashboardID() string
}

type DashboardIDHeader struct {
	DashboardID string `json:"dashboardID"`
}

func (idHeader DashboardIDHeader) getDashboardID() string {
	return idHeader.DashboardID
}

type DashboardPropUpdater interface {
	DashboardIDInterface
	updateProps(trackerDBHandle *sql.DB, dashboard *Dashboard) error
}

func updateDashboardProps(trackerDBHandle *sql.DB, propUpdater DashboardPropUpdater) (*Dashboard, error) {

	// Retrieve the bar chart from the data store
	dbForUpdate, getErr := GetDashboard(trackerDBHandle, propUpdater.getDashboardID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to get existing dashboard: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(trackerDBHandle, dbForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to update existing dashboard properties: %v", propUpdateErr)
	}

	updatedDb, updateErr := updateExistingDashboard(trackerDBHandle, propUpdater.getDashboardID(), dbForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to update existing dashboard properties: datastore update error =  %v", updateErr)
	}

	return updatedDb, nil
}

type SetDashboardNameParams struct {
	DashboardIDHeader
	NewName string `json:"newName"`
}

func (updateParams SetDashboardNameParams) updateProps(trackerDBHandle *sql.DB, db *Dashboard) error {

	if validateErr := validateDashboardName(trackerDBHandle,
		updateParams.DashboardID, updateParams.NewName); validateErr != nil {
		return validateErr
	}

	db.Name = updateParams.NewName

	return nil
}

type SetDashboardLayoutParams struct {
	DashboardIDHeader
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (updateParams SetDashboardLayoutParams) updateProps(trackerDBHandle *sql.DB, db *Dashboard) error {

	db.Properties.Layout = updateParams.Layout

	return nil
}

type SetIncludeInSidebarParams struct {
	DashboardIDHeader
	IncludeInSidebar bool `json:"includeInSidebar"`
}

func (updateParams SetIncludeInSidebarParams) updateProps(trackerDBHandle *sql.DB, db *Dashboard) error {

	db.Properties.IncludeInSidebar = updateParams.IncludeInSidebar

	return nil
}
