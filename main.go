package main

import (
	"fmt"
	"lox/lexer"
)

func main() {
	s := lexer.NewScanner("{}()*;  \t/\n\n!!====>=<=<>  //thisisatest\n}^")
	err := s.ReadTokens()
	if err != nil {
		fmt.Println(err)
	}
}
