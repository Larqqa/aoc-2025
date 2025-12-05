package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayfive_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 3 {
		t.Errorf("Expected 3, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 14 {
		t.Errorf("Expected 14, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayfive.txt")

	p1 := solvePartOne(data)
	if p1 != 761 {
		t.Errorf("Expected 761, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 345755049374932 {
		t.Errorf("Expected 345755049374932, got %d", p2)
	}
}
