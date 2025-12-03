package main

import (
	lib "aoc/2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daythree.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) []string {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")
	return parsed
}

func largestValidValue(s string, maxDepth int) (byte, int) {
	maxByte := s[0]
	maxIndex := 0
	for i := 1; i < len(s); i++ {
		if s[i] > maxByte && (len(s)-i) >= maxDepth {
			maxByte = s[i]
			maxIndex = i
		}
	}
	return maxByte, maxIndex
}

func findLargestNumberOfLength(s string, maxDepth int) int {
	acc := ""
	depth := 0
	index := 0
	for depth < maxDepth {
		largest, newIndex := largestValidValue(s[index:], maxDepth-depth)
		acc += string(largest)
		index = index + newIndex + 1
		depth++
	}
	intVal, _ := strconv.Atoi(acc)
	return intVal
}

func solvePartOne(data string) int {
	rows := parse(data)
	sum := 0

	for _, row := range rows {
		sum += findLargestNumberOfLength(row, 2)
	}

	return sum
}

func solvePartTwo(data string) int {
	rows := parse(data)
	sum := 0

	for _, row := range rows {
		sum += findLargestNumberOfLength(row, 12)
	}

	return sum
}
