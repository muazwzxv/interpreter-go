package lexer

import "interpreter/token"

// Lexer struct
type Lexer struct {
	input        string
	position     int  // Current position in input(points current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // Current char under examination
}

// New lexer
func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken function
func (l *Lexer) NextToken() token.Token {
	var key token.Token

	switch l.ch {
	case '=':
		key = newToken(token.ASSIGN, l.ch)
	case ';':
		key = newToken(token.SEMICOLON, l.ch)
	case '(':
		key = newToken(token.LPAREN, l.ch)
	case ')':
		key = newToken(token.RPAREN, l.ch)
	case ',':
		key = newToken(token.COMMA, l.ch)
	case '+':
		key = newToken(token.PLUS, l.ch)
	case '{':
		key = newToken(token.LBRACE, l.ch)
	case '}':
		key = newToken(token.RBRACE, l.ch)
	case 0:
		key.Literal = ""
		key.Type = token.EOF
	}

	l.readChar()
	return key
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
