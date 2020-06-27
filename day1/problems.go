package day1

import (
	"strconv"

	"../lib/std"
)

func ProbOne() int {
	sum := 0
	for _, line := range std.ReadFile("./day1/problem_1.txt") {
		val, err := strconv.Atoi(line)
		std.Check(err)
		sum += core(val)
	}
	return sum
}

func ProbTwo() int {
	sum := 0
	for _, line := range std.ReadFile("./day1/problem_2.txt") {
		val, err := strconv.Atoi(line)
		std.Check(err)
		sum += coreExtension(val)
	}
	return sum
}

func core(mass int) int {
	return ((mass / 3) - 2)
}

func coreExtension(mass int) int {
	calc := core(mass)
	if calc > 0 {
		return calc + coreExtension(calc)
	} else {
		return 0
	}
}
