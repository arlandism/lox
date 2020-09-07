package main

import (
	"lox/lexer"
)

func main() {
	s := lexer.NewScanner("Hola World!")
	s.ReadTokens()
}
