// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/generic"
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

func GetFormComponentFormID(trackerDBHandle *sql.DB, componentID string) (string, error) {

	formID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT form_id FROM form_components
		 WHERE component_id=$1 LIMIT 1`,
		componentID).Scan(&formID)
	if getErr != nil {
		return "", fmt.Errorf("GetFormComponentFormID: Unabled to get table id for column: id = %v: datastore err=%v",
			componentID, getErr)
	}
	return formID, nil

}

func GetFormComponent(trackerDBHandle *sql.DB, componentType string, parentFormID string, componentID string, properties interface{}) error {

	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(`SELECT properties FROM form_components
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
	defer rows.Close()

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

func UpdateFormComponent(trackerDBHandle *sql.DB, componentType string, parentFormID string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateFormComponent: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE form_components 
				SET properties=$1
				WHERE form_id=$2 AND component_id=$3`,
		encodedProps, parentFormID, componentID); updateErr != nil {
		return fmt.Errorf("UpdateFormComponent: Can't update form component %v: error = %v",
			componentType, updateErr)
	}

	return nil

}

func DeleteFormComponent(trackerDBHandle *sql.DB, parentFormID string, componentID string) error {
	if _, deleteErr := trackerDBHandle.Exec(`DELETE FROM form_components 
				WHERE form_id=$1 AND component_id=$2`, parentFormID, componentID); deleteErr != nil {
		return fmt.Errorf("DeleteFormComponent: Can't delete form component %v: error = %v",
			componentID, deleteErr)
	}
	return nil
}
