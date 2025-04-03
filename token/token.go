package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	NUM   = "NUM"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIV    = "/"
	BANG   = "!"
	EQ     = "=="
	NEQ    = "!="
	LT     = "<"
	GT     = ">"
	LE     = "<="
	GE     = ">="
	PLUS_ASSIGN = "+="
	MINUS_ASSIGN = "-="
	MULT_ASSIGN = "*="
	DIV_ASSIGN = "/="
	PLUS_PLUS = "++"
	MINUS_MINUS = "--"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if":  IF,
	"else": ELSE,
	"return": RETURN,
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}