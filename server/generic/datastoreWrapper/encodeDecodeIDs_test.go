package datastoreWrapper

import (
	"appengine/aetest"
	"appengine/datastore"
	"testing"
)

type TestData struct {
	Name string
}

func TestIDEncodeDecode(t *testing.T) {
	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testData := TestData{Name: "Foo"}
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

	decodedID, decodeErr := DecodeUniqueEntityIDStrToInt(encodedID)
	if decodeErr != nil {
		t.Fatal(err)
	}

	if decodedID != putKey.IntID() {
		t.Errorf("Error decoding: expecting %v, got %v", putKey.IntID(), decodedID)
	}

}
