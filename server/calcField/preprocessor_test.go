package calcField

import (
	"testing"
)

func testOnePreprocess(t *testing.T, inputStr string, fieldReplMap IdentReplacementMap,
	expectedOutput string, whatTest string) {

	globalReplMap := IdentReplacementMap{}

	if preprocessedStr, err := preprocessFormulaInput(inputStr, fieldReplMap, globalReplMap); err != nil {
		t.Error(err)
	} else {
		if preprocessedStr != expectedOutput {
			t.Errorf("testOnePreprocess (fail): Unexpected preprocessor output: %v: input=%v, got=%v, expected=%v",
				whatTest, inputStr, preprocessedStr, expectedOutput)
		} else {
			t.Logf("testOnePreprocess (pass): input=%v, preprocessed=%v", inputStr, preprocessedStr)
		}
	}
}

func testOnePreprocessFail(t *testing.T, inputStr string, fieldReplMap IdentReplacementMap, whatTest string) {

	globalReplMap := IdentReplacementMap{}

	if preprocessedStr, err := preprocessFormulaInput(inputStr, fieldReplMap, globalReplMap); err == nil {
		t.Errorf("testOnePreprocess (fail): Expecting failure, but preprocessing succeeded: %v: input=%v, output=%v",
			whatTest, inputStr, preprocessedStr)
	} else {
		t.Logf("testOnePreprocess (pass): Preprocessing failed as expected: %v: input=%v, error=%v",
			whatTest, inputStr, err)
	}
}

func TestPreprocess(t *testing.T) {

	fieldReplMap := IdentReplacementMap{}

	fieldReplMap["fieldRef1"] = "fieldID1"
	fieldReplMap["fieldRef2"] = "fieldID2"

	testOnePreprocess(t, "[fieldRef1]", fieldReplMap, "[fieldID1]", "basic field")
	testOnePreprocess(t, "[ fieldRef2 ]", fieldReplMap, "[ fieldID2 ]", "whitespace between left bracket and field ref name")
	testOnePreprocess(t, "[fieldRef2] [fieldRef1]", fieldReplMap, "[fieldID2] [fieldID1]", "multiple field references")
	testOnePreprocess(t, "variable = sum([fieldRef1])", fieldReplMap, "variable = sum([fieldID1])", "field ref mixed in with other syntax")

	testOnePreprocessFail(t, "[fieldRef3]", fieldReplMap, "no field ID in map for fieldRef3")
	testOnePreprocessFail(t, "[()]", fieldReplMap, "invalid token inside brackets - needs to be an identifyer")
	testOnePreprocessFail(t, "[]", fieldReplMap, "no token inside brackets - needs to be an identifyer")
	testOnePreprocessFail(t, "[%]", fieldReplMap, "unrecognized token % - needs to be an identifyer")

}
