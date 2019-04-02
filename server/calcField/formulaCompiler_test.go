// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package calcField

import (
	"github.com/resultra/resultra/server/common/testUtil"
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
	verifyOneFormulaCompile(t, " - 42.5")
	verifyOneFormulaCompile(t, "10 + 20")
	verifyOneFormulaCompile(t, "10 - 20")
	verifyOneFormulaCompile(t, "10 + 20 + 30")
	verifyOneFormulaCompile(t, "10 * 30")
	verifyOneFormulaCompile(t, "10 + 20 * 30")

	verifyOneFormulaCompile(t, "10 + 20 - 30")
	verifyOneFormulaCompile(t, "10 + 20 / 30")
	verifyOneFormulaCompile(t, "10 + (20 / 30)")
}

func TestFormulaFuncCompile(t *testing.T) {

	verifyOneFormulaCompileVsExpected(t, "SQRT(42.5)",
		`{"funcName":"SQRT","funcArgs":[{"numberVal":42.5}]}`)

	verifyOneFormulaCompileVsExpected(t, "POW(10,2)",
		`{"funcName":"POW","funcArgs":[{"numberVal":10},{"numberVal":2}]}`)

	// Mix functions into an expression
	verifyOneFormulaCompileVsExpected(t, "SQRT(42.5) * 30",
		`{"funcName":"MULTIPLY","funcArgs":[{"funcName":"SQRT","funcArgs":[{"numberVal":42.5}]},{"numberVal":30}]}`)

}

func TestFormulaFieldRefCompile(t *testing.T) {
	verifyOneFormulaCompileVsExpected(t, "[FieldRefA]",
		`{"fieldID":"FieldRefA"}`)

	verifyOneFormulaCompileVsExpected(t, "[FieldRefA] + [FieldRefB]",
		`{"funcName":"ADD","funcArgs":[{"fieldID":"FieldRefA"},{"fieldID":"FieldRefB"}]}`)

	verifyOneFormulaCompileVsExpected(t, "[FieldRefB] * 2.5",
		`{"funcName":"MULTIPLY","funcArgs":[{"fieldID":"FieldRefB"},{"numberVal":2.5}]}`)

}

func TestFormulaTextCompile(t *testing.T) {

	verifyOneFormulaCompileVsExpected(t, ` "Hello world"`, `{"textVal":"Hello world"}`)

	// Test escaped quotes within the string. The lexer should un-escape the escaped quotes
	verifyOneFormulaCompileVsExpected(t, ` "Hello \"world\"!"`, `{"textVal":"Hello \"world\"!"}`)

}

func verifyOneEqnParsing(t *testing.T, inputStr string) {
	if eqn, err := compileFormula(inputStr); err != nil {
		t.Error(err)
	} else {
		t.Logf("Matched equation: %v", testUtil.EncodeJSONString(t, eqn))
	}
}

