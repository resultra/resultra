package calcField

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/table"
	"strings"
)

func compileFormula(inputStr string) (*EquationNode, error) {

	lexer, lexerErr := newFormulaLexer(inputStr)
	if lexerErr != nil {
		return nil, lexerErr
	}

	parseResult := formulaParse(lexer)
	if parseResult != 0 {
		return nil, fmt.Errorf("Parse error")
	} else {
		// As a work-around to a yacc limitation, the lexer is used as local storage
		// for the parser. See the comment in formulaParser.y for details.
		rootEqnNode := lexer.rootEqnNode
		return rootEqnNode, nil
	}

}

func preprocessCalcFieldFormula(compileParams formulaCompileParams) (string, error) {

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(compileParams.appEngContext,
		field.GetFieldListParams{ParentTableID: compileParams.parentTableID})
	if indexErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: Unable to retrieve fields list for table: tableID=%v, error=%v ",
			compileParams.parentTableID, indexErr)
	}

	fieldRefFieldIDMap := FieldNameReplacementMap{}
	for fieldRefName, fieldRef := range fieldRefIndex.FieldRefsByRefName {
		fieldRefFieldIDMap[fieldRefName] = fieldRef.FieldID
	}

	preprocessOutput, preprocessErr := preprocessFormulaInput(compileParams.formulaText, fieldRefFieldIDMap)
	if preprocessErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: Error preprocessing formula: error=%v ", preprocessErr)
	}

	return preprocessOutput, nil

}

// reverseProcessCalcFieldFormula does the opposite of preprocessCalcFieldFormula. It takes a preprocessed formula
// with embedded permanent field IDs and replaces them with the most up to date field reference names. This
// allows the user to change the reference name after saving a formula. Since the formula is stored with a
// field's permanent & unique ID, when the formula's text is retrieved again it will have the most up to
// date reference name.
func reverseProcessCalcFieldFormula(compileParams formulaCompileParams) (string, error) {

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(compileParams.appEngContext,
		field.GetFieldListParams{ParentTableID: compileParams.parentTableID})
	if indexErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: Unable to retrieve fields list for table: tableID=%v, error=%v ",
			compileParams.parentTableID, indexErr)
	}

	fieldRefFieldIDMap := FieldNameReplacementMap{}
	for _, fieldRef := range fieldRefIndex.FieldRefsByRefName {
		fieldRefFieldIDMap[fieldRef.FieldID] = fieldRef.FieldInfo.RefName
	}

	log.Printf("reverseProcessCalcFieldFormula: Starting reverse processing of pre-processed formula text: %v",
		compileParams.formulaText)

	reverseProcessOutput, err := preprocessFormulaInput(compileParams.formulaText, fieldRefFieldIDMap)
	if err != nil {
		return "", fmt.Errorf("reverseProcessCalcFieldFormula: Error loading formula: error=%v ", err)
	}

	log.Printf("reverseProcessCalcFieldFormula: Reverse processed formula text: %v", reverseProcessOutput)

	return reverseProcessOutput, nil

}

type GetRawFormulaParams struct {
	FieldID string `json:"fieldID"`
}

type GetRawFormulaResult struct {
	FieldID        string `json:"fieldID"`
	RawFormulaText string `json:"rawFormulaText"`
}

func getRawFormulaText(appEngContext appengine.Context, params GetRawFormulaParams) (*GetRawFormulaResult, error) {

	parentTableID, getParentErr := datastoreWrapper.GetParentID(params.FieldID, table.TableEntityKind)
	if getParentErr != nil {
		return nil, fmt.Errorf("getRawFormulaText: Unable to get parent table for field: field id =%v, error=%v ",
			params.FieldID, getParentErr)
	}

	calcField, getFieldErr := field.GetField(appEngContext, params.FieldID)
	if getFieldErr != nil {
		return nil, fmt.Errorf("getRawFormulaText: Unable to get calculated field field: field id =%v, error=%v ",
			params.FieldID, getFieldErr)
	}

	compileParams := formulaCompileParams{
		appEngContext:      appEngContext,
		formulaText:        calcField.PreprocessedFormulaText,
		parentTableID:      parentTableID,
		expectedResultType: calcField.Type,
		resultFieldID:      params.FieldID}

	rawFormulaText, reverseProcessErr := reverseProcessCalcFieldFormula(compileParams)
	if reverseProcessErr != nil {
		return nil, fmt.Errorf("getRawFormulaText: Unable to read calculated field field: field id =%v, error=%v ",
			params.FieldID, reverseProcessErr)
	}

	return &GetRawFormulaResult{FieldID: params.FieldID, RawFormulaText: rawFormulaText}, nil
}

