package main

import (
	lib "aoc/2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daytwo.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) [][2]int {
	ranges := strings.Split(data, ",")

	parsed := [][2]int{}
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])
		parsed = append(parsed, [2]int{lower, upper})
	}

	return parsed
}

func solvePartOne(data string) int {
	ranges := parse(data)
	sum := 0
	for _, r := range ranges {
		start, end := r[0], r[1]
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			half := len(str) / 2
			if str[:half] == str[half:] {
				sum += i
			}
		}
	}

	return sum
}

func solvePartTwo(data string) int {
	ranges := parse(data)
	sum := 0
	for _, r := range ranges {
		start, end := r[0], r[1]
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			half := len(str) / 2
			for j := 1; j <= half; j++ {
				chunkStrings := lib.ChunkString(str, j)
				allEqual := true
				for _, chunk := range chunkStrings {
					if chunk != chunkStrings[0] {
						allEqual = false
						break
					}
				}
				if allEqual {
					sum += i
					break
				}
			}
		}
	}

	return sum
}
