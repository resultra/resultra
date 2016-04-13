package common

import (
	"testing"
)

func TestNameSanitize(t *testing.T) {

	// Leading or trailing whitespace will be stripped
	_, err := SanitizeName("ABC 123")
	if err != nil {
		t.Error(err)
	}

	// Empty names or names with newlines, tabs, or formfeeds are not OK
	_, err = SanitizeName("")
	if err == nil {
		t.Error(err)
	}

	_, err = SanitizeName("N\r\nF")
	if err == nil {
		t.Error(err)
	}

	_, err = SanitizeName("N\t\fF")
	if err == nil {
		t.Error(err)
	}

}
