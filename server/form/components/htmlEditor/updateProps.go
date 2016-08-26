package htmlEditor

import (
	"fmt"
	"resultra/datasheet/server/common"
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
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams HtmlEditorResizeParams) updateProps(htmlEditor *HtmlEditor) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set html editor dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	htmlEditor.Properties.Geometry = updateParams.Geometry

	return nil
}
