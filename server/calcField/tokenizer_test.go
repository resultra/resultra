package calcField

import (
	"strings"
	"testing"
)

func TestTokens(t *testing.T) {

	inputStr := `    HelloWorldFunc  ( arg1, arg3,true ,-32.43,22,"hello \" world" )   `
	if matchSeq, err := tokenizeInput(inputStr); err != nil {
		t.Fatal(err)
	} else {
		tokenIDs := strings.Join(matchSeq.tokenIDs(), ",")
		t.Logf("Match sequence: %v", tokenIDs)
		expectedTokenIDs := strings.Join([]string{
			tokenIdent.ID, tokenLParen.ID,
			tokenIdent.ID, tokenComma.ID,
			tokenIdent.ID, tokenComma.ID,
			tokenBool.ID, tokenComma.ID,
			tokenNumber.ID, tokenComma.ID,
			tokenNumber.ID, tokenComma.ID,
			tokenText.ID,
			tokenRParen.ID,
		}, ",")
		if tokenIDs != expectedTokenIDs {
			t.Errorf("Unexpected token sequence: got=%v, expected=%v", tokenIDs, expectedTokenIDs)
		}
	}

}
