package datamodel

import (
	"testing"
)

func TestNameSanitize(t *testing.T) {

	// Leading or trailing whitespace will be stripped
	_, err := sanitizeName("ABC 123")
	if err != nil {
		t.Error(err)
	}

	// Empty names or names with newlines, tabs, or formfeeds are not OK
	_, err = sanitizeName("")
	if err == nil {
		t.Error(err)
	}

	_, err = sanitizeName("N\r\nF")
	if err == nil {
		t.Error(err)
	}

	_, err = sanitizeName("N\t\fF")
	if err == nil {
		t.Error(err)
	}

}
