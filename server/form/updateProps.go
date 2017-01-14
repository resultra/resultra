package form

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
)

type FormIDInterface interface {
	getFormID() string
}

type FormIDHeader struct {
	FormID string `json:"formID"`
}

func (idHeader FormIDHeader) getFormID() string {
	return idHeader.FormID
}

type FormPropUpdater interface {
	FormIDInterface
	updateProps(form *Form) error
}

func updateFormProps(propUpdater FormPropUpdater) (*Form, error) {

	// Retrieve the bar chart from the data store
	formForUpdate, getErr := GetForm(propUpdater.getFormID())
	if getErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to get existing form: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(formForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to update existing form properties: %v", propUpdateErr)
	}

	form, updateErr := updateExistingForm(propUpdater.getFormID(), formForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to update existing form properties: datastore update error =  %v", updateErr)
	}

	return form, nil
}

type SetFormNameParams struct {
	FormIDHeader
	NewFormName string `json:"newFormName"`
}

func (updateParams SetFormNameParams) updateProps(form *Form) error {

	// TODO - Validate name

	form.Name = updateParams.NewFormName

	return nil
}

type SetLayoutParams struct {
	FormIDHeader
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (updateParams SetLayoutParams) updateProps(form *Form) error {

	form.Properties.Layout = updateParams.Layout

	return nil
}

type SetAddNewParams struct {
	FormIDHeader
	AddNewItemFromForm bool `json:"addNewItemFromForm"`
}

func (updateParams SetAddNewParams) updateProps(form *Form) error {

	form.Properties.AddNewItemFromForm = updateParams.AddNewItemFromForm

	return nil
}
