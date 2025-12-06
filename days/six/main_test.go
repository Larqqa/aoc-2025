package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/daysix_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 4277556 {
		t.Errorf("Expected 4277556, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 3263827 {
		t.Errorf("Expected 3263827, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/daysix.txt")

	p1 := solvePartOne(data)
	if p1 != 5524274308182 {
		t.Errorf("Expected 5524274308182, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 8843673199391 {
		t.Errorf("Expected 8843673199391, got %d", p2)
	}
}
