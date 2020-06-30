package day2

import (
	"reflect"

	"../lib/optcode"
)

type kv map[int]int
type path struct {
}

func ProbOne() int {
	intcode := optcode.GetIntcode(currentDir() + "/problem_1.txt")
	return optcode.IntcodeInterpreter(intcode, kv{1: 12, 2: 2})[0]
}

func ProbTwo() int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			intcode := optcode.GetIntcode(currentDir() + "/problem_1.txt")
			if optcode.IntcodeInterpreter(intcode, kv{1: i, 2: j})[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1 // error occured, return bad value
}

func currentDir() string {
	path := reflect.TypeOf(path{}).PkgPath()
	if path[0] == '_' {
		path = path[1:]
	}
	return path
}
