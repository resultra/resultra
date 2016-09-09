package calcField

import (
	"fmt"
	"strings"
)

type IdentReplacementMap map[string]string

func skipWhiteTokens(startIndex int, matchSeq TokenMatchSequence) int {
	if startIndex >= len(matchSeq) {
		return startIndex
	}

	currIndex := startIndex
	for matchSeq[currIndex].TokenID == tokenWhiteSpace.ID && currIndex < len(matchSeq) {
		currIndex++
	}
	return currIndex
}

func replaceNextIdent(startIndex int, matchSeq TokenMatchSequence,
	identReplMap IdentReplacementMap, prefixTokID int) (int, error) {

	if startIndex >= len(matchSeq) {
		return startIndex, nil
	}

	for currTokIndex := startIndex; currTokIndex < len(matchSeq); currTokIndex++ {
		if matchSeq[currTokIndex].TokenID == prefixTokID {
			nextTokIndex := currTokIndex + 1
			nextTokIndex = skipWhiteTokens(nextTokIndex, matchSeq) // skip any leading whitespace
			if nextTokIndex >= len(matchSeq) {
				return nextTokIndex, nil
			} else {
				if matchSeq[nextTokIndex].TokenID == tokenIdent.ID {

					replText, foundRepl := identReplMap[matchSeq[nextTokIndex].matchedStr]
					if !foundRepl {
						return nextTokIndex, fmt.Errorf(
							"Unexpected error prepocessing formula: Can't find substitution for '%v', valid matches = %+v (len=%v)",
							matchSeq[nextTokIndex].matchedStr, identReplMap, len(identReplMap))
					} else {
						matchSeq[nextTokIndex].matchedStr = replText
						return nextTokIndex + 1, nil // skip over closing bracket
					}
				} else {
					return nextTokIndex, fmt.Errorf("Unexpected input, expecting an identifier, got '%v'",
						matchSeq[nextTokIndex].matchedStr)
				}
			}
		}
	}

	// No matches found, return the final index
	return len(matchSeq), nil

}

func preprocessTokenSeq(tokSeq TokenMatchSequence, identReplMap IdentReplacementMap, prefixTokID int) error {
	currTokIndex := 0
	for currTokIndex < len(tokSeq) {
		tokIndexAfterRepl, err := replaceNextIdent(currTokIndex, tokSeq, identReplMap, prefixTokID)
		if err != nil {
			return fmt.Errorf("Error preprocessing formula input: %v", err)
		} else {
			currTokIndex = tokIndexAfterRepl
		}
	}
	return nil

}

// preprocessFormula takes equation/formula source and replaces field references inside [] with
// the corresponding value in fieldNameReplMap. The user types field names using a "field reference name".
// However, the preprocessed source is stored using the permanent field ID; i.e. the user can easily
// change the field reference name, in which case we can't have the stored formulas then have invalid
// references. After retrieving pre-processed formula source from the datastore, the reverse mapping
// can be done using this same function to allow the user to edit the source again. Or, if the formula
// is being retrieve for calculations, the field ID in the pre-processed source can be used directly.
func preprocessFormulaInput(inputStr string, fieldNameReplMap IdentReplacementMap) (string, error) {

	// Tokenize the input, keeping whitespace and comments as tokens
	tokenizeWhiteOrComment := true
	matchSeq, err := tokenizeInput(inputStr, tokenizeWhiteOrComment)
	if err != nil {
		return "", fmt.Errorf("Error preprocessing formula input: %v", err)
	}

	if err := preprocessTokenSeq(matchSeq, fieldNameReplMap, tokenLBracket.ID); err != nil {
		return "", fmt.Errorf("Error preprocessing formula input: %v", err)
	}

	// Re-assemble the pre-processed tokens back into a single string
	preprocessedTokText := []string{}
	for tokIndex := 0; tokIndex < len(matchSeq); tokIndex++ {
		preprocessedTokText = append(preprocessedTokText, matchSeq[tokIndex].matchedStr)
	}
	return strings.Join(preprocessedTokText, ""), nil

}
