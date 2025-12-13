package main

import (
	lib "aoc/2025"
	"fmt"
	"testing"
)

func TestDay(t *testing.T) {
	testData := lib.ReadFile("inputs/daytwelve_test.txt")

	p1 := solvePartOne(testData)
	if p1 != 2 {
		t.Errorf("Expected 2, got %d", p1)
	}
}

func TestCorrect(t *testing.T) {
	data := lib.ReadFile("inputs/daytwelve.txt")

	p1 := solvePartOne(data)
	if p1 != 541 {
		t.Errorf("Expected 541, got %d", p1)
	}
}


// PASS in 2.6s
func BenchmarkSolvePartOne(b *testing.B) {
	data := lib.ReadFile("inputs/daytwelve.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePartOne(data)
	}
}

func TestPlace(t *testing.T) {
	data := lib.ReadFile("inputs/daytwelve_test.txt")
	presents, regions := parse(data)

	placeTile(&regions[1].grid, presents[0].grid, lib.NewCoord(0, 0), '#')

	regions[1].grid.Print()

	fmt.Println(canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(0, 0)))
	fmt.Println(canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(5, 5)))

	fmt.Println()

	placeTile(&regions[1].grid, presents[0].grid, lib.NewCoord(3, 0), '#')
	placeTile(&regions[1].grid, presents[0].grid, lib.NewCoord(6, 0), '#')
	placeTile(&regions[1].grid, presents[0].grid, lib.NewCoord(9, 0), '#')

	regions[1].grid.Print()

	fmt.Println(canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(11, 0)))
	fmt.Println(canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(12, 0)))
	fmt.Println(canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(0, 3)))

	if canPlaceTile(regions[1].grid, presents[0].grid, lib.NewCoord(0, 0)) {
		t.Errorf("Expected collision at (0,0)")
	}
}
