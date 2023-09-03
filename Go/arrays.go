package main

import (
	"time"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
func RemoveDuplicates(nums []int) int {
	defer timeTrack(time.Now(), "removeDuplicates")
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	nums = nums[:i+1]

	return i + 1
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
// func MaxProfit(prices []int)  int {

// }

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/
func isValidSudoku(board [][]byte) bool {
	defer timeTrack(time.Now(), "isValidSudoku")
	var row [9][9]bool
	var col [9][9]bool
	var box [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row[i][j] = false
			col[i][j] = false
			box[i][j] = false
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] != '.' {
				num := board[i][j] - '1'
				k := (i/3)*3 + j/3
				if row[i][num] || col[j][num] || box[k][num] {
					return false
				}
				row[i][num] = true
				col[j][num] = true
				box[k][num] = true
			}
		}
	}
	return true
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/770/
func rotate(matrix [][]int) {
	defer timeTrack(time.Now(), "rotate image")

	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]         // left -> top
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1] // bottom -> left
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1] // right -> bottom
			matrix[j][n-i-1] = temp                 // top -> right
		}

	}
}
