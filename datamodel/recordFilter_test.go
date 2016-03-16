package datamodel

import (
	"testing"
)

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
