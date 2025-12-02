package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDayTwo(t *testing.T) {
	testData := lib.ReadFile("inputs/daytwo_test.txt")
	p1 := solvePartOne(testData)

	if p1 != 1227775554 {
		t.Errorf("Expected 1227775554, got %d", p1)
	}

	p2 := solvePartTwo(testData)

	if p2 != 4174379265 {
		t.Errorf("Expected 4174379265, got %d", p2)
	}
}

func TestSimpleCase(t *testing.T) {
	p1 := solvePartOne(`11-22`)

	if p1 != 33 {
		t.Errorf("Expected 33, got %d", p1)
	}

	p2 := solvePartTwo(`11-22`)

	if p2 != 33 {
		t.Errorf("Expected 33, got %d", p2)
	}
}

func TestHardCase(t *testing.T) {
	p1 := solvePartOne(`1188511880-1188511890`)

	if p1 != 1188511885 {
		t.Errorf("Expected 1188511885, got %d", p1)
	}

	p2 := solvePartTwo(`1188511880-1188511890`)

	if p2 != 1188511885 {
		t.Errorf("Expected 1188511885, got %d", p2)
	}
}

func TestTwoCase(t *testing.T) {
	{
		p2 := solvePartTwo(`222220-222224`)

		if p2 != 222222 {
			t.Errorf("Expected 222222, got %d", p2)
		}
	}
	{
		p2 := solvePartTwo(`998-1012`)

		if p2 != 2009 {
			t.Errorf("Expected 2009, got %d", p2)
		}
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/daytwo.txt")

	p1 := solvePartOne(data)
	if p1 != 35367539282 {
		t.Errorf("Expected 35367539282, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 45814076230 {
		t.Errorf("Expected 45814076230, got %d", p2)
	}
}
