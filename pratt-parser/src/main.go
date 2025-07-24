package main

import (
	"errors"
	"fmt"
	"mguzm4n/pratt-parser/src/lexer"
	"mguzm4n/pratt-parser/src/parser"
)

/*
Ex. -1 => - binds 1 to the right
*/
func prefixBindingPower(atom parser.Atom) (interface{}, uint8) {
	switch atom.Char {
	case '+', '-':
		return nil, 5
	default:
		fmt.Printf("%+v\n", atom)
		panic("bad operation atom value")
	}
}

func infixBindingPower(atom parser.Atom) (uint8, uint8, error) {
	switch atom.Char {
	case '+', '-':
		return 1, 2, nil
	case '*', '/':
		return 3, 4, nil
	default: // catch closing parenthesis
		return 1, 2, errors.New("")
	}
}

func expr(input string) parser.Node {
	lex := lexer.New(input)
	// lex.DbgPrintTokens()

	return exprBp(lex, 0)
}

func exprBp(lex *lexer.Lexer, minBp uint8) parser.Node {
	var lhs parser.Node

	switch tok := lex.Next(); {
	case tok.Type == lexer.Atom:
		lhs = parser.Atom{
			Char: tok.Value,
		}
	case tok.Type == lexer.Op && tok.Value == '(':
		lhs = exprBp(lex, 0) // keep parsing inside parenthesis
		if next := lex.Next(); !(next.Type == lexer.Op && next.Value == ')') {
			panic("bad token for op open parenthesis")
		}
	case tok.Type == lexer.Op: // possible unary operator
		op := parser.Atom{Char: tok.Value}
		_, rBp := prefixBindingPower(op)
		rhs := exprBp(lex, rBp)
		lhs = parser.Cons{
			Head: op.Char,
			Tail: []parser.Node{rhs},
		}
	default:
		fmt.Printf("%+v\n", tok)
		panic("bad token for lhs")
	}

	for {
		onEof := false
		var op parser.Atom

		switch next := lex.Peek(); next.Type {
		case lexer.Eof:
			onEof = true
		case lexer.Op:
			op = parser.Atom{
				Char: next.Value,
			}
		default:
			fmt.Printf("%+v\n", next)
			panic("bad token for op")
		}

		if onEof {
			break
		}

		lBp, rBp, err := infixBindingPower(op)
		if err != nil {
			break
		}

		if lBp < minBp {
			break
		}

		lex.Next()
		rhs := exprBp(lex, rBp)
		lhs = parser.Cons{
			Head: op.Char,
			Tail: []parser.Node{
				lhs, rhs,
			},
		}
	}

	return lhs
}

func main() {
	s := expr("1 + 2")
	fmt.Print(s.String())
}
