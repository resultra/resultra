package form

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

const formEntityKind string = "Form"

type Form struct {
	FormID        string `json:"formID"`
	ParentTableID string `json:"parentTableID"`
	Name          string
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
		FormID: databaseWrapper.GlobalUniqueID(),
		Name:   sanitizedName}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO forms (table_id,form_id,name) VALUES ($1,$2,$3)`,
		newForm.ParentTableID, newForm.FormID, newForm.Name); insertErr != nil {
		return nil, fmt.Errorf("newForm: Can't create form: error = %v", insertErr)
	}

	log.Printf("NewForm: Created new form: %+v", newForm)

	return &newForm, nil
}

type GetFormParams struct {
	ParentTableID string `json:"parentTableID"`
	FormID        string `json:"formID"`
}

func GetForm(params GetFormParams) (*Form, error) {

	formName := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT name FROM forms
		 WHERE table_id=$1 AND form_id=$2 LIMIT 1`,
		params.ParentTableID, params.FormID).Scan(&formName)
	if getErr != nil {
		return nil, fmt.Errorf("GetForm: Unabled to get form: params = %+v: datastore err=%v",
			params, getErr)
	}

	getForm := Form{
		ParentTableID: params.ParentTableID,
		FormID:        params.FormID,
		Name:          formName}

	return &getForm, nil
}

func getAllForms(parentTableID string) ([]Form, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT table_id,form_id,name FROM forms WHERE table_id = $1`,
		parentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllForms: Failure querying database: %v", queryErr)
	}

	forms := []Form{}
	for rows.Next() {
		var currForm Form
		if scanErr := rows.Scan(&currForm.ParentTableID, &currForm.FormID, &currForm.Name); scanErr != nil {
			return nil, fmt.Errorf("getAllForms: Failure querying database: %v", scanErr)
		}
		forms = append(forms, currForm)
	}

	return forms, nil

}
