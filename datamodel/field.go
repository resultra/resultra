package datamodel

import (
	"appengine"
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

func NewField(appEngContext appengine.Context, newField Field) (string, error) {

	sanitizedName, sanitizeErr := sanitizeName(newField.Name)
	if sanitizeErr != nil {
		return "", fmt.Errorf("Can't create new field: invalid name: '%v'", sanitizeErr)
	}
	newField.Name = sanitizedName

	if !((newField.Type == fieldTypeText) ||
		(newField.Type == fieldTypeNumber) ||
		(newField.Type == fieldTypeDate)) {
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
