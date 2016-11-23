package database

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

type DatabaseProperties struct {
	FormsOrder      []string `json:"formsOrder"`
	DashboardsOrder []string `json:"dashboardsOrder"`
}

type Database struct {
	DatabaseID string             `json:"databaseID"`
	Name       string             `json:"name"`
	Properties DatabaseProperties `json:"properties"`
}

func SaveNewDatabase(newDatabase Database) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newDatabase.Properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO databases VALUES ($1,$2,$3)`,
		newDatabase.DatabaseID, newDatabase.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveNewDatabase: insert failed: error = %v", insertErr)
	}

	return nil

}

func CloneDatabase(remappedIDs map[string]string, newName string, srcDatabaseID string) (*Database, error) {

	srcDatabase, err := GetDatabase(srcDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	destDatabaseID := uniqueID.GenerateSnowflakeID()
	remappedIDs[srcDatabaseID] = destDatabaseID

	dest := *srcDatabase
	dest.DatabaseID = destDatabaseID
	dest.Name = newName
	// TODO - handle properties

	if err := SaveNewDatabase(dest); err != nil {
		return nil, fmt.Errorf("Clone database: Can't save database: %v", err)
	}

	return &dest, nil

}

type NewDatabaseParams struct {
	Name string `json:"name"`
}

func SaveNewEmptyDatabase(params NewDatabaseParams) (*Database, error) {

	sanitizedDbName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	databaseID := uniqueID.GenerateSnowflakeID()

	dbProps := DatabaseProperties{}

	newDatabase := Database{
		DatabaseID: databaseID,
		Name:       sanitizedDbName,
		Properties: dbProps}

	if err := SaveNewDatabase(newDatabase); err != nil {
		return nil, fmt.Errorf("SaveNewEmptyDatabase: %v", err)
	}

	return &newDatabase, nil
}

func GetDatabase(databaseID string) (*Database, error) {

	dbName := ""
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT name,properties FROM databases
		 WHERE database_id=$1 LIMIT 1`, databaseID).Scan(&dbName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabase: Unabled to get database: db ID = %v: datastore err=%v",
			databaseID, getErr)
	}

	var dbProps DatabaseProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &dbProps); decodeErr != nil {
		return nil, fmt.Errorf("getDatabase: can't decode properties: %v, err=%v", encodedProps, decodeErr)
	}

	getDb := Database{
		DatabaseID: databaseID,
		Name:       dbName,
		Properties: dbProps}

	return &getDb, nil
}

func updateExistingDatabase(databaseID string, updatedDB *Database) (*Database, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedDB.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE databases 
				SET properties=$1, name=$2
				WHERE database_id=$3`,
		encodedProps, updatedDB.Name, databaseID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: Can't update database properties %v: error = %v",
			databaseID, updateErr)
	}

	return updatedDB, nil

}

func validateDatabaseName(databaseID string, databaseName string) error {
	if !stringValidation.WellFormedItemName(databaseName) {
		return fmt.Errorf("Invalid database name")
	}
	return nil
}
