package parser

import (
	"fmt"
)

var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

const DELIMITERS = "+-*/^%()"

func Parse(tokens []string) {
	countEquals := occurances(tokens, "=")
	if countEquals > 1 {
		panic("Expressions may not have more than one equal sign.")
	}

	if countEquals == 1 {
		// Analyze equality
		index := contains(tokens, "=")
		fmt.Println(index)
		fmt.Println(tokens[index])

		left := tokens[0:index]
		right := tokens[index+1 : len(tokens)]
		fmt.Println(left, right)
		fmt.Println("\n----- LEFT -----")
		parseTokens(left)
		fmt.Println("\n----- RIGHT -----")
		parseTokens(right)

	} else {
		fmt.Println("expression not equation")
	}

}

func parseInfix(tokens []string) (rpn string, rpn_tokens []string) {
	var stack []string // holds operators and left parenthesis

	for _, tok := range tokens {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
				rpn_tokens = append(rpn_tokens, op)
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
					rpn_tokens = append(rpn_tokens, op)
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
				rpn_tokens = append(rpn_tokens, tok)
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		rpn_tokens = append(rpn_tokens, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return rpn, rpn_tokens
}

func parseTokens(tokens []string) {
	rpn, stack := parseInfix(tokens)
	fmt.Println(rpn)
	fmt.Println(stack)
	// e := types.Expression{
	// 	Left:  nil,
	// 	Op:    types.PLUS,
	// 	Right: nil,
	// }

	// tokenReader := Tokens{tokens, 0}

	// fmt.Println(e)
	// fmt.Println(tokenReader.Peek())

	// /*
	// 	Algorithm:
	// 	 - read until operator found
	// 	 - recurse on stuff to left of operator
	// 	 - consume said operator
	// 	 - continue
	// */

}

func contains(slice []string, str string) int {
	for i, a := range slice {
		if a == str {
			return i
		}
	}
	return -1
}

func occurances(slice []string, str string) int {
	count := 0
	for _, a := range slice {
		if a == str {
			count++
		}
	}
	return count
}
