// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package calcField

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/global"
	"log"
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

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(compileParams.trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: compileParams.databaseID})
	if indexErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: %v", indexErr)
	}

	fieldRefFieldIDMap := IdentReplacementMap{}
	for fieldRefName, currField := range fieldRefIndex.FieldsByRefName {
		fieldRefFieldIDMap[fieldRefName] = currField.FieldID
	}

	globals, getGlobalsErr := global.GetGlobals(compileParams.trackerDBHandle, compileParams.databaseID)
	if getGlobalsErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: Unable to retrieve globals: databaseID=%v, error=%v ",
			compileParams.databaseID, getGlobalsErr)
	}
	globalRefGlobalIDMap := IdentReplacementMap{}
	for _, currGlobal := range globals {
		globalRefGlobalIDMap[currGlobal.RefName] = currGlobal.GlobalID
	}

	preprocessOutput, preprocessErr := preprocessFormulaInput(compileParams.formulaText, fieldRefFieldIDMap, globalRefGlobalIDMap)
	if preprocessErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: Error preprocessing formula: error=%v ", preprocessErr)
	}

	return preprocessOutput, nil

}

// ClonePreprocessedFormula is used to duplicate a formula when saving an existing database as a template or when cloning from a
// template to a new database.
func ClonePreprocessedFormula(trackerDBHandle *sql.DB, srcDatabaseID string,
	remappedIDs uniqueID.UniqueIDRemapper, preProcessedFormula string) (string, error) {

	getFieldParams := field.GetFieldListParams{ParentDatabaseID: srcDatabaseID}
	fields, err := field.GetAllFields(trackerDBHandle, getFieldParams)
	if err != nil {
		return "", fmt.Errorf("cloneFields: %v", err)
	}

	fieldIDRemappedIDMap := IdentReplacementMap{}
	for _, currField := range fields {
		remappedID, err := remappedIDs.GetExistingRemappedID(currField.FieldID)
		if err != nil {
			return "", fmt.Errorf("ClonePreprocessedFormula: Can't find mapped ID for field id = %v", currField.FieldID)
		}
		fieldIDRemappedIDMap[currField.FieldID] = remappedID
	}

	globals, err := global.GetGlobals(trackerDBHandle, srcDatabaseID)
	if err != nil {
		return "", fmt.Errorf("ClonePreprocessedFormula: Unable to retrieve globals: databaseID=%v, error=%v ",
			srcDatabaseID, err)
	}
	globalIDRemappedIDMap := IdentReplacementMap{}
	for _, currGlobal := range globals {
		remappedID, err := remappedIDs.GetExistingRemappedID(currGlobal.GlobalID)
		if err != nil {
			return "", fmt.Errorf("ClonePreprocessedFormula: Can't find mapped ID for global id = %v", currGlobal.GlobalID)
		}
		globalIDRemappedIDMap[currGlobal.GlobalID] = remappedID
	}

	remappedFormulaOutput, err := preprocessFormulaInput(preProcessedFormula, fieldIDRemappedIDMap, globalIDRemappedIDMap)
	if err != nil {
		return "", fmt.Errorf("ClonePreprocessedFormula: Error mapping ids to clone: error=%v ", err)
	}

	return remappedFormulaOutput, nil

}

