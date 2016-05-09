package field

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

type GetFieldListParams struct {
	ParentTableID string `json:"parentTableID"`
}

func GetAllFieldRefs(appEngContext appengine.Context, params GetFieldListParams) ([]FieldRef, error) {

	var allFields []Field
	fieldIDs, err := datastoreWrapper.GetAllChildEntities(appEngContext, params.ParentTableID, FieldEntityKind, &allFields)
	if err != nil {
		return nil, fmt.Errorf("GetFieldsByType: Unable to retrieve fields from datastore: datastore error =%v", err)
	}

	fieldRefs := make([]FieldRef, len(allFields))
	for i, currField := range allFields {
		fieldID := fieldIDs[i]
		fieldRefs[i] = FieldRef{fieldID, currField}
	}
	return fieldRefs, nil
}

type FieldsByType struct {
	TextFields   []FieldRef `json:"textFields"`
	TimeFields   []FieldRef `json:"timeFields"`
	NumberFields []FieldRef `json:"numberFields"`
	BoolFields   []FieldRef `json:"boolFields"`
}

func GetFieldsByType(appEngContext appengine.Context, params GetFieldListParams) (*FieldsByType, error) {

	fieldRefs, getErr := GetAllFieldRefs(appEngContext, params)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldsByType: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	fieldsByType := FieldsByType{}
	for fieldRefIndex := range fieldRefs {
		fieldRef := fieldRefs[fieldRefIndex]
		switch fieldRef.FieldInfo.Type {
		case FieldTypeText:
			fieldsByType.TextFields = append(fieldsByType.TextFields, fieldRef)
		case FieldTypeTime:
			fieldsByType.TimeFields = append(fieldsByType.TimeFields, fieldRef)
		case FieldTypeNumber:
			fieldsByType.NumberFields = append(fieldsByType.NumberFields, fieldRef)
		case FieldTypeBool:
			fieldsByType.BoolFields = append(fieldsByType.BoolFields, fieldRef)
		default:
			return nil, fmt.Errorf(
				"GetFieldsByType: Unable to retrieve fields from datastore: Invalid field type %v",
				fieldRef.FieldInfo.Type)
		}
	}
	return &fieldsByType, nil

}

type StringFieldRefMap map[string]FieldRef

type FieldRefIDIndex struct {
	FieldRefsByID      StringFieldRefMap
	FieldRefsByRefName StringFieldRefMap
}

func (fieldRefIDIndex FieldRefIDIndex) getFieldRefByID(fieldID string) (*FieldRef, error) {
	fieldRef, fieldRefFound := fieldRefIDIndex.FieldRefsByID[fieldID]
	if fieldRefFound != true {
		return nil, fmt.Errorf("getFieldRefByID: Unable to retrieve field for field with ID = %v ", fieldID)
	}
	return &fieldRef, nil

}

func GetFieldRefIDIndex(appEngContext appengine.Context, params GetFieldListParams) (*FieldRefIDIndex, error) {

	fieldRefs, getErr := GetAllFieldRefs(appEngContext, params)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	log.Printf("GetFieldRefIDIndex: Indexing %v fields", len(fieldRefs))

	fieldRefsByRefName := StringFieldRefMap{}
	fieldRefsByID := StringFieldRefMap{}
	for _, fieldRef := range fieldRefs {

		if _, keyExists := fieldRefsByRefName[fieldRef.FieldInfo.RefName]; keyExists == true {
			return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: "+
				" found duplicate reference name for field = %+v", fieldRef)
		}

		if _, keyExists := fieldRefsByID[fieldRef.FieldID]; keyExists == true {
			return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: "+
				" found duplicate key for field = %+v", fieldRef)
		}

		log.Printf("GetFieldRefIDIndex: Indexed field: %+v", fieldRef)

		fieldRefsByRefName[fieldRef.FieldInfo.RefName] = fieldRef
		fieldRefsByID[fieldRef.FieldID] = fieldRef

	}

	return &FieldRefIDIndex{fieldRefsByID, fieldRefsByRefName}, nil

}
