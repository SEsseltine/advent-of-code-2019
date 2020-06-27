package day3

import (
	"strconv"
	"strings"

	"../lib/std"
)

type coordinate struct {
	x int
	y int
}

func ProbOne() int {
	data := readData()
	minimumDistance := 0 // 0 is not possible so if 0 returns it is an error
	for k := range data {
		if len(data[k]) == 2 {
			distance := std.Abs(k.x) + std.Abs(k.y)
			if distance < minimumDistance || minimumDistance == 0 {
				minimumDistance = distance
			}
		}
	}
	return minimumDistance
}

func ProbTwo() int {
	data := readData()
	minimumSteps := 0 // 0 is not possible so if 0 returns it is an error
	for k := range data {
		if len(data[k]) == 2 {
			steps := 0
			for i := 0; i < 2; i++ {
				steps += std.Abs(data[k][i])
			}
			if steps < minimumSteps || minimumSteps == 0 {
				minimumSteps = steps
			}
		}
	}
	return minimumSteps
}

func readData() map[coordinate]map[int]int {
	wires := std.ReadFile("./day3/problem_1.txt")
	diagram := make(map[coordinate]map[int]int)
	for i, wire := range wires {
		pos := coordinate{0, 0}
		steps := 0
		for _, move := range strings.Split(wire, ",") {
			direction := string(move[0])
			num, err := strconv.Atoi(move[1:])
			std.Check(err)
			switch direction {
			case "L":
				for num > 0 {
					steps++
					pos.x--
					num--
					diagram = recordsStepsAtCoordinate(diagram, pos, i, steps)
				}
			case "R":
				for num > 0 {
					steps++
					pos.x++
					num--
					diagram = recordsStepsAtCoordinate(diagram, pos, i, steps)
				}
			case "U":
				for num > 0 {
					steps++
					pos.y++
					num--
					diagram = recordsStepsAtCoordinate(diagram, pos, i, steps)
				}
			case "D":
				for num > 0 {
					steps++
					pos.y--
					num--
					diagram = recordsStepsAtCoordinate(diagram, pos, i, steps)
				}
			}

		}
	}
	return diagram
}

func recordsStepsAtCoordinate(data map[coordinate]map[int]int, coord coordinate, wire int, steps int) map[coordinate]map[int]int {
	if data[coord] == nil {
		data[coord] = make(map[int]int)
	}
	if _, exists := data[coord][wire]; !exists {
		data[coord][wire] = steps
	}
	return data
}
