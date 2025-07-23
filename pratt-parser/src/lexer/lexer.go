package lexer

import (
	"fmt"
	"mguzm4n/pratt-parser/src/sliceutils"
	"unicode"
)

type Lexer struct {
	tokens []Token
}

func New(input string) *Lexer {
	tokens := []Token{}

	for _, r := range input {
		switch {
		case unicode.IsSpace(r):
			continue
		case unicode.IsDigit(r):
			tokens = append(tokens, Token{
				Type:  Atom,
				Value: Char(r),
			})
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			tokens = append(tokens, Token{
				Type:  Op,
				Value: Char(r),
			})
		}
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
