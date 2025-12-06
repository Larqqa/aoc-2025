package main

import (
	lib "aoc/2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daysix.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) ([]string, []string) {
	rows := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")
	instructions := strings.Fields(strings.TrimSpace(rows[len(rows)-1]))

	return rows[:len(rows)-1], instructions
}

func doInstruction(a int, b int, instruction string) int {
	switch instruction {
	case "*":
		if a == 0 {
			return b
		}
		return a * b
	case "+":
		return a + b
	default:
		return a
	}
}

func solvePartOne(data string) int {
	values, instructions := parse(data)

	cells := []int{}
	for _, r := range values {
		values := strings.FieldsSeq(strings.TrimSpace(r))
		for v := range values {
			num, _ := strconv.Atoi(v)
			cells = append(cells, num)
		}
	}

	grid := lib.Grid[int]{
		Cells: cells, Width: len(instructions), Height: len(values),
	}

	sum := 0
	for X := 0; X < grid.Width; X++ {
		acc := 0
		for Y := 0; Y < grid.Height; Y++ {
			ind := lib.GetIndexOfCoord(lib.Coord{X: X, Y: Y}, grid.Width)
			acc = doInstruction(acc, grid.Cells[ind], instructions[X])
		}
		sum += acc
	}
	return sum
}

func solvePartTwo(data string) int {
	values, instructions := parse(data)

	rotated := make([]string, len(values[0]))
	for i := range rotated {
		var sb strings.Builder
		for j := range values {
			if i < len(values[j]) {
				sb.WriteByte(values[j][i])
			}
		}
		rotated[i] = sb.String()
	}

	// assumed data does not contain 0 and used as block separator
	rotated = append(rotated, "0")

	sum := 0
	acc := 0
	insIndex := 0
	for _, r := range rotated {
		val, _ := strconv.Atoi(strings.TrimSpace(r))

		if val != 0 {
			acc = doInstruction(acc, val, instructions[insIndex])
			continue
		}

		sum += acc
		acc = 0
		insIndex++
	}

	return sum
}
