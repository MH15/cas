package parser

import (
	"cas/types"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

func ParseEquation(tokens []string) (types.BinaryNode, types.BinaryNode) {
	// Analyze equality
	index := Contains(tokens, "=")
	fmt.Println(index)
	fmt.Println(tokens[index])

	left := tokens[0:index]
	right := tokens[index+1 : len(tokens)]
	// fmt.Println("\n----- LEFT -----")
	leftTree := ParseExpression(left)
	// fmt.Println("\n----- RIGHT -----")
	rightTree := ParseExpression(right)

	return leftTree, rightTree
}

func ParseExpression(tokens []string) types.BinaryNode {
	_, stack := parseInfix(tokens)
	// fmt.Println(rpn)
	fmt.Println(stack)

	tree := assembleTree(stack)
	// fmt.Println(tree)
	spew.Dump(tree)

	return tree
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

func assembleTree(tokens []string) types.BinaryNode {
	var stack []types.BinaryNode
	for _, token := range tokens {
		if isOperator(token) {
			right := stack[len(stack)-1] // pop last from stack
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1] // pop another from stack
			stack = stack[:len(stack)-1]

			node := types.BinaryNode{
				Left:  &left,
				Data:  token,
				Right: &right,
			}
			stack = append(stack, node) // push to stack
		} else {
			// psuh to stack
			stack = append(stack, types.BinaryNode{Data: token})
		}
	}
	fmt.Println("len: ", len(stack))
	return stack[0]
}

func isOperator(token string) bool {
	return strings.Contains("+-*/^%", token)

}

func assembleTree1(tokens []string) types.BinaryNode {
	l := len(tokens)
	if l > 2 {
		// split into left, right and operator
		data := tokens[l-1]
		right := tokens[l-2]
		left := tokens[0 : l-2]
		fmt.Println("data: ", data)
		// fmt.Println("right: ", right)
		// fmt.Println("left: ", left)

		// recurse
		leftTree := assembleTree(left)

		rightTree := types.BinaryNode{
			Left:  nil,
			Data:  right,
			Right: nil,
		}

		// assemble tree with result from recursion
		node := types.BinaryNode{
			Left:  &leftTree,
			Data:  data,
			Right: &rightTree,
		}
		return node
	} else {
		return types.BinaryNode{}
	}
}

func height(root *types.BinaryNode) int {
	if root == nil {
		return 0
	} else {
		// compute the height of each subtree
		lheight := height(root.Left)
		rheight := height(root.Right)

		// use the larger one
		if lheight > rheight {
			return lheight + 1
		} else {
			return rheight + 1
		}
	}
}

func printTree(root *types.BinaryNode, indent int) {
	h := height(root)
	for i := 1; i <= h; i++ {
		fmt.Println("call")
		printGivenLevel(root, i)
	}
}

func printGivenLevel(root *types.BinaryNode, level int) {
	if *root == (types.BinaryNode{}) {
		return
	}
	if level == 1 {
		fmt.Printf("%s \n", root.Data)
	} else if level > 1 {
		printGivenLevel(root.Left, level-1)
		printGivenLevel(root.Right, level-1)
		fmt.Println()
	}
}

// TODO: combine these functions
func Contains(slice []string, str string) int {
	for i, a := range slice {
		if a == str {
			return i
		}
	}
	return -1
}

func Occurances(slice []string, str string) int {
	count := 0
	for _, a := range slice {
		if a == str {
			count++
		}
	}
	return count
}
