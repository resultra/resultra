package recordUpdate

import (
	"appengine"
	"appengine/aetest"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"testing"
)

// TODO - This unit test probably doesn't belong here. However, based upon package dependencies
// after refactoring the package organization, the test needed to be here. The calculated
// field equations are expected to be rewritten as more of a scripting language, at which
// point this test will then become obsolete.

// Helper functions for testing records

func newTestRecord(appEngContext appengine.Context, t *testing.T) *record.RecordRef {
	newRecordRef, insertErr := record.NewRecord(appEngContext)
	if insertErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully created new record: id = %v", newRecordRef.RecordID)
	}

	return newRecordRef
}

func setTestTextFieldVal(appEngContext appengine.Context, t *testing.T, recordRef *record.RecordRef, fieldID string, val string) {
	// Set the record value, the updated record is returned, but is not needed for this testing
	setParams := SetRecordTextValueParams{RecordUpdateHeader{recordRef.RecordID, fieldID}, val}
	if _, setErr := UpdateRecordValue(appEngContext, setParams); setErr != nil {
		t.Fatal(setErr)
	}

}

func TestNewRecord(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testField := field.NewFieldParams{Name: "Test Field", Type: "text", RefName: "TestRef1"}
	fieldID, err := field.NewField(appEngCntxt, testField)
	if err != nil {
		t.Fatal(err)
	}

	newRecordRef, insertErr := record.NewRecord(appEngCntxt)
	if insertErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully created new record: id = %v", newRecordRef.RecordID)
	}

	testVal := "Dummy Test Val"
	// Set the record value, the updated record is returned, but is not needed for this testing
	setInitialValParams := SetRecordTextValueParams{RecordUpdateHeader{newRecordRef.RecordID, fieldID}, testVal}
	if _, setErr := UpdateRecordValue(appEngCntxt, setInitialValParams); setErr != nil {
		t.Errorf("Error setting value: %v", setErr)
	}

	recordID := newRecordRef.RecordID

	if recordRef, getErr := record.GetRecord(appEngCntxt, record.RecordID{recordID}); getErr != nil {
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
	setValParams := SetRecordTextValueParams{RecordUpdateHeader{recordID, fieldID}, setVal}
	if _, setErr := UpdateRecordValue(appEngCntxt, setValParams); setErr != nil {
		t.Errorf("Error setting value: %v", setErr)
	}

	if recordRef2, getErr2 := record.GetRecord(appEngCntxt, record.RecordID{recordID}); getErr2 != nil {
		t.Fatal(getErr2)
	} else {
		t.Logf("Successfully retrieved new record (2nd time): rec = %+v", recordRef2)
		getVal := recordRef2.FieldValues[fieldID]
		if getVal != setVal {
			t.Errorf("Value mismatch/missing in retrieved result: expecting '%v', got '%v'", setVal, getVal)
		}
	}

}
