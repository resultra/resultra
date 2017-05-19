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

type HtmlEditor struct {
	ParentFormID string               `json:"parentFormID"`
	HtmlEditorID string               `json:"htmlEditorID"`
	Properties   HtmlEditorProperties `json:"properties"`
}

type NewHtmlEditorParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validHtmlEditorFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLongText {
		return true
	} else {
		return false
	}
}

func saveHtmlEditor(newHtmlEditor HtmlEditor) error {

	if saveErr := common.SaveNewFormComponent(htmlEditorEntityKind,
		newHtmlEditor.ParentFormID, newHtmlEditor.HtmlEditorID, newHtmlEditor.Properties); saveErr != nil {
		return fmt.Errorf("saveNewHtmlEditor: Unable to save html editor: error = %v", saveErr)
	}
	return nil

}

func saveNewHtmlEditor(params NewHtmlEditorParams) (*HtmlEditor, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validHtmlEditorFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultEditorProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newHtmlEditor := HtmlEditor{ParentFormID: params.ParentFormID,
		HtmlEditorID: uniqueID.GenerateSnowflakeID(),
		Properties:   properties}

	if err := saveHtmlEditor(newHtmlEditor); err != nil {
		return nil, fmt.Errorf("saveNewHtmlEditor: Unable to save html editor with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: %+v", newHtmlEditor)

	return &newHtmlEditor, nil

}

func getHtmlEditor(parentFormID string, htmlEditorID string) (*HtmlEditor, error) {

	editorProps := newDefaultEditorProperties()
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
		editorProps := newDefaultEditorProperties()
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

func CloneHTMLEditors(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcHtmlEditors, err := GetHtmlEditors(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneHTMLEditors: %v", err)
	}

	for _, srcHtmlEditor := range srcHtmlEditors {
		remappedHtmlEditorID := remappedIDs.AllocNewOrGetExistingRemappedID(srcHtmlEditor.HtmlEditorID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcHtmlEditor.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destProperties, err := srcHtmlEditor.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destHtmlEditor := HtmlEditor{
			ParentFormID: remappedFormID,
			HtmlEditorID: remappedHtmlEditorID,
			Properties:   *destProperties}
		if err := saveHtmlEditor(destHtmlEditor); err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
	}

	return nil
}

func updateExistingHtmlEditor(htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditor, error) {

	if updateErr := common.UpdateFormComponent(htmlEditorEntityKind, updatedHtmlEditor.ParentFormID,
		updatedHtmlEditor.HtmlEditorID, updatedHtmlEditor.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: error updating existing date editor: %v", updateErr)
	}

	return updatedHtmlEditor, nil

}
