/*

Link: https://neetcode.io/problems/sliding-window-maximum/question
You are given an array of integers nums and an integer k. There is a sliding window of size k that starts at the left edge of the array. The window slides one position to the right until it reaches the right edge of the array.

Return a list that contains the maximum element in the window at each step.

Input: nums = [1,2,1,0,4,2,6], k = 3

Output: [2,2,4,4,6]

*/

package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	left, right := 0, 0
	q := []int{}

	for right < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[right] {
			q = q[:len(q)-1]
		}
		q = append(q, right)

		if left > q[0] {
			q = q[1:]
		}

		if right+1 >= k {
			output = append(output, nums[q[0]])
			left++
		}
		right++
	}
	return output
}

func main() {
	nums := []int{1, 2, 1, 0, 4, 2, 6}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
