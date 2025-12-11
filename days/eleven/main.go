package main

import (
	lib "aoc/2025"
	"fmt"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayeleven.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) map[string][]string {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")

	graph := map[string][]string{}
	for _, line := range parsed {
		v := strings.Split(line, ":")
		values := strings.Split(strings.TrimSpace(v[1]), " ")
		graph[v[0]] = values
	}

	return graph
}

func findPaths(start string, graph map[string][]string, mustContain []string, cache map[string]int) int {
	if start == "out" {
		if len(mustContain) > 0 {
			return 0
		}
		return 1
	}

	key := start + "|" + strings.Join(mustContain, ",")
	if val, ok := cache[key]; ok {
		return val
	}

	targets := graph[start]
	totalPaths := 0
	for _, t := range targets {
		newMustContain := []string{}
		for _, m := range mustContain {
			if m != t {
				newMustContain = append(newMustContain, m)
			}
		}
		totalPaths += findPaths(t, graph, newMustContain, cache)
	}

	cache[key] = totalPaths
	return totalPaths
}

func solvePartOne(data string) int {
	graph := parse(data)
	return findPaths("you", graph, []string{}, make(map[string]int))
}

func solvePartTwo(data string) int {
	graph := parse(data)
	return findPaths("svr", graph, []string{"fft", "dac"}, make(map[string]int))
}
