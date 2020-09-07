package lexer

// 1 + 2
// 1 * 2
// 3 / 4
// 5 * 7

type Token struct {
	literal string
}

func NewToken(literal string) Token {
	return Token{literal}
}

func (t Token) String() string {
	return t.literal
}
