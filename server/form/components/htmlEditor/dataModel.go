package htmlEditor

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const htmlEditorEntityKind string = "HtmlEditor"

type HtmlEditor struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type HtmlEditorRef struct {
	HtmlEditorID string                `json:"htmlEditorID"`
	FieldRef     field.FieldRef        `json:"fieldRef"`
	Geometry     common.LayoutGeometry `json:"geometry"`
}

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

func saveNewHtmlEditor(appEngContext appengine.Context, params NewHtmlEditorParams) (*HtmlEditorRef, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("saveNewHtmlEditor: Can't get field with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validHtmlEditorFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("saveNewHtmlEditor: Invalid field type: expecting time field, got %v", fieldRef.FieldInfo.Type)
	}

	newHtmlEditor := HtmlEditor{Field: fieldKey, Geometry: params.Geometry}

	htmlEditorID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentID, htmlEditorEntityKind, &newHtmlEditor)
	if insertErr != nil {
		return nil, insertErr
	}

	htmlEditorRef := HtmlEditorRef{
		HtmlEditorID: htmlEditorID,
		FieldRef:     *fieldRef,
		Geometry:     params.Geometry}

	log.Printf("INFO: API: New HtmlEditor: Created new html editor container: id=%v params=%+v", htmlEditorID, params)

	return &htmlEditorRef, nil

}

func getHtmlEditor(appEngContext appengine.Context, htmlEditorID string) (*HtmlEditor, error) {

	var htmlEditor HtmlEditor
	if getErr := datastoreWrapper.GetEntity(appEngContext, htmlEditorID, &htmlEditor); getErr != nil {
		return nil, fmt.Errorf("getHtmlEditor: Unable to get html editor from datastore: error = %v", getErr)
	}
	return &htmlEditor, nil
}

func GetHtmlEditors(appEngContext appengine.Context, parentFormID string) ([]HtmlEditorRef, error) {

	var htmlEditors []HtmlEditor
	htmlEditorIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentFormID, htmlEditorEntityKind, &htmlEditors)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve html editors: form id=%v", parentFormID)
	}

	htmlEditorRefs := make([]HtmlEditorRef, len(htmlEditors))
	for htmlEditorIter, currHtmlEditor := range htmlEditors {

		htmlEditorID := htmlEditorIDs[htmlEditorIter]

		fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currHtmlEditor.Field)
		if fieldErr != nil {
			return nil, fmt.Errorf("GetHtmlEditors: Error retrieving field for html editor: error = %v", fieldErr)
		}

		htmlEditorRefs[htmlEditorIter] = HtmlEditorRef{
			HtmlEditorID: htmlEditorID,
			FieldRef:     *fieldRef,
			Geometry:     currHtmlEditor.Geometry}

	} // for each check box
	return htmlEditorRefs, nil

}

func updateExistingHtmlEditor(appEngContext appengine.Context, htmlEditorID string, updatedHtmlEditor *HtmlEditor) (*HtmlEditorRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntity(appEngContext, htmlEditorID, updatedHtmlEditor); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating html editor: error = %v", updateErr)
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedHtmlEditor.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error retrieving field for html editor: error = %v", fieldErr)
	}

	htmlEditorRef := HtmlEditorRef{
		HtmlEditorID: htmlEditorID,
		FieldRef:     *fieldRef,
		Geometry:     updatedHtmlEditor.Geometry}

	return &htmlEditorRef, nil

}
