package recordValue

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/record"
	"time"
)

type RecordValueResults struct {
	ParentDatabaseID     string                `json:"parentDatabaseID"`
	RecordID             string                `json:"recordID"`
	FieldValues          record.RecFieldValues `json:"fieldValues"`
	HiddenFormComponents []string              `json:"hiddenFormComponents"`
	UpdateTimestamp      time.Time             `json:"updateTimestamp"`
}

func SaveRecordValueResults(recValResults RecordValueResults) error {

	encodedValues, encodeErr := generic.EncodeJSONString(recValResults.FieldValues)
	if encodeErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v",
			recValResults.FieldValues, encodeErr)
	}

	encodedHiddenComponents, encodeErr := generic.EncodeJSONString(recValResults.HiddenFormComponents)
	if encodeErr != nil {
		return fmt.Errorf("saveRecordValueResults: Unable to encode record value results %+v: error = %v",
			recValResults.FieldValues, encodeErr)
	}

	if _, delPrevErr := databaseWrapper.DBHandle().Exec(`DELETE FROM record_val_results WHERE database_id=$1 and record_id=$2`,
		recValResults.ParentDatabaseID, recValResults.RecordID); delPrevErr != nil {
		return fmt.Errorf("saveRecordValueResults: delete previous record failed: error = %v", delPrevErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO record_val_results 
					(database_id, record_id, field_vals,hidden_form_components, update_timestamp_utc) 
					VALUES ($1,$2,$3,$4,$5)`,
		recValResults.ParentDatabaseID, recValResults.RecordID, encodedValues, encodedHiddenComponents, time.Now().UTC()); insertErr != nil {
		return fmt.Errorf("saveRecordValueResults: insert failed: error = %v", insertErr)
	}

	return nil
}

func GetAllRecordValueResults(parentDatabaseID string) ([]RecordValueResults, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT record_val_results.record_id,record_val_results.field_vals,record_val_results.hidden_form_components,record_val_results.update_timestamp_utc 
		 FROM record_val_results,records WHERE 
		 record_val_results.record_id = records.record_id AND
		 records.is_draft_record = $1 AND
		 record_val_results.database_id = $2`, false, parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: Failure querying database: %v", queryErr)
	}
	recValResults := []RecordValueResults{}
	for rows.Next() {
		var currValResults RecordValueResults
		encodedFieldVals := ""
		encodedHiddenComponents := ""
		if scanErr := rows.Scan(&currValResults.RecordID, &encodedFieldVals, &encodedHiddenComponents, &currValResults.UpdateTimestamp); scanErr != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: Failure querying database: %v", scanErr)
		}
		if err := generic.DecodeJSONString(encodedFieldVals, &currValResults.FieldValues); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
		}
		if err := generic.DecodeJSONString(encodedHiddenComponents, &currValResults.HiddenFormComponents); err != nil {
			return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding hidden component values: %v", err)
		}

		currValResults.ParentDatabaseID = parentDatabaseID
		recValResults = append(recValResults, currValResults)
	}

	return recValResults, nil

}

type GetRecordValResultParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RecordID         string `json:"recordID"`
}

func getRecordValueResults(params GetRecordValResultParams) (*RecordValueResults, error) {

	var valResults RecordValueResults
	valResults.ParentDatabaseID = params.ParentDatabaseID
	valResults.RecordID = params.RecordID
	encodedFieldVals := ""
	encodedHiddenComponents := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT field_vals,hidden_form_components,update_timestamp_utc 
		FROM record_val_results 
		WHERE database_id=$1 and record_id=$2 LIMIT 1`,
		params.ParentDatabaseID, params.RecordID).Scan(&encodedFieldVals, &encodedHiddenComponents, &valResults.UpdateTimestamp)
	if getErr != nil {
		return nil, fmt.Errorf("getRecordValueResults: Unabled to get record results: datastore err=%v", getErr)
	}
	if err := generic.DecodeJSONString(encodedFieldVals, &valResults.FieldValues); err != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding field values: %v", err)
	}
	if err := generic.DecodeJSONString(encodedHiddenComponents, &valResults.HiddenFormComponents); err != nil {
		return nil, fmt.Errorf("GetAllRecordValueResults: failure decoding hidden component values: %v", err)
	}

	return &valResults, nil

}
