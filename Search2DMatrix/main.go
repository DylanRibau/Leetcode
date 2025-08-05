package main

import "fmt"

func main() {
	RunUnitTest([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true)
	RunUnitTest([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false)
	RunUnitTest([][]int{{1}}, 1, true)
	RunUnitTest([][]int{{1}}, 2, false)
	RunUnitTest([][]int{{1}, {3}}, 1, true)
	RunUnitTest([][]int{{1}, {3}}, 3, true)
	RunUnitTest([][]int{{1}}, 0, false)
	RunUnitTest([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 30, true)
}

func RunUnitTest(matrix [][]int, target int, answer bool) {
	exists := searchMatrix(matrix, target)
	if exists == answer {
		fmt.Printf("\nFinal Result: (%v)", exists)
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: %v, Expected: %v", exists, answer)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	y := searchY(matrix, target, 0, len(matrix)-1)
	if y < 0 {
		return false
	}
	x := searchX(matrix, y, target, 0, len(matrix[y]))
	return x >= 0
}

func searchY(matrix [][]int, target int, min int, max int) int {
	index := (min + max) / 2
	middle := matrix[index][0]

	if middle == target || (middle <= target && index+1 >= len(matrix)) || (index+1 < len(matrix) && (matrix[index+1][0] > target && middle < target)) {
		return index
	}

	if min > max {
		if target >= middle && (target <= matrix[index][len(matrix[index])-1] || index <= 0) {
			return index
		}
		return -1
	}

	if middle < target && matrix[index+1][0] <= target {
		return searchY(matrix, target, index+1, max)
	}
	return searchY(matrix, target, min, index-1)
}

func searchX(matrix [][]int, y int, target int, min int, max int) int {
	if min > max || min < 0 {
		return -1
	}

	index := (min + max) / 2
	if index >= len(matrix[y]) {
		return -1
	}
	middle := matrix[y][index]

	if middle == target {
		return index
	}

	if middle < target {
		return searchX(matrix, y, target, index+1, max)
	}
	return searchX(matrix, y, target, min, index-1)
}
