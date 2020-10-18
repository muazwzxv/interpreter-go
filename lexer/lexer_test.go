package lexer

import (
	"testing"
	"interpreter/token"
)

type struct Token {
	expectedToken token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=_(){},;`

	tests := []Token {{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

}
