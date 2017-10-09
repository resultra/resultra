package common

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/common/databaseWrapper"
)

func SaveNewDashboardComponent(componentType string, parentDashboard string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewDashboardComponent: Unable to save %v: error = %v", componentType, encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO dashboard_components (dashboard_id,component_id,type,properties) 
				VALUES ($1,$2,$3,$4)`,
		parentDashboard, componentID, componentType, encodedProps); insertErr != nil {
		return fmt.Errorf("SaveNewDashboardComponent: Can't save %v: error = %v", componentType, insertErr)
	}

	return nil
}

func GetDashboardComponent(componentType string, parentDashboardID string, componentID string, properties interface{}) error {

	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT properties FROM dashboard_components
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

func GetDashboardComponents(componentType string, parentDashboardID string, addComponentFunc addComponentCallbackFunc) error {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT component_id,properties
			FROM dashboard_components 
			WHERE dashboard_id=$1 AND type=$2`,
		parentDashboardID, componentType)
	if queryErr != nil {
		return fmt.Errorf("GetDashboardComponents: Failure querying database: %v", queryErr)
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

func UpdateDashboardComponent(componentType string, parentDashboardID string, componentID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateDashboardComponent: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE dashboard_components 
				SET properties=$1
				WHERE dashboard_id=$2 AND component_id=$3`,
		encodedProps, parentDashboardID, componentID); updateErr != nil {
		return fmt.Errorf("UpdateDashboardComponent: Can't update dashboard component %v: error = %v",
			componentType, updateErr)
	}

	return nil

}
