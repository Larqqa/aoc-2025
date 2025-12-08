package main

import (
	lib "aoc/2025"
	"fmt"
	"slices"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayseven.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) (lib.Grid[rune], lib.Coord) {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")

	width := len(parsed[0])
	height := len(parsed)
	grid := make([]rune, height*width)

	start := lib.Coord{
		X: 0,
		Y: 0,
	}

	for y, row := range parsed {
		for x, char := range row {
			index := lib.GetIndexOfCoord(lib.Coord{X: x, Y: y}, width)
			grid[index] = char

			if char == 'S' {
				start = lib.Coord{X: x, Y: y}
			}
		}
	}

	return lib.Grid[rune]{
		Width:  width,
		Height: height,
		Cells:  grid,
	}, start
}

func solvePartOne(data string) int {
	grid, start := parse(data)
	sum := 0

	curY := 0
	tachyons := []lib.Coord{start}
	for {
		for i := len(tachyons) - 1; i >= 0; i-- {
			t := tachyons[i]
			t.Y++

			curCell := grid.Cells[lib.GetIndexOfCoord(t, grid.Width)]

			if curCell != '^' {
				if slices.Contains(tachyons, t) {
					tachyons = slices.Delete(tachyons, i, i+1)
				} else {
					tachyons[i] = t
				}
				continue
			}

			sum++

			// Move current to left
			t.X--
			if t.X < 0 {
			}
			if slices.Contains(tachyons, t) || t.X < 0 {
				tachyons = slices.Delete(tachyons, i, i+1)
			} else {
				tachyons[i] = t
			}

			// Add new to right
			split := lib.Coord{X: t.X + 2, Y: t.Y}
			if split.X >= grid.Width-1 {
				continue
			}
			if slices.Contains(tachyons, split) {
				continue
			}

			tachyons = append(tachyons, split)
		}

		curY++
		if curY >= grid.Height-1 {
			break
		}
	}

	return sum
}

func branch(c lib.Coord, grid *lib.Grid[rune], memo *map[lib.Coord]int) int {
	if cached, ok := (*memo)[c]; ok {
		return cached
	}

	if c.Y >= grid.Height-1 {
		return 1
	}

	next := lib.Coord{X: c.X, Y: c.Y + 1}
	nextCell := grid.Cells[lib.GetIndexOfCoord(next, grid.Width)]

	var result int
	if nextCell == '^' {
		result = branch(lib.Coord{X: next.X - 1, Y: next.Y}, grid, memo) + branch(lib.Coord{X: next.X + 1, Y: next.Y}, grid, memo)
	} else {
		result = branch(next, grid, memo)
	}

	(*memo)[c] = result
	return result
}

func solvePartTwo(data string) int {
	grid, start := parse(data)
	memo := make(map[lib.Coord]int)
	return branch(start, &grid, &memo)
}
