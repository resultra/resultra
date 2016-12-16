package database

import (
	"fmt"
	"resultra/datasheet/server/generic/stringValidation"
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
