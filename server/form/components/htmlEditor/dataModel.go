package htmlEditor

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveHtmlEditor(destDBHandle *sql.DB, newHtmlEditor HtmlEditor) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, htmlEditorEntityKind,
		newHtmlEditor.ParentFormID, newHtmlEditor.HtmlEditorID, newHtmlEditor.Properties); saveErr != nil {
		return fmt.Errorf("saveNewHtmlEditor: Unable to save html editor: error = %v", saveErr)
	}
	return nil

}

func saveNewHtmlEditor(trackerDBHandle *sql.DB, params NewHtmlEditorParams) (*HtmlEditor, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validHtmlEditorFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultEditorProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newHtmlEditor := HtmlEditor{ParentFormID: params.ParentFormID,
		HtmlEditorID: uniqueID.GenerateSnowflakeID(),
		Properties:   properties}

	if err := saveHtmlEditor(trackerDBHandle, newHtmlEditor); err != nil {
		return nil, fmt.Errorf("saveNewHtmlEditor: Unable to save html editor with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: %+v", newHtmlEditor)

	return &newHtmlEditor, nil

}

func getHtmlEditor(trackerDBHandle *sql.DB, parentFormID string, htmlEditorID string) (*HtmlEditor, error) {

	editorProps := newDefaultEditorProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, htmlEditorEntityKind, parentFormID, htmlEditorID, &editorProps); getErr != nil {
		return nil, fmt.Errorf("getHtmlEditor: Unable to retrieve html editor: %v", getErr)
	}

	htmlEditor := HtmlEditor{
		ParentFormID: parentFormID,
		HtmlEditorID: htmlEditorID,
		Properties:   editorProps}

	return &htmlEditor, nil
}

func getHtmlEditorsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]HtmlEditor, error) {

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
	if getErr := common.GetFormComponents(srcDBHandle, htmlEditorEntityKind, parentFormID, addEditor); getErr != nil {
		return nil, fmt.Errorf("GetHtmlEditors: Can't get html editors: %v")
	}

	return htmlEditors, nil

}

func GetHtmlEditors(trackerDBHandle *sql.DB, parentFormID string) ([]HtmlEditor, error) {
	return getHtmlEditorsFromSrc(trackerDBHandle, parentFormID)
}

func CloneHTMLEditors(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcHtmlEditors, err := getHtmlEditorsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneHTMLEditors: %v", err)
	}

	for _, srcHtmlEditor := range srcHtmlEditors {
		remappedHtmlEditorID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcHtmlEditor.HtmlEditorID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcHtmlEditor.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destProperties, err := srcHtmlEditor.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destHtmlEditor := HtmlEditor{
			ParentFormID: remappedFormID,
			HtmlEditorID: remappedHtmlEditorID,
			Properties:   *destProperties}
		if err := saveHtmlEditor(cloneParams.DestDBHandle, destHtmlEditor); err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
	}

	return nil
}

func updateExistingHtmlEditor(trackerDBHandle *sql.DB, htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditor, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, htmlEditorEntityKind, updatedHtmlEditor.ParentFormID,
		updatedHtmlEditor.HtmlEditorID, updatedHtmlEditor.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: error updating existing date editor: %v", updateErr)
	}

	return updatedHtmlEditor, nil

}
