%{

package calcField

import (
	"fmt"
)
		
%}

// Fields inside this union end up as the fields in a structure known
// as FormulaSymType. This struct is passed to the tokenizer/lexer
// to be populated during parsing.
%union{
	number float64
	text string
	eqnNode *EquationNode
	args []*EquationNode
}

%token<number> TOK_NUMBER
%token TOK_PLUS
%token TOK_TIMES
%token TOK_WHITE
%token<text> TOK_IDENT
%token TOK_ASSIGN
%token TOK_EQUAL
%token TOK_LPAREN
%token TOK_RPAREN
%token TOK_DOUBLE_LBRACKET
%token TOK_DOUBLE_RBRACKET
%token TOK_LBRACKET
%token TOK_RBRACKET
%token TOK_COMMA
%token TOK_BOOL
%token TOK_COMMENT
%token<text> TOK_TEXT

%left TOK_PLUS
%left TOK_TIMES

// Any non-terminal which returns a value needs a type. This type
// needs to be one of the field names in the %union above.
%type <eqnNode> expr
%type <eqnNode> func
%type <eqnNode> fieldRef
%type <eqnNode> globalRef
%type <args> funcArg
%type <args> funcArgs
	
%% // start of parser
			
root : expr 
		{ 
			// There's currently no "clean" way to return the root node
			// from the parser (and have the parser still be reentrant). 
			// The work-around is to set a value on the lexer (see https://goo.gl/NdKNYI)
			formulalex.(*formulaLexerImpl).rootEqnNode = $1
			fmt.Printf("\nRoot equation node: %+v\n",$1) 
		}		

expr	:   expr TOK_PLUS expr
		{ 
			funcArgs := []*EquationNode{$1,$3}
			$$  =  FuncEqnNode(FuncNameSum,funcArgs)
		}
		| expr TOK_TIMES expr
		{
			funcArgs := []*EquationNode{$1,$3}
			$$  =  FuncEqnNode(FuncNameProduct,funcArgs)
		}
		| fieldRef
		{ 
			$$ = $1
		}
		| globalRef
		{ 
			$$ = $1
		}
		| func
		{
			$$ = $1
		}
		| TOK_NUMBER
		{ 
			$$ = NumberEqnNode($1) 
		}
		| TOK_TEXT
		{ 
			$$ = TextEqnNode($1) 
		}

fieldRef : TOK_LBRACKET TOK_IDENT TOK_RBRACKET
		{
			$$ = FieldRefEqnNode($2)
		}

globalRef : TOK_DOUBLE_LBRACKET TOK_IDENT TOK_DOUBLE_RBRACKET
		{
			$$ = GlobalRefEqnNode($2)
		}


		
func : TOK_IDENT TOK_LPAREN funcArgs TOK_RPAREN
		{
			/* 1 or more arguments */
			$$ = FuncEqnNode($1,$3)
		}
		| TOK_IDENT TOK_LPAREN TOK_RPAREN
		{
			/* No arguments */
			$$ = FuncEqnNode($1,[]*EquationNode{})
		}

funcArgs: funcArg
		{
			$$ = $1
		}
		| funcArg TOK_COMMA funcArgs
		{
			$$ = []*EquationNode{}
			$$ = append($$,$1...)
			$$ = append($$,$3...)
		}

funcArg: expr
		{
			$$ = []*EquationNode{$1}
		}
		

%% // end of parser