/*

Link: https://neetcode.io/problems/evaluate-reverse-polish-notation/question?list=neetcode150
You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.

Return the integer that represents the evaluation of the expression.

The operands may be integers or the results of other operations.
The operators include '+', '-', '*', and '/'.
Assume that division between integers always truncates toward zero.
Example 1:

Input: tokens = ["1","2","+","3","*","4","-"]

Output: 5

Explanation: ((1 + 2) * 3) - 4 = 5
Constraints:

1 <= tokens.length <= 1000.
tokens[i] is "+", "-", "*", or "/", or a string representing an integer in the range [-200, 200].

*/

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		convert, _ := strconv.Atoi(tokens[0])
		return convert
	}
	val := 0
	stack := []int{}
	for _, item := range tokens {
		convert, error := strconv.Atoi(item)
		if error == nil {
			stack = append(stack, convert)
		} else {
			firstElement := stack[len(stack)-2]
			secondElement := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch item {
			case "+":
				val = firstElement + secondElement
			case "-":
				val = firstElement - secondElement
			case "*":
				val = firstElement * secondElement
			case "/":
				val = firstElement / secondElement
			}
			stack = append(stack, val)
		}
	}
	return val
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	fmt.Println(evalRPN(tokens))

}
