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

	recordID, insertErr := SaveNewRecord(appEngCntxt, testRecord)
	if insertErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully created new record: id = %v", recordID)
	}

	if recordRef, getErr := GetRecord(appEngCntxt, GetRecordParams{recordID}); getErr != nil {
		t.Fatal(getErr)
	} else {
		t.Logf("Successfully retrieved new record: rec = %+v", recordRef)
		getVal := recordRef.FieldValues[fieldID]
		if getVal != testVal {
			t.Errorf("Value mismatch/missing in retrieved result: expecting '%v', got '%v'", testVal, getVal)
		}
	}

	setVal := "Another value for SetRecordValue()"
	// Set the record value, the updated record is returned, but is not needed for this testing
	setValParams := SetRecordValueParams{RecordID: recordID, FieldID: fieldID, Value: setVal}
	if _, setErr := SetRecordValue(appEngCntxt, setValParams); setErr != nil {
		t.Errorf("Error setting value: %v", setErr)
	}

	if recordRef2, getErr2 := GetRecord(appEngCntxt, GetRecordParams{recordID}); getErr2 != nil {
		t.Fatal(getErr2)
	} else {
		t.Logf("Successfully retrieved new record (2nd time): rec = %+v", recordRef2)
		getVal := recordRef2.FieldValues[fieldID]
		if getVal != setVal {
			t.Errorf("Value mismatch/missing in retrieved result: expecting '%v', got '%v'", setVal, getVal)
		}
	}

}
