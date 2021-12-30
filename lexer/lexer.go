package lexer

import (
	"github.com/domterion/yapl/token"
)

// A struct to represent and hold data for the Lexer
type Lexer struct {
	input        string
	position     int  // The current position in the input, a pointer to current char
	readPosition int  // The current reading position in the input, current char + 1
	ch           byte // Current char being examined
}

// Instantiates a new lexer with the given input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// Reads the next char then advances the read position
// If we reach the end of the input then it sets the current char to 0 or NUL indicating the EOF or we havent reached anything
// If not then it sets the current char to the next char
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code for NUL
		l.ch = 0
	} else {
		// Sets the current char to the char at the read position
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition

	l.readPosition += 1
}

// Examines the current char and return a token that matches that char then we advance the char
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		// Lets look into the future to make sense of the code!
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// We can look into the future without going 88 MPH
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
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
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
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
			// Early returns because we already advanced the char above and dont need to do it again
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			// Same reason for early return as above
			return tok
		} else {
			// We dont know how to handle it therefore the char is illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Advance the char since we have no examined the current
	l.readChar()

	return tok
}

// Reads an identifier andd advances the lexers position till it finds a non number (!0-9) char
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Reads an identifier and advances the lexers position till it finds a non letter char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// This isnt Python, whitespace means nothing to us.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// A helper function to check if a char is a number or not
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// A helper function to check if a char is a letter or not
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// A helper function to help with token creation
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// A helper function that peeks at the next char without advancing position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
