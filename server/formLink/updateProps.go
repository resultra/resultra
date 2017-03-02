package formLink

import (
	"fmt"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/record"
)

type FormLinkIDInterface interface {
	getFormLinkID() string
}

type FormLinkIDHeader struct {
	FormLinkID string `json:"formLinkID"`
}

func (idHeader FormLinkIDHeader) getFormLinkID() string {
	return idHeader.FormLinkID
}

type FormLinkPropUpdater interface {
	FormLinkIDInterface
	updateProps(button *FormLink) error
}

func updateFormLinkProps(propUpdater FormLinkPropUpdater) (*FormLink, error) {

	// Retrieve the bar chart from the data store
	linkForUpdate, getErr := GetFormLink(propUpdater.getFormLinkID())
	if getErr != nil {
		return nil, fmt.Errorf("updateFormLinkProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(linkForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormLinkProps: Unable to update existing form link properties: %v", propUpdateErr)
	}

	updatedLink, updateErr := updateExistingFormLink(linkForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf(
			"updateFormLinkProps: Unable to update existing form link properties: datastore update error =  %v", updateErr)
	}

	return updatedLink, nil
}

type FormLinkDefaultValParams struct {
	FormLinkIDHeader
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func (updateParams FormLinkDefaultValParams) updateProps(linkForUpdate *FormLink) error {

	if validateErr := record.ValidateWellFormedDefaultValues(updateParams.DefaultValues); validateErr != nil {
		return fmt.Errorf("updateProps: invalid default value(s): %v", validateErr)
	}

	linkForUpdate.Properties.DefaultValues = updateParams.DefaultValues

	return nil
}

type FormLinkNameParams struct {
	FormLinkIDHeader
	NewName string `json:"newName"`
}

func (updateParams FormLinkNameParams) updateProps(linkForUpdate *FormLink) error {

	if !stringValidation.WellFormedItemLabel(updateParams.NewName) {
		return fmt.Errorf("Can't update form link name: invalid name: %v", updateParams.NewName)
	}

	linkForUpdate.Name = updateParams.NewName

	return nil
}

type FormLinkFormParams struct {
	FormLinkIDHeader
	FormID string `json:"formID"`
}

func (updateParams FormLinkFormParams) updateProps(linkForUpdate *FormLink) error {

	newForm, getErr := form.GetForm(updateParams.FormID)
	if getErr != nil {
		return fmt.Errorf("Update form properties: Unable to get form specified as new form: %v", getErr)
	}

	oldForm, getErr := form.GetForm(linkForUpdate.FormID)
	if getErr != nil {
		return fmt.Errorf("Update form properties: Unable to get form specified as new form: %v", getErr)
	}

	if newForm.ParentDatabaseID != oldForm.ParentDatabaseID {
		return fmt.Errorf("Update form properties: Database mismatch for new form: %v", newForm.ParentDatabaseID)
	}

	linkForUpdate.FormID = updateParams.FormID

	return nil
}

type FormLinkIncludeInSidebarParams struct {
	FormLinkIDHeader
	IncludeInSidebar bool `json:"includeInSidebar"`
}

func (updateParams FormLinkIncludeInSidebarParams) updateProps(linkForUpdate *FormLink) error {

	linkForUpdate.IncludeInSidebar = updateParams.IncludeInSidebar

	return nil
}

type FormLinkEnableSharedLinkParams struct {
	FormLinkIDHeader
}

func (updateParams FormLinkEnableSharedLinkParams) updateProps(linkForUpdate *FormLink) error {

	if len(linkForUpdate.SharedLinkID) <= 0 {
		// Generate a new shared link ID if there insn't one already (from previously enabling the shared link)
		linkForUpdate.SharedLinkID = uniqueID.GenerateSnowflakeID()
	}
	linkForUpdate.SharedLinkEnabled = true

	return nil
}

type FormLinkDisableSharedLinkParams struct {
	FormLinkIDHeader
}

func (updateParams FormLinkDisableSharedLinkParams) updateProps(linkForUpdate *FormLink) error {

	linkForUpdate.SharedLinkEnabled = false

	return nil
}
