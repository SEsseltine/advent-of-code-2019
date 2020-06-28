package optcode_test

import (
	"reflect"
	"testing"

	"../optcode"
)

type kv map[int]int
type testSamples struct {
	code []int
	kv
	expected []int
}

var day2ProbOneSamples = []*testSamples{
	{[]int{1, 0, 0, 0, 99}, kv{}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, kv{}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, kv{}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, kv{}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestProbOne(t *testing.T) {
	for _, test := range day2ProbOneSamples {
		got := optcode.IntcodeInterpreter(test.code, test.kv)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("day1.ProbOne() = %d; want %d", got, test.expected)
		}
	}
}
