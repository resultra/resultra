package calcField

import (
	"fmt"
	"regexp"
)

var regexpLeadingWhite = regexp.MustCompile("^[[:space:]]+")

type TokenDef struct {
	Regexp *regexp.Regexp
	ID     int
}

var tokenWhiteSpace = TokenDef{regexpLeadingWhite, TOK_WHITE}
var tokenComment = TokenDef{regexp.MustCompile("^//.*"), TOK_COMMENT}

// The parser currently expects both UUIDs and identifyers to return TOK_IDENT. More work is needed to
// distinguish between both cases.
var tokenIdent = TokenDef{regexp.MustCompile("^[[:alpha:]][[:word:]\\-\\_]*"), TOK_IDENT}
var tokenUUID = TokenDef{regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"), TOK_IDENT}
var tokenAssign = TokenDef{regexp.MustCompile("^="), TOK_ASSIGN}
var tokenEqual = TokenDef{regexp.MustCompile("^=="), TOK_EQUAL}
var tokenPlus = TokenDef{regexp.MustCompile("^\\+"), TOK_PLUS}
var tokenMinus = TokenDef{regexp.MustCompile("^\\-"), TOK_MINUS}
var tokenDivide = TokenDef{regexp.MustCompile("^\\/"), TOK_DIVIDE}
var tokenTimes = TokenDef{regexp.MustCompile("^\\*"), TOK_TIMES}
var tokenLParen = TokenDef{regexp.MustCompile("^\\("), TOK_LPAREN}
var tokenRParen = TokenDef{regexp.MustCompile("^\\)"), TOK_RPAREN}
var tokenLBracket = TokenDef{regexp.MustCompile("^\\["), TOK_LBRACKET}
var tokenRBracket = TokenDef{regexp.MustCompile("^\\]"), TOK_RBRACKET}

var tokenDoubleLBracket = TokenDef{regexp.MustCompile("^\\[\\["), TOK_DOUBLE_LBRACKET}
var tokenDoubleRBracket = TokenDef{regexp.MustCompile("^\\]\\]"), TOK_DOUBLE_RBRACKET}
var tokenComma = TokenDef{regexp.MustCompile("^\\,"), TOK_COMMA}
var tokenBool = TokenDef{regexp.MustCompile("^(true)|(false)|(TRUE)|(FALSE)"), TOK_BOOL}
var tokenNumber = TokenDef{regexp.MustCompile("^[-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?"), TOK_NUMBER}
var tokenText = TokenDef{regexp.MustCompile(`"(?:[^"\\]|\\.)*"`), TOK_TEXT}

var tokenDefs = []TokenDef{
	tokenWhiteSpace,
	tokenComment,
	tokenBool,
	tokenUUID,
	tokenIdent,
	tokenEqual,
	tokenPlus,
	tokenMinus,
	tokenTimes,
	tokenDivide,
	tokenAssign,
	tokenLParen,
	tokenRParen,
	tokenDoubleLBracket,
	tokenDoubleRBracket,
	tokenLBracket,
	tokenRBracket,
	tokenComma,
	tokenNumber,
	tokenText,
}

type TokenMatch struct {
	TokenID    int
	matchedStr string
}

type TokenMatchSequence []TokenMatch

func (matchSeq TokenMatchSequence) tokenIDs() []int {
	matchIDs := []int{}
	for _, match := range matchSeq {
		matchIDs = append(matchIDs, match.TokenID)
	}
	return matchIDs
}

func matchToken(inputStr string, tokenRegexp *regexp.Regexp, tokenID int) (*TokenMatch, string, bool) {
	matchIndices := tokenRegexp.FindStringIndex(inputStr)
	if matchIndices != nil {
		remaining := inputStr[matchIndices[1]:len(inputStr)]
		matchStr := inputStr[matchIndices[0]:matchIndices[1]]
		return &TokenMatch{tokenID, matchStr}, remaining, true
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
			return nil, fmt.Errorf("Error tokening input: overall input = %v, token error = %v", inputStr, tokErr)
		} else {
			if tokenIsWhitespaceOrComment(*nextToken) {
				// Only add whitespace and comments to token list if tokenizeWhiteAndComment is true
				if tokenizeWhiteAndComment {
					matches = append(matches, *nextToken)
				}
			} else { // Not whitespace or comment => always add it to token list
				matches = append(matches, *nextToken)
			}
		}
	}

	return matches, nil

}
