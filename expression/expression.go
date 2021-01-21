package expression

import (
	"cas/parser"
	"cas/tokenizer"
	"cas/types"
)

type Expression struct {
	tree types.BinaryNode
}

func NewExpression(infix string) Expression {
	tokens := tokenizer.Tokens(infix)
	return Expression{
		tree: parser.ParseExpression(tokens),
	}
}

func (e *Expression) ToString() string {
	s := e.tree.ToString()
	return s
}
