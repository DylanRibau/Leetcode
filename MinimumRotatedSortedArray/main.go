package main

import "fmt"

func main() {
	RunUnitTest([]int{3, 4, 5, 1, 2}, 1)
	RunUnitTest([]int{3, 4, 5, 6, 1, 2}, 1)
	RunUnitTest([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	RunUnitTest([]int{11, 13, 15, 17}, 11)
	RunUnitTest([]int{5, 1, 2, 3, 4}, 1)
}

func RunUnitTest(nums []int, target int) {
	minimum := findMin(nums)
	if minimum == target {
		fmt.Printf("\nFinal Result: (%v)", minimum)
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: %v, Expected: %v", minimum, target)
	}
}

func findMin(nums []int) int {
	return search(nums, 0, len(nums))
}

func search(nums []int, min int, max int) int {
	index := (min + max) / 2
	num := nums[index]

	if min >= max || (index != 0 && num < nums[index-1]) {
		return num
	}

	if num >= nums[0] && num > nums[len(nums)-1] {
		return search(nums, index+1, max)
	}
	return search(nums, min, index-1)
}
