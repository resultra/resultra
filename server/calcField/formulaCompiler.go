package calcField

import (
	"fmt"
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
