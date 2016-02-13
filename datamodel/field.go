package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
)

const fieldEntityKind string = "Field"

const fieldTypeText = "text"
const fieldTypeNumber = "number"
const fieldTypeDate = "date"

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FieldRef struct {
	FieldID   string `json:"fieldID"`
	FieldInfo Field  `json:"fieldInfo"`
}

type FieldsByType struct {
	TextFields   []FieldRef `json:"textFields"`
	DateFields   []FieldRef `json:"dateFields"`
	NumberFields []FieldRef `json:"numberFields"`
}

func validFieldType(fieldType string) bool {
	switch fieldType {
	case fieldTypeText:
		return true
	case fieldTypeNumber:
		return true
	case fieldTypeDate:
		return true
	default:
		return false
	}
}

func NewField(appEngContext appengine.Context, newField Field) (string, error) {

	sanitizedName, sanitizeErr := sanitizeName(newField.Name)
	if sanitizeErr != nil {
		return "", fmt.Errorf("Can't create new field: invalid name: '%v'", sanitizeErr)
	}
	newField.Name = sanitizedName

	if !validFieldType(newField.Type) {
		return "", fmt.Errorf("Can't create new field: invalid field type: '%v'", newField.Type)
	}

	fieldID, insertErr := insertNewEntity(appEngContext, fieldEntityKind, nil, &newField)
	if insertErr != nil {
		return "", fmt.Errorf("Can't create new field: error inserting into datastore: %v", insertErr)
	}

	// TODO - verify IntID != 0
	log.Printf("NewField: Created new field: id= %v, field='%+v'", fieldID, newField)

	return fieldID, nil

}

type GetFieldParams struct {
	// TODO - There will be more parameters once a field is
	// tied to a database table (i.e. TableID)
	FieldID string `json:"fieldID"`
}

func GetField(appEngContext appengine.Context, fieldParams GetFieldParams) (*FieldRef, error) {

	fieldGetDest := Field{}
	if getErr := getRootEntityByID(appEngContext, fieldEntityKind, fieldParams.FieldID, &fieldGetDest); getErr != nil {
		return nil, fmt.Errorf("Unabled to get field: params = %+v: datastore err=%v", fieldParams, getErr)
	} else {
		return &FieldRef{fieldParams.FieldID, fieldGetDest}, nil
	}
}

func GetAllFieldRefs(appEngContext appengine.Context) ([]FieldRef, error) {

	var allFields []Field
	fieldQuery := datastore.NewQuery(fieldEntityKind)
	keys, err := fieldQuery.GetAll(appEngContext, &allFields)

	if err != nil {
		return nil, fmt.Errorf("GetFieldsByType: Unable to retrieve fields from datastore: datastore error =%v", err)
	}

	fieldRefs := make([]FieldRef, len(allFields))
	for i, currField := range allFields {
		fieldKey := keys[i]
		fieldID, encodeErr := encodeUniqueEntityIDToStr(fieldKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("Failed to encode unique ID for field: key=%+v, encode err=%v", fieldKey, encodeErr)
		}
		fieldRefs[i] = FieldRef{fieldID, Field{currField.Name, currField.Type}}
	}
	return fieldRefs, nil
}

func GetFieldsByType(appEngContext appengine.Context) (*FieldsByType, error) {

	fieldRefs, getErr := GetAllFieldRefs(appEngContext)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldsByType: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	fieldsByType := FieldsByType{}
	for fieldRefIndex := range fieldRefs {
		fieldRef := fieldRefs[fieldRefIndex]
		switch fieldRef.FieldInfo.Type {
		case fieldTypeText:
			fieldsByType.TextFields = append(fieldsByType.TextFields, fieldRef)
		case fieldTypeDate:
			fieldsByType.DateFields = append(fieldsByType.DateFields, fieldRef)
		case fieldTypeNumber:
			fieldsByType.NumberFields = append(fieldsByType.NumberFields, fieldRef)
		default:
			return nil, fmt.Errorf(
				"GetFieldsByType: Unable to retrieve fields from datastore: Invalid field type %v",
				fieldRef.FieldInfo.Type)
		}
	}
	return &fieldsByType, nil

}
