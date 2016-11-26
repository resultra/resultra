package common

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
)

const LinkedValTypeGlobal string = "global"
const LinkedValTypeField string = "field"

type ComponentLink struct {
	LinkedValType string `json:"linkedValType"`
	FieldID       string `json:"fieldID"`
	GlobalID      string `json:"globalID"`
}

func validLinkedValType(valType string) bool {
	if valType == LinkedValTypeGlobal || valType == LinkedValTypeField {
		return true
	} else {
		return false
	}
}

type FieldTypeValidationFunc func(string) bool

func ValidateComponentLink(compLink ComponentLink, fieldTypeValidationFunc FieldTypeValidationFunc) error {

	if !validLinkedValType(compLink.LinkedValType) {
		return fmt.Errorf("verifyComponentLink: Invalid linked value type: %v", compLink.LinkedValType)
	}

	if compLink.LinkedValType == LinkedValTypeField {
		field, fieldErr := field.GetFieldWithoutTableID(compLink.FieldID)
		if fieldErr != nil {
			return fmt.Errorf("verifyComponentLink: Can't get field with field ID = '%v': datastore error=%v",
				compLink.FieldID, fieldErr)
		}

		if !fieldTypeValidationFunc(field.Type) {
			return fmt.Errorf("verifyComponentLink: Invalid field type %v", field.Type)
		}

	}

	return nil

}

func (srcLink ComponentLink) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ComponentLink, error) {

	destLink := srcLink

	if srcLink.LinkedValType == LinkedValTypeGlobal {
		remappedGlobalID, err := remappedIDs.GetExistingRemappedID(srcLink.GlobalID)
		if err != nil {
			return nil, fmt.Errorf("ComponentLink.Clone: %v", err)
		}
		destLink.GlobalID = remappedGlobalID
	} else {
		remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcLink.FieldID)
		if err != nil {
			return nil, fmt.Errorf("ComponentLink.Clone: %v", err)
		}
		destLink.FieldID = remappedFieldID
	}

	return &destLink, nil

}
