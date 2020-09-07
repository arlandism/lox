package lexer

import (
	"fmt"
	"unicode/utf8"
)

type Scanner struct {
	current int
	source  string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		0,
		source,
	}
}

func (s *Scanner) ReadTokens() {
	for s.current < len(s.source) {
		token := s.readNextToken()
		fmt.Printf("%s\n", token)
	}
}

func (s *Scanner) readNextToken() Token {
	r := s.advance()
	return NewToken(string(r))
}

// Iterating through a string returns raw byte values
// not 'characters', so we have to use the internal utf-8
// lib to grab the next Unicode code point and the byte-width
// of that code point, incrementing the 'current' index pointer
// by the byte-width.
func (s *Scanner) advance() rune {
	r, width := utf8.DecodeRuneInString(s.source[s.current:])
	s.current += width
	return r
}
