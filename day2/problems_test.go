package day2_test

import (
	"testing"

	"../day2"
)

func TestProbOne(t *testing.T) {
	got, expected := day2.ProbOne(), 5866663
	if got != expected {
		t.Errorf("day1.ProbOne() = %d; want %d", got, expected)
	}
}

func TestProbTwo(t *testing.T) {
	got, expected := day2.ProbTwo(), 4259
	if got != expected {
		t.Errorf("day1.ProbOne() = %d; want %d", got, expected)
	}
}
