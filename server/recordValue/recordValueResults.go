package recordValue

import (
	"resultra/datasheet/server/record"
)

type RecordValueResults struct {
	ParentTableID string                `jsaon:"parentTableID"`
	RecordID      string                `json:"recordID"`
	FieldValues   record.RecFieldValues `json:"fieldValues" datastore:"-"`
}
