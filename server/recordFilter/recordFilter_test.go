package recordFilter

import (
	"appengine"
	"appengine/aetest"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordUpdate"
	"testing"
)

// Helper functions for testing fields

func newTestNumField(appEngContext appengine.Context, t *testing.T, refName string) string {

	testNumField := field.NewFieldParams{Name: refName, Type: field.FieldTypeNumber, RefName: refName}
	testNumFieldID, numFieldErr := field.NewField(appEngContext, testNumField)
	if numFieldErr != nil {
		t.Fatal(numFieldErr)
	}

	return testNumFieldID
}

func newTestTextField(appEngContext appengine.Context, t *testing.T, refName string) string {

	testField := field.NewFieldParams{Name: "Test Text Field", Type: field.FieldTypeText, RefName: "TestRef1"}
	testFieldID, err := field.NewField(appEngContext, testField)
	if err != nil {
		t.Fatal(err)
	}

	return testFieldID
}

func setTestTextFieldVal(appEngContext appengine.Context, t *testing.T, recordRef *record.RecordRef, fieldID string, val string) {
	// Set the record value, the updated record is returned, but is not needed for this testing
	setParams := recordUpdate.SetRecordTextValueParams{recordUpdate.RecordUpdateHeader{RecordID: recordRef.RecordID, FieldID: fieldID}, val}
	if _, setErr := recordUpdate.UpdateRecordValue(appEngContext, setParams); setErr != nil {
		t.Fatal(setErr)
	}

}

func newTestRecord(appEngContext appengine.Context, t *testing.T) *record.RecordRef {
	newRecordRef, insertErr := record.NewRecord(appEngContext)
	if insertErr != nil {
		t.Fatal(insertErr)
	} else {
		t.Logf("Successfully created new record: id = %v", newRecordRef.RecordID)
	}

	return newRecordRef
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

func TestRecordFilter(t *testing.T) {

	appEngCntxt := newTestStronglyConsistentAppEngContext(t)

	//	testNumFieldID := newTestNumField(appEngCntxt, t, "TestNumField")
	testTextFieldID := newTestTextField(appEngCntxt, t, "TestTextField")

	record1 := newTestRecord(appEngCntxt, t)
	newTestRecord(appEngCntxt, t) // test record with no value set, shouldn't match

	setTestTextFieldVal(appEngCntxt, t, record1, testTextFieldID, "non blank value for record 1")

	newTestFilterRuleNoParams(appEngCntxt, t, filterRuleIDNotBlank, testTextFieldID)

	filteredRecs, filterErr := GetFilteredRecords(appEngCntxt)

	if filterErr != nil {
		t.Fatal(filterErr)
	}

	t.Logf("Number of records matched by filter: %v", len(filteredRecs))

	if len(filteredRecs) != 1 {
		t.Fatalf("Number of records matched by filter: %v", len(filteredRecs))
	}

	filteredRec := filteredRecs[0]
	t.Logf("Matched record: %+v", filteredRec)

	if filteredRec.RecordID != record1.RecordID {
		t.Errorf("Unexpected filtered record: got id = %v, expecting id = %v", filteredRec.RecordID, record1.RecordID)
	}

}
