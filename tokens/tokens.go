package token

type TokenType string

type TokenStruct struct{
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

	FUNCTION = "FUNCTION"
	LET = "LET"
)

