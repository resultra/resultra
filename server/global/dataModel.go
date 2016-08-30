package global

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const GlobalTypeText string = "text"
const GlobalTypeNumber string = "number"
const GlobalTypeTime string = "time"
const GlobalTypeBool string = "bool"
const GlobalTypeLongText string = "longText"
const GlobalTypeFile string = "file"

func validGlobalType(globalType string) bool {
	switch globalType {
	case GlobalTypeText:
		return true
	case GlobalTypeNumber:
		return true
	case GlobalTypeTime:
		return true
	case GlobalTypeBool:
		return true
	case GlobalTypeLongText:
		return true
	case GlobalTypeFile:
		return true
	default:
		return false
	}
}

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
