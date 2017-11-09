package urlLink

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type UrlLinkIDInterface interface {
	getUrlLinkID() string
	getParentTableID() string
}

type UrlLinkIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	UrlLinkID     string `json:"urlLinkID"`
}

func (idHeader UrlLinkIDHeader) getUrlLinkID() string {
	return idHeader.UrlLinkID
}

func (idHeader UrlLinkIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type UrlLinkPropUpdater interface {
	UrlLinkIDInterface
	updateProps(urlLink *UrlLink) error
}

func updateUrlLinkProps(trackerDBHandle *sql.DB, propUpdater UrlLinkPropUpdater) (*UrlLink, error) {

	// Retrieve the bar chart from the data store
	urlLinkForUpdate, getErr := getUrlLink(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getUrlLinkID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(urlLinkForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	urlLink, updateErr := updateExistingUrlLink(trackerDBHandle, propUpdater.getUrlLinkID(), urlLinkForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return urlLink, nil
}

type UrlLinkLabelFormatParams struct {
	UrlLinkIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams UrlLinkLabelFormatParams) updateProps(urlLink *UrlLink) error {

	// TODO - Validate format is well-formed.

	urlLink.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type UrlLinkPermissionParams struct {
	UrlLinkIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams UrlLinkPermissionParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.Permissions = updateParams.Permissions

	return nil
}

type UrlLinkValidationParams struct {
	UrlLinkIDHeader
	Validation UrlLinkValidationProperties `json:"validation"`
}

func (updateParams UrlLinkValidationParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.Validation = updateParams.Validation

	return nil
}

type UrlLinkClearValueSupportedParams struct {
	UrlLinkIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams UrlLinkClearValueSupportedParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	UrlLinkIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
