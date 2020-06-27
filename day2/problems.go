package day2

import (
	"../lib/optcode"
)

func ProbOne() int {
	return optcode.IntcodeInterpreter(optcode.GetIntcode("./day2/problem_1.txt"), 12, 2)[0]
}

func ProbTwo() int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if optcode.IntcodeInterpreter(optcode.GetIntcode("./day2/problem_1.txt"), i, j)[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1 // error occured, return bad value
}
