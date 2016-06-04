package htmlEditor

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const htmlEditorEntityKind string = "HtmlEditor"

type HtmlEditor struct {
	ParentFormID string                `json:"parentID"`
	HtmlEditorID string                `json:"htmlEditorID"`
	FieldID      string                `json:"fieldID"`
	Geometry     common.LayoutGeometry `json:"geometry"`
}

const htmlEditorParentFormIDFieldName string = "ParentFormID"
const htmlEditorIDFieldName string = "HtmlEditorID"

type NewHtmlEditorParams struct {
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func validHtmlEditorFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLongText {
		return true
	} else {
		return false
	}
}

func saveNewHtmlEditor(appEngContext appengine.Context, params NewHtmlEditorParams) (*HtmlEditor, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	field, fieldErr := field.GetField(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create image with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validHtmlEditorFieldType(field.Type) {
		return nil, fmt.Errorf("saveNewHtmlEditor: Invalid field type: expecting time field, got %v", field.Type)
	}

	newHtmlEditor := HtmlEditor{ParentFormID: params.ParentID,
		FieldID:      params.FieldID,
		HtmlEditorID: uniqueID.GenerateUniqueID(),
		Geometry:     params.Geometry}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, htmlEditorEntityKind, &newHtmlEditor)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new html editor component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: %+v", newHtmlEditor)

	return &newHtmlEditor, nil

}

func getHtmlEditor(appEngContext appengine.Context, htmlEditorID string) (*HtmlEditor, error) {

	var htmlEditor HtmlEditor

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, htmlEditorEntityKind,
		htmlEditorIDFieldName, htmlEditorID, &htmlEditor); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to image container from datastore: error = %v", getErr)
	}
	return &htmlEditor, nil
}

func GetHtmlEditors(appEngContext appengine.Context, parentFormID string) ([]HtmlEditor, error) {

	var htmlEditors []HtmlEditor

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentFormID,
		htmlEditorEntityKind, htmlEditorParentFormIDFieldName, &htmlEditors)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve html editors: form id=%v", parentFormID)
	}

	return htmlEditors, nil

}

func updateExistingHtmlEditor(appEngContext appengine.Context, htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditor, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		htmlEditorID, htmlEditorEntityKind, htmlEditorIDFieldName, updatedHtmlEditor); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating html editor: error = %v", updateErr)
	}

	return updatedHtmlEditor, nil

}
