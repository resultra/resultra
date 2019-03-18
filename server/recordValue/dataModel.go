// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
