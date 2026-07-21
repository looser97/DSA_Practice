/*
Link: https://neetcode.io/problems/daily-temperatures/question?list=neetcode150
You are given an array of integers temperatures where temperatures[i] represents the daily temperatures on the ith day.

Return an array result where result[i] is the number of days after the ith day before a warmer temperature appears on a future day. If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

Example 1:

Input: temperatures = [30,38,30,36,35,40,28]

Output: [1,4,1,2,1,0,0]
Example 2:

Input: temperatures = [22,21,20]

Output: [0,0,0]
Constraints:

1 <= temperatures.length <= 1000.
1 <= temperatures[i] <= 100
*/

package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))

	stack := []int{}

	for i, item := range temperatures {
		if len(stack) > 0 {
			for len(stack) > 0 && item > temperatures[stack[len(stack)-1]] {
				results[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		} else {
			stack = append(stack, i)
		}
	}

	return results
}

func main() {
	temperatures := []int{22, 21, 20}

	fmt.Println(dailyTemperatures(temperatures))
}
