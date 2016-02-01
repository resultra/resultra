package datamodel

import (
	"appengine"
	"appengine/datastore"
	"log"
	"strings"
)

const LayoutEntityKindSuffix string = "Layout"
const EntityKindPartsSep string = "_"

// Dummied-up user and database IDs - TODO replace with
// real ones later.
const dummyUserID string = "dummyUserID"
const dummyDatabaseID string = "dummyDBID"

type Layout struct {
	Name string `json:"name"`
}

func layoutEntityKind(userID string, databaseID string) string {
	// TODO - Verify userID and database ID (panic or throw error if not valid)
	kindParts := []string{userID, databaseID, LayoutEntityKindSuffix}
	return strings.Join(kindParts, EntityKindPartsSep)

}

func NewLayout(appEngContext appengine.Context, layoutName string) (string, error) {

	sanitizedLayoutName, sanitizeErr := sanitizeName(layoutName)
	if sanitizeErr != nil {
		return "", sanitizeErr
	}

	var newLayout = Layout{sanitizedLayoutName}

	entityKind := layoutEntityKind(dummyUserID, dummyDatabaseID)
	newKey := datastore.NewIncompleteKey(appEngContext, entityKind, nil)
	key, err := datastore.Put(appEngContext, newKey, &newLayout)
	if err != nil {
		return "", err
	}

	layoutID, encodeErr := encodeUniqueEntityIDToStr(key)
	if encodeErr != nil {
		return "", encodeErr
	}

	// TODO - verify IntID != 0
	log.Printf("NewLayout: Created new Layout: id= %v, name='%v'", layoutID, sanitizedLayoutName)

	return layoutID, nil

}
