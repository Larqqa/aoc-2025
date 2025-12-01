package main

import (
	lib "aoc/2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayone.txt")
	p1, p2 := solvePartOne(data)
	fmt.Println("Part One:", p1, "Part Two:", p2)
}

func solvePartOne(data string) (int, int) {
	position := 50
	dialLength := 100
	stopsAtZero := 0
	rotationsToZero := 0
	commands := strings.SplitSeq(data, "\n")

	for command := range commands {
		direction := command[0]
		value, _ := strconv.Atoi(command[1:])
		rotations := value / dialLength
		steps := value % dialLength

		if direction == 'L' {
			if position != 0 && position-steps < 0 {
				rotations++
			}
			position = ((position-steps)%dialLength + dialLength) % dialLength
		}

		if direction == 'R' {
			if position != 0 && position+steps > dialLength {
				rotations++
			}
			position = ((position+steps)%dialLength + dialLength) % dialLength
		}

		rotationsToZero += rotations

		if position == 0 {
			stopsAtZero++
		}
	}

	return stopsAtZero, rotationsToZero + stopsAtZero
}
