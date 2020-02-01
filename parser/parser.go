package parser

import (
	"cas/types"
	"fmt"
)

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

func parseTokens(tokens []string) {
	e := types.Expression{
		Left:  nil,
		Op:    types.PLUS,
		Right: nil,
	}

	tokenReader := Tokens{tokens, 0}

	fmt.Println(e)
	fmt.Println(tokenReader.Peek())

	/*
		Algorithm:
		 - read until operator found
		 - recurse on stuff to left of operator
		 - consume said operator
		 - continue
	*/

	tokenQueue := make([]string, 0)
	for tokenReader.HasNext() {
		// token := tokenReader.Peek()
		// if contains(strings.Split(DELIMITERS, ""), token) > 0 {
		// 	// this is an operator, recurse on stuff prior to this
		// } else {
		// 	tokenQueue = append(tokenQueue, token)
		// }
	}

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
