package form

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const formEntityKind string = "Form"

type Form struct {
	FormID        string `json:"formID"`
	ParentTableID string `json:"parentTableID"`
	Name          string
	Properties    FormProperties `json:"properties"`
}

type NewFormParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
}

func newForm(params NewFormParams) (*Form, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newForm := Form{ParentTableID: params.ParentTableID,
		FormID: uniqueID.GenerateSnowflakeID(),
		Name:   sanitizedName}

	formProps := FormProperties{}
	encodedFormProps, encodeErr := generic.EncodeJSONString(formProps)
	if encodeErr != nil {
		return nil, fmt.Errorf("newForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO forms (table_id,form_id,name,properties) VALUES ($1,$2,$3)`,
		newForm.ParentTableID, newForm.FormID, newForm.Name, encodedFormProps); insertErr != nil {
		return nil, fmt.Errorf("newForm: Can't create form: error = %v", insertErr)
	}

	log.Printf("NewForm: Created new form: %+v", newForm)

	return &newForm, nil
}

func GetForm(formID string) (*Form, error) {

	formName := ""
	encodedProps := ""
	tableID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT table_id,name,properties FROM forms
		 WHERE form_id=$1 LIMIT 1`, formID).Scan(&tableID, &formName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetForm: Unabled to get form: form ID = %v: datastore err=%v",
			formID, getErr)
	}

	var formProps FormProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &formProps); decodeErr != nil {
		return nil, fmt.Errorf("GetForm: can't decode properties: %v", encodedProps)
	}

	getForm := Form{
		ParentTableID: tableID,
		FormID:        formID,
		Name:          formName,
		Properties:    formProps}

	return &getForm, nil
}

func getAllForms(parentTableID string) ([]Form, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT table_id,form_id,name,properties FROM forms WHERE table_id = $1`,
		parentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllForms: Failure querying database: %v", queryErr)
	}

	forms := []Form{}
	for rows.Next() {
		var currForm Form
		encodedProps := ""

		if scanErr := rows.Scan(&currForm.ParentTableID, &currForm.FormID, &currForm.Name, encodedProps); scanErr != nil {
			return nil, fmt.Errorf("getAllForms: Failure querying database: %v", scanErr)
		}

		var formProps FormProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &formProps); decodeErr != nil {
			return nil, fmt.Errorf("GetForm: can't decode properties: %v", encodedProps)
		}
		currForm.Properties = formProps

		forms = append(forms, currForm)
	}

	return forms, nil

}

func updateExistingForm(formID string, updatedForm *Form) (*Form, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedForm.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE forms 
				SET properties=$1
				WHERE form_id=$2`,
		encodedProps, formID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingForm: Can't update form properties %v: error = %v",
			formID, updateErr)
	}

	return updatedForm, nil

}
