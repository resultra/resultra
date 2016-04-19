package datastoreWrapper

import (
	"appengine/datastore"
	"errors"
	"fmt"
	"strconv"
)

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

func DecodeUniqueEntityIDStrToInt(encodedID string) (int64, error) {

	decodeVal, err := strconv.ParseInt(encodedID, 36, 64)
	if err != nil {
		return 0, fmt.Errorf("Can't decode datastore id: expecting base36 integer string, got '%v'", encodedID)
	}

	return decodeVal, nil
}
