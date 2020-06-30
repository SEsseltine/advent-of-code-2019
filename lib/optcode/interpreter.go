package optcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"../std"
)

type binaryOperator func(a int, b int) int
type unaryOperator func(a int) int

var add binaryOperator = func(a, b int) int { return a + b }
var mul binaryOperator = func(a, b int) int { return a * b }
var leq binaryOperator = func(a, b int) int {
	if a < b {
		return 1
	} else {
		return 0
	}
}
var eq binaryOperator = func(a, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}

var binOpMap = map[int]binaryOperator{1: add, 2: mul, 7: leq, 8: eq}

func IntcodeInterpreter(intcode []int, overwrites map[int]int) []int {
	for index, value := range overwrites {
		if index > 0 {
			intcode[index] = value
		}
	}
	pos := 0
	for intcode[pos] != 99 {
		instruction, firstInputMode, secondInputMode, thirdInputMode := getOpcodeParamModes(intcode[pos])
		switch instruction {
		case 1, 2, 7, 8:
			a, err := getValue(&intcode, firstInputMode, pos+1)
			std.Check(err)
			b, err := getValue(&intcode, secondInputMode, pos+2)
			std.Check(err)
			std.Check(setValue(&intcode, pos+3, binOpMap[instruction](a, b)))
			pos += 4
		case 3:
			std.Check(setValue(&intcode, pos+1, overwrites[-1]))
			pos += 2
		case 4:
			val, err := getValue(&intcode, firstInputMode, pos+1)
			std.Check(err)
			fmt.Print(val, " ")
			pos += 2
		case 5:
			a, err := getValue(&intcode, firstInputMode, pos+1)
			std.Check(err)
			b, err := getValue(&intcode, secondInputMode, pos+2)
			std.Check(err)
			if a != 0 {
				pos = b
			} else {
				pos += 3
			}
		case 6:
			a, err := getValue(&intcode, firstInputMode, pos+1)
			std.Check(err)
			b, err := getValue(&intcode, secondInputMode, pos+2)
			std.Check(err)
			if a == 0 {
				pos = b
			} else {
				pos += 3
			}
		case 99:
			return intcode
		default:
			println(thirdInputMode)
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

func getOpcodeParamModes(opcode int) (int, int, int, int) {
	instruction := opcode % 100
	opcode /= 100
	firstInputMode := opcode % 10
	opcode /= 10
	secondInputMode := opcode % 10
	opcode /= 10
	return instruction, firstInputMode, secondInputMode, opcode % 10
}

func getValue(intcode *[]int, mode int, argIndex int) (int, error) {
	switch mode {
	case 0:
		return (*intcode)[(*intcode)[argIndex]], nil
	case 1:
		return (*intcode)[argIndex], nil
	}
	return 0, errors.New("Argument mode not recognized")
}

func setValue(intcode *[]int, argIndex int, val int) error {
	(*intcode)[(*intcode)[argIndex]] = val
	return nil
}
