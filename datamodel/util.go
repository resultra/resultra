package datamodel

import (
	"appengine/datastore"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Names can include anything except newlines, form feeds, and tabs - spaces are OK,
// but leading and trailing whitespace (including spaces) is stripped (see below)
var validNameRegexp = regexp.MustCompile("^[^\t\n\f\r ][^\t\n\f\r]*$")

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

func SanitizeName(unsanitizedName string) (string, error) {

	stripWhite := strings.TrimSpace(unsanitizedName) // strip leading & trailing whitespace

	if !validNameRegexp.MatchString(stripWhite) {
		return "", errors.New("Invalid name: Cannot be empty and must not contain newlines or tabs")
	}
	return stripWhite, nil
}
