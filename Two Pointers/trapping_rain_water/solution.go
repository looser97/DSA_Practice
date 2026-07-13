/*
Link: https://leetcode.com/problems/trapping-rain-water/description/
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
*/

package main

import "fmt"

func trapWithExtraSpace(height []int) int {

	leftTallest := make([]int, len(height))
	rightTallest := make([]int, len(height))

	leftTallest[0] = height[0]
	rightTallest[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height)-1; i++ {
		leftTallest[i] = max(height[i], leftTallest[i-1])
		rightTallest[len(height)-i-1] = max(rightTallest[len(height)-i], height[len(height)-i-1])
	}

	result := 0

	for i := 1; i < len(height)-1; i++ {
		val := min(leftTallest[i], rightTallest[i]) - height[i]
		if val > 0 {
			result += val
		}
	}

	return result

}

func trap(height []int) int {
	left := 0
	leftMax := height[left]
	right := len(height) - 1
	rightMax := height[right]
	count := 0
	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				count += leftMax - height[left]
			}
		} else {
			right--
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				count += rightMax - height[right]
			}
		}
	}

	return count
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
