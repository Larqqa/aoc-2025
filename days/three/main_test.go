package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDayThree(t *testing.T) {
	testData := lib.ReadFile("inputs/daythree_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 357 {
		t.Errorf("Expected 357, got %d", p1)
	}

	p2 := solvePartTwo(testData)
	if p2 != 3121910778619 {
		t.Errorf("Expected 3121910778619, got %d", p2)
	}
}

func TestP2(t *testing.T) {
	p2 := solvePartTwo("1111111111111111111111818181911112111")
	if p2 != 888911112111 {
		t.Errorf("Expected 888911112111, got %d", p2)
	}
}

func TestP2Small(t *testing.T) {
	p2 := solvePartTwo("1212123213")
	if p2 != 323 {
		t.Errorf("Expected 323, got %d", p2)
	}
}

func TestP2Large(t *testing.T) {
	p2 := solvePartTwo("2252522232121122125212322341424262435212421332333533223124122242222222112222222222423222112211212132")
	if p2 != 12345 {
		t.Errorf("Expected 1234, got %d", p2)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/daythree.txt")

	p1 := solvePartOne(data)
	if p1 != 17142 {
		t.Errorf("Expected 17142, got %d", p1)
	}

	p2 := solvePartTwo(data)
	if p2 != 169935154100102 {
		t.Errorf("Expected 169935154100102, got %d", p2)
	}
}
