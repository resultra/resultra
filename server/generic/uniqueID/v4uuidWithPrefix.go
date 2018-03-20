package uniqueID

import (
	"fmt"
	"github.com/twinj/uuid"
	"resultra/datasheet/server/generic/timestamp"
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
func GenerateUniqueID() string {
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
