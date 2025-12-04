package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayfour_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 13 {
		t.Errorf("Expected 13, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 43 {
		t.Errorf("Expected 43, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayfour.txt")

	p1 := solvePartOne(data)
	if p1 != 1537 {
		t.Errorf("Expected 1537, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 8707 {
		t.Errorf("Expected 8707, got %d", p2)
	}
}
