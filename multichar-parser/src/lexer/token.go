package lexer

import "fmt"

type TokenType int

const (
	Atom TokenType = iota
	Op
	Eof
	LParen
	RParen
	Plus
	Minus
	Star
	Div
	Str
	Num
)

var tokNames = [...]string{
	"Atom",
	"Op",
	"Eof",
	"LParen",
	"RParen",
	"Plus",
	"Minus",
	"Star",
	"Div",
	"Str",
	"Num",
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
}

func (tt TokenType) String() string {
	return tokNames[tt]
}

func (t Token) String() string {
	return fmt.Sprintf("{ Type: %s, Value: %s, Literal: %s }", t.Type, t.Lexeme, t.Literal)
}
