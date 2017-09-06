package emailAddr

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type EmailAddrIDInterface interface {
	getEmailAddrID() string
	getParentFormID() string
}

type EmailAddrIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	EmailAddrID  string `json:"emailAddrID"`
}

func (idHeader EmailAddrIDHeader) getEmailAddrID() string {
	return idHeader.EmailAddrID
}

func (idHeader EmailAddrIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type EmailAddrPropUpdater interface {
	EmailAddrIDInterface
	updateProps(emailAddr *EmailAddr) error
}

func updateEmailAddrProps(propUpdater EmailAddrPropUpdater) (*EmailAddr, error) {

	// Retrieve the bar chart from the data store
	emailAddrForUpdate, getErr := getEmailAddr(propUpdater.getParentFormID(), propUpdater.getEmailAddrID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(emailAddrForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	emailAddr, updateErr := updateExistingEmailAddr(propUpdater.getEmailAddrID(), emailAddrForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateEmailAddrProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return emailAddr, nil
}

type EmailAddrResizeParams struct {
	EmailAddrIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams EmailAddrResizeParams) updateProps(emailAddr *EmailAddr) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	emailAddr.Properties.Geometry = updateParams.Geometry

	return nil
}

type EmailAddrLabelFormatParams struct {
	EmailAddrIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams EmailAddrLabelFormatParams) updateProps(emailAddr *EmailAddr) error {

	// TODO - Validate format is well-formed.

	emailAddr.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type EmailAddrVisibilityParams struct {
	EmailAddrIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams EmailAddrVisibilityParams) updateProps(emailAddr *EmailAddr) error {

	// TODO - Validate conditions

	emailAddr.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type EmailAddrPermissionParams struct {
	EmailAddrIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams EmailAddrPermissionParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.Permissions = updateParams.Permissions

	return nil
}

type EmailAddrValidationParams struct {
	EmailAddrIDHeader
	Validation EmailAddrValidationProperties `json:"validation"`
}

func (updateParams EmailAddrValidationParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.Validation = updateParams.Validation

	return nil
}

type EmailAddrClearValueSupportedParams struct {
	EmailAddrIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams EmailAddrClearValueSupportedParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	EmailAddrIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(emailAddr *EmailAddr) error {

	emailAddr.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
