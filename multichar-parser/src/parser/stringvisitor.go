package parser

import (
	"fmt"
	"strings"
)

type StringVisitor struct{}

func (v *StringVisitor) Print(expr Expr) string {
	return expr.Accept(v).(string)
}

func (v *StringVisitor) parenthesize(name string, expr ...Expr) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)

	for _, e := range expr {
		sb.WriteString(" ")
		sb.WriteString(
			e.Accept(v).(string),
		)
	}

	sb.WriteString(")")
	return sb.String()
}

func (v *StringVisitor) VisitBinaryExpr(expr *Binary) any {
	return v.parenthesize(
		expr.Operator.Lexeme,
		expr.Left,
		expr.Right,
	)
}

func (v *StringVisitor) VisitGroupingExpr(expr *Grouping) any {
	return v.parenthesize(
		"group",
		expr.Expression,
	)
}

func (v *StringVisitor) VisitLiteralExpr(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", expr.Value)
}

func (v *StringVisitor) VisitUnaryExpr(expr *Unary) any {
	return v.parenthesize(
		expr.Operator.Lexeme,
		expr.Right,
	)
}
