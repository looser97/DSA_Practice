/*

Link: https://leetcode.com/problems/container-with-most-water/description/
ou are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
*/

package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	area := 0

	for left < right {
		calculatedArea := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if calculatedArea > area {
			area = calculatedArea
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}

func main() {
	numbers := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(numbers))
}
