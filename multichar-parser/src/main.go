package main

import (
	"fmt"
	"mguzm4n/multichar-parser/src/lexer"
	"mguzm4n/multichar-parser/src/parser"
)

func main() {
	lex := lexer.New(`"0.5 - 100" + "b"`)
	lex.Scan()
	lex.DbgPrintTokens()

	// expr := parser.NewBinary(
	// 	parser.NewUnary(
	// 		lexer.Token{Type: lexer.Minus, Lexeme: "-", Literal: nil},
	// 		parser.NewLiteral(11),
	// 	),
	// 	lexer.Token{Type: lexer.Star, Lexeme: "*", Literal: nil},
	// 	parser.NewGrouping(parser.NewLiteral(22.5)),
	// )

	rdParser := parser.NewParser(lex.Tokens())
	expr := rdParser.Parse()

	interpreter := parser.NewInterpreter()
	err := interpreter.Interpret(expr)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// visitor := parser.StringVisitor{}
	// fmt.Println(visitor.Print(expr))
}
