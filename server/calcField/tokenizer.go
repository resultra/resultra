package calcField

import (
	"fmt"
	"log"
	"regexp"
)

var regexpLeadingWhite = regexp.MustCompile("^[[:space:]]+")

type TokenDef struct {
	Regexp *regexp.Regexp
	ID     string
}

var tokenWhiteSpace = TokenDef{regexpLeadingWhite, "wspace"}
var tokenComment = TokenDef{regexp.MustCompile("^//.*"), "comment"}

var tokenIdent = TokenDef{regexp.MustCompile("^[[:alpha:]][[:word:]]+"), "ident"}
var tokenAssign = TokenDef{regexp.MustCompile("^="), "assign"}
var tokenEqual = TokenDef{regexp.MustCompile("^=="), "equal"}
var tokenLParen = TokenDef{regexp.MustCompile("^\\("), "lparen"}
var tokenRParen = TokenDef{regexp.MustCompile("^\\)"), "rparen"}
var tokenLBracket = TokenDef{regexp.MustCompile("^\\["), "lbrack"}
var tokenRBracket = TokenDef{regexp.MustCompile("^\\]"), "rbrack"}
var tokenComma = TokenDef{regexp.MustCompile("^\\,"), "comma"}
var tokenBool = TokenDef{regexp.MustCompile("^(true)|(false)|(TRUE)|(FALSE)"), "bool"}
var tokenNumber = TokenDef{regexp.MustCompile("^[-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?"), "number"}
var tokenText = TokenDef{regexp.MustCompile(`"(?:[^"\\]|\\.)*"`), "text"}

var tokenDefs = []TokenDef{
	tokenWhiteSpace,
	tokenComment,
	tokenBool,
	tokenIdent,
	tokenEqual,
	tokenAssign,
	tokenLParen,
	tokenRParen,
	tokenLBracket,
	tokenRBracket,
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

func tokenIsWhitespaceOrComment(match TokenMatch) bool {
	if match.TokenID == tokenWhiteSpace.ID {
		return true
	} else if match.TokenID == tokenComment.ID {
		return true
	} else {
		return false
	}
}

// tokenizeInput supports 2 modes of tokenizing. If tokenizeWhiteAndComment is set to true, then
// comments and whitespace are treated like regular tokens and included in the returned
// TokenMatchSequence. Otherwise, whitespace and comments are stripped. These 2 different modes
// are needed so the intput can either be parsed as an "executable" equation or pre-processed
// to replace field references with actual unique field IDs.
func tokenizeInput(inputStr string, tokenizeWhiteAndComment bool) (TokenMatchSequence, error) {

	var nextToken *TokenMatch
	var remaining string
	var tokErr error

	matches := TokenMatchSequence{}
	remaining = inputStr

	for len(remaining) > 0 {
		nextToken, remaining, tokErr = matchNextToken(remaining)
		if tokErr != nil {
			// error handling here
			log.Printf("no token found, aborting: remaining = '%v'", remaining)
			return nil, fmt.Errorf("Error tokening input: %v", tokErr)
		} else {
			if tokenIsWhitespaceOrComment(*nextToken) {
				// Only add whitespace and comments to token list if tokenizeWhiteAndComment is true
				if tokenizeWhiteAndComment {
					matches = append(matches, *nextToken)
				}
			} else { // Not whitespace or comment => always add it to token list
				matches = append(matches, *nextToken)
			}
			log.Printf("found token: %v '%v' , remaining = '%v'", nextToken.TokenID, nextToken.matchedStr, remaining)
		}
	}
	log.Printf("tokenizeInput: done tokenizing input: found  %v tokens", len(matches))

	return matches, nil

}
