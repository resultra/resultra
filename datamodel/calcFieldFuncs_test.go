package datamodel

import (
	"appengine/aetest"
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

	evalContext := EqnEvalContext{appEngCntxt, calcFieldDefinedFuncs}

	if funcResult, err := sumEvalFunc(&evalContext, []EquationNode{eqn1, eqn2}); err != nil {
		t.Error(err)
	} else {
		resultVal, resultErr := funcResult.getNumberResult()
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
