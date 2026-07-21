/*
Link: https://neetcode.io/problems/validate-parentheses/question?list=neetcode150
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.

Example 1:

Input: s = "[]"

Output: true
Example 2:

Input: s = "([{}])"

Output: true
Example 3:

Input: s = "[(])"

Output: false

*/

// Can optimize by pushing the closing brackets instead of opening

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else {
			l := len(stack)
			if l == 0 {
				return false
			}
			if ch == '}' && stack[l-1] != '{' {
				return false
			}
			if ch == ']' && stack[l-1] != '[' {
				return false
			}
			if ch == ')' && stack[l-1] != '(' {
				return false
			}
			stack = stack[:l-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	s := "[(])"

	fmt.Println(isValid(s))
}
