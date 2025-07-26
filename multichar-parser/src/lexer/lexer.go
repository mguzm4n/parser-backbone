package lexer

import (
	"fmt"
	"strconv"
	"unicode"
)

type Lexer struct {
	source  string
	tokens  []Token
	start   int64
	current int64
}

func New(input string) *Lexer {
	return &Lexer{
		source: input,
	}
}

func (l *Lexer) Tokens() []Token {
	return l.tokens
}

func (l *Lexer) addToken(tt TokenType) {
	txt := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{
		Type:   tt,
		Lexeme: txt,
	})
}

func (l *Lexer) addTokenLiteral(tt TokenType, literal any) {
	txt := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{
		Type:    tt,
		Literal: literal,
		Lexeme:  txt,
	})
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= int64(len(l.source))
}

func (l *Lexer) advance() rune {
	ch := rune(l.source[l.current])
	l.current++
	return ch
}

// look ahead 1 char
func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return rune('\000')
	}

	return rune(l.source[l.current])
}

// look ahead two chars
func (l *Lexer) peekNext() rune {
	if l.current+1 >= int64(len(l.source)) {
		return rune('0')
	}

	return rune(l.source[l.current+1])
}

func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		l.advance()
	}

	if l.isAtEnd() {
		panic("[string scanner]: unterminated")
	}

	l.advance()
	val := l.source[l.start+1 : l.current-1] // here, trim the quotes
	l.addTokenLiteral(Str, val)
}

func (l *Lexer) number() {
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && unicode.IsDigit(l.peekNext()) { // here, we need 2 char lookahead
		l.advance()

		for unicode.IsDigit(l.peek()) {
			l.advance()
		}
	}

	val := l.source[l.start:l.current]
	parsedVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("[number scanner] tried to parse '%s'\n", val)
		panic("[number scanner] incorrect number")
	}

	l.addTokenLiteral(
		Num,
		parsedVal,
	)
}

func (l *Lexer) scanToken() {
	c := l.advance()
	switch c {
	case '+':
		l.addToken(Plus)
	case '-':
		l.addToken(Minus)
	case ':':
		l.addToken(Div)
	case '*':
		l.addToken(Star)
	case '(':
		l.addToken(LParen)
	case ')':
		l.addToken(RParen)
	case '"':
		l.string()
	case ' ', '\r', '\t':
		break
	default:
		if unicode.IsDigit(c) {
			l.number()
		} else {
			panic("[scanner] unrecognized character")
		}
	}
}

func (l *Lexer) Scan() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, Token{Type: Eof})
	return l.tokens
}

func (l *Lexer) DbgPrintTokens() {
	for _, it := range l.tokens {
		fmt.Printf("%+v\n", it)
	}
}
