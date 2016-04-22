package calcField

import (
	"appengine/aetest"
	"resultra/datasheet/server/record"
	"testing"
)

func TestSumFunc(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	val1 := 24.2
	eqn1 := EquationNode{NumberVal: &val1}

	val2 := 30.0
	eqn2 := EquationNode{NumberVal: &val2}

	// This test doesn't retrieve record values, so a dummy record will suffice
	dummyRecordRef := record.RecordRef{"dummyFieldID", record.Record{}}
	evalContext := EqnEvalContext{appEngCntxt, CalcFieldDefinedFuncs, dummyRecordRef}

	if funcResult, err := sumEvalFunc(&evalContext, []EquationNode{eqn1, eqn2}); err != nil {
		t.Error(err)
	} else {
		resultVal, resultErr := funcResult.GetNumberResult()
		if resultErr != nil {
			t.Error(err)
		} else {
			t.Logf("SUM result: %v", resultVal)
			if resultVal != 54.2 {
				t.Errorf("TestSumFunc: expected SUM() to return 54.2, got %v", resultVal)
			}
		}
	}

}
