package parser

import "mguzm4n/multichar-parser/src/lexer"

type Parser struct {
	tokens  []lexer.Token
	current int64
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens,
		0,
	}
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.Eof
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) checkNextAgainst(tkn lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tkn
}

func (p *Parser) match(tokens ...lexer.TokenType) bool {
	for tkn := range tokens {
		if p.checkNextAgainst(lexer.TokenType(tkn)) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(lexer.Bang_Eq, lexer.Eq_Eq) {
		op := p.previous()
		right := p.comparison()
		expr = NewBinary(expr, op, right)
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for p.match(
		lexer.Gt, lexer.Geq,
		lexer.Ls, lexer.Leq,
	) {
		op := p.previous()
		right := p.term()
		expr = NewBinary(expr, op, right)
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(lexer.Minus, lexer.Plus) {
		op := p.previous()
		right := p.factor()
		expr = NewBinary(expr, op, right)
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(lexer.Slash, lexer.Star) {
		op := p.previous()
		right := p.unary()
		expr = NewBinary(expr, op, right)
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(lexer.Bang, lexer.Minus) {
		op := p.previous()
		right := p.unary()
		return NewUnary(op, right)
	}
	return primary()
}

func (p *Parser) primary() Expr {
	if p.match(lexer.False) return NewLiteral(false)
	if p.match(lexer.True) return NewLiteral(true)
}
