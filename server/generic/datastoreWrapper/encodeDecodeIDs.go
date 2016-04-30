package datastoreWrapper

import (
	"appengine/datastore"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// UniqueID stores the opaque/encoded database keys for an entity.
type decodedChildID struct {
	parentID string
	childID  string
}

// Return ID as a string: An integer representation of the database
// ID is an internal implementation, so clients of this package
// have no need to manipulate this ID as an integer. The "base 36"
// encoding is also more compact than an base 10 format.
func EncodeUniqueEntityIDToStr(key *datastore.Key) (string, error) {

	if key == nil {
		return "", errors.New("Error decoding datastore ID: cannot decode nil key")
	}

	id := key.IntID()
	if id == 0 {
		return "", errors.New("Error encoding datastore ID: cannot encode 0")
	}

	encodedID := strconv.FormatInt(key.IntID(), 36)
	return encodedID, nil
}

func decodeUniqueEntityIDStrToInt(encodedID string) (int64, error) {

	decodeVal, err := strconv.ParseInt(encodedID, 36, 64)
	if err != nil {
		return 0, fmt.Errorf("Can't decode datastore id: expecting base36 integer string, got '%v'", encodedID)
	}

	return decodeVal, nil
}

func encodeChildEntityIDToStr(parentID string, childID string) string {
	return parentID + "-" + childID
}

func decodeUniqueChildID(encodedChildID string) (*decodedChildID, error) {
	splitParentChildIDs := strings.Split(encodedChildID, "-")
	if len(splitParentChildIDs) != 2 {
		return nil, fmt.Errorf("Can't decode datastore id: unrecocognized child id = '%v'", encodedChildID)
	}

	parentID := splitParentChildIDs[0]
	if len(parentID) == 0 {
		return nil, fmt.Errorf("Can't decode datastore id: unrecocognized parent id = '%v'", encodedChildID)
	}

	childID := splitParentChildIDs[1]
	if len(childID) == 0 {
		return nil, fmt.Errorf("Can't decode datastore id: unrecocognized child id = '%v'", encodedChildID)
	}

	return &decodedChildID{parentID: parentID, childID: childID}, nil

}
