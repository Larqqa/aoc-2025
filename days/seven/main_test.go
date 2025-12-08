package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/dayseven_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 21 {
		t.Errorf("Expected 21, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 40 {
		t.Errorf("Expected 40, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayseven.txt")

	p1 := solvePartOne(data)
	if p1 != 1698 {
		t.Errorf("Expected 1698, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 95408386769474 {
		t.Errorf("Expected 95408386769474, got %d", p2)
	}
}
