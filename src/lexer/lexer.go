package lexer

import "monkey/token"

type Lexer struct {
	input string // only ascii characters are
	position int // current position in input (points to curret char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination, must be an ascii character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // readChar immediately so we populate readPosition and position correctly.
	return l
} 

// read the next ascii char in the Lexer's input string. 
// this is not idempotent. Reading the next progresses the Lexer's positions through the input string.
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

// Peak the next char in the Lexer's input string, without moving the lexer's positions. 
// this is idempotent. peakChar() can be called repeatedly without changing the state of the Lexer.
func (l *Lexer) peakChar() byte {
	if l.position >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=': 
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+': 
			tok = newToken(token.PLUS, l.ch)
	case '-': 
			tok = newToken(token.MINUS, l.ch)
	case '!': 
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/': 
			tok = newToken(token.SLASH, l.ch)
	case '*': 
			tok = newToken(token.ASTERISK, l.ch)
	case '<': 
			tok = newToken(token.LT, l.ch)
	case '>': 
			tok = newToken(token.GT, l.ch)
	case ';': 
		tok = newToken(token.SEMICOLON, l.ch)
	case '(': 
		tok = newToken(token.LPAREN, l.ch)
	case ')': 
		tok = newToken(token.RPAREN, l.ch)
	case ',': 
		tok = newToken(token.COMMA, l.ch)
	case '{': 
		tok = newToken(token.LBRACE, l.ch)
	case '}': 
		tok = newToken(token.RBRACE, l.ch)
	case 0: 
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// return early, we've already progressed our lexer during readIdentifier
			return tok
			} else if isDigit(l.ch) {
				tok.Literal = l.readNumber()
				tok.Type = token.INT
				// return early, we've already progressed our lexer during readNumber
				return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// progress the lexer until we read a char that isn't blank, tab, newline, carriage return
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// determines the valid characters that can be used in keywords/identifiers. 
// These EXPLICITLY should not be bytes that are mapped to already existing tokens.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' // add ch == '?' and ch == '!'?
}

func (l *Lexer) readIdentifier() string {
	// we cache the starting position
	position := l.position
	// progress the lexer until we read a character that isn't a letter
	for isLetter(l.ch) {
		l.readChar()
	}
	// return the substring that is the starting position to the character that isn't a letter (exclusive)
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	// we cache the starting position
	position := l.position
	// progress the lexer until we read a character that isn't a letter
	for isDigit(l.ch) {
		l.readChar()
	}
	// return the substring that is the starting position to the character that isn't a letter (exclusive)
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
