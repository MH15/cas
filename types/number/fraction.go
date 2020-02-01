package number

type Fraction struct {
	Numerator   int
	Denominator int
}

func (n Fraction) print() string {
	return "print for Fraction"
}

func (n Fraction) dump() map[string]interface{} {
	return make(map[string]interface{})
}

func (n Fraction) float() float64 {
	return float64(n.Numerator) / float64(n.Denominator)
}

func (n Fraction) Equals(a Number) bool {
	// nF := n.float()
	// aF := a.float()
	// fmt.Println(nF, aF)
	return n.float() == a.float()
}

func (n Fraction) CompareTo(a Number) int {
	if a.float() < n.float() {
		return -1
	}
	if a.float() > n.float() {
		return 1
	}
	return 0
}
