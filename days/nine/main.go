package main

import (
	lib "aoc/2025"
	"fmt"
	"math"
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
	for _, coord := range coords {
		for _, coord2 := range coords {
			width := int(math.Abs(float64(coord.X-coord2.X))) + 1
			height := int(math.Abs(float64(coord.Y-coord2.Y))) + 1
			area := width * height
			if area > largest {
				largest = area
			}
		}
	}

	return largest
}

func solvePartTwo(data string) int {
	return 0
}
