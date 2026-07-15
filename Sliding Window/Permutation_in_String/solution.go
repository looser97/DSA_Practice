/*
Link: https://leetcode.com/problems/permutation-in-string/description/
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.



Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false


Constraints:

1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.

*/

package main

import (
	"fmt"
	"slices"
)

func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}
	s1Slice := make([]int, 26)
	s2Slice := make([]int, 26)
	for _, ch := range s1 {
		s1Slice[ch-'a']++
	}

	left := 0
	right := 0

	for right < len(s1)-1 {
		s2Slice[s2[right]-'a']++
		right++
	}

	for right < len(s2) {
		s2Slice[s2[right]-'a']++
		if slices.Equal(s1Slice, s2Slice) {
			return true
		}
		s2Slice[s2[left]-'a']--
		left++
		right++
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println(checkInclusion(s1, s2))

}
