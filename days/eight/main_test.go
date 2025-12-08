package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayeight_test.txt")

	p1 := solvePartOne(testData, 10)
	if p1 != 40 {
		t.Errorf("Expected 40, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 25272 {
		t.Errorf("Expected 25272, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayeight.txt")

	p1 := solvePartOne(data, 1000)
	if p1 != 24360 {
		t.Errorf("Expected 24360, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 2185817796 {
		t.Errorf("Expected 2185817796, got %d", p2)
	}
}
