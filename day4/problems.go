package day4

import (
	"strconv"

	"../lib/std"
)

// Given
var MIN_VALUE, MAX_VALUE int = 165432, 707912

func ProbOne() int {
	count := 0
	for i := MIN_VALUE; i <= MAX_VALUE; i++ {
		if validate(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}

func ProbTwo() int {
	count := 0
	var reduced []string
	for i := MIN_VALUE; i <= MAX_VALUE; i++ {
		str := strconv.Itoa(i)
		if validate(str) {
			reduced = append(reduced, str)
		}
	}
	for _, str := range reduced {
		if doubleNotInLargerSet(str) {
			count++
		}
	}
	return count
}

func validate(input string) bool {
	intValue, err := strconv.Atoi(input)
	std.Check(err)
	return len(input) == 6 &&
		hasDouble(input) &&
		increasing(input) &&
		intValue >= MIN_VALUE &&
		intValue <= MAX_VALUE
}

func hasDouble(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			return true
		}
	}
	return false
}

func doubleNotInLargerSet(input string) bool {
	start := 0
	for start < len(input)-2 {
		end := start + 1
		for end < len(input) && input[start] == input[end] {
			end++
		}
		if end-start > 2 {
			input = input[:start] + input[end:]
		} else {
			start++
		}
	}
	return hasDouble(input)
}

func increasing(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
