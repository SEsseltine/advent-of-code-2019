package optcode

import (
	"strconv"
	"strings"

	"../std"
)

func IntcodeInterpreter(intcode []int, overwrites map[int]int) []int {
	for index, value := range overwrites {
		intcode[index] = value
	}
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

func GetIntcode(path string) []int {
	lines := std.ReadFile(path)
	var intcode []int
	for _, val := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(val)
		std.Check(err)
		intcode = append(intcode, num)
	}
	return intcode
}
