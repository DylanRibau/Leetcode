package main

import (
	"fmt"
	"reflect"
)

func main() {
	RunUnitTest(3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	RunUnitTest(1, []string{"()"})
}

func RunUnitTest(n int, answer []string) {
	combinations := make([]string, 0)
	combinations = backtrack(n, combinations, "")
	if reflect.DeepEqual(combinations, answer) {
		fmt.Printf("\nFinal Result: %v", combinations)
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: Combinations: %v, EXPECTED:%v", combinations, answer)
	}
}

func backtrack(n int, combinations []string, current string) []string {
	if len(current) >= n*2 {
		combinations = append(combinations, current)
		return combinations
	}

	for _, i := range []string{"(", ")"} {
		if isValid(n, current, i) {
			current += i
			combinations = backtrack(n, combinations, current)
			current = current[:len(current)-1]
		}
	}

	return combinations
}

func isValid(n int, current string, char string) bool {
	if len(current) > n*2 {
		return false
	}

	open := 0
	close := 0
	for _, i := range current {
		if string(i) == "(" {
			open++
			continue
		}
		close++
	}

	if (char == "(" && open < n) || (char == ")" && close < open) {
		return true
	}

	return false
}
