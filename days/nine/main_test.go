package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/daynine_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 50 {
		t.Errorf("Expected 50, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 24 {
		t.Errorf("Expected 24, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/daynine.txt")

	p1 := solvePartOne(data)
	if p1 != 4781546175 {
		t.Errorf("Expected 4781546175, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 2185817796 {
		t.Errorf("Expected 2185817796, got %d", p2)
	}
}
