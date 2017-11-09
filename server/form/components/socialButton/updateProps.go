package socialButton

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type SocialButtonIDInterface interface {
	getSocialButtonID() string
	getParentFormID() string
}

type SocialButtonIDHeader struct {
	ParentFormID   string `json:"parentFormID"`
	SocialButtonID string `json:"socialButtonID"`
}

func (idHeader SocialButtonIDHeader) getSocialButtonID() string {
	return idHeader.SocialButtonID
}

func (idHeader SocialButtonIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type SocialButtonPropUpdater interface {
	SocialButtonIDInterface
	updateProps(socialButton *SocialButton) error
}

func updateSocialButtonProps(trackerDBHandle *sql.DB, propUpdater SocialButtonPropUpdater) (*SocialButton, error) {

	// Retrieve the bar chart from the data store
	socialButtonForUpdate, getErr := getSocialButton(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getSocialButtonID())
	if getErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to get existing socialButton: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(socialButtonForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to update existing socialButton properties: %v", propUpdateErr)
	}

	updatedSocialButton, updateErr := updateExistingSocialButton(trackerDBHandle, socialButtonForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateSocialButtonProps: Unable to update existing socialButton properties: datastore update error =  %v", updateErr)
	}

	return updatedSocialButton, nil
}

type SocialButtonResizeParams struct {
	SocialButtonIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams SocialButtonResizeParams) updateProps(socialButton *SocialButton) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	socialButton.Properties.Geometry = updateParams.Geometry

	return nil
}

type SocialButtonTooltipParams struct {
	SocialButtonIDHeader
	Tooltips []string `json:"tooltips"`
}

func (updateParams SocialButtonTooltipParams) updateProps(socialButton *SocialButton) error {

	socialButton.Properties.Tooltips = updateParams.Tooltips

	return nil
}

type SocialButtonIconParams struct {
	SocialButtonIDHeader
	Icon string `json:"icon"`
}

func (updateParams SocialButtonIconParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate icon is a valid name

	socialButton.Properties.Icon = updateParams.Icon

	return nil
}

type SocialButtonLabelFormatParams struct {
	SocialButtonIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams SocialButtonLabelFormatParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate format is well-formed.

	socialButton.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type SocialButtonVisibilityParams struct {
	SocialButtonIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams SocialButtonVisibilityParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate conditions

	socialButton.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type SocialButtonPermissionParams struct {
	SocialButtonIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams SocialButtonPermissionParams) updateProps(socialButton *SocialButton) error {

	// TODO - Validate conditions

	socialButton.Properties.Permissions = updateParams.Permissions

	return nil
}

type HelpPopupMsgParams struct {
	SocialButtonIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(socialButton *SocialButton) error {

	socialButton.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
