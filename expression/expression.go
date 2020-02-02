package expression

import (
	"cas/parser"
	"cas/types"
)

type Expression struct {
	tree types.BinaryNode
}

func NewExpression(tokens []string) Expression {
	return Expression{
		tree: parser.ParseExpression(tokens),
	}
}

func (e *Expression) ToString() string {
	s := e.tree.ToString()
	return s
}
