package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	COMMA = ","
	SEMICOLON = ";"
	ASSIGN = "="
	PLUS = "+"

	FUNCTION = "FUNCTION"
	LET = "LET"
)

