package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type TokenType string
type Token struct {
	Type TokenType
	LiteralType string
}
