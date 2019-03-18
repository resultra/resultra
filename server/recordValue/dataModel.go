package recordValue

import (
	"resultra/tracker/server/record"
	"time"
)

type RecordValueResults struct {
	ParentDatabaseID     string                `json:"parentDatabaseID"`
	RecordID             string                `json:"recordID"`
	FieldValues          record.RecFieldValues `json:"fieldValues"`
	HiddenFormComponents []string              `json:"hiddenFormComponents"`
	UpdateTimestamp      time.Time             `json:"updateTimestamp"`
}
