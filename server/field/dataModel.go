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
func CreateNewFieldFromRawInputs(parentTableID string, newField Field) (string, error) {

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
	newField.FieldID = uniqueID.GenerateSnowflakeID()

	if !generic.WellFormedFormulaReferenceName(newField.RefName) {
		return "", fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			newField.RefName)

	}

	// TODO: Validate the reference name is unique versus the other names field names already in use.

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO fields (table_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		newField.ParentTableID,
		newField.FieldID,
		newField.Name,
		newField.Type,
		newField.RefName,
		newField.CalcFieldEqn,
		newField.IsCalcField,
		newField.PreprocessedFormulaText); insertErr != nil {
		return "", fmt.Errorf("CreateNewFieldFromRawInputs: insert failed: error = %v", insertErr)
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

func NewField(fieldParams NewFieldParams) (string, error) {
	newField := Field{
		Name:                    fieldParams.Name,
		Type:                    fieldParams.Type,
		RefName:                 fieldParams.RefName,
		CalcFieldEqn:            "",
		PreprocessedFormulaText: "",
		IsCalcField:             false} // always set calculated field to false

	return CreateNewFieldFromRawInputs(fieldParams.ParentTableID, newField)
}

// Getting an individual field doesn't require the table ID, since the field ID is unique.
// TODO - Refactor existing code to remove dependency on parent table ID to retrieve an individiual field.
func GetFieldWithoutTableID(fieldID string) (*Field, error) {
	var fieldGetDest Field

	if getErr := databaseWrapper.DBHandle().QueryRow(`SELECT table_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
			FROM fields WHERE field_id=$1 LIMIT 1`, fieldID).Scan(
		&fieldGetDest.ParentTableID,
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

func GetField(tableID string, fieldID string) (*Field, error) {

	return GetFieldWithoutTableID(fieldID)
}

func UpdateExistingField(updatedField *Field) (*Field, error) {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE fields 
			SET name=$1,type=$2,ref_name=$3,calc_field_eqn=$4,preprocessed_formula_text=$5,is_calc_field=$6 
			WHERE table_id=$7 AND field_id=$8`,
		updatedField.Name,
		updatedField.Type,
		updatedField.RefName,
		updatedField.CalcFieldEqn,
		updatedField.PreprocessedFormulaText,
		updatedField.IsCalcField,
		updatedField.ParentTableID,
		updatedField.FieldID); updateErr != nil {
		return nil, fmt.Errorf("UpdateExistingField: Error updating field %v: error = %v", updatedField.FieldID, updateErr)
	}

	return updatedField, nil

}

type GetFieldListParams struct {
	ParentTableID string `json:"parentTableID"`
}

func GetAllFields(params GetFieldListParams) ([]Field, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT table_id,field_id,name,type,ref_name,calc_field_eqn,is_calc_field,preprocessed_formula_text 
		FROM fields WHERE table_id=$1`, params.ParentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
	allFields := []Field{}
	for rows.Next() {
		var currField Field
		if scanErr := rows.Scan(&currField.ParentTableID,
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
