package types

type Operator int

const (
	PLUS Operator = 0
	MINUS
	MULT
	DIVIDE
	POW
	EXP
	MOD
	LPAREN
	RPAREN
)

type BinaryNode struct {
	Left  *BinaryNode
	Data  string
	Right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}
