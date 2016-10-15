package header

import (
	"fmt"
	"resultra/datasheet/server/generic/stringValidation"
)

type HeaderIDInterface interface {
	getHeaderID() string
	getParentFormID() string
}

type HeaderIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	HeaderID     string `json:"headerID"`
}

func (idHeader HeaderIDHeader) getHeaderID() string {
	return idHeader.HeaderID
}

func (idHeader HeaderIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type HeaderPropUpdater interface {
	HeaderIDInterface
	updateProps(header *Header) error
}

func updateHeaderProps(propUpdater HeaderPropUpdater) (*Header, error) {

	// Retrieve the bar chart from the data store
	headerForUpdate, getErr := getHeader(propUpdater.getParentFormID(), propUpdater.getHeaderID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to get existing header: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(headerForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header properties: %v", propUpdateErr)
	}

	updatedHeader, updateErr := updateExistingHeader(headerForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header properties: datastore update error =  %v", updateErr)
	}

	return updatedHeader, nil
}

type HeaderLabelParams struct {
	HeaderIDHeader
	Label string `json:"label"`
}

func (updateParams HeaderLabelParams) updateProps(headerForUpdate *Header) error {

	if !stringValidation.WellFormedItemLabel(updateParams.Label) {
		return fmt.Errorf("Update header label: invalid label: %v", updateParams.Label)
	}

	headerForUpdate.Properties.Label = updateParams.Label

	return nil
}
