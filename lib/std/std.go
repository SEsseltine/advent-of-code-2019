package std

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	Check(err)
	return data
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
