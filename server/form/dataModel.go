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
	FormID        string         `json:"formID"`
	ParentTableID string         `json:"parentTableID"`
	Name          string         `json:"name"`
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
				SET properties=$1, name=$2
				WHERE form_id=$3`,
		encodedProps, updatedForm.Name, formID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingForm: Can't update form properties %v: error = %v",
			formID, updateErr)
	}

	return updatedForm, nil

}

func getFormDatabaseID(formID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id 
			FROM data_tables, forms 
			WHERE forms.form_id=$1 
				AND forms.table_id=data_tables.table_id LIMIT 1`,
		formID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getFormDatabaseID: can't get database for form = %v: err=%v",
			formID, getErr)
	}

	return databaseID, nil

}

type FormNameValidationInfo struct {
	Name string
	ID   string
}

func validateFormName(formID string, formName string) error {

	if !generic.WellFormedItemName(formName) {
		return fmt.Errorf("Invalid form name")
	}

	databaseID, err := getFormDatabaseID(formID)
	if err != nil {
		return fmt.Errorf("System error validating form name (%v)", err)
	}

	// Query to validate the name is unique:
	// 1. Select all the forms in the same database
	// 2. Include forms with the same name.
	// 3. Exclude forms with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT forms.form_id,forms.name 
			FROM forms,data_tables,databases
			WHERE databases.database_id=$1 AND
				data_tables.database_id=databases.database_id AND 
				forms.table_id=data_tables.table_id AND
				forms.name=$2 AND forms.form_id<>$3`,
		databaseID, formName, formID)
	if queryErr != nil {
		return fmt.Errorf("System error validating form name (%v)", queryErr)
	}

	existingFormNameUsedByAnotherForm := rows.Next()
	if existingFormNameUsedByAnotherForm {
		return fmt.Errorf("Invalid form name - names must be unique")
	}

	return nil
}
