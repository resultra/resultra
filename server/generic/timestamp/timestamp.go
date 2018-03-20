package timestamp

import (
	"time"
)

func CurrentTimestampUTC() time.Time {
	return time.Now().UTC()
}
