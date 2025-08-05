package main

import "fmt"

func main() {
	RunUnitTest([]int{5, 7, 7, 8, 8, 10}, 8, 3, 4)
	RunUnitTest([]int{5, 7, 7, 8, 8, 10}, 6, -1, -1)
	RunUnitTest([]int{}, 0, -1, -1)
	RunUnitTest([]int{1}, 1, 0, 0)
}

func RunUnitTest(nums []int, target int, answerMin int, answerMax int) {
	ranges := searchRange(nums, target)
	if ranges[0] == answerMin && ranges[1] == answerMax {
		fmt.Printf("\nFinal Result: (%v,%v)", ranges[0], ranges[1])
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: min: %v, max:%v EXPECTED: min:%v, max:%v", ranges[0], ranges[1], answerMin, answerMax)
	}
}

func searchRange(nums []int, target int) []int {
	index := binSearch(nums, target, 0, len(nums)-1)
	if index == -1 {
		return []int{-1, -1}
	}
	return MinMax(nums, index)
}

func binSearch(nums []int, target int, start int, end int) int {
	if end < 0 {
		return -1
	}
	index := (start + end) / 2
	middle := nums[index]

	if index >= len(nums) || index < 0 || start > end {
		return -1
	}
	if target == middle {
		return index
	}

	if target > middle {
		return binSearch(nums, target, index+1, end)
	} else {
		return binSearch(nums, target, start, index-1)
	}
}

func MinMax(nums []int, index int) []int {
	target := nums[index]
	return []int{getMin(nums, target, index), getMax(nums, target, index)}
}

func getMin(nums []int, target int, index int) int {
	if index < 0 {
		return 0
	}

	if nums[index] == target {
		return getMin(nums, target, index-1)
	}
	return index + 1
}

func getMax(nums []int, target int, index int) int {
	if index >= len(nums) {
		return len(nums) - 1
	}
	if nums[index] == target {
		return getMax(nums, target, index+1)
	}
	return index - 1
}
