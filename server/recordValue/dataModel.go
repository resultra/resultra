package recordValue

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
	"resultra/datasheet/server/record"
	"time"
)

type RecordValueResults struct {
	ParentTableID   string                `json:"parentTableID"`
	RecordID        string                `json:"recordID"`
	FieldValues     record.RecFieldValues `json:"fieldValues"`
	UpdateTimestamp time.Time             `json:"updateTimestamp"`
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
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v", encodeErr)
	}

	if insertErr := dbSession.Query(`INSERT INTO record_val_results 
					(table_id, record_id, field_vals,update_timestamp_utc) 
					VALUES (?,?,?,toTimestamp(now()))`,
		recValResults.ParentTableID, recValResults.RecordID, encodedValues).Exec(); insertErr != nil {
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

	valResultsIter := dbSession.Query(`SELECT record_id,field_vals,update_timestamp_utc 
		FROM record_val_results WHERE table_id = ?`,
		parentTableID).Iter()

	var currValResults RecordValueResults
	recValResults := []RecordValueResults{}
	encodedFieldVals := ""
	for valResultsIter.Scan(&currValResults.RecordID, &encodedFieldVals, &currValResults.UpdateTimestamp) {
		var currFieldVals record.RecFieldValues
		if err := generic.DecodeJSONString(encodedFieldVals, &currFieldVals); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		currValResults.ParentTableID = parentTableID
		currValResults.FieldValues = currFieldVals
		recValResults = append(recValResults, currValResults)
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
		return nil, fmt.Errorf("GetAllRecordValueResults: Can't create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	var valResults RecordValueResults
	valResults.ParentTableID = params.ParentTableID
	valResults.RecordID = params.RecordID
	encodedFieldVals := ""
	getErr := dbSession.Query(`SELECT field_vals,update_timestamp_utc 
		FROM record_val_results 
		WHERE table_id=? and record_id=? LIMIT 1`,
		params.ParentTableID, params.RecordID).Scan(&encodedFieldVals, &valResults.UpdateTimestamp)
	if getErr != nil {
		return nil, fmt.Errorf("GetForm: Unabled to get dashboard: datastore err=%v", getErr)
	}
	var fieldVals record.RecFieldValues
	if err := generic.DecodeJSONString(encodedFieldVals, &fieldVals); err != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
	}
	valResults.FieldValues = fieldVals

	return &valResults, nil

}
