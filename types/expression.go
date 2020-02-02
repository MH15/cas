package types

type Comparison int

const (
	Equal Comparison = 0
	LessThan
	LessThanOrEqualTo
	GreaterThan
	GreaterThanOrEqualTo
)

type Expression struct {
	tree     BinaryNode
	equality Comparison
}
