package form

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
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

	if err := uniqueID.ValidatedWellFormedID(params.ParentTableID); err != nil {
		return nil, fmt.Errorf("Can't create new form: invalid parent table: %v", err)
	}

	newForm := Form{ParentTableID: params.ParentTableID,
		FormID: uniqueID.GenerateUniqueID(),
		Name:   sanitizedName}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, formEntityKind, &newForm)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new image component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("NewForm: Created new form: %+v", newForm)

	return &newForm, nil
}

type GetFormParams struct {
	FormID string `json:"formID"`
}

func GetForm(appEngContext appengine.Context, params GetFormParams) (*Form, error) {

	var form Form

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, formEntityKind,
		formIDFieldName, params.FormID, &form); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to get form from datastore: error = %v", getErr)
	}

	return &form, nil
}

func getAllForms(appEngContext appengine.Context, parentTableID string) ([]Form, error) {

	var forms []Form

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentTableID,
		formEntityKind, formParentTableFieldName, &forms)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve form: table id=%v", parentTableID)
	}

	return forms, nil

}
