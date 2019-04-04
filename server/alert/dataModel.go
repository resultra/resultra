// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package alert

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/stringValidation"
	"github.com/resultra/resultra/server/generic/timestamp"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
	"log"
	"time"
)

const formEntityKind string = "Form"

type Alert struct {
	AlertID          string          `json:"alertID"`
	ParentDatabaseID string          `json:"parentDatabaseID"`
	Name             string          `json:"name"`
	Properties       AlertProperties `json:"properties"`
}

type NewAlertParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
	FormID           string `json:"formID"`
}

func saveAlert(destDBHandle *sql.DB, newAlert Alert) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newAlert.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveAlert: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO alerts (database_id,alert_id,name,properties) VALUES ($1,$2,$3,$4)`,
		newAlert.ParentDatabaseID, newAlert.AlertID, newAlert.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveForm: Can't create form: error = %v", insertErr)
	}
	return nil

}

func newAlert(trackerDBHandle *sql.DB, params NewAlertParams) (*Alert, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newAlertProps := newDefaultAlertProperties()
	// TODO - Verify form ID is valid
	newAlertProps.FormID = params.FormID

	newAlert := Alert{ParentDatabaseID: params.ParentDatabaseID,
		AlertID:    uniqueID.GenerateUniqueID(),
		Name:       sanitizedName,
		Properties: newAlertProps}

	if err := saveAlert(trackerDBHandle, newAlert); err != nil {
		return nil, fmt.Errorf("newAlert: error saving form: %v", err)
	}

	return &newAlert, nil
}

func GetAlert(trackerDBHandle *sql.DB, alertID string) (*Alert, error) {

	alertName := ""
	encodedProps := ""
	databaseID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT database_id,name,properties FROM alerts
		 WHERE alert_id=$1 LIMIT 1`, alertID).Scan(&databaseID, &alertName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("getAlert: Unabled to get form: form ID = %v: datastore err=%v",
			alertID, getErr)
	}

	var alertProps AlertProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &alertProps); decodeErr != nil {
		return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
	}

	getAlert := Alert{
		ParentDatabaseID: databaseID,
		AlertID:          alertID,
		Name:             alertName,
		Properties:       alertProps}

	return &getAlert, nil
}

type GetAlertListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllAlertsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]Alert, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT database_id,alert_id,name,properties FROM alerts WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllAlerts: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	alerts := []Alert{}
	for rows.Next() {
		var currAlert Alert
		encodedProps := ""

		if scanErr := rows.Scan(&currAlert.ParentDatabaseID, &currAlert.AlertID, &currAlert.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllForms: Failure querying database: %v", scanErr)
		}

		var alertProps AlertProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &alertProps); decodeErr != nil {
			return nil, fmt.Errorf("GetAllForms: can't decode properties: %v", encodedProps)
		}
		currAlert.Properties = alertProps

		alerts = append(alerts, currAlert)
	}

	return alerts, nil

}

type AdvanceNotificationParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func advanceNotificationTime(trackerDBHandle *sql.DB, userID string, parentDatabaseID string) error {

	currTimestampUTC := timestamp.CurrentTimestampUTC()

	_, insertErr := trackerDBHandle.Exec(`INSERT into alert_notification_times 
		(database_id,user_id,latest_alert_timestamp_utc) VALUES ($1,$2,$3)`, parentDatabaseID, userID, currTimestampUTC)
	if insertErr != nil {
		return fmt.Errorf("advanceNotificationTime: %v", insertErr)
	}

	_, deleteErr := trackerDBHandle.Exec(`DELETE from alert_notification_times 
		where database_id=$1 AND user_id=$2 AND latest_alert_timestamp_utc<$3`, parentDatabaseID, userID, currTimestampUTC)
	if deleteErr != nil {
		return fmt.Errorf("advanceNotificationTime: %v", deleteErr)
	}

	return nil
}

