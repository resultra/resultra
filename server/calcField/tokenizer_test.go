package calcField

import (
	"strings"
	"testing"
)

func testOneTokenize(t *testing.T, tokenizeWhiteComment bool, inputStr string, expectedTokenIDs []string, whatTest string) {
	if matchSeq, err := tokenizeInput(inputStr, tokenizeWhiteComment); err != nil {
		t.Fatal(err)
	} else {
		tokenIDs := strings.Join(matchSeq.tokenIDs(), ",")
		expectedIDs := strings.Join(expectedTokenIDs, ",")
		if tokenIDs != expectedIDs {
			t.Errorf("testOneTokenString: Unexpected token sequence: %v: got=[%v], expected=[%v]", whatTest, tokenIDs, expectedIDs)
		}
	}
}

func TestTokens(t *testing.T) {

	tokenizeWhiteOrComment := false

	testOneTokenize(t, tokenizeWhiteOrComment, `  == = HelloWorldFunc  ( arg1, arg3,true ,-32.43,22,"hello \" world" )   `, []string{
		tokenEqual.ID, tokenAssign.ID,
		tokenIdent.ID, tokenLParen.ID,
		tokenIdent.ID, tokenComma.ID,
		tokenIdent.ID, tokenComma.ID,
		tokenBool.ID, tokenComma.ID,
		tokenNumber.ID, tokenComma.ID,
		tokenNumber.ID, tokenComma.ID,
		tokenText.ID,
		tokenRParen.ID,
	}, "kitchen sink")

	testOneTokenize(t, tokenizeWhiteOrComment, `  [ foo ] " hello [ world ]" `,
		[]string{tokenLBracket.ID, tokenIdent.ID, tokenRBracket.ID, tokenText.ID}, "brackets inside and outside quoted text")

	testOneTokenize(t, tokenizeWhiteOrComment, ` funcName(" hello \" world ]") `,
		[]string{tokenIdent.ID, tokenLParen.ID, tokenText.ID, tokenRParen.ID}, "escaped quote inside text")

}

func TestCommentTokens(t *testing.T) {

	tokenizeWhiteOrComment := true

	testOneTokenize(t, tokenizeWhiteOrComment, `// stuff after comment `,
		[]string{tokenComment.ID}, "comment")

	tokenizeWhiteOrComment = false
	testOneTokenize(t, tokenizeWhiteOrComment, `// stuff after comment `,
		[]string{}, "comment stripped")

}

func TestWhitespaceTokens(t *testing.T) {

	tokenizeWhiteOrComment := true

	testOneTokenize(t, tokenizeWhiteOrComment, `   ident1   ident2  // stuff after comment `,
		[]string{tokenWhiteSpace.ID, tokenIdent.ID,
			tokenWhiteSpace.ID, tokenIdent.ID,
			tokenWhiteSpace.ID, tokenComment.ID}, "comment")

	tokenizeWhiteOrComment = false

	testOneTokenize(t, tokenizeWhiteOrComment, `   ident1   ident2  // stuff after comment `,
		[]string{tokenIdent.ID, tokenIdent.ID}, "comment")

}
