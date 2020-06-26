package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 1:", day1())
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

func day1_algorithm(mass int) int {
	return ((mass / 3) - 2)
}

func day1() int {
	sum := 0
	for _, line := range readFile("./day1_given.txt") {
		val, err := strconv.Atoi(line)
		check(err)
		sum += day1_algorithm(val)
	}
	return sum
}
