package global

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

type Global struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	GlobalID         string `json:"globalID"`
	Name             string `json:"name"`
	Type             string `json:"type"`
}

type NewGlobalParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
	Type             string `json:"type"`
}

func newGlobal(params NewGlobalParams) (*Global, error) {

	validateErr := validateNewGlobalName(params.ParentDatabaseID, params.Name)
	if validateErr != nil {
		return nil, validateErr
	}

	if !validGlobalType(params.Type) {
		return nil, fmt.Errorf("newGlobal: Invalid type = %v", params.Type)
	}

	newGlobal := Global{ParentDatabaseID: params.ParentDatabaseID,
		GlobalID: uniqueID.GenerateSnowflakeID(),
		Name:     params.Name,
		Type:     params.Type}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO globals (database_id,global_id,name,type) VALUES ($1,$2,$3,$4)`,
		newGlobal.ParentDatabaseID, newGlobal.GlobalID, newGlobal.Name, newGlobal.Type); insertErr != nil {
		return nil, fmt.Errorf("newGlobal: Can't create global: error = %v", insertErr)
	}

	log.Printf("newGlobal: Created new global: %+v", newGlobal)

	return &newGlobal, nil
}

func getGlobal(globalID string) (*Global, error) {

	var theGlobal Global
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id,global_id,name,type 
			FROM globals 
			WHERE globals.global_id=$1 LIMIT 1`,
		globalID).Scan(&theGlobal.ParentDatabaseID, &theGlobal.GlobalID,
		&theGlobal.Name, &theGlobal.Type)
	if getErr != nil {
		return nil, fmt.Errorf(
			"getGlobal: can't get database for global = %v: err=%v",
			globalID, getErr)
	}

	return &theGlobal, nil

}

type GetGlobalsParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getGlobals(parentDatabaseID string) ([]Global, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT global_id,name,type FROM globals WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getGlobals: Failure querying database: %v", queryErr)
	}

	globals := []Global{}
	for rows.Next() {

		var currGlobal Global
		currGlobal.ParentDatabaseID = parentDatabaseID

		if scanErr := rows.Scan(&currGlobal.GlobalID, &currGlobal.Name, &currGlobal.Type); scanErr != nil {
			return nil, fmt.Errorf("getGlobals: Failure querying database: %v", scanErr)
		}

		globals = append(globals, currGlobal)
	}

	return globals, nil

}

func getGlobalDatabaseID(globalID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id 
			FROM globals 
			WHERE globals.global_id=$1 LIMIT 1`,
		globalID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getGlobalDatabaseID: can't get database for global = %v: err=%v",
			globalID, getErr)
	}

	return databaseID, nil

}

type GlobalValUpdate struct {
	UpdateID        string `json:"updateID"`
	GlobalID        string `json:"globalID"`
	UpdateTimestamp time.Time
	Value           string `json:value`
}

func saveValUpdate(globalID string, encodedValue string) (*GlobalValUpdate, error) {

	valUpdate := GlobalValUpdate{
		UpdateID:        uniqueID.GenerateSnowflakeID(),
		GlobalID:        globalID,
		UpdateTimestamp: time.Now().UTC(),
		Value:           encodedValue}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO global_updates (update_id,global_id,update_timestamp_utc,value) VALUES ($1,$2,$3,$4)`,
		valUpdate.UpdateID, valUpdate.GlobalID, valUpdate.UpdateTimestamp, valUpdate.Value); insertErr != nil {
		return nil, fmt.Errorf("saveValUpdate: Can't save global value: error = %v", insertErr)
	}

	return &valUpdate, nil

}

// getValUpdates retrieves a list of value updates for all the globals in the database.
func getValUpdates(parentDatabaseID string) ([]GlobalValUpdate, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT update_id,global_updates.global_id,update_timestamp_utc,value
			FROM globals,global_updates
			WHERE globals.database_id=$1 and globals.global_id=global_updates.global_id
			ORDER BY global_id,update_timestamp_utc`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getValUpdates: Failure querying database: %v", queryErr)
	}
	valUpdates := []GlobalValUpdate{}
	for rows.Next() {
		var currValUpdate GlobalValUpdate
		if scanErr := rows.Scan(&currValUpdate.UpdateID,
			&currValUpdate.GlobalID,
			&currValUpdate.UpdateTimestamp, &currValUpdate.Value); scanErr != nil {
			return nil, fmt.Errorf("getValUpdates: Failure querying database: %v", scanErr)

		}
		valUpdates = append(valUpdates, currValUpdate)
	}

	return valUpdates, nil
}
