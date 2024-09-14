package token

type TokenType string

const (
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	SLASH    TokenType = "/"
	ASTERISK TokenType = "*"
	BANG     TokenType = "!"
	ASSIGN   TokenType = "="
	EQ       TokenType = "=="
	NOT_EQ   TokenType = "!="
	ARROW    TokenType = "->"
	COLON    TokenType = ":"
	SEMI     TokenType = ";"
	COMMA    TokenType = ","
	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LT       TokenType = "<"
	GT       TokenType = ">"

	IDENT  TokenType = "IDENT"
	LET    TokenType = "LET"
	CONST  TokenType = "CONST"
	FUNC   TokenType = "FUNC"
	RETURN TokenType = "RETURN"
	TRUE   TokenType = "TRUE"
	FALSE  TokenType = "FALSE"
	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"

	INT    TokenType = "INT"
	FLOAT  TokenType = "FLOAT"
	STRING TokenType = "STRING"

	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":    LET,
	"const":  CONST,
	"func":   FUNC,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
