package field

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"regexp"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
	"strings"
)

const FieldEntityKind string = "Field"

// A "reference name" for a field can only contain
// TODO - Can't start with "true or false" - add this when supporting boolean values
var validRefNameRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

const FieldTypeText string = "text"
const FieldTypeNumber string = "number"
const FieldTypeTime string = "time"
const FieldTypeBool string = "bool"
const FieldTypeLongText string = "longText"
const FieldTypeFile string = "file"

type Field struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`

	Name string `json:"name"`
	Type string `json:"type"`

	// Short name for referencing the field in a calculated fields
	RefName string `json:"refName"`

	// If IsCalcField is true, then the field is a calculated field. The
	// equation will be used to determine the values for this field.
	// The datastore can't store recursively nested structs, so
	// if field is a calculated field, then store a JSON representation
	// of the equation.
	CalcFieldEqn            string `json:"calcFieldEqn"`
	IsCalcField             bool   `json:"isCalcField"` // defaults to false
	PreprocessedFormulaText string `json:"calcFieldFormulaText"`
}

const fieldIDFieldName string = "FieldID"

func validFieldType(fieldType string) bool {
	switch fieldType {
	case FieldTypeText:
		return true
	case FieldTypeLongText:
		return true
	case FieldTypeNumber:
		return true
	case FieldTypeTime:
		return true
	case FieldTypeBool:
		return true
	case FieldTypeFile:
		return true
	default:
		return false
	}
}

// Internal function for creating new fields given raw inputs. Should only be called by
// other "NewField" functions with well-formed parameters for either a regular (non-calculated)
// or calculated field.
func CreateNewFieldFromRawInputs(appEngContext appengine.Context, parentTableID string, newField Field) (string, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(newField.Name)
	if sanitizeErr != nil {
		return "", fmt.Errorf("Can't create new field: invalid name: '%v'", sanitizeErr)
	}
	newField.Name = sanitizedName

	if !validFieldType(newField.Type) {
		return "", fmt.Errorf("Can't create new field: invalid field type: '%v'", newField.Type)
	}

	newField.ParentTableID = parentTableID
	// The UUID for fields is substituted for the fields reference name when stored in the preprocessed formula.
	// The tokenizer for the formula compiler could potentially read the UUID as a number literal if there isn't a distinct prefix.
	newField.FieldID = gocql.TimeUUID().String()

	newField.RefName = strings.TrimSpace(newField.RefName) // strip leading & trailing whitespace
	if !validRefNameRegexp.MatchString(newField.RefName) {
		return "", fmt.Errorf("Invalid reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			newField.RefName)
	}

	// TODO: Validate the reference name is unique versus the other names field names already in use.

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return "", fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(
		`INSERT INTO field (tableID,fieldID,name,type,refname,calcFieldEqn,isCalcField,preprocessedFormulaText) VALUES (?,?,?,?,?,?,?,?)`,
		newField.ParentTableID,
		newField.FieldID,
		newField.Name,
		newField.Type,
		newField.RefName,
		newField.CalcFieldEqn,
		newField.IsCalcField,
		newField.PreprocessedFormulaText).Exec(); insertErr != nil {
		fmt.Errorf("CreateNewFieldFromRawInputs: Can't create field: unable to insert into database: error = %v", insertErr)
	}

	// TODO - verify IntID != 0
	log.Printf("CreateNewFieldFromRawInputs: Created new field: id= %v, field='%+v'", newField.FieldID, newField)

	return newField.FieldID, nil

}

type NewFieldParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	RefName       string `json:"refName"`
}

func NewField(appEngContext appengine.Context, fieldParams NewFieldParams) (string, error) {
	newField := Field{
		Name:                    fieldParams.Name,
		Type:                    fieldParams.Type,
		RefName:                 fieldParams.RefName,
		CalcFieldEqn:            "",
		PreprocessedFormulaText: "",
		IsCalcField:             false} // always set calculated field to false

	return CreateNewFieldFromRawInputs(appEngContext, fieldParams.ParentTableID, newField)
}

func GetField(appEngContext appengine.Context, fieldID string) (*Field, error) {

	var fieldGetDest Field

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("CreateNewFieldFromRawInputs: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	getErr := dbSession.Query(`SELECT tableID,fieldID,name,type,refname,calcFieldEqn,isCalcField,preprocessedFormulaText FROM field WHERE fieldID = ? LIMIT 1`,
		fieldID).Scan(&fieldGetDest.ParentTableID,
		&fieldGetDest.FieldID,
		&fieldGetDest.Name,
		&fieldGetDest.Type,
		&fieldGetDest.RefName,
		&fieldGetDest.CalcFieldEqn,
		&fieldGetDest.IsCalcField,
		&fieldGetDest.PreprocessedFormulaText)
	if getErr != nil {
		return nil, fmt.Errorf("Unabled to get field: id = %+v: datastore err=%v", fieldID, getErr)
	}

	return &fieldGetDest, nil
}

func UpdateExistingField(appEngContext appengine.Context, updatedField *Field) (*Field, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("UpdateExistingField: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if updateErr := dbSession.Query(`UPDATE field SET name=?,type=?,refName=?,calcFieldEqn=?,preprocessedFormulaText=?,isCalcField=? where fieldID=?`,
		updatedField.Name,
		updatedField.Type,
		updatedField.RefName,
		updatedField.CalcFieldEqn,
		updatedField.IsCalcField,
		updatedField.PreprocessedFormulaText,
		updatedField.FieldID).Exec(); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingField: Error updating field %v: error = %v", updatedField.FieldID, updateErr)
	}

	return updatedField, nil

}
