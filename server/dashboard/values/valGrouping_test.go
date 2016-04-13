package values

import (
	"appengine"
	"appengine/aetest"
	"resultra/datasheet/server/field"
	"testing"
)

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

func verifyOneNewValGroupingFail(appEngContext appengine.Context, t *testing.T, params NewValGroupingParams, whyShouldFail string) {
	if _, newGrouping, err := NewValGrouping(appEngContext, params); err == nil {
		t.Errorf("verifyOneNewValGroupingFail: value group creation should have failed: %v, grouping = %+v", whyShouldFail, *newGrouping)
	} else {
		t.Logf("verifyOneNewValGroupingFail: val grouping failed as expected: why fail = %v, err = %v", whyShouldFail, err)
	}
}

func verifyOneNewValGrouping(appEngContext appengine.Context, t *testing.T, params NewValGroupingParams, whatTested string) {
	if _, newGrouping, err := NewValGrouping(appEngContext, params); err != nil {
		t.Errorf("verifyOneNewValGroupingFail: value group creation should have succeeded: what tested = %v, err = %v", whatTested, err)
	} else {
		t.Logf("verifyOneNewValGroupingFail: val grouping succeeded as expected: what tested = %v, grouping = %+v", whatTested, *newGrouping)
	}
}

func TestNewValGrouping(t *testing.T) {
	appEngCntxt, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}

	testNumField := newTestNumField(appEngCntxt, t, "NumField")
	verifyOneNewValGroupingFail(appEngCntxt, t, NewValGroupingParams{testNumField, valGroupByDay, 0.0}, "Group non-date field by day")
	verifyOneNewValGrouping(appEngCntxt, t, NewValGroupingParams{testNumField, valGroupByNone, 0.0}, "Group number field by none")

	verifyOneNewValGrouping(appEngCntxt, t, NewValGroupingParams{testNumField, valGroupByBucket, 1.0}, "Group number field by bucket")
	verifyOneNewValGroupingFail(appEngCntxt, t, NewValGroupingParams{testNumField, valGroupByBucket, 0.0},
		"Group number field by bucket (invalid bucket size)")
	verifyOneNewValGroupingFail(appEngCntxt, t, NewValGroupingParams{testNumField, valGroupByBucket, -10.0},
		"Group number field by bucket (invalid bucket size)")

	testTextField := newTestTextField(appEngCntxt, t, "TextField")
	verifyOneNewValGrouping(appEngCntxt, t, NewValGroupingParams{testTextField, valGroupByNone, 0.0}, "Group text field by none")
	verifyOneNewValGroupingFail(appEngCntxt, t, NewValGroupingParams{testTextField, valGroupByDay, 0.0}, "Group text field by day")
	verifyOneNewValGroupingFail(appEngCntxt, t, NewValGroupingParams{testTextField, valGroupByBucket, 0.0}, "Group text field by bucket")
}
