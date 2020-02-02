package expression

import (
	"cas/parser"
	"cas/tokenizer"
)

type Comparison string

const (
	Equal                Comparison = "="
	LessThan                        = "<"
	LessThanOrEqualTo               = "<="
	GreaterThan                     = ">"
	GreaterThanOrEqualTo            = ">="
)

type Equation struct {
	left     *Expression
	right    *Expression
	equality Comparison
}

func NewEquation(infix string) (equation Equation) {
	tokens := tokenizer.Tokens(infix)
	equalsCount := parser.Occurances(tokens, "=")
	equation.equality = Equal
	if equalsCount == 1 {
		left, right := parser.ParseEquation(tokens)
		equation.left = &Expression{left}
		equation.right = &Expression{right}

	} else {
		panic("Expressions may not have more than one equal sign.")
	}
	return equation
}

func (e *Equation) ToString() string {
	s := e.left.ToString() + string(e.equality) + e.right.ToString()
	return s
}
