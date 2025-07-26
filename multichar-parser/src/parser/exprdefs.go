package parser

import (
	"mguzm4n/multichar-parser/src/lexer"
)

type Expr interface {
	isExpr()
	Accept(v Visitor) any
}
type Binary struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (*Binary) isExpr() {}
func (b *Binary) Accept(v Visitor) any {
	return v.VisitBinaryExpr(b)
}
func NewBinary(Left Expr, Operator lexer.Token, Right Expr) *Binary {
	return &Binary{
		Left, Operator, Right,
	}
}

type Grouping struct {
	Expression Expr
}

func (*Grouping) isExpr() {}
func (g *Grouping) Accept(v Visitor) any {
	return v.VisitGroupingExpr(g)
}
func NewGrouping(Expression Expr) *Grouping {
	return &Grouping{
		Expression,
	}
}

type Literal struct {
	Value any
}

func (*Literal) isExpr() {}
func (l *Literal) Accept(v Visitor) any {
	return v.VisitLiteralExpr(l)
}
func NewLiteral(Value any) *Literal {
	return &Literal{
		Value,
	}
}

type Unary struct {
	Operator lexer.Token
	Right    Expr
}

func (*Unary) isExpr() {}
func (u *Unary) Accept(v Visitor) any {
	return v.VisitUnaryExpr(u)
}
func NewUnary(Operator lexer.Token, Right Expr) *Unary {
	return &Unary{
		Operator, Right,
	}
}

type Visitor interface {
	VisitBinaryExpr(expr *Binary) any
	VisitGroupingExpr(expr *Grouping) any
	VisitLiteralExpr(expr *Literal) any
	VisitUnaryExpr(expr *Unary) any
}
