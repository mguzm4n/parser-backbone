package main

import "mguzm4n/multichar-parser/src/lexer"

func main() {
	lex := lexer.New("1200 + 23.1001")
	lex.Scan()
	lex.DbgPrintTokens()
}
