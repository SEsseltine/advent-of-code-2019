package main

import (
	"./day1"
	"./day2"
	"./day3"
	"./day4"
	"./day5"
)

func main() {
	println("Day 1")
	println("===========")
	println("Problem 1:", day1.ProbOne())
	println("Problem 2:", day1.ProbTwo())
	println()
	println("Day 2")
	println("===========")
	println("Problem 1:", day2.ProbOne())
	println("Problem 2:", day2.ProbTwo())
	println()
	println("Day 3")
	println("===========")
	println("Problem 1:", day3.ProbOne())
	println("Problem 2:", day3.ProbTwo())
	println()
	println("Day 4")
	println("===========")
	println("Problem 1:", day4.ProbOne())
	println("Problem 1:", day4.ProbTwo())
	println()
	println("Day 5")
	println("===========")
	print("Problem 1: ")
	day5.ProbOne()
	println()
	print("Problem 2: ")
	day5.ProbTwo()
	println()
}
