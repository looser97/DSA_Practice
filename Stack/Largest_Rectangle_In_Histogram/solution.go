/*
Link: https://neetcode.io/problems/largest-rectangle-in-histogram/question?list=neetcode150
You are given an array of integers heights where heights[i] represents the height of a bar. The width of each bar is 1.

Return the area of the largest rectangle that can be formed among the bars.

Note: This chart is known as a histogram.

Example 1:

Input: heights = [7,1,7,2,2,4]

Output: 8
Example 2:

Input: heights = [1,3,7]

Output: 7
Constraints:

1 <= heights.length <= 1000.
0 <= heights[i] <= 1000

*/

package main

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := [][2]int{}
	maxArea := 0

	heights = append(heights, 0)

	for i, item := range heights {
		if len(stack) == 0 || stack[len(stack)-1][1] <= item {
			stack = append(stack, [2]int{i, item})
		} else {
			indexToUse := i
			for len(stack) > 0 && stack[len(stack)-1][1] > item {
				topElement := stack[len(stack)-1]
				area := topElement[1] * (i - topElement[0])
				maxArea = max(area, maxArea)
				indexToUse = topElement[0]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, [2]int{indexToUse, item})
		}
	}
	return maxArea
}

func main() {
	heights := []int{7, 1, 7, 2, 2, 4}

	fmt.Println(largestRectangleArea(heights))

}
