/*
DESCRIPTION
Given an array of integers nums and an integer k, find the maximum sum of any contiguous subarray of size k.

Example 1: Input:

nums = [2, 1, 5, 1, 3, 2]
k = 3
Output:

9
Explanation: The subarray with the maximum sum is [5, 1, 3] with a sum of 9.
*/

package main

import (
	"fmt"
	"math"
)

func maxSum(nums []int, k int) int {
	windowSum := 0
	maxi := math.MinInt32
	start := 0
	end := 0

	for end < len(nums) {
		windowSum += nums[end]

		if end-start+1 == k {
			maxi = max(maxi, windowSum)
			windowSum -= nums[start]
			start++
		}

		end++
	}
	return maxi
}

func main() {
	nums := []int{-1, -2, -3, -4, -5}
	k := 2

	fmt.Println(maxSum(nums, k))
}
