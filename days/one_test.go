package main

import (
	lib "aoc/2025"
	"testing"
)

func TestDayOne(t *testing.T) {
	testData := lib.ReadFile("inputs/dayone_test.txt")
	p1, p2 := solvePartOne(testData)

	if p1 != 3 {
		t.Errorf("Expected 3, got %d", p1)
	}
	if p2 != 6 {
		t.Errorf("Expected 6, got %d", p2)
	}
}

func TestLeftRot(t *testing.T) {
	{
		_, p2 := solvePartOne("L60")
		if p2 != 1 {
			t.Errorf("Expected 1, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("L1000")
		if p2 != 10 {
			t.Errorf("Expected 10, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("L490")
		if p2 != 5 {
			t.Errorf("Expected 5, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("L123")
		if p2 != 1 {
			t.Errorf("Expected 1, got %d", p2)
		}
	}
}

func TestRightRot(t *testing.T) {
	{
		_, p2 := solvePartOne("R60")
		if p2 != 1 {
			t.Errorf("Expected 1, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("R1000")
		if p2 != 10 {
			t.Errorf("Expected 10, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("R490")
		if p2 != 5 {
			t.Errorf("Expected 5, got %d", p2)
		}
	}
	{
		_, p2 := solvePartOne("R123")
		if p2 != 1 {
			t.Errorf("Expected 1, got %d", p2)
		}
	}
}

func TestRightLeftRot(t *testing.T) {
	{
		p1, p2 := solvePartOne(`L50
R60
L120
R120
L60`)

		if p1 != 2 {
			t.Errorf("Expected 2, got %d", p1)
		}

		if p2 != 4 {
			t.Errorf("Expected 4, got %d", p2)
		}
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/dayone.txt")
	p1, p2 := solvePartOne(data)

	if p1 != 1129 {
		t.Errorf("Expected 1129, got %d", p1)
	}
	if p2 != 6638 {
		t.Errorf("Expected 6638, got %d", p2)
	}
}
