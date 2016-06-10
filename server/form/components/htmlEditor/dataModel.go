package htmlEditor

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	geometry "resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
)

const htmlEditorEntityKind string = "html_editor"

type HtmlEditorProperties struct {
	FieldID  string                  `json:"fieldID"`
	Geometry geometry.LayoutGeometry `json:"geometry"`
}

type HtmlEditor struct {
	ParentFormID string               `json:"parentID"`
	HtmlEditorID string               `json:"htmlEditorID"`
	Properties   HtmlEditorProperties `json:"properties"`
}

type NewHtmlEditorParams struct {
	ParentFormID       string                  `json:"parentFormID"`
	FieldParentTableID string                  `json:"fieldParentTableID"`
	FieldID            string                  `json:"fieldID"`
	Geometry           geometry.LayoutGeometry `json:"geometry"`
}

func validHtmlEditorFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLongText {
		return true
	} else {
		return false
	}
}

func saveNewHtmlEditor(appEngContext appengine.Context, params NewHtmlEditorParams) (*HtmlEditor, error) {

	if !geometry.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validHtmlEditorFieldType(field.Type) {
		return nil, fmt.Errorf("saveNewHtmlEditor: Invalid field type: expecting time field, got %v", field.Type)
	}

	properties := HtmlEditorProperties{
		Geometry: params.Geometry,
		FieldID:  params.FieldID}

	newHtmlEditor := HtmlEditor{ParentFormID: params.ParentFormID,
		HtmlEditorID: gocql.TimeUUID().String(),
		Properties:   properties}

	if saveErr := common.SaveNewFormComponent(htmlEditorEntityKind,
		newHtmlEditor.ParentFormID, newHtmlEditor.HtmlEditorID, newHtmlEditor.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewHtmlEditor: Unable to save html editor with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: %+v", newHtmlEditor)

	return &newHtmlEditor, nil

}

func getHtmlEditor(appEngContext appengine.Context, parentFormID string, htmlEditorID string) (*HtmlEditor, error) {

	editorProps := HtmlEditorProperties{}
	if getErr := common.GetFormComponent(htmlEditorEntityKind, parentFormID, htmlEditorID, &editorProps); getErr != nil {
		return nil, fmt.Errorf("getHtmlEditor: Unable to retrieve html editor: %v", getErr)
	}

	htmlEditor := HtmlEditor{
		ParentFormID: parentFormID,
		HtmlEditorID: htmlEditorID,
		Properties:   editorProps}

	return &htmlEditor, nil
}

func GetHtmlEditors(appEngContext appengine.Context, parentFormID string) ([]HtmlEditor, error) {

	htmlEditors := []HtmlEditor{}

	addEditor := func(editorID string, encodedProps string) error {
		var editorProps HtmlEditorProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &editorProps); decodeErr != nil {
			return fmt.Errorf("GetHtmlEditors: can't decode properties: %v", encodedProps)
		}

		currEditor := HtmlEditor{
			ParentFormID: parentFormID,
			HtmlEditorID: editorID,
			Properties:   editorProps}

		htmlEditors = append(htmlEditors, currEditor)

		return nil
	}
	if getErr := common.GetFormComponents(htmlEditorEntityKind, parentFormID, addEditor); getErr != nil {
		return nil, fmt.Errorf("GetHtmlEditors: Can't get html editors: %v")
	}

	return htmlEditors, nil

}

func updateExistingHtmlEditor(appEngContext appengine.Context, htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditor, error) {

	if updateErr := common.UpdateFormComponent(htmlEditorEntityKind, updatedHtmlEditor.ParentFormID,
		updatedHtmlEditor.HtmlEditorID, updatedHtmlEditor.Properties); updateErr != nil {
	}

	return updatedHtmlEditor, nil

}
