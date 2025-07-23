package main

import (
	"fmt"
	"mguzm4n/pratt-parser/src/lexer"
)

func infixBindingPower(token lexer.Token) (uint8, uint8) {
	switch token.Value {
	case '+', '-':
		return 1, 2
	case '*', '/':
		return 3, 4
	default:
		fmt.Printf("%+v\n", token)
		panic("bad operation token value")
	}
}

func expr(input string) lexer.SNot {
	lex := lexer.New(input)
	lex.DbgPrintTokens()

	return exprBp(lex, 0)
}

func exprBp(lex *lexer.Lexer, minBp uint8) lexer.SNot {
	var lhs lexer.SNot

	switch next := lex.Next(); next.Type {
	case lexer.Atom:
		lhs.Head = lexer.Token{
			Type:  lexer.Atom,
			Value: next.Value,
		}
	default:
		fmt.Printf("%+v\n", next)
		panic("bad token for lhs")
	}

	for {
		onEof := false
		var op lexer.Token
		_ = op

		switch next := lex.Peek(); next.Type {
		case lexer.Eof:
			onEof = true
		case lexer.Op:
			op = lexer.Token{
				Type:  lexer.Op,
				Value: next.Value,
			}
		default:
			fmt.Printf("%+v\n", next)
			panic("bad token for op")
		}

		if onEof {
			break
		}

		lBp, rBp := infixBindingPower(op)
		if lBp < minBp {
			break
		}

		lex.Next()
		rhs := exprBp(lex, rBp)
		lhs = lexer.SNot{
			Head: op,
			Tail: []lexer.SNot{
				lhs, rhs,
			},
		}
	}

	return lhs
}

func main() {
	s := expr("1 + 2")
	s.String()
}
