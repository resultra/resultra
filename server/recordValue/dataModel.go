package recordValue

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordFilter"
	"time"
)

type RecordValueResults struct {
	ParentTableID   string                             `json:"parentTableID"`
	RecordID        string                             `json:"recordID"`
	FieldValues     record.RecFieldValues              `json:"fieldValues"`
	FilterMatches   recordFilter.RecFilterMatchResults `json:"filterMatches"`
	UpdateTimestamp time.Time                          `json:"updateTimestamp"`
}

func saveRecordValueResults(recValResults RecordValueResults) error {

	log.Printf("saveRecordValueResults: Saving results: %+v", recValResults)

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("saveRecordValueResults: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedValues, encodeErr := generic.EncodeJSONString(recValResults.FieldValues)
	if encodeErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v",
			recValResults.FieldValues, encodeErr)
	}

	encodedMatches, encodeMatchErr := generic.EncodeJSONString(recValResults.FilterMatches)
	if encodeMatchErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v", encodeMatchErr)
	}

	if insertErr := dbSession.Query(`INSERT INTO record_val_results 
					(table_id, record_id, field_vals,filter_matches, update_timestamp_utc) 
					VALUES (?,?,?,?,toTimestamp(now()))`,
		recValResults.ParentTableID, recValResults.RecordID, encodedValues, encodedMatches).Exec(); insertErr != nil {
		return fmt.Errorf("saveRecordValueResults: Error saving results %+v: error = %v", recValResults, insertErr)
	}

	return nil
}

func GetAllRecordValueResults(parentTableID string) ([]RecordValueResults, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: Can't create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	valResultsIter := dbSession.Query(`SELECT record_id,field_vals,filter_matches,update_timestamp_utc 
		FROM record_val_results WHERE table_id = ?`,
		parentTableID).Iter()

	var currValResults RecordValueResults
	recValResults := []RecordValueResults{}
	encodedFieldVals := ""
	encodedMatches := ""
	for valResultsIter.Scan(&currValResults.RecordID, &encodedFieldVals, &encodedMatches, &currValResults.UpdateTimestamp) {
		if err := generic.DecodeJSONString(encodedFieldVals, &currValResults.FieldValues); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		if err := generic.DecodeJSONString(encodedMatches, &currValResults.FilterMatches); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		currValResults.ParentTableID = parentTableID
		recValResults = append(recValResults, currValResults)
		encodedFieldVals = ""
		encodedMatches = ""
		currValResults = RecordValueResults{}
	}
	if closeErr := valResultsIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: Failure querying database: %v", closeErr)
	}

	return recValResults, nil

}

type GetRecordValResultParams struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
}

func getRecordValueResults(params GetRecordValResultParams) (*RecordValueResults, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("getRecordValueResults: Can't create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	var valResults RecordValueResults
	valResults.ParentTableID = params.ParentTableID
	valResults.RecordID = params.RecordID
	encodedFieldVals := ""
	encodedMatches := ""
	getErr := dbSession.Query(`SELECT field_vals,filter_matches, update_timestamp_utc 
		FROM record_val_results 
		WHERE table_id=? and record_id=? LIMIT 1`,
		params.ParentTableID, params.RecordID).Scan(&encodedFieldVals, &encodedMatches, &valResults.UpdateTimestamp)
	if getErr != nil {
		return nil, fmt.Errorf("getRecordValueResults: Unabled to get record results: datastore err=%v", getErr)
	}
	if err := generic.DecodeJSONString(encodedFieldVals, &valResults.FieldValues); err != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
	}
	if err := generic.DecodeJSONString(encodedMatches, &valResults.FilterMatches); err != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
	}

	return &valResults, nil

}
