package main

import (
	lib "aoc/2025"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daynine.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)

	// too low 123179664
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
	minX, minY := math.MaxInt, math.MaxInt
	for _, coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
		if coord.X < minX {
			minX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
	}

	getNextInDirection := func(c lib.Coord, d lib.Direction) lib.Coord {
		furthest := lib.Coord{}
		for {
			n := c.GetNextCoord(d)
			if n.X < minX || n.X > maxX || n.Y < minY || n.Y > maxY {
				break
			}
			if slices.Contains(coords, n) {
				furthest = n
			}
			c = n
		}
		return furthest
	}

	largest := 0
	for i, coord := range coords {
		fmt.Println(i)
		width, height := 0, 0

		up := getNextInDirection(coord, lib.Up)
		down := getNextInDirection(coord, lib.Down)

		uDist := coord.ManhattanDistance(up)
		dDist := coord.ManhattanDistance(down)

		if up != (lib.Coord{}) && uDist > height {
			height = uDist
		}

		if down != (lib.Coord{}) && dDist > height {
			height = dDist
		}

		left := getNextInDirection(coord, lib.Left)
		right := getNextInDirection(coord, lib.Right)

		lDist := coord.ManhattanDistance(left)
		rDist := coord.ManhattanDistance(right)

		if left != (lib.Coord{}) && lDist > width {
			width = lDist
		}

		if right != (lib.Coord{}) && rDist > width {
			width = rDist
		}

		area := width * height
		if area > largest {
			largest = area
		}
	}

	return largest
}
