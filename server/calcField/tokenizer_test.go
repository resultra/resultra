package calcField

import (
	"fmt"
	"strings"
	"testing"
)

func tokenIDsToList(ids []int) string {
	tokenList := []string{}
	for id := range ids {
		tokenList = append(tokenList, fmt.Sprintf("%v", id))
	}
	return strings.Join(tokenList, ",")

}

func testOneTokenize(t *testing.T, tokenizeWhiteComment bool, inputStr string, expectedTokenIDs []int, whatTest string) {
	if matchSeq, err := tokenizeInput(inputStr, tokenizeWhiteComment); err != nil {
		t.Fatal(err)
	} else {
		tokenIDs := tokenIDsToList(matchSeq.tokenIDs())
		expectedIDs := tokenIDsToList(expectedTokenIDs)
		if tokenIDs != expectedIDs {
			t.Errorf("testOneTokenString: Unexpected token sequence: %v: got=[%v], expected=[%v]", whatTest, tokenIDs, expectedIDs)
		}
	}
}

func TestTokens(t *testing.T) {

	tokenizeWhiteOrComment := false

	testOneTokenize(t, tokenizeWhiteOrComment, `  == = HelloWorldFunc  ( arg1, arg3,true ,-32.43,22,"hello \" world" )   `, []int{
		tokenEqual.ID, tokenAssign.ID,
		tokenIdent.ID, tokenLParen.ID,
		tokenIdent.ID, tokenComma.ID,
		tokenIdent.ID, tokenComma.ID,
		tokenTrue.ID, tokenComma.ID,
		tokenMinus.ID, tokenNumber.ID, tokenComma.ID,
		tokenNumber.ID, tokenComma.ID,
		tokenText.ID,
		tokenRParen.ID,
	}, "kitchen sink")

	testOneTokenize(t, tokenizeWhiteOrComment, `  [ foo ] " hello [ world ]" `,
		[]int{tokenLBracket.ID, tokenIdent.ID, tokenRBracket.ID, tokenText.ID}, "brackets inside and outside quoted text")

	testOneTokenize(t, tokenizeWhiteOrComment, ` funcName(" hello \" world ]") `,
		[]int{tokenIdent.ID, tokenLParen.ID, tokenText.ID, tokenRParen.ID}, "escaped quote inside text")

}

func TestIdentTokens(t *testing.T) {

	tokenizeWhiteOrComment := false

	testOneTokenize(t, tokenizeWhiteOrComment, `a B b_ a_b c-d a_b_c_d `,
		[]int{tokenIdent.ID, tokenIdent.ID, tokenIdent.ID,
			tokenIdent.ID, tokenIdent.ID, tokenIdent.ID}, "identifyers - single letter idents and mixed with non-alpha characters")

	testOneTokenize(t, tokenizeWhiteOrComment, `F20160603204641834_cb3a9fd2-a942-450b-a93c-3fee5b58e6e5`,
		[]int{tokenIdent.ID}, "UUID - used in place of field after preprocessing")
}

func TestCommentTokens(t *testing.T) {

	tokenizeWhiteOrComment := true

	testOneTokenize(t, tokenizeWhiteOrComment, `// stuff after comment `,
		[]int{tokenComment.ID}, "comment")

	tokenizeWhiteOrComment = false
	testOneTokenize(t, tokenizeWhiteOrComment, `// stuff after comment `,
		[]int{}, "comment stripped")

}

func TestWhitespaceTokens(t *testing.T) {

	tokenizeWhiteOrComment := true

	testOneTokenize(t, tokenizeWhiteOrComment, `   ident1   ident2  // stuff after comment `,
		[]int{tokenWhiteSpace.ID, tokenIdent.ID,
			tokenWhiteSpace.ID, tokenIdent.ID,
			tokenWhiteSpace.ID, tokenComment.ID}, "comment")

	tokenizeWhiteOrComment = false

	testOneTokenize(t, tokenizeWhiteOrComment, `   ident1   ident2  // stuff after comment `,
		[]int{tokenIdent.ID, tokenIdent.ID}, "comment")

}

func TestGlobalReferenceTokens(t *testing.T) {

	tokenizeWhiteOrComment := false

	testOneTokenize(t, tokenizeWhiteOrComment, `[[ident2]] `,
		[]int{tokenDoubleLBracket.ID, tokenIdent.ID, tokenDoubleRBracket.ID},
		"global references")

	testOneTokenize(t, tokenizeWhiteOrComment, `[ident1] [[ident2]] `,
		[]int{tokenLBracket.ID, tokenIdent.ID, tokenRBracket.ID,
			tokenDoubleLBracket.ID, tokenIdent.ID, tokenDoubleRBracket.ID},
		"field vs global references")

}
