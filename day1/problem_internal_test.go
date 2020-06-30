package day1

import (
	"testing"
)

type testMassCalcs struct {
	mass     int
	expected int
}

var probOneSamples = []testMassCalcs{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestSampleProbOne(t *testing.T) {
	for _, test := range probOneSamples {
		got, expected := core(test.mass), test.expected
		if got != expected {
			t.Errorf("calc(%d) = %d; want %d", test.mass, got, expected)
		}
	}
}

var probTwoSamples = []testMassCalcs{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestSampleProbTwo(t *testing.T) {
	for _, test := range probTwoSamples {
		got, expected := coreExtension(test.mass), test.expected
		if got != expected {
			t.Errorf("calc(%d) = %d; want %d", test.mass, got, expected)
		}
	}
}
