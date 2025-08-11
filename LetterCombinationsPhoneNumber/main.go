package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	digitNums := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	RunUnitTest("23", digitNums, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	RunUnitTest("", digitNums, []string{})
	RunUnitTest("2", digitNums, []string{"a", "b", "c"})
}

func RunUnitTest(digits string, mapping map[int][]string, answer []string) {
	combinations := make([]string, 0)
	combinations = backtrack(digits, mapping, combinations, 0, "")
	if reflect.DeepEqual(combinations, answer) {
		fmt.Printf("\nFinal Result: %v", combinations)
	} else {
		fmt.Printf("\n\n!!!!!!!!!!!!!!!!!!INCORRECT RESULT: Combinations: %v, EXPECTED:%v", combinations, answer)
	}
}

func backtrack(digits string, mapping map[int][]string, combinations []string, index int, current string) []string {
	if index >= len(digits) {
		if len(current) > 0 {
			combinations = append(combinations, current)
		}
		return combinations
	}

	i, _ := strconv.Atoi(string(digits[index]))

	for _, letter := range mapping[i] {
		current += letter
		combinations = backtrack(digits, mapping, combinations, index+1, current)
		current = current[:len(current)-1]
	}

	return combinations
}
