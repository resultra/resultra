// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package htmlEditor

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/form/components/common"
)

type HtmlEditorIDInterface interface {
	getHtmlEditorID() string
	getParentFormID() string
}

type HtmlEditorIDHeader struct {
	HtmlEditorID string `json:"htmlEditorID"`
	ParentFormID string `json:"parentFormID"`
}

func (idHeader HtmlEditorIDHeader) getHtmlEditorID() string {
	return idHeader.HtmlEditorID
}

func (idHeader HtmlEditorIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type HtmlEditorPropUpdater interface {
	HtmlEditorIDInterface
	updateProps(htmlEditor *HtmlEditor) error
}

func updateHtmlEditorProps(trackerDBHandle *sql.DB, propUpdater HtmlEditorPropUpdater) (*HtmlEditor, error) {

	// Retrieve the bar chart from the data store
	htmlEditorForUpdate, getErr := getHtmlEditor(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getHtmlEditorID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to get existing html editor: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(htmlEditorForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to update existing html editor properties: %v", propUpdateErr)
	}

	htmlEditor, updateErr := updateExistingHtmlEditor(trackerDBHandle, propUpdater.getHtmlEditorID(), htmlEditorForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to update existing html editor properties: datastore update error =  %v", updateErr)
	}

	return htmlEditor, nil
}

type HtmlEditorResizeParams struct {
	HtmlEditorIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams HtmlEditorResizeParams) updateProps(htmlEditor *HtmlEditor) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set html editor dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	htmlEditor.Properties.Geometry = updateParams.Geometry

	return nil
}

type EditorLabelFormatParams struct {
	HtmlEditorIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams EditorLabelFormatParams) updateProps(htmlEditor *HtmlEditor) error {

	// TODO - Validate format is well-formed.

	htmlEditor.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type EditorVisibilityParams struct {
	HtmlEditorIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams EditorVisibilityParams) updateProps(editor *HtmlEditor) error {

	// TODO - Validate conditions

	editor.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type EditorPermissionParams struct {
	HtmlEditorIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams EditorPermissionParams) updateProps(editor *HtmlEditor) error {

	editor.Properties.Permissions = updateParams.Permissions

	return nil
}

type EditorValidationParams struct {
	HtmlEditorIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams EditorValidationParams) updateProps(editor *HtmlEditor) error {

	editor.Properties.Validation = updateParams.Validation

	return nil
}

type HelpPopupMsgParams struct {
	HtmlEditorIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(editor *HtmlEditor) error {

	editor.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