func getLatestNotificationTime(trackerDBHandle *sql.DB, userID string, parentDatabaseID string) time.Time {

	latestTime := time.Time{}

	getErr := trackerDBHandle.QueryRow(`SELECT latest_alert_timestamp_utc FROM alert_notification_times
		 WHERE database_id=$1 AND user_id=$2 LIMIT 1`, parentDatabaseID, userID).Scan(&latestTime)
	if getErr != nil {
		log.Printf("getLatestNotificationTime: error: %v", getErr)
		return latestTime
	}
	return latestTime

}

func getAllAlerts(trackerDBHandle *sql.DB, parentDatabaseID string) ([]Alert, error) {
	return getAllAlertsFromSrc(trackerDBHandle, parentDatabaseID)
}

func CloneAlerts(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	remappedDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneAlerts: Error getting remapped table ID: %v", err)
	}

	alerts, err := getAllAlertsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneAlerts: Error getting alerts for parent database ID = %v: %v",
			cloneParams.SourceDatabaseID, err)
	}

	for _, currAlert := range alerts {

		destAlert := currAlert
		destAlert.ParentDatabaseID = remappedDatabaseID

		destAlertID, err := cloneParams.IDRemapper.AllocNewRemappedID(currAlert.AlertID)
		if err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
		destAlert.AlertID = destAlertID

		destProps, err := currAlert.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
		destAlert.Properties = *destProps

		if err := saveAlert(cloneParams.DestDBHandle, destAlert); err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
	}

	return nil

}

func updateExistingAlert(trackerDBHandle *sql.DB, alertID string, updatedAlert *Alert) (*Alert, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedAlert.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE alerts 
				SET properties=$1, name=$2
				WHERE alert_id=$3`,
		encodedProps, updatedAlert.Name, alertID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingAlert: Can't update form properties %v: error = %v",
			alertID, updateErr)
	}

	return updatedAlert, nil
}

func getAlertDatabaseID(trackerDBHandle *sql.DB, alertID string) (string, error) {

	theAlert, err := GetAlert(trackerDBHandle, alertID)
	if err != nil {
		return "", nil
	}
	return theAlert.ParentDatabaseID, nil
}

type AlertNameValidationInfo struct {
	Name string
	ID   string
}

func validateUniqueAlertName(trackerDBHandle *sql.DB, databaseID string, alertID string, alertName string) error {
	// Query to validate the name is unique:
	// 1. Select all the alerts in the same database
	// 2. Include alerts with the same name.
	// 3. Exclude alerts with the same alert ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT alerts.alert_id,alerts.name 
			FROM alerts,databases
			WHERE databases.database_id=$1 AND
			alerts.database_id=databases.database_id AND
				alerts.name=$2 AND alerts.alert_id<>$3`,
		databaseID, alertName, alertID)
	if queryErr != nil {
		return fmt.Errorf("System error validating alert name (%v)", queryErr)
	}
	defer rows.Close()

	existingFormNameUsedByAnotherAlert := rows.Next()
	if existingFormNameUsedByAnotherAlert {
		return fmt.Errorf("Invalid form name - names must be unique")
	}

	return nil

}

func validateAlertName(trackerDBHandle *sql.DB, alertID string, alertName string) error {

	if !stringValidation.WellFormedItemName(alertName) {
		return fmt.Errorf("Invalid alert name")
	}

	databaseID, err := getAlertDatabaseID(trackerDBHandle, alertID)
	if err != nil {
		return fmt.Errorf("System error validating form name (%v)", err)
	}

	if uniqueErr := validateUniqueAlertName(trackerDBHandle, databaseID, alertID, alertName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFormName(trackerDBHandle *sql.DB, databaseID string, alertName string) error {

	if !stringValidation.WellFormedItemName(alertName) {
		return fmt.Errorf("Invalid alert name")
	}

	// No form will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	alertID := ""
	if uniqueErr := validateUniqueAlertName(trackerDBHandle,
		databaseID, alertID, alertName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
