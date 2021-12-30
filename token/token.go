package token

// A base type for all tokens. The type is a string because it is a flexible type in regards to what its values can be
type TokenType string

// The struct to represent tokens
type Token struct {
	Type    TokenType // The type the token represents
	Literal string    // The literal repreesentatopm
}

// All valid tokens that the lexer can recognize
const (
	// Special
	ILLEGAL = "ILLEGAL" // An invalid token, one the lexer doesnt recognize
	EOF     = "EOF"     // Tells the lexer it is done with the file and can stop
	COMMENT = "COMMENT" // well... this is obvious

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1243456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// A map of keyword to identifier
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Checks if the ident is a valid keyword and if so returns the type, if not returns IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