type formulaCompileParams struct {
	appEngContext      appengine.Context
	formulaText        string
	parentTableID      string
	expectedResultType string

	// This is the fieldID being assigned to by the formula. This is used to check for
	// circular references in the semantic analyzer. This can be left
	// as an empty string if only validating the equation for a new field. In other words,
	// if validating the formula for a new calculated field, there by definition can't be
	// any circular references to the field, since the field is new.
	resultFieldID string
}

type formulaCompileResults struct {
	preprocessedFormula string
	jsonEncodedEqn      string
	compiledFormula     *EquationNode
}

func compileAndEncodeFormula(params formulaCompileParams) (*formulaCompileResults, error) {

	// Run the prepocessor on formula text to replace any field references with their
	// permanent field IDs.
	preprocessedFormulaText, preprocessErr := preprocessCalcFieldFormula(params)
	if preprocessErr != nil {
		return nil, preprocessErr
	}

	log.Printf("compileAndEncodeFormula: preprocessor succeeded: %v", preprocessedFormulaText)

	compiledFormulaEqn, compileErr := compileFormula(preprocessedFormulaText)
	if compileErr != nil {
		return nil, compileErr
	} else if compiledFormulaEqn == nil {
		return nil, fmt.Errorf("Unexpected formula compile err: formula compiler returned nil compile result")
	}

	semanticAnalysisResults, semAnalysisErr := analyzeSemantics(params, compiledFormulaEqn)
	if semAnalysisErr != nil {
		return nil, fmt.Errorf("Unexpected formula compile err: semantic analyzer error = %v", semAnalysisErr)
	}
	if semanticAnalysisResults.hasErrors() {
		// TODO - The compiler needs to support returning multiple errors. The semantic analyzer already support
		// multiple errors, so we concatenate them until multiple errors can be returned.
		errMsgs := strings.Join(semanticAnalysisResults.analyzeErrors, " -- ")
		return nil, fmt.Errorf("Semantic analyzer error(s) = %v", errMsgs)
	}

	jsonEncodeEqn, encodeErr := generic.EncodeJSONString(compiledFormulaEqn)
	if encodeErr != nil {
		return nil, encodeErr
	}

	log.Printf("compileAndEncodeFormula: parsing succeeded: %v", jsonEncodeEqn)

	compileResults := formulaCompileResults{
		jsonEncodedEqn:      jsonEncodeEqn,
		preprocessedFormula: preprocessedFormulaText,
		compiledFormula:     compiledFormulaEqn}
	log.Printf("compileAndEncodeFormula:Formula compilation succeeded: %+v", compileResults)

	return &compileResults, nil

}

type ValidateFormulaParams struct {
	FieldID     string `json:fieldID`
	FormulaText string `json:formulaText`
}

type ValidationResponse struct {
	IsValidFormula bool   `json:"isValidFormula"`
	ErrorMsg       string `json:"errorMsg"`
}

func validateFormulaText(appEngContext appengine.Context, validationParams ValidateFormulaParams) *ValidationResponse {

	parentTableID, getParentErr := datastoreWrapper.GetParentID(validationParams.FieldID, table.TableEntityKind)
	if getParentErr != nil {
		errWithPrefixMsg := fmt.Errorf("validateFormulaText: Unable to get parent table for field: field id =%v, error=%v ",
			validationParams.FieldID, getParentErr)
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: errWithPrefixMsg.Error()}
	}

	fieldRef, getFieldErr := field.GetFieldRef(appEngContext, validationParams.FieldID)
	if getFieldErr != nil {
		errMsg := fmt.Sprintf("validateFormulaText: Unable to get  retrieve field: error=%v ", getFieldErr)
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: errMsg}
	} else {
		if !fieldRef.FieldInfo.IsCalcField {
			errorMsg := fmt.Sprintf("Formulas only work with calculated fields, got a regular field: %v",
				fieldRef.FieldInfo.Name)
			return &ValidationResponse{IsValidFormula: false, ErrorMsg: errorMsg}
		}
	}

	compileParams := formulaCompileParams{
		appEngContext:      appEngContext,
		formulaText:        validationParams.FormulaText,
		parentTableID:      parentTableID,
		expectedResultType: fieldRef.FieldInfo.Type,
		resultFieldID:      fieldRef.FieldID}

	_, compileErr := compileAndEncodeFormula(compileParams)
	if compileErr != nil {
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: compileErr.Error()}
	} else {
		return &ValidationResponse{IsValidFormula: true, ErrorMsg: ""}
	}

}
