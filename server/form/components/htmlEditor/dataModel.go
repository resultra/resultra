package htmlEditor

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const htmlEditorEntityKind string = "html_editor"

type HtmlEditorProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type HtmlEditor struct {
	ParentFormID string               `json:"parentID"`
	HtmlEditorID string               `json:"htmlEditorID"`
	Properties   HtmlEditorProperties `json:"properties"`
}

type NewHtmlEditorParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validHtmlEditorFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLongText {
		return true
	} else {
		return false
	}
}

func saveNewHtmlEditor(params NewHtmlEditorParams) (*HtmlEditor, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validHtmlEditorFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", compLinkErr)
	}

	properties := HtmlEditorProperties{
		Geometry:      params.Geometry,
		ComponentLink: params.ComponentLink}

	newHtmlEditor := HtmlEditor{ParentFormID: params.ParentFormID,
		HtmlEditorID: uniqueID.GenerateSnowflakeID(),
		Properties:   properties}

	if saveErr := common.SaveNewFormComponent(htmlEditorEntityKind,
		newHtmlEditor.ParentFormID, newHtmlEditor.HtmlEditorID, newHtmlEditor.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewHtmlEditor: Unable to save html editor with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: %+v", newHtmlEditor)

	return &newHtmlEditor, nil

}

func getHtmlEditor(parentFormID string, htmlEditorID string) (*HtmlEditor, error) {

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

func GetHtmlEditors(parentFormID string) ([]HtmlEditor, error) {

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

func updateExistingHtmlEditor(htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditor, error) {

	if updateErr := common.UpdateFormComponent(htmlEditorEntityKind, updatedHtmlEditor.ParentFormID,
		updatedHtmlEditor.HtmlEditorID, updatedHtmlEditor.Properties); updateErr != nil {
	}

	return updatedHtmlEditor, nil

}
