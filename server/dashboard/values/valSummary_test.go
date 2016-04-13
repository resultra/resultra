package values

import (
	"appengine"
	"appengine/aetest"
	"testing"
)

func verifyOneNewValSummaryFail(appEngContext appengine.Context, t *testing.T, params NewValSummaryParams, whyShouldFail string) {
	if _, newSummary, err := NewValSummary(appEngContext, params); err == nil {
		t.Errorf("verifyOneNewValSummaryFail: value group creation should have failed: %v, grouping = %+v", whyShouldFail, *newSummary)
	} else {
		t.Logf("verifyOneNewValSummaryFail: val grouping failed as expected: why fail = %v, err = %v", whyShouldFail, err)
	}
}

func verifyOneNewValSummary(appEngContext appengine.Context, t *testing.T, params NewValSummaryParams, whatTested string) {
	if _, newSummary, err := NewValSummary(appEngContext, params); err != nil {
		t.Errorf("verifyOneNewValSummary: value group creation should have succeeded: what tested = %v, err = %v", whatTested, err)
	} else {
		t.Logf("verifyOneNewValSummary: val grouping succeeded as expected: what tested = %v, grouping = %+v", whatTested, *newSummary)
	}
}

func TestNewValSummary(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}

	testNumField := newTestNumField(appEngCntxt, t, "NumField")
	verifyOneNewValSummary(appEngCntxt, t, NewValSummaryParams{testNumField, valSummaryCount}, "Count of numerical values")
	verifyOneNewValSummary(appEngCntxt, t, NewValSummaryParams{testNumField, valSummaryAvg}, "Average of numerical values")
	verifyOneNewValSummary(appEngCntxt, t, NewValSummaryParams{testNumField, valSummarySum}, "Sum of numerical values")

	testTextField := newTestTextField(appEngCntxt, t, "TextField")
	verifyOneNewValSummary(appEngCntxt, t, NewValSummaryParams{testTextField, valSummaryCount}, "Count of numerical values")
	verifyOneNewValSummaryFail(appEngCntxt, t, NewValSummaryParams{testTextField, valSummaryAvg}, "Average of text values")
	verifyOneNewValSummaryFail(appEngCntxt, t, NewValSummaryParams{testTextField, valSummarySum}, "Sum of text values")

}