func verifyOneEqnParsingVsExpected(t *testing.T, inputStr string, expectedJSON string) {
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

func verifyOneEqnParseFail(t *testing.T, inputStr string, whyShouldFail string) {
	if eqn, err := compileFormula(inputStr); err != nil {
		t.Logf("Got an expected parse error for input=%v (why should fail=%v), error=%v", inputStr, whyShouldFail, err)
	} else {
		t.Errorf("Matched equation when parsing should have failed: %v, input string = %v, why should fail = %v",
			testUtil.EncodeJSONString(t, eqn), inputStr, whyShouldFail)
	}
}

func TestNumberParse(t *testing.T) {

	verifyOneEqnParsingVsExpected(t, "25.2", `{"numberVal":25.2}`)
	verifyOneEqnParsingVsExpected(t, "25", `{"numberVal":25}`)
	verifyOneEqnParsingVsExpected(t, "-25", `{"funcName":"MULTIPLY","funcArgs":[{"numberVal":-1},{"numberVal":25}]}`)
	verifyOneEqnParsingVsExpected(t, "10e-5", `{"numberVal":0.0001}`)
	verifyOneEqnParsing(t, "10e5")
	verifyOneEqnParseFail(t, `0x42`, "HEX numbers not yet supported") // TODO - Support hex numbers
}

func TestTextParse(t *testing.T) {

	verifyOneEqnParsing(t, ` "Hello world"`)
	verifyOneEqnParseFail(t, ` "Hello world""`, "extra quote after text")
}

func TestEmptyFormulaParse(t *testing.T) {
	verifyOneEqnParsing(t, ``)
	verifyOneEqnParsing(t, `   `)
}

func TestFieldRefParse(t *testing.T) {

	verifyOneEqnParsing(t, `[FieldRef1]`)
	verifyOneEqnParsing(t, ` [   FieldRef1  ] `) // extra whitespace
	verifyOneEqnParsing(t, `SUM([FieldRef1],[FieldRef2])`)
	verifyOneEqnParsing(t, `SUM([FieldRef1],SUM([FieldRef3],4))`)
	verifyOneEqnParseFail(t, `[FieldRef1],[FieldRef2]`, "argument list outside of function")
	verifyOneEqnParseFail(t, `[FieldRef1`, "missing closing bracket")
	verifyOneEqnParseFail(t, `FieldRef1]`, "missing opening bracket")
	verifyOneEqnParseFail(t, `[]`, "missing identifier")
	verifyOneEqnParseFail(t, `[ ]`, "missing identifier")
	verifyOneEqnParseFail(t, `[FieldRef1 245]`, "extra characters inside brackets")
}

func TestGlobalRefParse(t *testing.T) {

	verifyOneEqnParsing(t, `[[GlobalRef1]]`)
	verifyOneEqnParsing(t, ` [[   GlobalRef1  ]] `) // extra whitespace
	verifyOneEqnParsing(t, `SUM([[GlobalRef1]],[[GlobalRef2]])`)

	verifyOneEqnParseFail(t, `[[FieldRef1`, "missing closing double bracket")
	verifyOneEqnParseFail(t, `FieldRef1]]`, "missing opening double bracket")
	verifyOneEqnParseFail(t, `[[FieldRef1 245]]`, "extra characters inside double brackets")
	verifyOneEqnParseFail(t, `[[]]`, "missing identifier inside double brackets")
	verifyOneEqnParseFail(t, `[[ ]]`, "missing identifier inside double brackets")
	verifyOneEqnParseFail(t, `[[42.5]]`, "invalid characters inside double brackets")
}

func TestFunctionParse(t *testing.T) {

	//	verifyOneEqnParsing(t, `SUM()`)
	verifyOneEqnParsing(t, `SUM(1)`)
	verifyOneEqnParsingVsExpected(t, `SUM(1,2)`, `{"funcName":"SUM","funcArgs":[{"numberVal":1},{"numberVal":2}]}`)
	verifyOneEqnParsingVsExpected(t, `1+2`, `{"funcName":"ADD","funcArgs":[{"numberVal":1},{"numberVal":2}]}`)
	verifyOneEqnParsingVsExpected(t, `1/2`, `{"funcName":"DIVIDE","funcArgs":[{"numberVal":1},{"numberVal":2}]}`)
	verifyOneEqnParsingVsExpected(t, `1-2`, `{"funcName":"MINUS","funcArgs":[{"numberVal":1},{"numberVal":2}]}`)
	verifyOneEqnParsingVsExpected(t, `1*2`, `{"funcName":"MULTIPLY","funcArgs":[{"numberVal":1},{"numberVal":2}]}`)
	verifyOneEqnParsing(t, `SUM(1,2,"arg three")`)
	verifyOneEqnParsingVsExpected(t, `SUM(1,2,PRODUCT([FieldRef1],2.5))`, `{"funcName":"SUM","funcArgs":[{"numberVal":1},{"numberVal":2},{"funcName":"PRODUCT","funcArgs":[{"fieldID":"FieldRef1"},{"numberVal":2.5}]}]}`)
	verifyOneEqnParseFail(t, `SUM(1,,2)`, "extra comma between arguments")
}
