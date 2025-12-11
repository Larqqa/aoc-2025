package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayeleven_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 5 {
		t.Errorf("Expected 5, got %d", p1)
	}

	testData2 := lib.ReadFile("inputs/dayelevenp2_test.txt")
	p2 := solvePartTwo(testData2)
	if p2 != 2 {
		t.Errorf("Expected 2, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayeleven.txt")

	p1 := solvePartOne(data)
	if p1 != 448 {
		t.Errorf("Expected 448, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 553204221431080 {
		t.Errorf("Expected 553204221431080, got %d", p2)
	}
}
