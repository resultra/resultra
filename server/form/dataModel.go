package form

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

const formEntityKind string = "Form"

type Form struct {
	FormID        string `json:"formID"`
	ParentTableID string `json:"parentTableID"`
	Name          string
}

const formIDFieldName string = "FormID"
const formParentTableFieldName string = "ParentTableID"

type NewFormParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
}

func newForm(appEngContext appengine.Context, params NewFormParams) (*Form, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newForm := Form{ParentTableID: params.ParentTableID,
		FormID: gocql.TimeUUID().String(),
		Name:   sanitizedName}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(`INSERT INTO form (tableID,formID,name) VALUES (?,?,?)`,
		newForm.ParentTableID, newForm.FormID, newForm.Name).Exec(); insertErr != nil {
		return nil, fmt.Errorf("newForm: Can't create form: error = %v", insertErr)
	}

	log.Printf("NewForm: Created new form: %+v", newForm)

	return &newForm, nil
}

type GetFormParams struct {
	ParentTableID string `json:"parentTableID"`
	FormID        string `json:"formID"`
}

func GetForm(appEngContext appengine.Context, params GetFormParams) (*Form, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetForm: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	formName := ""
	getErr := dbSession.Query(`SELECT name FROM form
		 WHERE tableid=? AND formid=? LIMIT 1`,
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

func getAllForms(appEngContext appengine.Context, parentTableID string) ([]Form, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("getTableList: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	formIter := dbSession.Query(`SELECT tableID,formID,name FROM dataTable WHERE tableID = ?`,
		parentTableID).Iter()

	var currForm Form
	forms := []Form{}
	for formIter.Scan(&currForm.ParentTableID, &currForm.FormID, &currForm.Name) {
		forms = append(forms, currForm)
	}
	if closeErr := formIter.Close(); closeErr != nil {
		fmt.Errorf("getAllForms: Failure querying database: %v", closeErr)
	}

	return forms, nil

}
