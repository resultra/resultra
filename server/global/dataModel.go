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
	RefName          string `json:"refName"`
	Type             string `json:"type"`
}

type NewGlobalParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
	RefName          string `json:"refName"`
	Type             string `json:"type"`
}

func saveNewGlobal(newGlobal Global) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO globals (database_id,global_id,name,ref_name,type) VALUES ($1,$2,$3,$4,$5)`,
		newGlobal.ParentDatabaseID, newGlobal.GlobalID, newGlobal.Name, newGlobal.RefName, newGlobal.Type); insertErr != nil {
		return fmt.Errorf("newGlobal: Can't create global: error = %v", insertErr)
	}

	log.Printf("newGlobal: Created new global: %+v", newGlobal)

	return nil
}

func newGlobal(params NewGlobalParams) (*Global, error) {

	validateErr := validateNewGlobalName(params.ParentDatabaseID, params.Name)
	if validateErr != nil {
		return nil, validateErr
	}

	if !validGlobalType(params.Type) {
		return nil, fmt.Errorf("newGlobal: Invalid type = %v", params.Type)
	}

	if refNameErr := validateNewReferenceName(params.ParentDatabaseID, params.RefName); refNameErr != nil {
		return nil, fmt.Errorf("newGlobal: Invalid formula reference name = %v: %v", params.RefName, refNameErr)
	}

	newGlobal := Global{ParentDatabaseID: params.ParentDatabaseID,
		GlobalID: uniqueID.GenerateSnowflakeID(),
		Name:     params.Name,
		RefName:  params.RefName,
		Type:     params.Type}

	if err := saveNewGlobal(newGlobal); err != nil {
		return nil, fmt.Errorf("newGlobal: Can't create global: error = %v", err)
	}

	log.Printf("newGlobal: Created new global: %+v", newGlobal)

	return &newGlobal, nil
}

func getGlobal(globalID string) (*Global, error) {

	var theGlobal Global
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id,global_id,name,ref_name,type 
			FROM globals 
			WHERE globals.global_id=$1 LIMIT 1`,
		globalID).Scan(&theGlobal.ParentDatabaseID, &theGlobal.GlobalID,
		&theGlobal.Name, &theGlobal.RefName, &theGlobal.Type)
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

func GetGlobals(parentDatabaseID string) ([]Global, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT global_id,name,ref_name,type FROM globals WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getGlobals: Failure querying database: %v", queryErr)
	}

	globals := []Global{}
	for rows.Next() {

		var currGlobal Global
		currGlobal.ParentDatabaseID = parentDatabaseID

		if scanErr := rows.Scan(&currGlobal.GlobalID, &currGlobal.Name, &currGlobal.RefName,
			&currGlobal.Type); scanErr != nil {
			return nil, fmt.Errorf("getGlobals: Failure querying database: %v", scanErr)
		}

		globals = append(globals, currGlobal)
	}

	return globals, nil

}

func CloneGlobals(remappedIDs uniqueID.UniqueIDRemapper, srcDatabaseID string) error {

	destDatabaseID, err := remappedIDs.GetExistingRemappedID(srcDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneGlobals: Unable to get mapped ID for source database: %v", err)
	}

	globals, err := GetGlobals(srcDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneGlobals: Unable to retrieve globals: databaseID=%v, error=%v ",
			srcDatabaseID, err)
	}
	for _, currGlobal := range globals {

		remappedID := uniqueID.GenerateSnowflakeID()
		remappedIDs[currGlobal.GlobalID] = remappedID

		destGlobal := currGlobal
		destGlobal.ParentDatabaseID = destDatabaseID
		destGlobal.GlobalID = remappedID

		if err := saveNewGlobal(destGlobal); err != nil {
			return fmt.Errorf("CloneGlobals: Can't create global: error = %v", err)
		}
	}

	return nil

}

type GlobalIDGlobalIndex map[string]Global

func GetIndexedGlobals(parentDatabaseID string) (GlobalIDGlobalIndex, error) {

	globals, getGlobalErr := GetGlobals(parentDatabaseID)
	if getGlobalErr != nil {
		return nil, fmt.Errorf("GetIndexedGlobals: Failure getting globals: %v", getGlobalErr)

	}

	globalIDGlobalIndex := GlobalIDGlobalIndex{}

	for _, currGlobal := range globals {
		globalIDGlobalIndex[currGlobal.GlobalID] = currGlobal
	}

	return globalIDGlobalIndex, nil
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
			ORDER BY globals.global_id,update_timestamp_utc`,
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
