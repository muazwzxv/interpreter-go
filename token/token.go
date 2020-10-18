package token

type TokenType string

type Token struct {
	// Tokentype set to string for flexibility
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters 
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNTION"
	LET = "LET"

)
