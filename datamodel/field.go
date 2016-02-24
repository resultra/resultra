package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const fieldEntityKind string = "Field"

// A "reference name" for a field can only contain
// TODO - Can't start with "true or false" - add this when supporting boolean values
var validRefNameRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

const fieldTypeText = "text"
const fieldTypeNumber = "number"
const fieldTypeDate = "date"

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`

	// Short name for referencing the field in a calculated fields
	RefName string `json:"refName"`

	// If IsCalcField is true, then the field is a calculated field. The
	// equation will be used to determine the values for this field.
	// The datastore can't store recursively nested structs, so
	// if field is a calculated field, then store a JSON representation
	// of the equation.
	CalcFieldEqn string `json:"calcFieldEqn"`
	IsCalcField  bool   `json:"isCalcField"` // defaults to false
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

func (fieldRef FieldRef) evalEqn(evalContext *EqnEvalContext) (*EquationResult, error) {

	field := fieldRef.FieldInfo

	if field.IsCalcField {
		// Calculated field - return the result of the calculation
		decodedEqn, decodeErr := decodeEquation(field.CalcFieldEqn)
		if decodeErr != nil {
			return nil, fmt.Errorf("Failure decoding equation for evaluation: %v", decodeErr)
		} else {
			calcFieldResult, calcFieldErr := decodedEqn.evalEqn(evalContext)
			if calcFieldErr != nil {
				return calcFieldResult, calcFieldErr
			} else if calcFieldResult.ResultType != field.Type {
				return nil, fmt.Errorf("evalEqn: type mismatch in result calculated for field: "+
					" expecting %v, got %v: field = %+v", field.Type, calcFieldResult.ResultType, field)
			} else {
				return calcFieldResult, nil
			}
		}
	} else { // literal field value
		switch field.Type {
		case fieldTypeText:
			if textResult, err := evalContext.resultRecord.GetTextRecordValue(
				evalContext.appEngContext, fieldRef.FieldID); err != nil {
				return nil, err
			} else {
				return textEqnResult(textResult), nil
			}
		case fieldTypeNumber:
			if numberResult, err := evalContext.resultRecord.GetNumberRecordValue(
				evalContext.appEngContext, fieldRef.FieldID); err != nil {
				return nil, err
			} else {
				return numberEqnResult(numberResult), nil
			}
			//		case fieldTypeDate:
		default:
			return nil, fmt.Errorf("Unknown field result type: %v", field.Type)

		} // switch

	} // field value is a literal, just return it
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

	newField.RefName = strings.TrimSpace(newField.RefName) // strip leading & trailing whitespace
	if !validRefNameRegexp.MatchString(newField.RefName) {
		return "", fmt.Errorf("Invalid reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			newField.RefName)
	}

	// TODO: Validate the reference name is unique versus the other names field names already in use.

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
		fieldRefs[i] = FieldRef{fieldID, currField}
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

type StringFieldRefMap map[string]FieldRef

type FieldRefIDIndex struct {
	fieldRefsByID      StringFieldRefMap
	fieldRefsByRefName StringFieldRefMap
}

func (fieldRefIDIndex FieldRefIDIndex) getFieldRefByID(fieldID string) (*FieldRef, error) {
	fieldRef, fieldRefFound := fieldRefIDIndex.fieldRefsByID[fieldID]
	if fieldRefFound != true {
		return nil, fmt.Errorf("getFieldRefByID: Unable to retrieve field for field with ID = %v ", fieldID)
	}
	return &fieldRef, nil

}

func GetFieldRefIDIndex(appEngContext appengine.Context) (*FieldRefIDIndex, error) {

	fieldRefs, getErr := GetAllFieldRefs(appEngContext)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldsByRefName: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	log.Printf("GetFieldRefIDIndex: Indexing %v fields", len(fieldRefs))

	fieldRefsByRefName := StringFieldRefMap{}
	fieldRefsByID := StringFieldRefMap{}
	for _, fieldRef := range fieldRefs {

		if _, keyExists := fieldRefsByRefName[fieldRef.FieldInfo.RefName]; keyExists == true {
			return nil, fmt.Errorf("GetFieldsByRefName: Unable to retrieve fields from datastore: "+
				" found duplicate reference name for field = %+v", fieldRef)
		}

		if _, keyExists := fieldRefsByID[fieldRef.FieldID]; keyExists == true {
			return nil, fmt.Errorf("GetFieldsByRefName: Unable to retrieve fields from datastore: "+
				" found duplicate key for field = %+v", fieldRef)
		}

		log.Printf("GetFieldRefIDIndex: Indexed field: %+v", fieldRef)

		fieldRefsByRefName[fieldRef.FieldInfo.RefName] = fieldRef
		fieldRefsByID[fieldRef.FieldID] = fieldRef

	}

	return &FieldRefIDIndex{fieldRefsByID, fieldRefsByRefName}, nil

}