// reverseProcessCalcFieldFormula does the opposite of preprocessCalcFieldFormula. It takes a preprocessed formula
// with embedded permanent field IDs and replaces them with the most up to date field reference names. This
// allows the user to change the reference name after saving a formula. Since the formula is stored with a
// field's permanent & unique ID, when the formula's text is retrieved again it will have the most up to
// date reference name.
func reverseProcessCalcFieldFormula(compileParams formulaCompileParams) (string, error) {

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(compileParams.trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: compileParams.databaseID})
	if indexErr != nil {
		return "", fmt.Errorf("preprocessCalcFieldFormula: %v", indexErr)
	}

	fieldIDFieldRefMap := IdentReplacementMap{}
	for _, currField := range fieldRefIndex.FieldsByRefName {
		fieldIDFieldRefMap[currField.FieldID] = currField.RefName
	}

	globals, getGlobalsErr := global.GetGlobals(compileParams.trackerDBHandle, compileParams.databaseID)
	if getGlobalsErr != nil {
		return "", fmt.Errorf("reverseProcessCalcFieldFormula: Unable to retrieve globals: databaseID=%v, error=%v ",
			compileParams.databaseID, getGlobalsErr)
	}
	globalIDGlobalRefMap := IdentReplacementMap{}
	for _, currGlobal := range globals {
		globalIDGlobalRefMap[currGlobal.GlobalID] = currGlobal.RefName
	}

	log.Printf("reverseProcessCalcFieldFormula: Starting reverse processing of pre-processed formula text: %v",
		compileParams.formulaText)

	reverseProcessOutput, err := preprocessFormulaInput(compileParams.formulaText, fieldIDFieldRefMap, globalIDGlobalRefMap)
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

type formulaCompileParams struct {
	trackerDBHandle    *sql.DB
	formulaText        string
	databaseID         string
	expectedResultType string

	// This is the fieldID being assigned to by the formula. This is used to check for
	// circular references in the semantic analyzer. This can be left
	// as an empty string if only validating the equation for a new field. In other words,
	// if validating the formula for a new calculated field, there by definition can't be
	// any circular references to the field, since the field is new.
	resultFieldID string
}

func assembleCalcFieldCompileParams(trackerDBHandle *sql.DB, fieldID string) (*formulaCompileParams, error) {

	calcField, getFieldErr := field.GetField(trackerDBHandle, fieldID)
	if getFieldErr != nil {
		return nil, fmt.Errorf("assembleCalcFieldCompileParams: Unable to get calculated field field: field id =%v, error=%v ",
			fieldID, getFieldErr)
	}
	if !calcField.IsCalcField {
		return nil, fmt.Errorf("assembleCalcFieldCompileParams: Formulas only work with calculated fields, got a regular field: %v",
			calcField.Name)
	}

	compileParams := formulaCompileParams{
		trackerDBHandle:    trackerDBHandle,
		formulaText:        calcField.PreprocessedFormulaText,
		databaseID:         calcField.ParentDatabaseID,
		expectedResultType: calcField.Type,
		resultFieldID:      fieldID}

	return &compileParams, nil
}

func getRawFormulaText(trackerDBHandle *sql.DB, params GetRawFormulaParams) (*GetRawFormulaResult, error) {

	compileParams, paramErr := assembleCalcFieldCompileParams(trackerDBHandle, params.FieldID)
	if paramErr != nil {
		return nil, fmt.Errorf("getRawFormulaText: Unable to retrieve compilation parameters: error=%v ", paramErr)
	}

	rawFormulaText, reverseProcessErr := reverseProcessCalcFieldFormula(*compileParams)
	if reverseProcessErr != nil {
		return nil, fmt.Errorf("getRawFormulaText: Unable to read calculated field field: field id =%v, error=%v ",
			params.FieldID, reverseProcessErr)
	}

	return &GetRawFormulaResult{FieldID: params.FieldID, RawFormulaText: rawFormulaText}, nil
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

func validateFormulaText(trackerDBHandle *sql.DB, validationParams ValidateFormulaParams) *ValidationResponse {

	compileParams, paramErr := assembleCalcFieldCompileParams(trackerDBHandle, validationParams.FieldID)
	if paramErr != nil {
		errMsg := fmt.Sprintf("validateFormulaText: Unable to get  retrieve field: error=%v ", paramErr)
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: errMsg}
	}
	// By default, assembleCalcFieldCompileParams will return the parameters populated with the preprocessed formula text.
	// However, in this case,we want to compile with the given formula text.
	compileParams.formulaText = validationParams.FormulaText

	_, compileErr := compileAndEncodeFormula(*compileParams)
	if compileErr != nil {
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: compileErr.Error()}
	} else {
		return &ValidationResponse{IsValidFormula: true, ErrorMsg: ""}
	}

}
