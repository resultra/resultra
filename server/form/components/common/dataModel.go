package common

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

func SaveNewFormComponent(componentType string, parentForm string, componentID string, properties interface{}) error {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("saveNewCheckBox: Unable to save %v: error = %v", componentType, encodeErr)
	}

	if insertErr := dbSession.Query(`INSERT INTO form_components (form_id,component_id,type,properties) VALUES (?,?,?,?)`,
		parentForm, componentID, componentType, encodedProps).Exec(); insertErr != nil {
		return fmt.Errorf("saveNewFormComponent: Can't save %v: error = %v", componentType, insertErr)
	}

	return nil
}

func GetFormComponent(componentType string, parentFormID string, componentID string, properties interface{}) error {
	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProps := ""
	getErr := dbSession.Query(`SELECT properties FROM form_components
		 WHERE form_id=? AND component_id=? AND type=? LIMIT 1`,
		parentFormID, componentID, componentType).Scan(&encodedProps)
	if getErr != nil {
		return fmt.Errorf("GetRecord: Unabled to get form component %v: id = %v: datastore err=%v",
			componentType, componentID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedProps, properties); decodeErr != nil {
		return fmt.Errorf("GetFormComponent: Unabled to decode properties: encoded properties = %v: datastore err=%v",
			encodedProps, decodeErr)
	}

	return nil
}

type addComponentCallbackFunc func(string, string) error

func GetFormComponents(componentType string, parentFormID string, addComponentFunc addComponentCallbackFunc) error {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("GetFormComponents: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	componentIter := dbSession.Query(`SELECT component_id,properties
			FROM form_components 
			WHERE form_id=? AND type=?`,
		parentFormID, componentType).Iter()

	currComponentID := ""
	encodedProps := ""
	for componentIter.Scan(&currComponentID, &encodedProps) {
		log.Printf("GetFormComponents: Got form component: component id = %v, properties=%v",
			currComponentID, encodedProps)
		if err := addComponentFunc(currComponentID, encodedProps); err != nil {
			return err
		}
		currComponentID = ""
		encodedProps = ""
	}
	return nil
}

func UpdateFormComponent(componentType string, parentFormID string, componentID string, properties interface{}) error {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateFormComponent: failure encoding properties: error = %v", encodeErr)
	}

	if updateErr := dbSession.Query(`UPDATE form_components 
				SET properties=? 
				WHERE form_id=? AND component_id=?`,
		encodedProps, parentFormID, componentID).Exec(); updateErr != nil {
		return fmt.Errorf("UpdateFormComponent: Can't update form component %v: error = %v",
			componentType, updateErr)
	}

	return nil

}
