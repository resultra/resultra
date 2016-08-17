package database

import (
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

func updateDatabaseProps(propUpdater DatabasePropUpdater) (*Database, error) {

	// Retrieve the bar chart from the data store
	dbForUpdate, getErr := GetDatabase(propUpdater.getDatabaseID())
	if getErr != nil {
		return nil, fmt.Errorf("updateDatabaseProps: Unable to get existing database: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(dbForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to update existing database properties: %v", propUpdateErr)
	}

	updatedDb, updateErr := updateExistingDatabase(propUpdater.getDatabaseID(), dbForUpdate)
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
