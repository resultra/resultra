package datamodel

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var regexpLeadingWhite = regexp.MustCompile("^[[:space:]]+")

type TokenDef struct {
	Regexp *regexp.Regexp
	ID     string
}

var tokenIdent = TokenDef{regexp.MustCompile("^[[:alpha:]][[:word:]]+"), "ident"}
var tokenLParen = TokenDef{regexp.MustCompile("^\\("), "lparen"}
var tokenRParen = TokenDef{regexp.MustCompile("^\\)"), "rparen"}
var tokenComma = TokenDef{regexp.MustCompile("^\\,"), "comma"}
var tokenBool = TokenDef{regexp.MustCompile("^(true)|(false)|(TRUE)|(FALSE)"), "bool"}
var tokenNumber = TokenDef{regexp.MustCompile("^[-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?"), "number"}
var tokenText = TokenDef{regexp.MustCompile(`"(?:[^"\\]|\\.)*"`), "text"}

var tokenDefs = []TokenDef{
	tokenBool,
	tokenIdent,
	tokenLParen,
	tokenRParen,
	tokenComma,
	tokenNumber,
	tokenText,
}

type TokenMatch struct {
	TokenID    string
	matchedStr string
}

type TokenMatchSequence []TokenMatch

func (matchSeq TokenMatchSequence) tokenIDs() []string {
	matchIDs := []string{}
	for _, match := range matchSeq {
		matchIDs = append(matchIDs, match.TokenID)
	}
	return matchIDs
}

func matchToken(inputStr string, tokenRegexp *regexp.Regexp, tokenName string) (*TokenMatch, string, bool) {
	matchIndices := tokenRegexp.FindStringIndex(inputStr)
	if matchIndices != nil {
		remaining := inputStr[matchIndices[1]:len(inputStr)]
		matchStr := inputStr[matchIndices[0]:matchIndices[1]]
		return &TokenMatch{tokenName, matchStr}, remaining, true
	} else {
		return nil, inputStr, false
	}
}

func matchNextToken(inputStr string) (*TokenMatch, string, error) {
	for _, tokenDef := range tokenDefs {
		if nextToken, remaining, foundTok := matchToken(inputStr, tokenDef.Regexp, tokenDef.ID); foundTok == true {
			return nextToken, remaining, nil
		}
	}
	return nil, "", fmt.Errorf("No matching tokens: remaining equation text = %v", inputStr)
}

func skipLeadingWhite(inputStr string) string {
	matchIndices := regexpLeadingWhite.FindStringIndex(inputStr)
	if matchIndices != nil {
		remaining := inputStr[matchIndices[1]:len(inputStr)]
		return remaining
	} else {
		return inputStr
	}

}

func tokenizeInput(inputStr string) (TokenMatchSequence, error) {

	var nextToken *TokenMatch
	var remaining string
	var tokErr error

	matches := TokenMatchSequence{}
	remaining = skipLeadingWhite(inputStr)

	for len(remaining) > 0 {
		nextToken, remaining, tokErr = matchNextToken(remaining)
		if tokErr != nil {
			// error handling here
			log.Printf("no token found, aborting: remaining = '%v'", remaining)
			return nil, fmt.Errorf("Error tokening input: %v", tokErr)
		} else {
			matches = append(matches, *nextToken)
			log.Printf("found token: %v '%v' , remaining = '%v'", nextToken.TokenID, nextToken.matchedStr, remaining)
		}
		remaining = skipLeadingWhite(remaining)
	}
	return matches, nil

}

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
			return funcEqnNode(funcName, eqnArgs), currIndex + 1
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
				return fieldRefEqnNode(fieldID), startIndex + 1
			}

		case tokenBool.ID:
			return nil, startIndex // not supported yet - TODO: implement this

		case tokenNumber.ID:
			if numberVal, numErr := strconv.ParseFloat(nextTok.matchedStr, 64); numErr != nil {
				return nil, startIndex
			} else {
				return numberEqnNode(numberVal), startIndex + 1
			}

		case tokenText.ID:
			// TODO - strip quotes from the string before creating the equation node
			return textEqnNode(nextTok.matchedStr), startIndex + 1
		default:
			return nil, startIndex // no match
		}

	}
	return nil, startIndex
}

func parseCalcFieldEqn(inputStr string) (*EquationNode, error) {

	currTokIndex := 0

	tokens, tokErr := tokenizeInput(inputStr)
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
