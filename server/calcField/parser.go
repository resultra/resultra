package calcField

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"strconv"
)

func parseArgs(tokenSeq TokenMatchSequence, startIndex int) ([]EquationNode, int, error) {

	args := []EquationNode{}

	currIndex := startIndex

	if (currIndex < len(tokenSeq)) && (tokenSeq[currIndex].TokenID == tokenRParen.ID) {
		// zero-length argument list
		return args, startIndex, nil
	}

	firstArg, indexAfterFirstArg, parseErr := parseEqnRoot(tokenSeq, startIndex)
	if firstArg == nil {
		return nil, startIndex, fmt.Errorf("Error parsing function: Expecting function argument, got %v",
			tokenSeq[currIndex].TokenID)
	} else if parseErr != nil {
		return nil, startIndex, parseErr
	} else {
		args = append(args, *firstArg)
		currIndex = indexAfterFirstArg
	}

	for (currIndex < len(tokenSeq)) && (tokenSeq[currIndex].TokenID == tokenComma.ID) {

		currIndex = currIndex + 1 // skip the comma

		if currIndex < len(tokenSeq) {
			nextArg, indexAfterNextArg, parseErr := parseEqnRoot(tokenSeq, currIndex)
			if nextArg == nil {
				return nil, startIndex, fmt.Errorf("Error parsing function, expecting argument after comma, got %v",
					tokenSeq[currIndex].TokenID) // TODO - Should generate an error for missing argument
			} else if parseErr != nil {
				return nil, startIndex, parseErr
			} else {
				args = append(args, *nextArg)
				currIndex = indexAfterNextArg
			}
		} else {
			return nil, startIndex, fmt.Errorf("Error parsing function, expecting argument after comma, got end of input")
		}
	} // for each argument

	return args, currIndex, nil
}

func parseFuncEqn(tokenSeq TokenMatchSequence, startIndex int) (*EquationNode, int, error) {
	if startIndex >= len(tokenSeq) {
		return nil, startIndex, fmt.Errorf("Unexpected end of input parsing equation: expecting function name")
	} else {
		currIndex := startIndex
		nextTok := tokenSeq[currIndex]

		if nextTok.TokenID != tokenIdent.ID {
			return nil, startIndex,
				fmt.Errorf("Unexpected end of input parsing equation: expecting function name, got %v", nextTok.matchedStr)
		}
		funcName := nextTok.matchedStr

		currIndex++
		if currIndex >= len(tokenSeq) {
			return nil, startIndex, fmt.Errorf(
				"Unexpected end of input parsing equation: expecting left parentheses, got end of input")
		}
		nextTok = tokenSeq[currIndex]
		if nextTok.TokenID != tokenLParen.ID {
			return nil, startIndex, fmt.Errorf(
				"Unexpected input parsing equation: expecting left parentheses, got %v", nextTok.matchedStr)
		}
		currIndex++ // skip over lparen

		eqnArgs, indexAfterArgs, argsErr := parseArgs(tokenSeq, currIndex)
		if argsErr != nil {
			return nil, startIndex, argsErr
		}
		if eqnArgs != nil {
			currIndex = indexAfterArgs // skip over arguments
		} else {
			return nil, startIndex,
				fmt.Errorf("Unexpected error parsing function arguments: unexpected empty argument list")
		}

		if currIndex >= len(tokenSeq) {
			// expecting closing paren
			return nil, startIndex, fmt.Errorf("Unexpected end of input: expecting right parenthese")
		}
		nextTok = tokenSeq[currIndex]
		if nextTok.TokenID != tokenRParen.ID {
			return nil, startIndex, fmt.Errorf(
				"Unexpected input parsing equation: expecting right parentheses, got %v", nextTok.matchedStr)
		}

		return FuncEqnNode(funcName, eqnArgs), currIndex + 1, nil

	}
}

