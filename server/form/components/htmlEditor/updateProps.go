package htmlEditor

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
)

type HtmlEditorIDInterface interface {
	getHtmlEditorID() string
}

type HtmlEditorIDHeader struct {
	HtmlEditorID string `json:"htmlEditorID"`
}

func (idHeader HtmlEditorIDHeader) getHtmlEditorID() string {
	return idHeader.HtmlEditorID
}

type HtmlEditorPropUpdater interface {
	HtmlEditorIDInterface
	updateProps(htmlEditor *HtmlEditor) error
}

func updateHtmlEditorProps(appEngContext appengine.Context, propUpdater HtmlEditorPropUpdater) (*HtmlEditorRef, error) {

	// Retrieve the bar chart from the data store
	htmlEditorForUpdate, getErr := getHtmlEditor(appEngContext, propUpdater.getHtmlEditorID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to get existing html editor: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(htmlEditorForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to update existing html editor properties: %v", propUpdateErr)
	}

	htmlEditorRef, updateErr := updateExistingHtmlEditor(appEngContext, propUpdater.getHtmlEditorID(), htmlEditorForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateHtmlEditorProps: Unable to update existing html editor properties: datastore update error =  %v", updateErr)
	}

	return htmlEditorRef, nil
}

type HtmlEditorResizeParams struct {
	HtmlEditorIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams HtmlEditorResizeParams) updateProps(htmlEditor *HtmlEditor) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set html editor dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	htmlEditor.Geometry = updateParams.Geometry

	return nil
}

type HtmlEditorRepositionParams struct {
	HtmlEditorIDHeader
	Position common.LayoutPosition `json:"position"`
}

func (updateParams HtmlEditorRepositionParams) updateProps(htmlEditor *HtmlEditor) error {

	if err := htmlEditor.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for html editor: Invalid geometry: %v", err)
	}

	return nil
}
