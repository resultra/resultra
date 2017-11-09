package record

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
	"sync"
	"time"
)

var newRecordMutex = &sync.Mutex{}

type Record struct {
	ParentDatabaseID   string    `json:"parentDatabaseID"`
	RecordID           string    `json:"recordID"`
	IsDraftRecord      bool      `json:"isDraftRecord"`
	CreateTimestampUTC time.Time `json:"createTimestampUTC"`
	SequenceNum        int       `json:"sequenceNum"`
}

type NewRecordParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	IsDraftRecord    bool   `json:"isDraftRecord"`
}

func NewRecord(trackerDBHandle *sql.DB, params NewRecordParams) (*Record, error) {

	// Use a mutex when creating records. This is necessary, since a new sequence number is allocated for the record when it is
	// created. This also means only a single process can interact with the database for the purposes of creating records.
	newRecordMutex.Lock()
	defer newRecordMutex.Unlock()

	createTimestamp := time.Now().UTC()

	newRecord := Record{ParentDatabaseID: params.ParentDatabaseID,
		RecordID:           uniqueID.GenerateSnowflakeID(),
		IsDraftRecord:      params.IsDraftRecord,
		CreateTimestampUTC: createTimestamp}

	// Allocate the next sequence number, using the coalesce syntax.
	// Solution based upon the following: https://stackoverflow.com/questions/7452501/postgresql-turn-null-into-zero
	nextSequenceNum := 0
	if getErr := trackerDBHandle.QueryRow(
		`SELECT coalesce(max(sequence_num), 0)+1 as next_sequence_num FROM records WHERE database_id=$1`, params.ParentDatabaseID).Scan(
		&nextSequenceNum); getErr != nil {
		return nil, fmt.Errorf("NewRecord: Unabled to allocate sequence number: database id = %v: datastore err=%v",
			params.ParentDatabaseID, getErr)
	}

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO records (database_id, record_id,is_draft_record,create_timestamp_utc,sequence_num) VALUES ($1,$2,$3,$4,$5)`,
		newRecord.ParentDatabaseID,
		newRecord.RecordID,
		newRecord.IsDraftRecord,
		newRecord.CreateTimestampUTC,
		nextSequenceNum); insertErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", insertErr)
	}

	return &newRecord, nil

}

func GetRecord(trackerDBHandle *sql.DB, recordID string) (*Record, error) {

	getRecord := Record{}

	if getErr := trackerDBHandle.QueryRow(
		`SELECT database_id,record_id,is_draft_record,create_timestamp_utc,sequence_num
			 	FROM records WHERE record_id=$1 LIMIT 1`, recordID).Scan(
		&getRecord.ParentDatabaseID,
		&getRecord.RecordID,
		&getRecord.IsDraftRecord,
		&getRecord.CreateTimestampUTC,
		&getRecord.SequenceNum); getErr != nil {
		return nil, fmt.Errorf("GetRecord: Unabled to get record: id = %v: datastore err=%v", recordID, getErr)
	}

	return &getRecord, nil

}

func GetNonDraftRecords(trackerDBHandle *sql.DB, parentDatabaseID string) ([]Record, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT database_id,record_id,is_draft_record,create_timestamp_utc,sequence_num FROM records 
		WHERE database_id=$1 AND is_draft_record=0`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRecords: Failure querying database: %v", queryErr)
	}
	records := []Record{}
	for rows.Next() {
		var currRecord Record
		if scanErr := rows.Scan(&currRecord.ParentDatabaseID,
			&currRecord.RecordID,
			&currRecord.IsDraftRecord,
			&currRecord.CreateTimestampUTC,
			&currRecord.SequenceNum); scanErr != nil {
			return nil, fmt.Errorf("GetRecords: Failure querying database: %v", scanErr)

		}
		records = append(records, currRecord)
	}

	return records, nil
}

func GetNonDraftRecordIDRecordMap(trackingDBHandle *sql.DB, parentDatabaseID string) (map[string]Record, error) {

	records, err := GetNonDraftRecords(trackingDBHandle, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetNonDraftRecordIDRecordMap: %v", err)
	}

	recordIDRecordMap := map[string]Record{}
	for _, currRecord := range records {
		recordIDRecordMap[currRecord.RecordID] = currRecord
	}

	return recordIDRecordMap, nil
}

type SetDraftStatusParams struct {
	RecordID      string `json:"recordID"`
	IsDraftRecord bool   `json:"isDraftRecord"`
}

func setDraftStatus(trackerDBHandle *sql.DB, params SetDraftStatusParams) error {

	if _, updateErr := trackerDBHandle.Exec(`UPDATE records 
				SET is_draft_record=$1
				WHERE record_id=$2`,
		params.IsDraftRecord, params.RecordID); updateErr != nil {
		return fmt.Errorf("setDraftStatus: Can't update database properties %+v: error = %v",
			params, updateErr)
	}
	return nil
}

// Validate the field is of the correct type and not a calculated field (if allowCalcField not true). This is for validating
// the field when setting/getting values from regular "literal" fields which store values entered by end-users (as opposed to
// calculated fields)
func ValidateFieldForRecordValue(field field.Field, expectedFieldType string,
	allowCalcField bool) error {

	if field.Type != expectedFieldType {
		return fmt.Errorf("Can't update/set value:"+
			" Type mismatch with field: expecting %v: got %v: field=%+v", expectedFieldType, field.Type, field)

	} else if (!allowCalcField) && field.IsCalcField {
		return fmt.Errorf("Field is a calculated field, setting values directly not supported: field=%v",
			field.RefName)
	}
	return nil
}
