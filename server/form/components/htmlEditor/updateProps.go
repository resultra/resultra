package htmlEditor

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
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

func updateHtmlEditorProps(propUpdater HtmlEditorPropUpdater) (*HtmlEditor, error) {

	// Retrieve the bar chart from the data store
	htmlEditorForUpdate, getErr := getHtmlEditor(propUpdater.getParentFormID(), propUpdater.getHtmlEditorID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to get existing html editor: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(htmlEditorForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to update existing html editor properties: %v", propUpdateErr)
	}

	htmlEditor, updateErr := updateExistingHtmlEditor(propUpdater.getHtmlEditorID(), htmlEditorForUpdate)
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
