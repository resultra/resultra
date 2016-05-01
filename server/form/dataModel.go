package form

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const formEntityKind string = "Form"

type Form struct {
	Name string
}

type FormRef struct {
	FormID string `json:"formID"`
	Name   string `json:"name"`
}

type NewFormParams struct {
	TableID string `json:"tableID"`
	Name    string `json:"name"`
}

func newForm(appEngContext appengine.Context, params NewFormParams) (*FormRef, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newForm := Form{Name: sanitizedName}
	formID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext,
		params.TableID,
		formEntityKind, &newForm)
	if insertErr != nil {
		return nil, fmt.Errorf("NewForm: Unable to create new form: %v", insertErr)
	}

	log.Printf("NewForm: Created new form: id= %v, name='%v'", formID, sanitizedName)

	return &FormRef{FormID: formID, Name: sanitizedName}, nil

}

type GetFormParams struct {
	FormID string `json:"formID"`
}

func getForm(appEngContext appengine.Context, params GetFormParams) (*FormRef, error) {

	var form Form
	if getErr := datastoreWrapper.GetChildEntity(appEngContext, params.FormID, &form); getErr != nil {
		return nil, fmt.Errorf("GetForm: Unable to get form from datastore: error = %v", getErr)
	}

	formRef := FormRef{FormID: params.FormID, Name: form.Name}

	return &formRef, nil
}

func getAllForms(appEngContext appengine.Context, parentTableID string) ([]FormRef, error) {

	var forms []Form
	formIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentTableID, formEntityKind, &forms)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve forms: table id=%v", parentTableID)
	}

	formRefs := make([]FormRef, len(forms))
	for formIter, currForm := range forms {
		formID := formIDs[formIter]

		formRefs[formIter] = FormRef{
			FormID: formID,
			Name:   currForm.Name}
	} // for each form

	return formRefs, nil

}
