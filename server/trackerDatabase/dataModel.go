package trackerDatabase

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/generic/uniqueID"
)

type Database struct {
	DatabaseID      string             `json:"databaseID"`
	Name            string             `json:"name"`
	Properties      DatabaseProperties `json:"properties"`
	IsTemplate      bool               `json:"isTemplate"`
	IsActive        bool               `json:"isActive"`
	Description     string             `json:"description"`
	CreatedByUserID string             `json:"createdByUserID"`
}

func SaveNewDatabase(trackerDBHandle *sql.DB, newDatabase Database) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newDatabase.Properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewDatabase: failure encoding properties: error = %v", encodeErr)
	}

	isActive := true

	if _, insertErr := trackerDBHandle.Exec(`INSERT INTO databases 
			(database_id,name,properties,description,is_template,is_active,created_by_user_id)
			 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		newDatabase.DatabaseID,
		newDatabase.Name,
		encodedProps,
		newDatabase.Description,
		newDatabase.IsTemplate,
		isActive,
		newDatabase.CreatedByUserID); insertErr != nil {
		return fmt.Errorf("saveNewDatabase: insert failed: error = %v", insertErr)
	}

	return nil

}

type CloneDatabaseParams struct {
	NewName          string
	CreatedByUserID  string
	IsTemplate       bool
	SourceDatabaseID string
	SrcDBHandle      *sql.DB
	DestDBHandle     *sql.DB
	IDRemapper       uniqueID.UniqueIDRemapper
}

func CloneDatabase(cloneParams *CloneDatabaseParams) (*Database, error) {

	srcDatabase, err := GetDatabase(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}

	destDatabaseID, err := cloneParams.IDRemapper.AllocNewRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}

	dest := *srcDatabase
	dest.DatabaseID = destDatabaseID
	dest.Name = cloneParams.NewName
	dest.CreatedByUserID = cloneParams.CreatedByUserID
	dest.IsTemplate = cloneParams.IsTemplate

	destProps, err := srcDatabase.Properties.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CloneDatabase: %v", err)
	}
	dest.Properties = *destProps

	if err := SaveNewDatabase(cloneParams.DestDBHandle, dest); err != nil {
		return nil, fmt.Errorf("Clone database: Can't save database: %v", err)
	}

	return &dest, nil

}

const NewDatabaseTemplateSourceFactory string = "factory"
const NewDatabaseTemplateSourceAccount string = "account"

type NewDatabaseParams struct {
	Name               string  `json:"name"`
	TemplateDatabaseID *string `json:"templateDatabaseID"`
	TemplateSource     *string `json:"templateSource"`
	CreatedByUserID    string  `json:"createdByUserID"`
	IsTemplate         bool
}

func SaveNewEmptyDatabase(trackerDBHandle *sql.DB, params NewDatabaseParams) (*Database, error) {

	sanitizedDbName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	databaseID := uniqueID.GenerateUniqueID()

	dbProps := newDefaultDatabaseProperties()

	newDatabase := Database{
		DatabaseID:      databaseID,
		Name:            sanitizedDbName,
		IsTemplate:      params.IsTemplate,
		CreatedByUserID: params.CreatedByUserID,
		Properties:      dbProps}

	if err := SaveNewDatabase(trackerDBHandle, newDatabase); err != nil {
		return nil, fmt.Errorf("SaveNewEmptyDatabase: %v", err)
	}

	return &newDatabase, nil
}

func GetDatabase(trackerDBHandle *sql.DB, databaseID string) (*Database, error) {

	dbName := ""
	encodedProps := ""
	desc := ""
	isTemplate := false
	createdByUserID := ""
	isActive := false
	getErr := trackerDBHandle.QueryRow(`SELECT name,properties,description,is_template,created_by_user_id,is_active
		FROM databases
		 WHERE database_id=$1 LIMIT 1`, databaseID).Scan(&dbName, &encodedProps, &desc, &isTemplate, &createdByUserID, &isActive)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabase: Unabled to get database: db ID = %v: datastore err=%v",
			databaseID, getErr)
	}

	dbProps := newDefaultDatabaseProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &dbProps); decodeErr != nil {
		return nil, fmt.Errorf("getDatabase: can't decode properties: %v, err=%v", encodedProps, decodeErr)
	}

	getDb := Database{
		DatabaseID:      databaseID,
		Name:            dbName,
		Description:     desc,
		Properties:      dbProps,
		IsActive:        isActive,
		CreatedByUserID: createdByUserID,
		IsTemplate:      isTemplate}

	return &getDb, nil
}

func updateExistingDatabase(trackerDBHandle *sql.DB, databaseID string, updatedDB *Database) (*Database, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedDB.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE databases 
				SET properties=$1, name=$2, description=$3, is_active=$4
				WHERE database_id=$5`,
		encodedProps, updatedDB.Name, updatedDB.Description, updatedDB.IsActive, databaseID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingDatabase: Can't update database properties %v: error = %v",
			databaseID, updateErr)
	}

	return updatedDB, nil

}
