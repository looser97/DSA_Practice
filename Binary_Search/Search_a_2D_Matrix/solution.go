/*
Link: https://neetcode.io/problems/search-2d-matrix/question
You are given an m x n 2-D integer array matrix and an integer target.

Each row in matrix is sorted in non-decreasing order.
The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.

Can you write a solution that runs in O(log(m * n)) time?

Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 10

Output: true
Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 15

Output: false
Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10000 <= matrix[i][j], target <= 10000
*/

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	noOfRows := len(matrix)
	noOfColums := len(matrix[0])

	t := 0
	b := noOfRows - 1

	selectedRow := -1

	for t <= b {
		mid := t + (b-t)/2
		if matrix[mid][0] <= target && matrix[mid][noOfColums-1] >= target {
			selectedRow = mid
			break
		}

		if matrix[mid][0] > target {
			b = mid - 1
		} else {
			t = mid + 1
		}
	}

	if selectedRow == -1 {
		return false
	}

	l := 0
	r := noOfColums - 1

	for l <= r {
		mid := l + (r-l)/2

		if matrix[selectedRow][mid] == target {
			return true
		}

		if matrix[selectedRow][mid] > target {
			r = mid - 1
		} else {
			l = l + 1
		}
	}

	return false

}

func main() {
	metrix := [][]int{
		{1, 2, 4, 8},
		{10, 11, 12, 13},
		{14, 20, 30, 40},
	}
	target := 10

	fmt.Println(searchMatrix(metrix, target))
}
