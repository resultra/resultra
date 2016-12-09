package form

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

const formEntityKind string = "Form"

type Form struct {
	FormID           string         `json:"formID"`
	ParentDatabaseID string         `json:"parentDatabaseID"`
	Name             string         `json:"name"`
	Properties       FormProperties `json:"properties"`
}

type NewFormParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
}

func saveForm(newForm Form) error {
	encodedFormProps, encodeErr := generic.EncodeJSONString(newForm.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO forms (database_id,form_id,name,properties) VALUES ($1,$2,$3,$4)`,
		newForm.ParentDatabaseID, newForm.FormID, newForm.Name, encodedFormProps); insertErr != nil {
		return fmt.Errorf("saveForm: Can't create form: error = %v", insertErr)
	}
	return nil

}

func newForm(params NewFormParams) (*Form, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newForm := Form{ParentDatabaseID: params.ParentDatabaseID,
		FormID:     uniqueID.GenerateSnowflakeID(),
		Name:       sanitizedName,
		Properties: newDefaultFormProperties()}

	if err := saveForm(newForm); err != nil {
		return nil, fmt.Errorf("newForm: error saving form: %v", err)
	}

	return &newForm, nil
}

func GetForm(formID string) (*Form, error) {

	formName := ""
	encodedProps := ""
	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,name,properties FROM forms
		 WHERE form_id=$1 LIMIT 1`, formID).Scan(&databaseID, &formName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetForm: Unabled to get form: form ID = %v: datastore err=%v",
			formID, getErr)
	}

	var formProps FormProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &formProps); decodeErr != nil {
		return nil, fmt.Errorf("GetForm: can't decode properties: %v", encodedProps)
	}

	getForm := Form{
		ParentDatabaseID: databaseID,
		FormID:           formID,
		Name:             formName,
		Properties:       formProps}

	return &getForm, nil
}

type GetFormListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllForms(parentDatabaseID string) ([]Form, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_id,form_id,name,properties FROM forms WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllForms: Failure querying database: %v", queryErr)
	}

	forms := []Form{}
	for rows.Next() {
		var currForm Form
		encodedProps := ""

		if scanErr := rows.Scan(&currForm.ParentDatabaseID, &currForm.FormID, &currForm.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("getAllForms: Failure querying database: %v", scanErr)
		}

		var formProps FormProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &formProps); decodeErr != nil {
			return nil, fmt.Errorf("getAllForms: can't decode properties: %v", encodedProps)
		}
		currForm.Properties = formProps

		forms = append(forms, currForm)
	}

	return forms, nil

}

func CloneForms(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	remappedDatabaseID, err := remappedIDs.GetExistingRemappedID(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableForms: Error getting remapped table ID: %v", err)
	}

	forms, err := getAllForms(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableForms: Error getting forms for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currForm := range forms {

		destForm := currForm
		destForm.ParentDatabaseID = remappedDatabaseID

		destFormID, err := remappedIDs.AllocNewRemappedID(currForm.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destForm.FormID = destFormID

		destProps, err := currForm.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destForm.Properties = *destProps

		if err := saveForm(destForm); err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}

		if err := cloneFormComponents(remappedIDs, currForm.FormID); err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}

	}

	return nil

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

	theForm, err := GetForm(formID)
	if err != nil {
		return "", nil
	}
	return theForm.ParentDatabaseID, nil
}

type FormNameValidationInfo struct {
	Name string
	ID   string
}

func validateUniqueFormName(databaseID string, formID string, formName string) error {
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

func validateFormName(formID string, formName string) error {

	if !stringValidation.WellFormedItemName(formName) {
		return fmt.Errorf("Invalid form name")
	}

	databaseID, err := getFormDatabaseID(formID)
	if err != nil {
		return fmt.Errorf("System error validating form name (%v)", err)
	}

	if uniqueErr := validateUniqueFormName(databaseID, formID, formName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFormName(databaseID string, formName string) error {

	if !stringValidation.WellFormedItemName(formName) {
		return fmt.Errorf("Invalid form name")
	}

	// No form will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	formID := ""
	if uniqueErr := validateUniqueFormName(databaseID, formID, formName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
