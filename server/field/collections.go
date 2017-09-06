package field

import (
	"fmt"
	"sort"
	"strings"
)

type FieldsByType struct {
	TextFields      []Field `json:"textFields"`
	LongTextFields  []Field `json:"longTextFields"`
	TimeFields      []Field `json:"timeFields"`
	NumberFields    []Field `json:"numberFields"`
	BoolFields      []Field `json:"boolFields"`
	FileFields      []Field `json:"fileFields"`
	UserFields      []Field `json:"userFields"`
	CommentFields   []Field `json:"commentFields"`
	LabelFields     []Field `json:"labelFields"`
	EmailAddrFields []Field `json:"emailAddrFields"`
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
		case FieldTypeLabel:
			fieldsByType.LabelFields = append(fieldsByType.LabelFields, currField)
		case FieldTypeEmail:
			fieldsByType.EmailAddrFields = append(fieldsByType.EmailAddrFields, currField)
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

type ByFieldName []Field

func (s ByFieldName) Len() int {
	return len(s)
}
func (s ByFieldName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in reverse chronological order; i.e. the most recent dates come first.
func (s ByFieldName) Less(i, j int) bool {

	if strings.Compare(s[i].Name, s[j].Name) < 0 {
		return true
	} else {
		return false
	}
}

type GetSortedFieldListParams struct {
	ParentDatabaseID string   `json:"parentDatabaseID"`
	FieldTypes       []string `json:"fieldTypes"`
}

func getSortedFieldsByType(params GetSortedFieldListParams) ([]Field, error) {

	fieldsByType, getErr := GetFieldsByType(GetFieldListParams{ParentDatabaseID: params.ParentDatabaseID})
	if getErr != nil {
		return nil, fmt.Errorf("GetSortedFields: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	matchedFields := []Field{}

	for _, currFieldType := range params.FieldTypes {
		switch currFieldType {
		case FieldTypeText:
			matchedFields = append(matchedFields, fieldsByType.TextFields...)
		case FieldTypeLongText:
			matchedFields = append(matchedFields, fieldsByType.LongTextFields...)
		case FieldTypeTime:
			matchedFields = append(matchedFields, fieldsByType.TimeFields...)
		case FieldTypeNumber:
			matchedFields = append(matchedFields, fieldsByType.NumberFields...)
		case FieldTypeBool:
			matchedFields = append(matchedFields, fieldsByType.BoolFields...)
		case FieldTypeUser:
			matchedFields = append(matchedFields, fieldsByType.UserFields...)
		case FieldTypeFile:
			matchedFields = append(matchedFields, fieldsByType.FileFields...)
		case FieldTypeComment:
			matchedFields = append(matchedFields, fieldsByType.CommentFields...)
		case FieldTypeLabel:
			matchedFields = append(matchedFields, fieldsByType.LabelFields...)
		case FieldTypeEmail:
			matchedFields = append(matchedFields, fieldsByType.EmailAddrFields...)
		default:
			return nil, fmt.Errorf("GetSortedFields: unsupported field type: %v", currFieldType)
		}

	}

	sort.Sort(ByFieldName(matchedFields))

	return matchedFields, nil

}

func getAllSortedFields(params GetFieldListParams) ([]Field, error) {
	fields, getErr := GetAllFields(params)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldRefIDIndex: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}
	sort.Sort(ByFieldName(fields))

	return fields, nil

}
