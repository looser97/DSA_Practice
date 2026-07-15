/*
Link: https://leetcode.com/problems/minimum-window-substring/description/
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.

*/

package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	mapS := make([]int, 128)
	count := len(t)
	start, end := 0, 0
	minLen, startIndex := int(^uint(0)>>1), 0
	/// UPVOTE !
	for _, char := range t {
		mapS[char]++
	}

	for end < len(s) {
		if mapS[s[end]] > 0 {
			count--
		}
		mapS[s[end]]--
		end++

		for count == 0 {
			if end-start < minLen {
				startIndex = start
				minLen = end - start
			}

			if mapS[s[start]] == 0 {
				count++
			}
			mapS[s[start]]++
			start++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}

func main() {
	s := ""
	t := ""

	fmt.Println(minWindow(s, t))
}
