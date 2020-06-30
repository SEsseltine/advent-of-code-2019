package day1_test

import (
	"testing"

	"../day1"
)

func TestProbOne(t *testing.T) {
	got, expected := day1.ProbTwo(), 4845669
	if got != expected {
		t.Errorf("day1.ProbOne() = %d; want %d", got, expected)
	}
}

func TestProbTwo(t *testing.T) {
	got, expected := day1.ProbTwo(), 4845669
	if got != expected {
		t.Errorf("day1.ProbOne() = %d; want %d", got, expected)
	}
}
