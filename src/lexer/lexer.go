package lexer

import "monkey/token"

type Lexer struct {
	input string // only ascii characters are
	position int // current position in input (points to currnt char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination, must be an ascii character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
} 

// read the next ascii char in the Lexer's input string. 
func (l *Lexer) readChar() {
	// we only support ascii characters (for now). Supporting unicode characters would mean chars could no longer be represented as bytes
	// but instead as runes, complicating this "next logic" and being unable to traverse the string simply, since a rune could be multiple bytes.

	if l.readPosition >= len(l.input) {
		l.ch= 0 // 0 is ascii NUL character
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=': 
		tok = newToken(token.ASSIGN, l.ch)
	case ';': 
		tok = newToken(token.SEMICOLON, l.ch)
	case '(': 
		tok = newToken(token.LPAREN, l.ch)
	case ')': 
		tok = newToken(token.RPAREN, l.ch)
	case ',': 
		tok = newToken(token.COMMA, l.ch)
	case '+': 
		tok = newToken(token.PLUS, l.ch)
	case '{': 
		tok = newToken(token.LBRACE, l.ch)
	case '}': 
		tok = newToken(token.RBRACE, l.ch)
	case 0: 
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
