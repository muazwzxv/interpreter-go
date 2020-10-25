package lexer

type Lexer struct {
	input        string
	position     int  // Current position in input(index current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // Current char under examination
}
