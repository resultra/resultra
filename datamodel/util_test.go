package datamodel

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"testing"
)

func newTestAppEngContext(t *testing.T) appengine.Context {
	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	return appEngCntxt
}

func newTestStronglyConsistentAppEngContext(t *testing.T) appengine.Context {
	// Many of the entities don't currently have a parent in the datastore. Using the strongly
	// consistent option ensures the datastore remains consistent for testing.
	appEngCntxt, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	return appEngCntxt
}

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

func TestIDEncodeDecode(t *testing.T) {
	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testData := Layout{Name: "Foo"}
	testKey := datastore.NewIncompleteKey(appEngCntxt, "TestEntityKind", nil)
	putKey, putErr := datastore.Put(appEngCntxt, testKey, &testData)
	if putErr != nil {
		t.Fatal(putErr)
	}

	encodedID, err := EncodeUniqueEntityIDToStr(putKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("testIDEncodeDecode: key ID: %v encoded ID: %v", putKey.IntID(), encodedID)

	decodedID, decodeErr := decodeUniqueEntityIDStrToInt(encodedID)
	if decodeErr != nil {
		t.Fatal(err)
	}

	if decodedID != putKey.IntID() {
		t.Errorf("Error decoding: expecting %v, got %v", putKey.IntID(), decodedID)
	}

}
