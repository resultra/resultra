package alert

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
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

func saveAlert(newAlert Alert) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newAlert.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveAlert: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO alerts (database_id,alert_id,name,properties) VALUES ($1,$2,$3,$4)`,
		newAlert.ParentDatabaseID, newAlert.AlertID, newAlert.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveForm: Can't create form: error = %v", insertErr)
	}
	return nil

}

func newAlert(params NewAlertParams) (*Alert, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newAlertProps := newDefaultAlertProperties()
	newAlertProps.FormID = params.FormID

	newAlert := Alert{ParentDatabaseID: params.ParentDatabaseID,
		AlertID:    uniqueID.GenerateSnowflakeID(),
		Name:       sanitizedName,
		Properties: newAlertProps}

	if err := saveAlert(newAlert); err != nil {
		return nil, fmt.Errorf("newAlert: error saving form: %v", err)
	}

	return &newAlert, nil
}

func GetAlert(alertID string) (*Alert, error) {

	alertName := ""
	encodedProps := ""
	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,name,properties FROM alerts
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

func getAllAlerts(parentDatabaseID string) ([]Alert, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_id,alert_id,name,properties FROM alerts WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllAlerts: Failure querying database: %v", queryErr)
	}

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

func CloneAlerts(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	remappedDatabaseID, err := remappedIDs.GetExistingRemappedID(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableForms: Error getting remapped table ID: %v", err)
	}

	alerts, err := getAllAlerts(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneAlerts: Error getting alerts for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currAlert := range alerts {

		destAlert := currAlert
		destAlert.ParentDatabaseID = remappedDatabaseID

		destAlertID, err := remappedIDs.AllocNewRemappedID(currAlert.AlertID)
		if err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
		destAlert.AlertID = destAlertID

		destProps, err := currAlert.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
		destAlert.Properties = *destProps

		if err := saveAlert(destAlert); err != nil {
			return fmt.Errorf("CloneAlerts: %v", err)
		}
	}

	return nil

}

func updateExistingAlert(alertID string, updatedAlert *Alert) (*Alert, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedAlert.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE alerts 
				SET properties=$1, name=$2
				WHERE alert_id=$3`,
		encodedProps, updatedAlert.Name, alertID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingAlert: Can't update form properties %v: error = %v",
			alertID, updateErr)
	}

	return updatedAlert, nil
}

func getAlertDatabaseID(alertID string) (string, error) {

	theAlert, err := GetAlert(alertID)
	if err != nil {
		return "", nil
	}
	return theAlert.ParentDatabaseID, nil
}

type AlertNameValidationInfo struct {
	Name string
	ID   string
}

func validateUniqueAlertName(databaseID string, alertID string, alertName string) error {
	// Query to validate the name is unique:
	// 1. Select all the alerts in the same database
	// 2. Include alerts with the same name.
	// 3. Exclude alerts with the same alert ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT alerts.alert_id,alerts.name 
			FROM alerts,databases
			WHERE databases.database_id=$1 AND
			alerts.database_id=databases.database_id AND
				alerts.name=$2 AND alerts.alert_id<>$3`,
		databaseID, alertName, alertID)
	if queryErr != nil {
		return fmt.Errorf("System error validating alert name (%v)", queryErr)
	}

	existingFormNameUsedByAnotherAlert := rows.Next()
	if existingFormNameUsedByAnotherAlert {
		return fmt.Errorf("Invalid form name - names must be unique")
	}

	return nil

}

func validateAlertName(alertID string, alertName string) error {

	if !stringValidation.WellFormedItemName(alertName) {
		return fmt.Errorf("Invalid alert name")
	}

	databaseID, err := getAlertDatabaseID(alertID)
	if err != nil {
		return fmt.Errorf("System error validating form name (%v)", err)
	}

	if uniqueErr := validateUniqueAlertName(databaseID, alertID, alertName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFormName(databaseID string, alertName string) error {

	if !stringValidation.WellFormedItemName(alertName) {
		return fmt.Errorf("Invalid alert name")
	}

	// No form will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	alertID := ""
	if uniqueErr := validateUniqueAlertName(databaseID, alertID, alertName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
