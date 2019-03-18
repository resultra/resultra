// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
