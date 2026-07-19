/*
Link: https://neetcode.io/problems/eating-bananas/question?list=neetcode150

You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.

Return the minimum integer k such that you can eat all the bananas within h hours.

Example 1:

Input: piles = [1,4,3,2], h = 9

Output: 2
Explanation: With an eating rate of 2, you can eat the bananas in 6 hours. With an eating rate of 1, you would need 10 hours to eat all the bananas (which exceeds h=9), thus the minimum eating rate is 2.

Example 2:

Input: piles = [25,10,23,4], h = 4

Output: 25
Constraints:

1 <= piles.length <= 1,000
piles.length <= h <= 1,000,000
1 <= piles[i] <= 1,000,000,000


*/

package main

import (
	"fmt"
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	maxi := slices.Max(piles)

	if h == len(piles) {
		return maxi
	}

	l := 0
	r := maxi

	minimumCuts := math.MaxFloat64

	for l <= r {
		mid := l + (r-l)/2
		if mid == 0 {
			break
		}
		hoursNeeded := 0.0
		for _, item := range piles {
			if mid > 0 {
				val := math.Ceil(float64(item) / float64(mid))
				hoursNeeded += val
			}
		}

		if int(hoursNeeded) <= h {
			minimumCuts = min(minimumCuts, float64(mid))
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int(minimumCuts)
}

func main() {
	piles := []int{1, 4, 3, 2}
	h := 9

	fmt.Println(minEatingSpeed(piles, h))
}
