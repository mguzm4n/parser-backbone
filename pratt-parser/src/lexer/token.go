package lexer

import (
	"fmt"
)

type Char rune

type TokenType int

const (
	Atom TokenType = iota
	Op
	Eof
)

type Token struct {
	Type  TokenType
	Value Char
}

func EOF() Token {
	return Token{
		Type:  Eof,
		Value: Char(0),
	}
}

func (c Char) String() string {
	return fmt.Sprintf("%q", rune(c))
}
