package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println("===========")
	fmt.Println("Problem 1:", day1_1())
	fmt.Println("Problem 2:", day1_2())
	fmt.Println()
	fmt.Println("Day 2")
	fmt.Println("===========")
	x, y := 12, 2
	fmt.Println("Problem 1:", day2_1(x, y)[0])
	noun, verb := day2_2()
	fmt.Println("Problem 2:", 100*noun+verb)
}

func readFile(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	check(err)
	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func day1_1_algorithm(mass int) int {
	return ((mass / 3) - 2)
}

func day1_1() int {
	sum := 0
	for _, line := range readFile("./day1/problem_1.txt") {
		val, err := strconv.Atoi(line)
		check(err)
		sum += day1_1_algorithm(val)
	}
	return sum
}

func day1_2_algorithm(mass int) int {
	calc := ((mass / 3) - 2)
	if calc > 0 {
		return calc + day1_2_algorithm(calc)
	} else {
		return 0
	}
}

func day1_2() int {
	sum := 0
	for _, line := range readFile("./day1/problem_2.txt") {
		val, err := strconv.Atoi(line)
		check(err)
		sum += day1_2_algorithm(val)
	}
	return sum
}

func day2_1_algorithm(intcode []int, noun int, verb int) []int {
	// Reset the code as instructed
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

func day2_1(noun int, verb int) []int {
	lines := readFile("./day2/problem_1.txt")
	var intcode []int
	for _, val := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(val)
		check(err)
		intcode = append(intcode, num)
	}
	return day2_1_algorithm(intcode, noun, verb)
}

func day2_2() (int, int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if day2_1(i, j)[0] == 19690720 {
				return i, j
			}
		}
	}
	return -1, -1
}
