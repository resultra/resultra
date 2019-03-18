// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package trackerDatabase

import (
	"fmt"
	"resultra/tracker/server/generic/stringValidation"
)

func validateNewTrackerName(trackerName string) error {
	if !stringValidation.WellFormedItemName(trackerName) {
		return fmt.Errorf("Invalid tracker name")
	}
	return nil
}

func validateDatabaseName(databaseID string, databaseName string) error {
	if !stringValidation.WellFormedItemName(databaseName) {
		return fmt.Errorf("Invalid database name")
	}
	return nil
}
