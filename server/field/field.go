package field

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"regexp"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"strings"
)

const fieldEntityKind string = "Field"

// A "reference name" for a field can only contain
// TODO - Can't start with "true or false" - add this when supporting boolean values
var validRefNameRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

const FieldTypeText string = "text"
const FieldTypeNumber string = "number"
const FieldTypeDate string = "date"
const FieldTypeBool string = "bool"

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
	CalcFieldEqn         string `json:"calcFieldEqn"`
	IsCalcField          bool   `json:"isCalcField"` // defaults to false
	CalcFieldFormulaText string `json:"calcFieldFormulaText"`
}

type FieldRef struct {
	FieldID   string `json:"fieldID"`
	FieldInfo Field  `json:"fieldInfo"`
}

type FieldsByType struct {
	TextFields   []FieldRef `json:"textFields"`
	DateFields   []FieldRef `json:"dateFields"`
	NumberFields []FieldRef `json:"numberFields"`
	BoolFields   []FieldRef `json:"boolFields"`
}

func validFieldType(fieldType string) bool {
	switch fieldType {
	case FieldTypeText:
		return true
	case FieldTypeNumber:
		return true
	case FieldTypeDate:
		return true
	case FieldTypeBool:
		return true
	default:
		return false
	}
}

// Internal function for creating new fields given raw inputs. Should only be called by
// other "NewField" functions with well-formed parameters for either a regular (non-calculated)
// or calculated field.
func CreateNewFieldFromRawInputs(appEngContext appengine.Context, newField Field) (string, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(newField.Name)
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

	fieldID, insertErr := datastoreWrapper.InsertNewEntity(appEngContext, fieldEntityKind, nil, &newField)
	if insertErr != nil {
		return "", fmt.Errorf("Can't create new field: error inserting into datastore: %v", insertErr)
	}

	// TODO - verify IntID != 0
	log.Printf("NewField: Created new field: id= %v, field='%+v'", fieldID, newField)

	return fieldID, nil

}

type NewFieldParams struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	RefName string `json:"refName"`
}

func NewField(appEngContext appengine.Context, fieldParams NewFieldParams) (string, error) {
	newField := Field{
		Name:                 fieldParams.Name,
		Type:                 fieldParams.Type,
		RefName:              fieldParams.RefName,
		CalcFieldEqn:         "",
		CalcFieldFormulaText: "",
		IsCalcField:          false} // always set calculated field to false

	return CreateNewFieldFromRawInputs(appEngContext, newField)
}

type GetFieldParams struct {
	// TODO - There will be more parameters once a field is
	// tied to a database table (i.e. TableID)
	FieldID string `json:"fieldID"`
}

func GetField(appEngContext appengine.Context, fieldID datastoreWrapper.UniqueRootID) (*Field, error) {
	var fieldGetDest Field
	if getErr := datastoreWrapper.GetRootEntityByID(appEngContext, fieldEntityKind,
		fieldID.ObjectID, &fieldGetDest); getErr != nil {
		return nil, fmt.Errorf("Unabled to get field: id = %+v: datastore err=%v", fieldID, getErr)
	}
	return &fieldGetDest, nil
}

func UpdateExistingField(appEngContext appengine.Context, fieldID datastoreWrapper.UniqueRootID, updatedField *Field) (*FieldRef, error) {

	if updateErr := datastoreWrapper.UpdateExistingRootEntity(appEngContext, fieldEntityKind,
		fieldID.ObjectID, updatedField); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingField: Can't set value: Error updating existing field: err = %v", updateErr)
	}

	fieldRef := FieldRef{
		FieldID:   fieldID.ObjectID,
		FieldInfo: *updatedField}

	return &fieldRef, nil

}

func GetFieldRef(appEngContext appengine.Context, fieldParams GetFieldParams) (*FieldRef, error) {

	fieldID := datastoreWrapper.UniqueRootID{ObjectID: fieldParams.FieldID}

	if field, getErr := GetField(appEngContext, fieldID); getErr != nil {
		return nil, getErr
	} else {
		return &FieldRef{fieldParams.FieldID, *field}, nil
	}
}

func GetFieldFromKey(appEngContext appengine.Context, fieldKey *datastore.Key) (*FieldRef, error) {

	fieldGetDest := Field{}
	getErr := datastore.Get(appEngContext, fieldKey, &fieldGetDest)
	if getErr != nil {
		return nil, fmt.Errorf("GetFieldFromKey: Failed to retrieve field from datastore: key=%+v, datastore err=%v",
			fieldKey, getErr)
	}

	fieldID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(fieldKey)
	if encodeErr != nil {
		return nil, fmt.Errorf("GetFieldFromKey: Failed to encode unique ID for field: key=%+v, encode err=%v",
			fieldKey, encodeErr)
	}

	return &FieldRef{fieldID, fieldGetDest}, nil
}

func GetExistingFieldKey(appEngContext appengine.Context, fieldID string) (*datastore.Key, error) {
	if len(fieldID) == 0 {
		return nil, fmt.Errorf("GetExistingFieldKey: Can't get field's key: missing field ID ")
	}
	fieldKey, fieldErr := datastoreWrapper.GetExistingRootEntityKey(appEngContext, fieldEntityKind, fieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("GetExistingFieldKey: invalid field ID '%v': datastore error=%v",
			fieldID, fieldErr)
	}
	return fieldKey, nil
}

func GetExistingFieldRefAndKey(appEngContext appengine.Context, fieldParams GetFieldParams) (*datastore.Key, *FieldRef, error) {
	// TODO - combine key retrieval and field retrieval
	fieldRef, fieldErr := GetFieldRef(appEngContext, GetFieldParams{fieldParams.FieldID})
	if fieldErr != nil {
		return nil, nil, fmt.Errorf("GetFieldRefAndKey: Can't get field for filter: datastore error = %v", fieldErr)
	}
	fieldKey, fieldKeyErr := GetExistingFieldKey(appEngContext, fieldParams.FieldID)
	if fieldKeyErr != nil {
		return nil, nil, fmt.Errorf("GetFieldRefAndKey: Can't create filtering rule with field ID = '%v': datastore error=%v",
			fieldParams.FieldID, fieldKeyErr)
	}

	return fieldKey, fieldRef, nil
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
		fieldID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(fieldKey)
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
		case FieldTypeText:
			fieldsByType.TextFields = append(fieldsByType.TextFields, fieldRef)
		case FieldTypeDate:
			fieldsByType.DateFields = append(fieldsByType.DateFields, fieldRef)
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
