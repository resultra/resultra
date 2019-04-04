// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package uniqueID

import (
	"fmt"
	"github.com/resultra/resultra/server/generic/timestamp"
	"github.com/twinj/uuid"
	"strings"
)

func GenerateUniqueIDWithPrefix(prefix string) string {
	timestamp := timestamp.CurrentTimestampUTC()
	millisecondsPerNanosecond := 1000000
	timestampMilliseconds := timestamp.Nanosecond() / millisecondsPerNanosecond
	timestampStr := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d",
		timestamp.Year(), timestamp.Month(), timestamp.Day(),
		timestamp.Hour(), timestamp.Minute(), timestamp.Second(),
		timestampMilliseconds)
	uuidStr := uuid.NewV4().String()

	uniqueIDStr := prefix + timestampStr + "_" + uuidStr

	return uniqueIDStr

}

/* Return a UID which also includes the current time down to the millisecond */
func GenerateUniqueIDWithTime() string {
	return GenerateUniqueIDWithPrefix("")
}

func ValidatedWellFormedID(uniqueID string) error {
	if len(uniqueID) == 0 {
		return fmt.Errorf("ValidatedWellFormedID: Empty id")
	}
	return nil
}

func GenerateV4UUIDNoDashes() string {
	uuidStr := uuid.NewV4().String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}
