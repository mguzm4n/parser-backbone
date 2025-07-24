package lexer

import (
	"fmt"
	"mguzm4n/pratt-parser/src/sliceutils"
	"strings"
	"unicode"
)

type Lexer struct {
	tokens []Token
}

func isRoundParen(c rune) bool {
	return c == '(' || c == ')'
}

func isOperation(c rune) bool {
	return strings.ContainsRune("+-*/", c)
}

func New(input string) *Lexer {
	tokens := []Token{}

	for _, r := range input {
		tok := Token{Value: r}
		switch {
		case unicode.IsSpace(r):
			continue
		case unicode.IsDigit(r):
			tok.Type = Atom
		case isOperation(r), isRoundParen(r):
			tok.Type = Op
		}
		tokens = append(tokens, tok)
	}

	sliceutils.Reverse(tokens)

	return &Lexer{
		tokens,
	}
}

func (l *Lexer) Next() Token {
	if len(l.tokens) == 0 {
		return EOF()
	}

	x := l.tokens[len(l.tokens)-1]
	l.tokens = l.tokens[:len(l.tokens)-1]
	return x
}

func (l *Lexer) Peek() Token {
	if len(l.tokens) == 0 {
		return EOF()
	}

	return l.tokens[len(l.tokens)-1]
}

func (l *Lexer) DbgPrintTokens() {
	for _, it := range l.tokens {
		fmt.Printf("%+v\n", it)
	}
}
