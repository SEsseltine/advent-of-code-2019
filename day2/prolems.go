package day2

import (
	"strconv"
	"strings"

	"../lib/std"
)

func ProbOne() int {
	return intcodeInterpreter(getIntcode(), 12, 2)[0]
}

func ProbTwo() int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if intcodeInterpreter(getIntcode(), i, j)[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1 // error occured, return bad value
}

func intcodeInterpreter(intcode []int, noun int, verb int) []int {
	intcode[1] = noun
	intcode[2] = verb
	for pos := 0; pos < len(intcode); pos += 4 {
		if intcode[pos] == 1 {
			intcode[intcode[pos+3]] = intcode[intcode[pos+1]] + intcode[intcode[pos+2]]
		} else if intcode[pos] == 2 {
			intcode[intcode[pos+3]] = intcode[intcode[pos+1]] * intcode[intcode[pos+2]]
		} else if intcode[pos] == 99 {
			return intcode
		} else {
			panic("Something went wrong reading the intcode, bad opcode instruction")
		}
	}
	return intcode
}

func getIntcode() []int {
	lines := std.ReadFile("./day2/problem_1.txt")
	var intcode []int
	for _, val := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(val)
		std.Check(err)
		intcode = append(intcode, num)
	}
	return intcode
}
