package main

import (
	"fmt"
	"mguzm4n/multichar-parser/src/lexer"
	"mguzm4n/multichar-parser/src/parser"
)

func main() {
	lex := lexer.New(`("what" + 1) "espacio"`)
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

	visitor := parser.StringVisitor{}
	fmt.Println(visitor.Print(expr))

}
