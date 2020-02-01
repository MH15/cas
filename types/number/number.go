package number

type Number interface {
	print() string
	dump() map[string]interface{}
	float() float64
	// Equals(Number) bool
	CompareTo(Number) int
}
