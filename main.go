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
	fmt.Println()
	fmt.Println("Day 3")
	fmt.Println("===========")
	day_3_data := day3_read_data()
	fmt.Println("Problem 1:", day3_1(day_3_data, 2))
	fmt.Println("Problem 2:", day3_2(day_3_data, 2))
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

func intcodeInterpreter(intcode []int, noun int, verb int) []int {
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
	return intcodeInterpreter(intcode, noun, verb)
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

type Coordinate struct {
	x int
	y int
}

func day3_read_data() map[Coordinate]map[int]int {
	wires := readFile("./day3/problem_1.txt")
	diagram := make(map[Coordinate]map[int]int)
	for i, wire := range wires {
		pos := Coordinate{0, 0}
		steps := 0
		for _, move := range strings.Split(wire, ",") {
			direction := string(move[0])
			num, err := strconv.Atoi(move[1:])
			check(err)
			switch direction {
			case "L":
				for num > 0 {
					steps++
					pos.x--
					num--
					diagram = record_steps_at_coord(diagram, pos, i, steps)
				}
			case "R":
				for num > 0 {
					steps++
					pos.x++
					num--
					diagram = record_steps_at_coord(diagram, pos, i, steps)
				}
			case "U":
				for num > 0 {
					steps++
					pos.y++
					num--
					diagram = record_steps_at_coord(diagram, pos, i, steps)
				}
			case "D":
				for num > 0 {
					steps++
					pos.y--
					num--
					diagram = record_steps_at_coord(diagram, pos, i, steps)
				}
			}

		}
	}
	return diagram
}

func record_steps_at_coord(data map[Coordinate]map[int]int, coord Coordinate, wire int, steps int) map[Coordinate]map[int]int {
	if data[coord] == nil {
		data[coord] = make(map[int]int)
	}
	if _, exists := data[coord][wire]; !exists {
		data[coord][wire] = steps
	}
	return data
}

func day3_1(data map[Coordinate]map[int]int, intersects int) int {
	var intersect_distance []int
	for k := range data {
		if len(data[k]) >= intersects {
			x_distance := abs(k.x)
			y_distance := abs(k.y)
			intersect_distance = append(intersect_distance, x_distance+y_distance)
		}
	}
	return min(intersect_distance)
}

func day3_2(data map[Coordinate]map[int]int, intersects int) int {
	var required_steps []int
	for k := range data {
		if len(data[k]) == intersects {
			total_steps := 0
			for i := 0; i < intersects; i++ {
				total_steps += abs(data[k][i])
			}
			required_steps = append(required_steps, total_steps)
		}
	}
	return min(required_steps)
}

func min(values []int) int {
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
