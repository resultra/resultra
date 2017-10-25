package common

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic"
)

func SaveNewFormComponent(destDBHandle *sql.DB,
	componentType string, parentForm string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("saveNewCheckBox: Unable to save %v: error = %v", componentType, encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO form_components (form_id,component_id,type,properties) VALUES ($1,$2,$3,$4)`,
		parentForm, componentID, componentType, encodedProps); insertErr != nil {
		return fmt.Errorf("saveNewFormComponent: Can't save %v: error = %v", componentType, insertErr)
	}

	return nil
}

func GetFormComponentFormID(componentID string) (string, error) {

	formID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT form_id FROM form_components
		 WHERE component_id=$1 LIMIT 1`,
		componentID).Scan(&formID)
	if getErr != nil {
		return "", fmt.Errorf("GetFormComponentFormID: Unabled to get table id for column: id = %v: datastore err=%v",
			componentID, getErr)
	}
	return formID, nil

}

func GetFormComponent(componentType string, parentFormID string, componentID string, properties interface{}) error {

	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT properties FROM form_components
		 WHERE form_id=$1 AND component_id=$2 AND type=$3 LIMIT 1`,
		parentFormID, componentID, componentType).Scan(&encodedProps)
	if getErr != nil {
		return fmt.Errorf("GetFormComponent: Unabled to get form component %v: id = %v: datastore err=%v",
			componentType, componentID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedProps, properties); decodeErr != nil {
		return fmt.Errorf("GetFormComponent: Unabled to decode properties: encoded properties = %v: datastore err=%v",
			encodedProps, decodeErr)
	}

	return nil
}

type addComponentCallbackFunc func(string, string) error

func GetFormComponents(srcDBHandle *sql.DB, componentType string, parentFormID string, addComponentFunc addComponentCallbackFunc) error {

	rows, queryErr := srcDBHandle.Query(`SELECT component_id,properties
			FROM form_components 
			WHERE form_id=$1 AND type=$2`,
		parentFormID, componentType)
	if queryErr != nil {
		return fmt.Errorf("GetFormComponents: Failure querying database: %v", queryErr)
	}

	for rows.Next() {
		currComponentID := ""
		encodedProps := ""
		if scanErr := rows.Scan(&currComponentID, &encodedProps); scanErr != nil {
			return fmt.Errorf("GetFormComponents: Failure querying database: %v", scanErr)
		}
		if err := addComponentFunc(currComponentID, encodedProps); err != nil {
			return err
		}
	}

	return nil
}

func UpdateFormComponent(componentType string, parentFormID string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateFormComponent: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE form_components 
				SET properties=$1
				WHERE form_id=$2 AND component_id=$3`,
		encodedProps, parentFormID, componentID); updateErr != nil {
		return fmt.Errorf("UpdateFormComponent: Can't update form component %v: error = %v",
			componentType, updateErr)
	}

	return nil

}

func DeleteFormComponent(parentFormID string, componentID string) error {
	if _, deleteErr := databaseWrapper.DBHandle().Exec(`DELETE FROM form_components 
				WHERE form_id=$1 AND component_id=$2`, parentFormID, componentID); deleteErr != nil {
		return fmt.Errorf("DeleteFormComponent: Can't delete form component %v: error = %v",
			componentID, deleteErr)
	}
	return nil
}
