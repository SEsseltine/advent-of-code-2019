package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println("===========")
	fmt.Println("Problem 1:", day1_1())
	fmt.Println("Problem 2:", day1_2())
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
