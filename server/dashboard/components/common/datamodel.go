// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic"
)

func SaveNewDashboardComponent(destDBHandle *sql.DB, componentType string, parentDashboard string,
	componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewDashboardComponent: Unable to save %v: error = %v", componentType, encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO dashboard_components (dashboard_id,component_id,type,properties) 
				VALUES ($1,$2,$3,$4)`,
		parentDashboard, componentID, componentType, encodedProps); insertErr != nil {
		return fmt.Errorf("SaveNewDashboardComponent: Can't save %v: error = %v", componentType, insertErr)
	}

	return nil
}

func GetDashboardComponent(trackerDBHandle *sql.DB, componentType string, parentDashboardID string,
	componentID string, properties interface{}) error {

	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(`SELECT properties FROM dashboard_components
		 WHERE dashboard_id=$1 AND component_id=$2 AND type=$3 LIMIT 1`,
		parentDashboardID, componentID, componentType).Scan(&encodedProps)
	if getErr != nil {
		return fmt.Errorf("GetDashboardComponent: Unabled to get dashboard component: dashboard id=%v, type=%v, id=%v: datastore err=%v",
			parentDashboardID, componentType, componentID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedProps, properties); decodeErr != nil {
		return fmt.Errorf("GetFormComponent: Unabled to decode properties: encoded properties = %v: datastore err=%v",
			encodedProps, decodeErr)
	}

	return nil
}

type addComponentCallbackFunc func(string, string) error

func GetDashboardComponents(srcDBHandle *sql.DB, componentType string, parentDashboardID string, addComponentFunc addComponentCallbackFunc) error {

	rows, queryErr := srcDBHandle.Query(`SELECT component_id,properties
			FROM dashboard_components 
			WHERE dashboard_id=$1 AND type=$2`,
		parentDashboardID, componentType)
	if queryErr != nil {
		return fmt.Errorf("GetDashboardComponents: Failure querying database: %v", queryErr)
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

func UpdateDashboardComponent(trackerDBHandle *sql.DB, componentType string, parentDashboardID string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateDashboardComponent: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE dashboard_components 
				SET properties=$1
				WHERE dashboard_id=$2 AND component_id=$3`,
		encodedProps, parentDashboardID, componentID); updateErr != nil {
		return fmt.Errorf("UpdateDashboardComponent: Can't update dashboard component %v: error = %v",
			componentType, updateErr)
	}

	return nil

}

func DeleteDashboardComponent(trackerDBHandle *sql.DB, parentDashboardID string, componentID string) error {
	if _, deleteErr := trackerDBHandle.Exec(`DELETE FROM dashboard_components 
				WHERE dashboard_id=$1 AND component_id=$2`, parentDashboardID, componentID); deleteErr != nil {
		return fmt.Errorf("DeleteDashboardComponent: Can't delete form component %v: error = %v",
			componentID, deleteErr)
	}
	return nil
}
