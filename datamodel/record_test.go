package datamodel

import (
	"testing"

	"appengine/aetest"
)

func TestNewRecord(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	var testRecord Record

	testField := Field{Name: "Test Field", Type: "text"}

	fieldID, err := NewField(appEngCntxt, testField)
	if err != nil {
		t.Fatal(err)
	}

	testVal := "Dummy Test Val"
	testRecord = Record{fieldID: testVal}
	t.Logf("Saving new record: rec = %+v", testRecord)

	recordID, insertErr := insertNewEntity(appEngCntxt, recordEntityKind, nil, &testRecord)
	if insertErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully created new record: id = %v", recordID)
	}

	getRecord := Record{}
	if getErr := getEntityByID(recordID, appEngCntxt, recordEntityKind, &getRecord); getErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully retrieved new record: rec = %+v", getRecord)
		getVal := getRecord[fieldID]
		if getVal != testVal {
			t.Errorf("Value mismatch/missing in retrieved result: expecting '%v', got '%v'", testVal, getVal)
		}
	}

}
