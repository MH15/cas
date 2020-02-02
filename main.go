package main

import (
	"cas/parser"
	"cas/tokenizer"
	"fmt"
)

func main() {
	// n1 := number.Fraction{Numerator: 4, Denominator: 5}
	// n2 := number.Float{Float: 0.8}
	// fmt.Println(n1)
	// fmt.Println(n2)
	// fmt.Println(n2.Equals(n1))
	// fmt.Println(n1.Equals(n2))
	tokens := tokenizer.Tokens("x_1 + 4 = 2/3")
	tokens = tokenizer.Tokens("a+b*c = 2+6")
	tokens = tokenizer.Tokens("w_1^2 - w_2^2 = 2*alpha*omega")
	fmt.Println(tokens)
	parser.Parse(tokens)
}
