// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/generic/timestamp"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
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

func saveNewGlobal(destDBHandle *sql.DB, newGlobal Global) error {

	if _, insertErr := destDBHandle.Exec(
		`INSERT INTO globals (database_id,global_id,name,ref_name,type) VALUES ($1,$2,$3,$4,$5)`,
		newGlobal.ParentDatabaseID, newGlobal.GlobalID, newGlobal.Name, newGlobal.RefName, newGlobal.Type); insertErr != nil {
		return fmt.Errorf("newGlobal: Can't create global: error = %v", insertErr)
	}

	log.Printf("newGlobal: Created new global: %+v", newGlobal)

	return nil
}

func newGlobal(trackerDBHandle *sql.DB, params NewGlobalParams) (*Global, error) {

	validateErr := validateNewGlobalName(trackerDBHandle, params.ParentDatabaseID, params.Name)
	if validateErr != nil {
		return nil, validateErr
	}

	if !validGlobalType(params.Type) {
		return nil, fmt.Errorf("newGlobal: Invalid type = %v", params.Type)
	}

	if refNameErr := validateNewReferenceName(trackerDBHandle, params.ParentDatabaseID, params.RefName); refNameErr != nil {
		return nil, fmt.Errorf("newGlobal: Invalid formula reference name = %v: %v", params.RefName, refNameErr)
	}

	newGlobal := Global{ParentDatabaseID: params.ParentDatabaseID,
		GlobalID: uniqueID.GenerateUniqueID(),
		Name:     params.Name,
		RefName:  params.RefName,
		Type:     params.Type}

	if err := saveNewGlobal(trackerDBHandle, newGlobal); err != nil {
		return nil, fmt.Errorf("newGlobal: Can't create global: error = %v", err)
	}

	log.Printf("newGlobal: Created new global: %+v", newGlobal)

	return &newGlobal, nil
}

func getGlobal(trackerDBHandle *sql.DB, globalID string) (*Global, error) {

	var theGlobal Global
	getErr := trackerDBHandle.QueryRow(
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

func getGlobalsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]Global, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT global_id,name,ref_name,type FROM globals WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getGlobals: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

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

func GetGlobals(trackerDBHandle *sql.DB, parentDatabaseID string) ([]Global, error) {
	return getGlobalsFromSrc(trackerDBHandle, parentDatabaseID)
}

func CloneGlobals(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	destDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneGlobals: Unable to get mapped ID for source database: %v", err)
	}

	globals, err := getGlobalsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneGlobals: Unable to retrieve globals: databaseID=%v, error=%v ",
			cloneParams.SourceDatabaseID, err)
	}
	for _, currGlobal := range globals {

		remappedID := uniqueID.GenerateUniqueID()
		cloneParams.IDRemapper[currGlobal.GlobalID] = remappedID

		destGlobal := currGlobal
		destGlobal.ParentDatabaseID = destDatabaseID
		destGlobal.GlobalID = remappedID

		if err := saveNewGlobal(cloneParams.DestDBHandle, destGlobal); err != nil {
			return fmt.Errorf("CloneGlobals: Can't create global: error = %v", err)
		}
	}

	return nil

}

type GlobalIDGlobalIndex map[string]Global

func GetIndexedGlobals(trackerDBHandle *sql.DB, parentDatabaseID string) (GlobalIDGlobalIndex, error) {

	globals, getGlobalErr := GetGlobals(trackerDBHandle, parentDatabaseID)
	if getGlobalErr != nil {
		return nil, fmt.Errorf("GetIndexedGlobals: Failure getting globals: %v", getGlobalErr)

	}

	globalIDGlobalIndex := GlobalIDGlobalIndex{}

	for _, currGlobal := range globals {
		globalIDGlobalIndex[currGlobal.GlobalID] = currGlobal
	}

	return globalIDGlobalIndex, nil
}

func getGlobalDatabaseID(trackerDBHandle *sql.DB, globalID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
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

func saveValUpdate(trackerDBHandle *sql.DB, globalID string, encodedValue string) (*GlobalValUpdate, error) {

	valUpdate := GlobalValUpdate{
		UpdateID:        uniqueID.GenerateUniqueID(),
		GlobalID:        globalID,
		UpdateTimestamp: timestamp.CurrentTimestampUTC(),
		Value:           encodedValue}

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO global_updates (update_id,global_id,update_timestamp_utc,value) VALUES ($1,$2,$3,$4)`,
		valUpdate.UpdateID, valUpdate.GlobalID, valUpdate.UpdateTimestamp, valUpdate.Value); insertErr != nil {
		return nil, fmt.Errorf("saveValUpdate: Can't save global value: error = %v", insertErr)
	}

	return &valUpdate, nil

}

// getValUpdates retrieves a list of value updates for all the globals in the database.
func getValUpdates(trackerDBHandle *sql.DB, parentDatabaseID string) ([]GlobalValUpdate, error) {

	rows, queryErr := trackerDBHandle.Query(`SELECT update_id,global_updates.global_id,update_timestamp_utc,value
			FROM globals,global_updates
			WHERE globals.database_id=$1 and globals.global_id=global_updates.global_id
			ORDER BY globals.global_id,update_timestamp_utc`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getValUpdates: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

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
