package lexer

import "fmt"

type TokenType int

const (
	Atom TokenType = iota
	Op
	OpOpenParen
	OpCloseParen
	Eof
)

var tokNames = [...]string{
	"Atom",
	"Op",
	"OpOpenParen",
	"OpCloseParen",
	"Eof",
}

type Token struct {
	Type  TokenType
	Value rune
}

func EOF() Token {
	return Token{
		Type:  Eof,
		Value: rune('0'),
	}
}

func (tt TokenType) String() string {
	return tokNames[tt]
}

func (t Token) String() string {
	return fmt.Sprintf("{ Type: %s, Value: %q }", t.Type, t.Value)
}
