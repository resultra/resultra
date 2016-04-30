package checkBox

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const checkBoxEntityKind string = "CheckBox"

func checkBoxChildParentEntityRel() datastoreWrapper.ChildParentEntityRel {
	return datastoreWrapper.ChildParentEntityRel{ParentEntityKind: dataModel.LayoutEntityKind, ChildEntityKind: checkBoxEntityKind}
}

type CheckBox struct {
	Field    *datastore.Key
	Geometry common.LayoutGeometry
}

type CheckBoxRef struct {
	CheckBoxID string                `json:"checkBoxID"`
	FieldRef   field.FieldRef        `json:"fieldRef"`
	Geometry   common.LayoutGeometry `json:"geometry"`
}

type NewCheckBoxParams struct {
	// ContainerID is initially assigned a temporary ID assigned by the client. It is passed back
	// to the client after the real datastore ID is assigned, allowing the client
	// to swizzle/replace the placeholder ID with the real one.
	ParentID string                `json:"parentID"`
	FieldID  string                `json:"fieldID"`
	Geometry common.LayoutGeometry `json:"geometry"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveNewCheckBox(appEngContext appengine.Context, params NewCheckBoxParams) (*CheckBoxRef, error) {

	if !common.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Can't text box with field ID = '%v': datastore error=%v",
			params.FieldID, fieldErr)
	}

	if !validCheckBoxFieldType(fieldRef.FieldInfo.Type) {
		return nil, fmt.Errorf("saveNewCheckBox: Invalid field type: expecting bool field, got %v", fieldRef.FieldInfo.Type)
	}

	newCheckBox := CheckBox{Field: fieldKey, Geometry: params.Geometry}

	checkBoxID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentID, checkBoxChildParentEntityRel(), &newCheckBox)
	if insertErr != nil {
		return nil, insertErr
	}

	checkBoxRef := CheckBoxRef{
		CheckBoxID: checkBoxID,
		FieldRef:   *fieldRef,
		Geometry:   params.Geometry}

	log.Printf("INFO: API: New Checkbox: Created new check box container: id=%v params=%+v", checkBoxID, params)

	return &checkBoxRef, nil

}

func getCheckBox(appEngContext appengine.Context, checkBoxID string) (*CheckBox, error) {

	var checkBox CheckBox
	if getErr := datastoreWrapper.GetChildEntity(appEngContext, checkBoxID, checkBoxChildParentEntityRel(), &checkBox); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to get check box from datastore: error = %v", getErr)
	}
	return &checkBox, nil
}

func GetCheckBoxes(appEngContext appengine.Context, parentFormID string) ([]CheckBoxRef, error) {

	var checkBoxes []CheckBox
	checkBoxIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, parentFormID, checkBoxChildParentEntityRel(), &checkBoxes)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve check boxes: form id=%v", parentFormID)
	}

	checkBoxRefs := make([]CheckBoxRef, len(checkBoxes))
	for checkBoxIter, currCheckBox := range checkBoxes {

		checkBoxID := checkBoxIDs[checkBoxIter]

		fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, currCheckBox.Field)
		if fieldErr != nil {
			return nil, fmt.Errorf("GetCheckBoxes: Error retrieving field for check box: error = %v", fieldErr)
		}

		checkBoxRefs[checkBoxIter] = CheckBoxRef{
			CheckBoxID: checkBoxID,
			FieldRef:   *fieldRef,
			Geometry:   currCheckBox.Geometry}

	} // for each check box
	return checkBoxRefs, nil

}

func updateExistingCheckBox(appEngContext appengine.Context, checkBoxID string, updatedCheckBox *CheckBox) (*CheckBoxRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingChildEntity(appEngContext, checkBoxID,
		checkBoxChildParentEntityRel(), updatedCheckBox); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: Error updating check box: error = %v", updateErr)
	}

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, updatedCheckBox.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: Error retrieving field for check box: error = %v", fieldErr)
	}

	checkBoxRef := CheckBoxRef{
		CheckBoxID: checkBoxID,
		FieldRef:   *fieldRef,
		Geometry:   updatedCheckBox.Geometry}

	return &checkBoxRef, nil

}
