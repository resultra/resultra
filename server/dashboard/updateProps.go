package dashboard

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
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
	updateProps(dashboard *Dashboard) error
}

func updateDashboardProps(propUpdater DashboardPropUpdater) (*Dashboard, error) {

	// Retrieve the bar chart from the data store
	dbForUpdate, getErr := GetDashboard(propUpdater.getDashboardID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to get existing dashboard: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(dbForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to update existing dashboard properties: %v", propUpdateErr)
	}

	updatedDb, updateErr := updateExistingDashboard(propUpdater.getDashboardID(), dbForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateDashboardProps: Unable to update existing dashboard properties: datastore update error =  %v", updateErr)
	}

	return updatedDb, nil
}

type SetDashboardNameParams struct {
	DashboardIDHeader
	NewName string `json:"newName"`
}

func (updateParams SetDashboardNameParams) updateProps(db *Dashboard) error {

	if validateErr := validateDashboardName(updateParams.DashboardID, updateParams.NewName); validateErr != nil {
		return validateErr
	}

	db.Name = updateParams.NewName

	return nil
}

type SetDashboardLayoutParams struct {
	DashboardIDHeader
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (updateParams SetDashboardLayoutParams) updateProps(db *Dashboard) error {

	db.Properties.Layout = updateParams.Layout

	return nil
}
