package calcField

import (
	"fmt"
	"strings"
)

type FieldNameReplacementMap map[string]string

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

func replaceNextFieldName(startIndex int, matchSeq TokenMatchSequence,
	fieldNameReplMap FieldNameReplacementMap) (int, error) {

	if startIndex >= len(matchSeq) {
		return startIndex, nil
	}

	for currTokIndex := startIndex; currTokIndex < len(matchSeq); currTokIndex++ {
		if matchSeq[currTokIndex].TokenID == tokenLBracket.ID {
			nextTokIndex := currTokIndex + 1
			nextTokIndex = skipWhiteTokens(nextTokIndex, matchSeq) // skip any leading whitespace
			if nextTokIndex >= len(matchSeq) {
				return nextTokIndex, nil
			} else {
				if matchSeq[nextTokIndex].TokenID == tokenIdent.ID {

					replText, foundRepl := fieldNameReplMap[matchSeq[nextTokIndex].matchedStr]
					if !foundRepl {
						return nextTokIndex, fmt.Errorf(
							"Unexpected error prepocessing field name: Can't find substitution for '%v'",
							matchSeq[nextTokIndex].matchedStr)
					} else {
						matchSeq[nextTokIndex].matchedStr = replText
						return nextTokIndex + 1, nil // skip over closing bracket
					}
				} else {
					return nextTokIndex, fmt.Errorf("Unexpected input, expecting a field reference inside [], got '%v'",
						matchSeq[nextTokIndex].matchedStr)
				}
			}
		}
	}

	// No matches found, return the final index
	return len(matchSeq), nil

}

// preprocessFormula takes equation/formula source and replaces field references inside [] with
// the corresponding value in fieldNameReplMap. The user types field names using a "field reference name".
// However, the preprocessed source is stored using the permanent field ID; i.e. the user can easily
// change the field reference name, in which case we can't have the stored formulas then have invalid
// references. After retrieving pre-processed formula source from the datastore, the reverse mapping
// can be done using this same function to allow the user to edit the source again. Or, if the formula
// is being retrieve for calculations, the field ID in the pre-processed source can be used directly.
func preprocessFormulaInput(inputStr string, fieldNameReplMap FieldNameReplacementMap) (string, error) {

	// Tokenize the input, keeping whitespace and comments as tokens
	tokenizeWhiteOrComment := true
	matchSeq, err := tokenizeInput(inputStr, tokenizeWhiteOrComment)
	if err != nil {
		return "", fmt.Errorf("Error preprocessing formula input: %v", err)
	}

	currTokIndex := 0
	for currTokIndex < len(matchSeq) {
		tokIndexAfterRepl, err := replaceNextFieldName(currTokIndex, matchSeq, fieldNameReplMap)
		if err != nil {
			return "", fmt.Errorf("Error preprocessing formula input: %v", err)
		} else {
			currTokIndex = tokIndexAfterRepl
		}
	}

	preprocessedTokText := []string{}
	for tokIndex := 0; tokIndex < len(matchSeq); tokIndex++ {
		preprocessedTokText = append(preprocessedTokText, matchSeq[tokIndex].matchedStr)
	}
	return strings.Join(preprocessedTokText, ""), nil

}
