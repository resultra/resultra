package recordValue

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
	"resultra/datasheet/server/record"
)

type RecordValueResults struct {
	ParentTableID string                `json:"parentTableID"`
	RecordID      string                `json:"recordID"`
	FieldValues   record.RecFieldValues `json:"fieldValues"`
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
