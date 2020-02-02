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

func (n *BinaryNode) ToString() string {
	var left, data, right string
	if n.Left != nil {
		left = n.Left.ToString()
	} else {
		left = ""
	}
	if n.Data != "" {
		data = n.Data
	} else {
		data = ""
	}

	if n.Right != nil {
		right = n.Right.ToString()
	} else {
		right = ""
	}

	s := left + data + right
	return s
}
