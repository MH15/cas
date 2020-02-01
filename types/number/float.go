package number

type Float struct {
	Float float64
}

func (n Float) print() string {
	return "print for Float"
}

func (n Float) dump() map[string]interface{} {
	return make(map[string]interface{})
}

func (n Float) float() float64 {
	return n.Float
}

func (n Float) Equals(a Number) bool {
	// nF := n.float()
	// aF := a.float()
	// fmt.Println(nF, aF)
	return n.float() == a.float()
}

func (n Float) CompareTo(a Number) int {
	if a.float() < n.float() {
		return -1
	}
	if a.float() > n.float() {
		return 1
	}
	return 0
}
