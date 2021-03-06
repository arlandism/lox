package lexer

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

//TODO: just condense this into a special unknown token
type UnrecognizedTokenError struct {
	literal string
	line    int
}

func NewUnrecognizedTokenError(literal string, line int) UnrecognizedTokenError {
	return UnrecognizedTokenError{
		literal,
		line,
	}
}

func (e UnrecognizedTokenError) Error() string {
	return fmt.Sprintf("LoxSyntaxError: Unrecognized token: %s on line %d", e.literal, e.line)
}

type Scanner struct {
	current int
	source  string
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		0,
		source,
		1,
	}
}

func (s *Scanner) ReadTokens() error {
	for s.current < len(s.source) {
		token, err := s.readNextToken()
		if err != nil {
			return err
		}
		if token.tokenType == EOF {
			break
		}
		fmt.Printf("%s\n", token)
	}
	return nil
}

func (s *Scanner) readNextToken() (Token, error) {
	var token bytes.Buffer
	r := s.advance()

WhitespaceLoop:
	for { // handle pass-through and whitespace tokens
		switch r {
		case '\n':
			s.line += 1 // for error reporting
		case '\t':
		case ' ':
		default:
			break WhitespaceLoop
		}
		r = s.advance()
	}

	token.WriteRune(r)
	switch r {
	case '\n':
		s.line += 1
		return NewToken(UNKNOWN, token.String()), nil
	case '(':
		return NewToken(LEFT_PAREN, token.String()), nil
	case ')':
		return NewToken(RIGHT_PAREN, token.String()), nil
	case '{':
		return NewToken(LEFT_BRACE, token.String()), nil
	case '}':
		return NewToken(RIGHT_BRACE, token.String()), nil
	case ',':
		return NewToken(COMMA, token.String()), nil
	case '.':
		return NewToken(DOT, token.String()), nil
	case '-':
		return NewToken(MINUS, token.String()), nil
	case '+':
		return NewToken(PLUS, token.String()), nil
	case ';':
		return NewToken(SEMICOLON, token.String()), nil
	case '/':
		if s.peekNext() == '/' {
			// this doesn't consume the newline rune
			// because we want that token to be handled by the
			// main scanning loop
			s.advanceStreamTo('\n')
			return s.readNextToken()
		} else {
			return NewToken(SLASH, token.String()), nil
		}
	case '*':
		return NewToken(STAR, token.String()), nil
	case '!':
		if s.peekNext() == '=' {
			next := s.advance()
			token.WriteRune(next)
			return NewToken(BANG_EQUAL, token.String()), nil
		} else {
			return NewToken(BANG, token.String()), nil
		}
	case '=':
		if s.peekNext() == '=' {
			next := s.advance()
			token.WriteRune(next)
			return NewToken(EQUAL_EQUAL, token.String()), nil
		} else {
			return NewToken(EQUAL, token.String()), nil
		}
	case '>':
		if s.peekNext() == '=' {
			next := s.advance()
			token.WriteRune(next)
			return NewToken(GREATER_EQUAL, token.String()), nil
		} else {
			return NewToken(GREATER, token.String()), nil
		}
	case '<':
		if s.peekNext() == '=' {
			next := s.advance()
			token.WriteRune(next)
			return NewToken(LESS_EQUAL, token.String()), nil
		} else {
			return NewToken(LESS, token.String()), nil
		}
	case utf8.RuneError:
		return NewToken(EOF, ""), nil
	default:
		return NewToken(UNKNOWN, token.String()), NewUnrecognizedTokenError(token.String(), s.line)
	}
}

// Consumes all tokens up to but not including the targetRune
func (s *Scanner) advanceStreamTo(targetRune rune) {
	if s.peekNext() == targetRune {
		return
	} else if s.peekNext() == utf8.RuneError {
		// this usually signals EOF so
		// let RuneError bubble up in the main scanning loop
		// where that case is handled
		// handling invalid encoding is a TODO
		return
	} else {
		s.advance()
		s.advanceStreamTo(targetRune)
	}
}

// Iterating through a string returns raw byte values
// not 'characters', so we have to use the internal utf-8
// lib to grab the next Unicode code point and the byte-width
// of that code point, incrementing the 'current' index pointer
// by the byte-width.

// Returns the rune and the number of bytes written
// RuneError, 0 is a special code meaning we've reached EOF
// RuneError, 1 is a special code meaning the encoding is invalid
func (s *Scanner) advance() rune {
	r, width := utf8.DecodeRuneInString(s.source[s.current:])
	if r != utf8.RuneError {
		// no need to advance the pointer for errors
		s.current += width
	}
	return r
}

// Similar to advance but doesn't actually consume the next token
func (s *Scanner) peekNext() rune {
	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return r
}
