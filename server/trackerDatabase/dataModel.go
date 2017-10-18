package trackerDatabase

import (
	"fmt"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

type Database struct {
	DatabaseID      string             `json:"databaseID"`
	Name            string             `json:"name"`
	Properties      DatabaseProperties `json:"properties"`
	IsTemplate      bool               `json:"isTemplate"`
	Description     *string            `json:"description"`
	CreatedByUserID string             `json:"createdByUserID"`
}

func SaveNewDatabase(newDatabase Database) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newDatabase.Properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO databases VALUES ($1,$2,$3,$4,$5,$6)`,
		newDatabase.DatabaseID, newDatabase.Name, encodedProps,
		newDatabase.Description, newDatabase.IsTemplate, newDatabase.CreatedByUserID); insertErr != nil {
		return fmt.Errorf("saveNewDatabase: insert failed: error = %v", insertErr)
	}

	return nil

}

type CloneDatabaseParams struct {
	NewName          string
	CreatedByUserID  string
	IsTemplate       bool
	SourceDatabaseID string
}

func CloneDatabase(remappedIDs uniqueID.UniqueIDRemapper, cloneParams CloneDatabaseParams) (*Database, error) {

	srcDatabase, err := GetDatabase(cloneParams.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}

	destDatabaseID, err := remappedIDs.AllocNewRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}

	dest := *srcDatabase
	dest.DatabaseID = destDatabaseID
	dest.Name = cloneParams.NewName
	dest.CreatedByUserID = cloneParams.CreatedByUserID
	dest.IsTemplate = cloneParams.IsTemplate

	destProps, err := srcDatabase.Properties.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}
	dest.Properties = *destProps

	if err := SaveNewDatabase(dest); err != nil {
		return nil, fmt.Errorf("Clone database: Can't save database: %v", err)
	}

	return &dest, nil

}

type NewDatabaseParams struct {
	Name               string  `json:"name"`
	TemplateDatabaseID *string `json:"templateDatabaseID"`
	CreatedByUserID    string  `json:"createdByUserID"`
	IsTemplate         bool
}

func SaveNewEmptyDatabase(params NewDatabaseParams) (*Database, error) {

	sanitizedDbName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	databaseID := uniqueID.GenerateSnowflakeID()

	dbProps := newDefaultDatabaseProperties()

	newDatabase := Database{
		DatabaseID:      databaseID,
		Name:            sanitizedDbName,
		IsTemplate:      params.IsTemplate,
		CreatedByUserID: params.CreatedByUserID,
		Properties:      dbProps}

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

	dbProps := newDefaultDatabaseProperties()
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
