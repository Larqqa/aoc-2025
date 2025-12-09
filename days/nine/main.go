package main

import (
	lib "aoc/2025"
	"fmt"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daynine.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) []lib.Coord {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")
	coords := make([]lib.Coord, len(parsed))
	for i, line := range parsed {
		coords[i] = lib.CoordFromString(line)
	}

	return coords
}

func solvePartOne(data string) int {
	coords := parse(data)
	largest := 0
	for i, coord := range coords {
		for _, coord2 := range coords[i+1:] {
			area := coord.GetArea(coord2)
			if area > largest {
				largest = area
			}
		}
	}

	return largest
}

func solvePartTwo(data string) int {
	coords := parse(data)
	maxX, maxY := 0, 0
	for _, coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	fmt.Println(maxX, maxY)

	// grid := lib.Grid[rune]{
	// 	Width:  maxX + 1,
	// 	Height: maxY + 1,
	// 	Cells:  make([]rune, (maxX+1)*(maxY+1)),
	// }

	// for i := range grid.Cells {
	// 	grid.Cells[i] = '.'
	// }

	// for _, coord := range coords {
	// 	grid.Cells[coord.GetIndex(grid.Width)] = '#'
	// }

	return 0
}
