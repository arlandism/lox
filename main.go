package main

import (
	"fmt"
	"lox/lexer"
)

func main() {
	s := lexer.NewScanner("{}()*;/\\")
	err := s.ReadTokens()
	if err != nil {
		fmt.Println(err)
	}
}