func parseFieldRef(tokenSeq TokenMatchSequence, startIndex int) (*EquationNode, int, error) {
	currIndex := startIndex + 1
	if currIndex >= len(tokenSeq) {
		return nil, startIndex,
			fmt.Errorf("Unexpected error parsing field reference: got end of input, expecing a field reference name")
	}

	fieldRefTok := tokenSeq[currIndex]

	if fieldRefTok.TokenID != tokenIdent.ID {
		return nil, startIndex,
			fmt.Errorf("Unexpected error parsing field reference: expecting a field reference, got %v", fieldRefTok.matchedStr)
	}
	// When actually parsing the equation input, the input will have been pre-processed to turn field references
	// defined by end-users into unique & permanent field IDs.
	fieldID := fieldRefTok.matchedStr

	currIndex++
	if currIndex >= len(tokenSeq) {
		return nil, startIndex,
			fmt.Errorf("Unexpected error parsing field reference: got end of input, expecing a closing bracket")
	}
	if tokenSeq[currIndex].TokenID != tokenRBracket.ID {
		return nil, startIndex,
			fmt.Errorf("Unexpected error parsing field reference: got %v, expecing a closing bracket",
				tokenSeq[currIndex].matchedStr)
	}

	indexAfterFieldRef := currIndex + 1
	return FieldRefEqnNode(fieldID), indexAfterFieldRef, nil
}

// Parse an equation root - i.e. a function, or literal number/ext/bool
// Should not see parentheses or commas (delimiters) here
func parseEqnRoot(tokenSeq TokenMatchSequence, startIndex int) (*EquationNode, int, error) {
	if startIndex >= len(tokenSeq) {
		return nil, startIndex, fmt.Errorf(
			"Unexpected end of input parsing calculated field equation: " +
				" expecting a function or function call, field reference or literal value, got end of input")
	} else {

		nextTok := tokenSeq[startIndex]

		switch nextTok.TokenID {
		case tokenIdent.ID:
			return parseFuncEqn(tokenSeq, startIndex)
		case tokenLBracket.ID:
			return parseFieldRef(tokenSeq, startIndex)
		case tokenBool.ID:
			return nil, startIndex, fmt.Errorf("Boolean literal not yet supported") // not supported yet - TODO: implement this

		case tokenNumber.ID:
			if numberVal, numErr := strconv.ParseFloat(nextTok.matchedStr, 64); numErr != nil {
				return nil, startIndex, numErr
			} else {
				return NumberEqnNode(numberVal), startIndex + 1, nil
			}

		case tokenText.ID:
			// TODO - strip quotes from the string before creating the equation node
			return TextEqnNode(nextTok.matchedStr), startIndex + 1, nil
		default:
			return nil, startIndex, fmt.Errorf(
				"Unexpected input parsing calculated field equation: "+
					" expecting a function or function call, field reference or literal value, got %v", nextTok.matchedStr)
		}

	}
}

func parseCalcFieldEqn(inputStr string) (*EquationNode, error) {

	currTokIndex := 0

	tokenizeWhiteComment := false
	tokens, tokErr := tokenizeInput(inputStr, tokenizeWhiteComment)
	if tokErr != nil {
		return nil, tokErr
	}

	if currTokIndex >= len(tokens) {
		// TBD - If the input is initially empty, should it return an "undefined" result?
		return nil, fmt.Errorf("parseCalcFieldEqn: Empty input (no tokens found)")
	}

	eqnRoot, endIndex, parseErr := parseEqnRoot(tokens, currTokIndex)
	if parseErr != nil {
		return nil, parseErr
	} else if eqnRoot == nil {
		return nil, fmt.Errorf("parseCalcFieldEqn: parse error: expecting a function or literal, got %v", tokens[currTokIndex].TokenID)
	}

	if endIndex < len(tokens) {
		return nil, fmt.Errorf("parseCalcFieldEqn: parse error: expecting end of equation, got %v", tokens[currTokIndex].TokenID)
	}

	return eqnRoot, nil

}

func validateCalcFieldEqnText(appEngContext appengine.Context, eqnStr string) error {

	_, fieldRefErr := field.GetFieldRefIDIndex(appEngContext)
	if fieldRefErr != nil {
		return fmt.Errorf("ValidateCalcFieldEqn: Error getting field references to validate calculated field equation: %v", fieldRefErr)
	}

	log.Printf("ValidateCalcFieldEqn: validating calc field eqn: %v", eqnStr)
	_, parseErr := parseCalcFieldEqn(eqnStr)
	if parseErr != nil {
		return fmt.Errorf("ValidateCalcFieldEqn: Error validating calculated field equation: parse error = %v", parseErr)
	}

	return nil
}
