package calcField

import (
	"testing"
)

func testOnePreprocess(t *testing.T, inputStr string, fieldReplMap IdentReplacementMap,
	globalReplMap IdentReplacementMap,
	expectedOutput string, whatTest string) {

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

func testOnePreprocessFail(t *testing.T, inputStr string, fieldReplMap IdentReplacementMap,
	globalReplMap IdentReplacementMap, whatTest string) {

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
	globalReplMap := IdentReplacementMap{}

	fieldReplMap["fieldRef1"] = "fieldID1"
	fieldReplMap["fieldRef2"] = "fieldID2"

	globalReplMap["globalRef1"] = "globalID1"
	globalReplMap["globalRef2"] = "globalID2"

	testOnePreprocess(t, "[fieldRef1]", fieldReplMap, globalReplMap, "[fieldID1]", "basic field")
	testOnePreprocess(t, "[ fieldRef2 ]", fieldReplMap, globalReplMap, "[ fieldID2 ]", "whitespace between left bracket and field ref name")
	testOnePreprocess(t, "[fieldRef2] [fieldRef1]", fieldReplMap, globalReplMap, "[fieldID2] [fieldID1]", "multiple field references")
	testOnePreprocess(t, "variable = sum([fieldRef1])", fieldReplMap, globalReplMap, "variable = sum([fieldID1])", "field ref mixed in with other syntax")

	testOnePreprocessFail(t, "[fieldRef3]", fieldReplMap, globalReplMap, "no field ID in map for fieldRef3")
	testOnePreprocessFail(t, "[()]", fieldReplMap, globalReplMap, "invalid token inside brackets - needs to be an identifyer")
	testOnePreprocessFail(t, "[]", fieldReplMap, globalReplMap, "no token inside brackets - needs to be an identifyer")
	testOnePreprocessFail(t, "[%]", fieldReplMap, globalReplMap, "unrecognized token % - needs to be an identifyer")

	testOnePreprocess(t, "[[globalRef1]]", fieldReplMap, globalReplMap, "[[globalID1]]", "basic global")
	testOnePreprocess(t, "[[globalRef1]]*[[globalRef2]]", fieldReplMap, globalReplMap, "[[globalID1]]*[[globalID2]]", "basic global - multiple globals")
	testOnePreprocess(t, "[[ globalRef1 ]]", fieldReplMap, globalReplMap, "[[ globalID1 ]]", "basic global - leading/trailing whitespace OK")
	testOnePreprocessFail(t, "[[]]", fieldReplMap, globalReplMap, "no token inside brackets - needs to be an identifyer")
	testOnePreprocessFail(t, "[[25]]", fieldReplMap, globalReplMap, "no identifier token inside brackets - needs to be an identifier")
	testOnePreprocessFail(t, "[[-32]]", fieldReplMap, globalReplMap, "no identifier token inside brackets - needs to be an identifier")
	testOnePreprocessFail(t, "[[globalRef3]]", fieldReplMap, globalReplMap, "no global named globalRef3")

}
