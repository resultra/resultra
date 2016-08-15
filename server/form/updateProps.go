package form

import (
	"fmt"
)

type FormProperties struct {
	DefaultFilterIDs   []string `json:"defaultFilterIDs"`
	AvailableFilterIDs []string `json:"availableFilterIDs"`
}

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

type FormDefaultFilterParams struct {
	FormIDHeader
	DefaultFilterIDs []string `json:"defaultFilterIDs"`
}

func (updateParams FormDefaultFilterParams) updateProps(form *Form) error {

	form.Properties.DefaultFilterIDs = updateParams.DefaultFilterIDs

	return nil
}

type FormAvailableFilterParams struct {
	FormIDHeader
	AvailableFilterIDs []string `json:"availableFilterIDs"`
}

func (updateParams FormAvailableFilterParams) updateProps(form *Form) error {

	form.Properties.AvailableFilterIDs = updateParams.AvailableFilterIDs

	return nil
}

type SetFormNameParams struct {
	FormIDHeader
	NewFormName string `json:"newFormName"`
}

func (updateParams SetFormNameParams) updateProps(form *Form) error {

	form.Name = updateParams.NewFormName

	return nil
}
