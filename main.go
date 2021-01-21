package main

import (
	"cas/expression"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// n1 := number.Fraction{Numerator: 4, Denominator: 5}
	// n2 := number.Float{Float: 0.8}
	// fmt.Println(n1)
	// fmt.Println(n2)
	// fmt.Println(n2.Equals(n1))
	// fmt.Println(n1.Equals(n2))

	// equation := expression.NewEquation("4=x/7")
	// equation = expression.NewEquation("(2*x+3)/4=(x+7)/3")
	// spew.Dump(equation)

	fmt.Println("--- DUMP ---")
	eq := expression.NewExpression("(a*4+5^7)/73")

	spew.Dump(eq)

	fmt.Println("--- RECONSTRUCTION ---")
	fmt.Println(eq.ToString())

	// fmt.Println(equation.ToString())

}
