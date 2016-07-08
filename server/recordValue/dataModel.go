package recordValue

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
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

	encodedValues, encodeErr := generic.EncodeJSONString(recValResults.FieldValues)
	if encodeErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v",
			recValResults.FieldValues, encodeErr)
	}

	encodedMatches, encodeMatchErr := generic.EncodeJSONString(recValResults.FilterMatches)
	if encodeMatchErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v", encodeMatchErr)
	}

	if _, delPrevErr := databaseWrapper.DBHandle().Exec(`DELETE FROM record_val_results WHERE table_id=$1 and record_id=$2`,
		recValResults.ParentTableID, recValResults.RecordID); delPrevErr != nil {
		return fmt.Errorf("saveRecordValueResults: delete previous record failed: error = %v", delPrevErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO record_val_results 
					(table_id, record_id, field_vals,filter_matches,update_timestamp_utc) 
					VALUES ($1,$2,$3,$4,$5)`,
		recValResults.ParentTableID, recValResults.RecordID, encodedValues, encodedMatches, time.Now().UTC()); insertErr != nil {
		return fmt.Errorf("saveRecordValueResults: insert failed: error = %v", insertErr)
	}

	return nil
}

func GetAllRecordValueResults(parentTableID string) ([]RecordValueResults, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT record_id,field_vals,filter_matches,update_timestamp_utc 
		FROM record_val_results WHERE table_id = $1`, parentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: Failure querying database: %v", queryErr)
	}
	recValResults := []RecordValueResults{}
	for rows.Next() {
		var currValResults RecordValueResults
		encodedFieldVals := ""
		encodedMatches := ""
		if scanErr := rows.Scan(&currValResults.RecordID, &encodedFieldVals, &encodedMatches, &currValResults.UpdateTimestamp); scanErr != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: Failure querying database: %v", scanErr)
		}
		if err := generic.DecodeJSONString(encodedFieldVals, &currValResults.FieldValues); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		if err := generic.DecodeJSONString(encodedMatches, &currValResults.FilterMatches); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		currValResults.ParentTableID = parentTableID
		recValResults = append(recValResults, currValResults)
	}

	return recValResults, nil

}

type GetRecordValResultParams struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
}

func getRecordValueResults(params GetRecordValResultParams) (*RecordValueResults, error) {

	var valResults RecordValueResults
	valResults.ParentTableID = params.ParentTableID
	valResults.RecordID = params.RecordID
	encodedFieldVals := ""
	encodedMatches := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT field_vals,filter_matches, update_timestamp_utc 
		FROM record_val_results 
		WHERE table_id=$1 and record_id=$2 LIMIT 1`,
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
