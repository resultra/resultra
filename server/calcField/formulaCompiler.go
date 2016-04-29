package calcField

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic"
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

type ValidateFormulaParams struct {
	FieldID     string `json:fieldID`
	FormulaText string `json:formulaText`
}

type ValidationResponse struct {
	IsValidFormula bool   `json:"isValidFormula"`
	ErrorMsg       string `json:"errorMsg"`
}

func validateFormulaText(appEngContext appengine.Context, validationParams ValidateFormulaParams) *ValidationResponse {

	if _, err := compileFormula(validationParams.FormulaText); err != nil {
		return &ValidationResponse{IsValidFormula: false, ErrorMsg: err.Error()}
	} else {
		return &ValidationResponse{IsValidFormula: true, ErrorMsg: ""}
	}
}

func compileAndEncodeFormula(formulaText string) (string, error) {

	// TODO - Need to run the prepocessor on formula text to replace any field references with their
	// permanent field IDs.

	compiledFormulaEqn, compileErr := compileFormula(formulaText)
	if compileErr != nil {
		return "", compileErr
	}

	jsonEncodeEqn, encodeErr := generic.EncodeJSONString(compiledFormulaEqn)
	if encodeErr != nil {
		return "", encodeErr
	}

	return jsonEncodeEqn, nil

}
