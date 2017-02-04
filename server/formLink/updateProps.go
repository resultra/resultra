package formLink

import (
	"fmt"
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
