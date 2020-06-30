package day1

import (
	"reflect"
	"strconv"

	"../lib/std"
)

type path struct {
}

type fn func(int) int

func ProbOne() int {
	return execute(core, "1")
}

func ProbTwo() int {
	return execute(coreExtension, "2")
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

func execute(f fn, prob string) int {
	path := reflect.TypeOf(path{}).PkgPath()
	if path[0] == '_' {
		path = path[1:]
	}
	sum := 0
	for _, line := range std.ReadFile(path + "/problem_" + prob + ".txt") {
		val, err := strconv.Atoi(line)
		std.Check(err)
		sum += f(val)
	}
	return sum
}
