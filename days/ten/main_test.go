package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayten_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 7 {
		t.Errorf("Expected 7, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 33 {
		t.Errorf("Expected 33, got %d", p2)
	}
}

func TestDayP2(t *testing.T) {
	testData := lib.ReadFile("inputs/dayten_test.txt")

	p2 := solvePartTwo(testData)
	if p2 != 33 {
		t.Errorf("Expected 33, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayten.txt")

	p1 := solvePartOne(data)
	if p1 != 520 {
		t.Errorf("Expected 520, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 20626 {
		t.Errorf("Expected 20626, got %d", p2)
	}
}
