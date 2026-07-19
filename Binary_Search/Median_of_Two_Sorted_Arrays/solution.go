/*
Link: https://neetcode.io/problems/median-of-two-sorted-arrays/question?list=neetcode150
You are given two integer arrays nums1 and nums2 of size m and n respectively, where each is sorted in ascending order. Return the median value among all elements of the two arrays.

Your solution must run in
O
(
l
o
g
(
m
+
n
)
)
O(log(m+n)) time.

Example 1:

Input: nums1 = [1,2], nums2 = [3]

Output: 2.0
Explanation: Among [1, 2, 3] the median is 2.

Example 2:

Input: nums1 = [1,3], nums2 = [2,4]

Output: 2.5
Explanation: Among [1, 2, 3, 4] the median is (2 + 3) / 2 = 2.5.

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
-10^6 <= nums1[i], nums2[i] <= 10^6

*/

package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	totalLen := (len(nums1) + len(nums2))

	medianPos := (totalLen + 1) / 2
	l := 0
	r := len(nums1)
	for l <= r {
		mid1 := l + (r-l)/2
		mid2 := medianPos - mid1

		l1 := math.MinInt32
		l2 := math.MinInt32
		r1 := math.MaxInt32
		r2 := math.MaxInt32

		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}
		if mid1 < len(nums1) {
			r1 = nums1[mid1]
		}

		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}
		if mid2 < len(nums2) {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			if totalLen%2 == 0 {
				return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2.0
			} else {
				return float64(max(float32(l1), float32(l2)))
			}
		} else if l1 > r2 {
			r = mid1 - 1
		} else {
			l = mid1 + 1
		}
	}
	return 0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
