package generic

import (
	"fmt"
	"github.com/twinj/uuid"
	"time"
)

/* Return a UID which also includes the current time down to the millisecond */
func GenerateUniqueID() string {
	timestamp := time.Now().UTC()
	millisecondsPerNanosecond := 1000000
	timestampMilliseconds := timestamp.Nanosecond() / millisecondsPerNanosecond
	timestampStr := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d",
		timestamp.Year(), timestamp.Month(), timestamp.Day(),
		timestamp.Hour(), timestamp.Minute(), timestamp.Second(),
		timestampMilliseconds)
	uuidStr := uuid.NewV4().String()

	uniqueIDStr := timestampStr + "_" + uuidStr

	return uniqueIDStr

}
