package calcField

import (
	"fmt"
	"log"
	"strconv"
)

const TOK_EOF = 0

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
type formulaLexerImpl struct {
	tokMatchSeq TokenMatchSequence
	currPos     int

	// rootEqnNode serves as local storage for the root node parsed by the
	// formula parser. See the comment in formulaParser.y for more details on
	// this work-around.
	rootEqnNode *EquationNode
}

func (lexer *formulaLexerImpl) Lex(lval *formulaSymType) int {

	if lexer.currPos >= len(lexer.tokMatchSeq) {
		return TOK_EOF
	} else {
		currTok := lexer.tokMatchSeq[lexer.currPos]
		lexer.currPos++
		switch currTok.TokenID {
		case TOK_NUMBER:
			if numberVal, numErr := strconv.ParseFloat(currTok.matchedStr, 64); numErr != nil {
				return TOK_EOF
			} else {
				lval.number = numberVal
				log.Printf("lexer: Number token: %v", lval.number)
				return TOK_NUMBER
			}
		default:

			return currTok.TokenID
		}
	}

}

func (lexer *formulaLexerImpl) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func newFormulaLexer(inputStr string) (*formulaLexerImpl, error) {
	tokenizeWhiteAndComment := false // skip whitespace and comments
	tokMatchSeq, err := tokenizeInput(inputStr, tokenizeWhiteAndComment)
	if err != nil {
		return nil, err
	}
	return &formulaLexerImpl{tokMatchSeq: tokMatchSeq, currPos: 0}, nil

}
