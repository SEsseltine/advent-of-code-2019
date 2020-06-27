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
			count += 1
		}
	}
	return count
}

func ProbTwo() int {
	count := 0
	for i := MIN_VALUE; i <= MAX_VALUE; i++ {
		str := strconv.Itoa(i)
		if validate(str) && doubleNotInLargerSet(str) {
			count += 1
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
	// TODO: Implement requirement ensuring there is a double not part of a larger set
	return true
}

func increasing(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
