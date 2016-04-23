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

	firstArg, indexAfterFirstArg := parseEqnRoot(tokenSeq, startIndex)
	if firstArg == nil {
		return nil, startIndex, fmt.Errorf("Error parsing function: Expecting function argument, got %v",
			tokenSeq[currIndex].TokenID)
	} else {
		args = append(args, *firstArg)
		currIndex = indexAfterFirstArg
	}

	for (currIndex < len(tokenSeq)) && (tokenSeq[currIndex].TokenID == tokenComma.ID) {

		currIndex = currIndex + 1 // skip the comma

		if currIndex < len(tokenSeq) {
			nextArg, indexAfterNextArg := parseEqnRoot(tokenSeq, currIndex)
			if nextArg == nil {
				return nil, startIndex, fmt.Errorf("Error parsing function, expecting argument after comma, got %v",
					tokenSeq[currIndex].TokenID) // TODO - Should generate an error for missing argument
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

func parseFuncEqn(tokenSeq TokenMatchSequence, startIndex int) (*EquationNode, int) {
	if startIndex >= len(tokenSeq) {
		return nil, startIndex
	} else {
		currIndex := startIndex
		nextTok := tokenSeq[currIndex]

		if nextTok.TokenID != tokenIdent.ID {
			return nil, startIndex
		}
		funcName := nextTok.matchedStr

		currIndex++
		if currIndex >= len(tokenSeq) {
			return nil, startIndex
		}
		nextTok = tokenSeq[currIndex]
		if nextTok.TokenID != tokenLParen.ID {
			return nil, startIndex
		}
		currIndex++ // skip over lparen

		eqnArgs, indexAfterArgs, argsErr := parseArgs(tokenSeq, currIndex)
		if argsErr != nil {
			return nil, startIndex
		}
		if eqnArgs != nil {
			currIndex = indexAfterArgs // skip over arguments
		} else {
			return nil, startIndex
		}

		if currIndex >= len(tokenSeq) {
			// expecting closing paren
			return nil, startIndex
		}
		nextTok = tokenSeq[currIndex]
		if nextTok.TokenID != tokenRParen.ID {
			return nil, startIndex
		} else {
			return FuncEqnNode(funcName, eqnArgs), currIndex + 1
		}

	}
	return nil, startIndex
}

// Parse an equation root - i.e. a function, or literal number/ext/bool
// Should not see parentheses or commas (delimiters) here
func parseEqnRoot(tokenSeq TokenMatchSequence, startIndex int) (*EquationNode, int) {
	if startIndex >= len(tokenSeq) {
		return nil, startIndex
	} else {

		nextTok := tokenSeq[startIndex]

		switch nextTok.TokenID {
		case tokenIdent.ID:

			// Either a function or field reference
			if funcEqn, indexAfterFunc := parseFuncEqn(tokenSeq, startIndex); funcEqn != nil {
				return funcEqn, indexAfterFunc
			} else {
				// Check for field reference
				fieldRefName := nextTok.matchedStr
				fieldID := fieldRefName // TODO - resolve field ID instead of dummying up with reference name
				return FieldRefEqnNode(fieldID), startIndex + 1
			}

		case tokenBool.ID:
			return nil, startIndex // not supported yet - TODO: implement this

		case tokenNumber.ID:
			if numberVal, numErr := strconv.ParseFloat(nextTok.matchedStr, 64); numErr != nil {
				return nil, startIndex
			} else {
				return NumberEqnNode(numberVal), startIndex + 1
			}

		case tokenText.ID:
			// TODO - strip quotes from the string before creating the equation node
			return TextEqnNode(nextTok.matchedStr), startIndex + 1
		default:
			return nil, startIndex // no match
		}

	}
	return nil, startIndex
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

	eqnRoot, endIndex := parseEqnRoot(tokens, currTokIndex)
	if eqnRoot == nil {
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
