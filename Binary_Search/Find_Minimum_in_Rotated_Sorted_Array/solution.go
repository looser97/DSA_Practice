/*
Link: https://neetcode.io/problems/find-minimum-in-rotated-sorted-array/question?list=neetcode150

You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:

[3,4,5,6,1,2] if it was rotated 4 times.
[1,2,3,4,5,6] if it was rotated 6 times.
Notice that rotating the array 4 times moves the last four elements of the array to the beginning. Rotating the array 6 times produces the original array.

Assuming all elements in the rotated sorted array nums are unique, return the minimum element of this array.

A solution that runs in O(n) time is trivial, can you write an algorithm that runs in O(log n) time?

Example 1:

Input: nums = [3,4,5,6,1,2]

Output: 1
Example 2:

Input: nums = [4,5,0,1,2,3]

Output: 0
*/

package main

import "fmt"

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	ans := nums[0]
	for l <= r {
		if nums[l] < nums[r] {
			ans = min(ans, nums[l])
			break
		}
		mid := l + (r-l)/2
		ans = min(ans, nums[mid])
		if nums[mid] >= nums[l] {
			// left side sorted

			l = mid + 1

		} else {
			// right side sorted
			r = mid - 1
		}
	}
	return ans
}

func main() {
	nums := []int{2, 1}
	fmt.Println(findMin(nums))
}
