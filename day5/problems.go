package day5

import (
	"reflect"

	"../lib/optcode"
)

type kv map[int]int
type path struct {
}

func ProbOne() {
	intcode := optcode.GetIntcode(currentDir() + "/problem_1.txt")
	optcode.IntcodeInterpreter(intcode, kv{-1: 1})
}

func ProbTwo() {
	intcode := optcode.GetIntcode(currentDir() + "/problem_1.txt")
	optcode.IntcodeInterpreter(intcode, kv{-1: 5})
}

func currentDir() string {
	path := reflect.TypeOf(path{}).PkgPath()
	if path[0] == '_' {
		path = path[1:]
	}
	return path
}
