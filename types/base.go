package types

type literal interface {
	print() string
	dump() string
}
