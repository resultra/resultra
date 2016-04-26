package calcField

import (
	"resultra/datasheet/server/common/testUtil"
	"testing"
)

func verifyOneFormulaCompile(t *testing.T, inputStr string) {
	if eqn, err := compileFormula(inputStr); err != nil {
		t.Error(err)
	} else {
		t.Logf("Matched equation: %v", testUtil.EncodeJSONString(t, eqn))
	}
}

func verifyOneFormulaCompileVsExpected(t *testing.T, inputStr string, expectedJSON string) {
	if eqn, err := compileFormula(inputStr); err != nil {
		t.Error(err)
	} else {
		eqnJSON := testUtil.EncodeJSONString(t, eqn)
		t.Logf("Matched equation: %v", testUtil.EncodeJSONString(t, eqn))
		if eqnJSON != expectedJSON {
			t.Errorf("Unexpected equation result: expected=%v, got=%v",
				expectedJSON, eqnJSON)
		}
	}
}

func TestFormulaExpressionCompile(t *testing.T) {
	verifyOneFormulaCompile(t, "42.5")
	verifyOneFormulaCompile(t, "10 + 20")
	verifyOneFormulaCompile(t, "10 + 20 + 30")
	verifyOneFormulaCompile(t, "10 * 30")
	verifyOneFormulaCompile(t, "10 + 20 * 30")
}

func TestFormulaFuncCompile(t *testing.T) {
	verifyOneFormulaCompileVsExpected(t, "SQRT(42.5)", `{"funcName":"SQRT","funcArgs":[{"numberVal":42.5}]}`)
	verifyOneFormulaCompileVsExpected(t, "POW(10,2)", `{"funcName":"POW","funcArgs":[{"numberVal":10},{"numberVal":2}]}`)
}
