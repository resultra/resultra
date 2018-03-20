package field

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/generic/uniqueID"
	"sync"
)

const FieldEntityKind string = "Field"

const FieldTypeText string = "text"
const FieldTypeNumber string = "number"
const FieldTypeTime string = "time"
const FieldTypeBool string = "bool"
const FieldTypeLongText string = "longText"
const FieldTypeAttachment string = "attachment"
const FieldTypeFile string = "file"
const FieldTypeUser string = "user"
const FieldTypeUsers string = "users"
const FieldTypeComment string = "comment"
const FieldTypeLabel string = "label"
const FieldTypeEmail string = "email"
const FieldTypeURL string = "url"
const FieldTypeImage string = "image"

type Field struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	FieldID          string `json:"fieldID"`

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
	case FieldTypeAttachment:
		return true
	case FieldTypeUser:
		return true
	case FieldTypeUsers:
		return true
	case FieldTypeComment:
		return true
	case FieldTypeLabel:
		return true
	case FieldTypeEmail:
		return true
	case FieldTypeFile:
		return true
	case FieldTypeURL:
		return true
	case FieldTypeImage:
		return true
	default:
		return false
	}
}

var newFieldMutex = &sync.Mutex{}

// Internal function for creating new fields given raw inputs. Should only be called by
// other "NewField" functions with well-formed parameters for either a regular (non-calculated)
// or calculated field.
func CreateNewFieldFromRawInputs(destDBHandle *sql.DB, newField Field) (*Field, error) {

	// Use a mutex when creating fields. This is necessary, so the validation of field properties against existing
	// fields can complete before inserting the new record with validated properties (notably a unique name and
	// formula reference name).
	newFieldMutex.Lock()
	defer newFieldMutex.Unlock()

	if !validFieldType(newField.Type) {
		return nil, fmt.Errorf("Can't create new field: invalid field type: '%v'", newField.Type)
	}

	if err := validateNewFieldRefName(destDBHandle, newField.ParentDatabaseID, newField.RefName); err != nil {
		return nil, fmt.Errorf("CreateNewFieldFromRawInputs: invalid formula reference name: '%v'", err)
	}

	if err := validateNewFieldName(destDBHandle, newField.ParentDatabaseID, newField.Name); err != nil {
		return nil, fmt.Errorf("CreateNewFieldFromRawInputs: invalid field name: '%v'", err)
	}

	if _, insertErr := destDBHandle.Exec(
		`INSERT INTO fields (database_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		newField.ParentDatabaseID,
		newField.FieldID,
		newField.Name,
		newField.Type,
		newField.RefName,
		newField.CalcFieldEqn,
		newField.IsCalcField,
		newField.PreprocessedFormulaText); insertErr != nil {
		return nil, fmt.Errorf("CreateNewFieldFromRawInputs: insert failed: error = %v", insertErr)
	}

	// TODO - verify IntID != 0
	log.Printf("CreateNewFieldFromRawInputs: Created new field: id= %v, field='%+v'", newField.FieldID, newField)

	return &newField, nil

}

type NewNonCalcFieldParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	RefName          string `json:"refName"`
}

func NewNonCalcField(trackerDBHandle *sql.DB, fieldParams NewNonCalcFieldParams) (*Field, error) {
	newField := Field{
		ParentDatabaseID:        fieldParams.ParentDatabaseID,
		FieldID:                 uniqueID.GenerateUniqueID(),
		Name:                    fieldParams.Name,
		Type:                    fieldParams.Type,
		RefName:                 fieldParams.RefName,
		CalcFieldEqn:            "",
		PreprocessedFormulaText: "",
		IsCalcField:             false} // always set calculated field to false

	return CreateNewFieldFromRawInputs(trackerDBHandle, newField)
}

// Getting an individual field doesn't require the table ID, since the field ID is unique.
// TODO - Refactor existing code to remove dependency on parent table ID to retrieve an individiual field.
func GetField(trackerDBHandle *sql.DB, fieldID string) (*Field, error) {
	var fieldGetDest Field

	if getErr := trackerDBHandle.QueryRow(`SELECT database_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
			FROM fields WHERE field_id=$1 LIMIT 1`, fieldID).Scan(
		&fieldGetDest.ParentDatabaseID,
		&fieldGetDest.FieldID,
		&fieldGetDest.Name,
		&fieldGetDest.Type,
		&fieldGetDest.RefName,
		&fieldGetDest.CalcFieldEqn,
		&fieldGetDest.IsCalcField,
		&fieldGetDest.PreprocessedFormulaText); getErr != nil {
		return nil, fmt.Errorf("GetField: Unabled to get field: id = %+v: datastore err=%v", fieldID, getErr)
	}

	return &fieldGetDest, nil

}

func UpdateExistingField(trackerDBHandle *sql.DB, updatedField *Field) (*Field, error) {

	if _, updateErr := trackerDBHandle.Exec(`UPDATE fields 
			SET name=$1,type=$2,ref_name=$3,calc_field_eqn=$4,preprocessed_formula_text=$5,is_calc_field=$6 
			WHERE database_id=$7 AND field_id=$8`,
		updatedField.Name,
		updatedField.Type,
		updatedField.RefName,
		updatedField.CalcFieldEqn,
		updatedField.PreprocessedFormulaText,
		updatedField.IsCalcField,
		updatedField.ParentDatabaseID,
		updatedField.FieldID); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingField: Error updating field %v: error = %v", updatedField.FieldID, updateErr)
	}

	return updatedField, nil

}

type GetFieldListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func GetAllFieldsFromSrc(srcDBHandle *sql.DB, params GetFieldListParams) ([]Field, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT database_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
		FROM fields WHERE database_id=$1`, params.ParentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	allFields := []Field{}
	for rows.Next() {
		var currField Field
		if scanErr := rows.Scan(&currField.ParentDatabaseID,
			&currField.FieldID,
			&currField.Name,
			&currField.Type,
			&currField.RefName,
			&currField.CalcFieldEqn,
			&currField.IsCalcField,
			&currField.PreprocessedFormulaText); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)

		}
		allFields = append(allFields, currField)
	}

	return allFields, nil
}

func GetAllFields(trackerDBHandle *sql.DB, params GetFieldListParams) ([]Field, error) {
	return GetAllFieldsFromSrc(trackerDBHandle, params)
}
