package day2

import (
	"../lib/optcode"
)

type kv map[int]int

func ProbOne() int {
	return optcode.IntcodeInterpreter(optcode.GetIntcode("day2/problem_1.txt"), kv{1: 12, 2: 2})[0]
}

func ProbTwo() int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			intcode := optcode.GetIntcode("day2/problem_1.txt")
			if optcode.IntcodeInterpreter(intcode, kv{1: i, 2: j})[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1 // error occured, return bad value
}
