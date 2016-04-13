package form

import (
	"appengine/aetest"
	"testing"
)

func TestNewLayout(t *testing.T) {
	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	layoutID, err := NewLayout(appEngCntxt, "My Layout")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully created new layout: id = %v", layoutID)
	}

	// Leading or trailing whitespace will be stripped
	layoutID, err = NewLayout(appEngCntxt, " My 2nd Layout ")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully created new layout: id = %v", layoutID)
	}
}
