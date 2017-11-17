package trackerDatabase

import (
	"database/sql"
	"fmt"
)

type DatabaseIDInterface interface {
	getDatabaseID() string
}

type DatabaseIDHeader struct {
	DatabaseID string `json:"databaseID"`
}

func (idHeader DatabaseIDHeader) getDatabaseID() string {
	return idHeader.DatabaseID
}

type DatabasePropUpdater interface {
	DatabaseIDInterface
	updateProps(db *Database) error
}

func updateDatabaseProps(trackerDBHandle *sql.DB, propUpdater DatabasePropUpdater) (*Database, error) {

	// Retrieve the bar chart from the data store
	dbForUpdate, getErr := GetDatabase(trackerDBHandle, propUpdater.getDatabaseID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDatabaseProps: Unable to get existing database: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(dbForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to update existing database properties: %v", propUpdateErr)
	}

	updatedDb, updateErr := updateExistingDatabase(trackerDBHandle, propUpdater.getDatabaseID(), dbForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateDatabaseProps: Unable to update existing database properties: datastore update error =  %v", updateErr)
	}

	return updatedDb, nil
}

type SetDatabaseNameParams struct {
	DatabaseIDHeader
	NewName string `json:"newName"`
}

func (updateParams SetDatabaseNameParams) updateProps(db *Database) error {

	if validateErr := validateDatabaseName(updateParams.DatabaseID, updateParams.NewName); validateErr != nil {
		return validateErr
	}

	db.Name = updateParams.NewName

	return nil
}

type SetListOrderParams struct {
	DatabaseIDHeader
	ListOrder []string `json:"listOrder"`
}

func (updateParams SetListOrderParams) updateProps(db *Database) error {

	// TODO - Validate all lists are in the database/tracker

	db.Properties.ListOrder = updateParams.ListOrder

	return nil
}

type SetDashboardOrderParams struct {
	DatabaseIDHeader
	DashboardOrder []string `json:"dashboardOrder"`
}

func (updateParams SetDashboardOrderParams) updateProps(db *Database) error {

	// TODO - Validate all lists are in the database/tracker

	db.Properties.DashboardOrder = updateParams.DashboardOrder

	return nil
}

type SetFormLinkOrderParams struct {
	DatabaseIDHeader
	FormLinkOrder []string `json:"formLinkOrder"`
}

func (updateParams SetFormLinkOrderParams) updateProps(db *Database) error {

	// TODO - Validate all lists are in the database/tracker

	db.Properties.FormLinkOrder = updateParams.FormLinkOrder

	return nil
}

type SetDescriptionParams struct {
	DatabaseIDHeader
	Description string `json:"description"`
}

func (updateParams SetDescriptionParams) updateProps(db *Database) error {

	db.Description = updateParams.Description

	return nil
}

type SetDatabaseActiveParams struct {
	DatabaseIDHeader
	IsActive bool `json:"isActive"`
}

func (updateParams SetDatabaseActiveParams) updateProps(db *Database) error {

	db.IsActive = updateParams.IsActive

	return nil
}
