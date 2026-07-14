/*
Link: https://leetcode.com/problems/valid-palindrome/description/
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import "fmt"

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left <= right {
		if !isAlphaNumeric(s[left]) {
			left++
		} else if !isAlphaNumeric(s[right]) {
			right--
		} else {
			if toLower(s[left]) != toLower(s[right]) {
				return false
			}
			left++
			right--
		}
	}

	return true

}

func main() {

	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
