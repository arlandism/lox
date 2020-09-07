package lexer

import "fmt"

const (
	UNKNOWN = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	IDENTIFIER
	STRING
	NUMBER

	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	WHILE

	EOF
)

type Token struct {
	tokenType int
	literal   string
}

func NewToken(tokenType int, literal string) Token {
	return Token{tokenType, literal}
}

func (t Token) String() string {
	return fmt.Sprintf("Token literal: %s\n", t.literal)
}
