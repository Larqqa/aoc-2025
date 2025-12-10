package main

import (
	lib "aoc/2025"
	"fmt"
	"slices"
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

	edges := make([]lib.Edge, 0)
	for i, coord := range coords {
		edges = append(edges, lib.Edge{A: coord, B: coords[(i+1)%len(coords)]})
	}

	largest := 0
	for i, c := range coords {
		for _, c2 := range coords[i+1:] {
			// Define the diagonals of the square
			squareEdges := []lib.Edge{
				{A: c, B: c2},
				{A: lib.Coord{X: c.X, Y: c2.Y}, B: lib.Coord{X: c2.X, Y: c.Y}},
			}

			// Check for intersections with existing edges
			intersects := false
			for _, se := range squareEdges {
				if slices.ContainsFunc(edges, se.Intersects) {
					intersects = true
				}
				if intersects {
					break
				}
			}

			if intersects {
				continue
			}

			// Find largest area from non colliding squares
			area := c.GetArea(c2)
			if area > largest {
				largest = area
			}
		}
	}

	return largest
}
