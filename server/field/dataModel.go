package field

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
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

// Internal function for creating new fields given raw inputs. Should only be called by
// other "NewField" functions with well-formed parameters for either a regular (non-calculated)
// or calculated field.
func CreateNewFieldFromRawInputs(newField Field) (*Field, error) {

	// TODO Validate field name

	if !validFieldType(newField.Type) {
		return nil, fmt.Errorf("Can't create new field: invalid field type: '%v'", newField.Type)
	}

	if !generic.WellFormedFormulaReferenceName(newField.RefName) {
		return nil, fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			newField.RefName)

	}

	// TODO: Validate the reference name is unique versus the other names field names already in use.

	if _, insertErr := databaseWrapper.DBHandle().Exec(
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

func NewNonCalcField(fieldParams NewNonCalcFieldParams) (*Field, error) {
	newField := Field{
		ParentDatabaseID:        fieldParams.ParentDatabaseID,
		FieldID:                 uniqueID.GenerateSnowflakeID(),
		Name:                    fieldParams.Name,
		Type:                    fieldParams.Type,
		RefName:                 fieldParams.RefName,
		CalcFieldEqn:            "",
		PreprocessedFormulaText: "",
		IsCalcField:             false} // always set calculated field to false

	return CreateNewFieldFromRawInputs(newField)
}

// Getting an individual field doesn't require the table ID, since the field ID is unique.
// TODO - Refactor existing code to remove dependency on parent table ID to retrieve an individiual field.
func GetField(fieldID string) (*Field, error) {
	var fieldGetDest Field

	if getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
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

func UpdateExistingField(updatedField *Field) (*Field, error) {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE fields 
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

func GetAllFields(params GetFieldListParams) ([]Field, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
		FROM fields WHERE database_id=$1`, params.ParentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
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
