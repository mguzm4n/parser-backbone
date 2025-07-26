package parser

import (
	"fmt"
	"mguzm4n/multichar-parser/src/lexer"
)

type RuntimeError struct {
	Token *lexer.Token
	Msg   string
}

func (r *RuntimeError) Error() string {
	// return fmt.Sprintf("[Line %d] Runtime Error: %s", r.Token.Line, r.Message)
	return fmt.Sprintf("[Line] Runtime Error: %s", r.Msg)
}

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (v *Interpreter) stringify(val any) string {
	return fmt.Sprintf("%+v", val)
}
func (v *Interpreter) Interpret(expr Expr) error {
	res, err := v.eval(expr)
	if err != nil {
		return err
	}
	fmt.Printf("Expr evaluates to: %v", v.stringify(res))
	return nil
}

func (v *Interpreter) VisitLiteralExpr(expr *Literal) (any, error) {
	return expr.Value, nil
}

func (v *Interpreter) eval(expr Expr) (any, error) {
	return expr.Accept(v)
}

func (v *Interpreter) VisitGroupingExpr(expr *Grouping) (any, error) {
	return v.eval(expr.Expression)
}

func (v *Interpreter) isTruthy(anyVal any) bool {
	switch anyVal.(type) {
	case nil:
		return false
	default:
		if anyVal.(bool) {
			return anyVal.(bool)
		}
	}
	return true
}

func (v *Interpreter) checkNumberOperand(operator lexer.Token, operand any) error {
	if _, ok := operand.(float64); ok {
		return nil
	}
	return &RuntimeError{Token: &operator, Msg: "Operand must be number"}
}

func (v *Interpreter) checkNumberOperands(operator lexer.Token, left, right any) error {
	if _, ok := left.(float64); ok {
		if _, ok := right.(float64); ok {
			return nil
		}
	}
	return &RuntimeError{Token: &operator, Msg: "Operands must be numbers"}
}

func (v *Interpreter) VisitUnaryExpr(expr *Unary) (any, error) {
	right, err := v.eval(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case lexer.Minus:
		if err := v.checkNumberOperand(expr.Operator, right); err != nil {
			return nil, err
		}
		return -1 * right.(float64), nil
	case lexer.Bang:
		return v.isTruthy(right), nil
	}

	return nil, nil // unreachable
}

func (v *Interpreter) isEqual(left, right any) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}

	return left == right
}

func (v *Interpreter) VisitBinaryExpr(expr *Binary) (any, error) {
	left, err := v.eval(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := v.eval(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case lexer.Bang_Eq:
		return !v.isEqual(left, right), nil
	case lexer.Eq_Eq:
		return v.isEqual(left, right), nil
	case lexer.Gt:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil
	case lexer.Geq:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) >= right.(float64), nil
	case lexer.Ls:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil
	case lexer.Leq:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) <= right.(float64), nil
	case lexer.Plus:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) + right.(float64), nil
	case lexer.Minus:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case lexer.Slash:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) / right.(float64), nil
	case lexer.Star:
		if err := v.checkNumberOperands(expr.Operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil
	}

	return nil, nil
}
