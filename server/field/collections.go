package field

import (
	"fmt"
	"log"
)

type FieldsByType struct {
	TextFields     []Field `json:"textFields"`
	LongTextFields []Field `json:"longTextFields"`
	TimeFields     []Field `json:"timeFields"`
	NumberFields   []Field `json:"numberFields"`
	BoolFields     []Field `json:"boolFields"`
	FileFields     []Field `json:"fileFields"`
	UserFields     []Field `json:"userFields"`
	CommentFields  []Field `json:"commentFields"`
}

func GetFieldsByType(params GetFieldListParams) (*FieldsByType, error) {

	fields, getErr := GetAllFields(params)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldsByType: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	fieldsByType := FieldsByType{}
	for fieldIndex := range fields {
		currField := fields[fieldIndex]
		switch currField.Type {
		case FieldTypeText:
			fieldsByType.TextFields = append(fieldsByType.TextFields, currField)
		case FieldTypeLongText:
			fieldsByType.LongTextFields = append(fieldsByType.LongTextFields, currField)
		case FieldTypeTime:
			fieldsByType.TimeFields = append(fieldsByType.TimeFields, currField)
		case FieldTypeNumber:
			fieldsByType.NumberFields = append(fieldsByType.NumberFields, currField)
		case FieldTypeBool:
			fieldsByType.BoolFields = append(fieldsByType.BoolFields, currField)
		case FieldTypeUser:
			fieldsByType.UserFields = append(fieldsByType.UserFields, currField)
		case FieldTypeFile:
			fieldsByType.FileFields = append(fieldsByType.FileFields, currField)
		case FieldTypeComment:
			fieldsByType.CommentFields = append(fieldsByType.CommentFields, currField)
		default:
			return nil, fmt.Errorf(
				"GetFieldsByType: Unable to retrieve fields from datastore: Invalid field type %v",
				currField.Type)
		}
	}
	return &fieldsByType, nil

}

type StringFieldMap map[string]Field

type FieldIDIndex struct {
	FieldsByID      StringFieldMap
	FieldsByRefName StringFieldMap
}

func (fieldIDIndex FieldIDIndex) GetFieldRefByID(fieldID string) (*Field, error) {
	field, fieldFound := fieldIDIndex.FieldsByID[fieldID]
	if fieldFound != true {
		return nil, fmt.Errorf("getFieldRefByID: Unable to retrieve field for field with ID = %v ", fieldID)
	}
	return &field, nil

}

func GetFieldRefIDIndex(params GetFieldListParams) (*FieldIDIndex, error) {

	fields, getErr := GetAllFields(params)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	log.Printf("GetFieldRefIDIndex: Indexing %v fields", len(fields))

	fieldsByRefName := StringFieldMap{}
	fieldsByID := StringFieldMap{}
	for _, field := range fields {

		if _, keyExists := fieldsByRefName[field.RefName]; keyExists == true {
			return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: "+
				" found duplicate reference name for field = %+v", field)
		}

		if _, keyExists := fieldsByID[field.FieldID]; keyExists == true {
			return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: "+
				" found duplicate key for field = %+v", field)
		}

		fieldsByRefName[field.RefName] = field
		fieldsByID[field.FieldID] = field

	}

	return &FieldIDIndex{fieldsByID, fieldsByRefName}, nil

}
