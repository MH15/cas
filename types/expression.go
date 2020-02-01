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

type Expression struct {
	Left  *Expression
	Op    Operator
	Right *Expression
}
