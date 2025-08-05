package main

import "fmt"

func main() {
	RunUnitTest([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4)
	RunUnitTest([]int{4, 5, 6, 7, 0, 1, 2}, 3, -1)
	RunUnitTest([]int{1}, 0, -1)
	RunUnitTest([]int{1, 3}, 0, -1)
	RunUnitTest([]int{3, 1}, 1, 1)
	RunUnitTest([]int{5, 1, 3}, 5, 0)
	RunUnitTest([]int{1, 3}, 3, 1)
	RunUnitTest([]int{4, 5, 6, 7, 0, 1, 2}, 5, 1)
	RunUnitTest([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4)
	RunUnitTest([]int{9, 0, 2, 7, 8}, 3, -1)
}

func RunUnitTest(nums []int, target int, answer int) {
	index := search(nums, target)
	if index == answer {
		fmt.Printf("\nFinal Result: %d", index)
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: INDEX: %v, EXPECTED:%v", index, answer)
	}
}

func search(nums []int, target int) int {
	first := nums[0]
	fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nCalling binSearch. Nums = %v, Target = %d \n", nums, target)
	return binSearch(nums, target, first, 0, len(nums)-1)
}

func binSearch(nums []int, target int, first int, start int, end int) int {
	fmt.Printf("\n RESULTS - target: %v, first: %d, start: %d, end: %d", target, first, start, end)
	middle := (start + end) / 2
	current := nums[middle]
	fmt.Printf(", current: %v, middle %v, target: %v", current, middle, target)

	if current == target {
		return middle
	}
	if start == end || end < 0 || start > end {
		return -1
	}

	// fmt.Printf("%v", "v1")
	if (current < target && target < first) || (current > target && first > target && current >= first) || (current < target && target > first && !(current < first)) {
		return binSearch(nums, target, first, middle+1, end) // go right
	} else {
		return binSearch(nums, target, first, start, middle-1) // go left
	}
}
