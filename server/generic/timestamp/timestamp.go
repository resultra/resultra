package timestamp

import (
	"log"
	"time"
)

type TimestampGenerator func() time.Time

var globalTimestampGen TimestampGenerator

func currentUTCTimestampGen() time.Time {
	return time.Now().UTC()
}

func init() {
	globalTimestampGen = currentUTCTimestampGen
}

func OverrideProductionTimestampGeneratorWithTestFunc(timestampGenFunc TimestampGenerator) {
	log.Printf("WARNING: Overriding production timestamp generator with test function. Use with cauton and only in development & test.")
	globalTimestampGen = timestampGenFunc
}

func CurrentTimestampUTC() time.Time {
	return globalTimestampGen()
}
