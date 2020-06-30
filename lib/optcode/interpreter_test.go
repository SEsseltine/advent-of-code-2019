package optcode_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"../optcode"
)

type kv map[int]int
type testSamples struct {
	code []int
	kv
	expected []int
}
type testExacts struct {
	code []int
	kv
	expected string
}

var day2ProbOneSamples = []*testSamples{
	{[]int{1, 0, 0, 0, 99}, kv{}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, kv{}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, kv{}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, kv{}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestBasicOptcodes(t *testing.T) {
	for _, test := range day2ProbOneSamples {
		got := optcode.IntcodeInterpreter(test.code, test.kv)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("day2.ProbOne() = %d; want %d", got, test.expected)
		}
	}
}

var day5ProbTwoSamples = []*testExacts{
	{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, kv{-1: -1}, "999"},
	{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, kv{-1: 7}, "999"},
	{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, kv{-1: 8}, "1000"},
	{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, kv{-1: 9}, "1001"},
	{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, kv{-1: 100}, "1001"},
}

func TestConditionalOptcodes(t *testing.T) {
	for _, test := range day5ProbTwoSamples {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		_ = optcode.IntcodeInterpreter(test.code, test.kv)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		outStr := strings.TrimSpace(string(out))
		if outStr != test.expected {
			t.Errorf("Unexpected output! Got %s, expected %s", outStr, test.expected)
		}
	}
}
