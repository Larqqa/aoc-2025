package main

import (
	lib "aoc/2025"
	"fmt"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayfour.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) lib.Grid {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")

	width := len(parsed[0])
	height := len(parsed)
	grid := make([]rune, height*width)

	for y, row := range parsed {
		for x, char := range row {
			index := lib.GetIndexOfCoord(lib.Coord{X: x, Y: y}, width)
			grid[index] = char
		}
	}

	return lib.Grid{
		Width:  width,
		Height: height,
		Cells:  grid,
	}
}

func findMoveable(grid lib.Grid) []int {
	moveable := []int{}
	for i, cell := range grid.Cells {
		if cell != '@' {
			continue
		}

		neighbours := 0
		for _, offset := range lib.FullAdjacencyMatrix {
			coord := lib.GetCoordOfIndex(i, grid.Width)
			coord.X += offset.X
			coord.Y += offset.Y

			if coord.X < 0 || coord.Y < 0 || coord.X >= grid.Width || coord.Y >= grid.Height {
				continue
			}

			neighbor := grid.Cells[lib.GetIndexOfCoord(coord, grid.Width)]
			if neighbor == '@' {
				neighbours++
			}
		}

		if neighbours < 4 {
			moveable = append(moveable, i)
		}
	}

	return moveable
}

func solvePartOne(data string) int {
	grid := parse(data)
	return len(findMoveable(grid))
}

func solvePartTwo(data string) int {
	grid := parse(data)
	sum := 0

	for {
		moveable := findMoveable(grid)

		if len(moveable) == 0 {
			break
		}

		sum += len(moveable)

		for _, index := range moveable {
			grid.Cells[index] = 'x'
		}
	}

	return sum
}
