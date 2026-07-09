/*

Link: https://leetcode.com/problems/valid-anagram/description/

Given two strings s and t, return true if t is an anagram of s, and false otherwise.



Example 1:

Input: s = "anagram", t = "nagaram"

Output: true

Example 2:

Input: s = "rat", t = "car"

Output: false



Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

*/

package main

import (
	"fmt"
	"slices"
)

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}

	sSlice := make([]int, 26)
	tSlice := make([]int, 26)

	for i := 0; i < sLen; i++ {
		sSlice[s[i]-'a']++
		tSlice[t[i]-'a']++
	}

	return slices.Equal(sSlice, tSlice)
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
