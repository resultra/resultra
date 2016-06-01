package datastoreWrapper

import (
	"appengine/datastore"
	"errors"
	"fmt"
)

// Return ID as a string: An integer representation of the database
// ID is an internal implementation, so clients of this package
// have no need to manipulate this ID as an integer. The "base 36"
// encoding is also more compact than an base 10 format.
func encodeUniqueEntityIDToStr(key *datastore.Key) (string, error) {

	if key == nil {
		return "", errors.New("Error decoding datastore ID: cannot decode nil key")
	}

	return key.Encode(), nil
}

func decodeUniqueEntityIDStrToKey(encodedID string) (*datastore.Key, error) {

	if len(encodedID) == 0 {
		return nil, fmt.Errorf("decodeUniqueEntityIDStrToKey: got an empty entity ID")
	}

	decodedKey, err := datastore.DecodeKey(encodedID)
	if err != nil {
		return nil, fmt.Errorf("Can't decode datastore id '%v': %v", encodedID, err)
	}
	return decodedKey, nil
}
