package main

import (
	lib "aoc/2025"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayfive.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

type Range struct {
	Min int
	Max int
}

func parse(data string) ([]Range, []int) {
	parts := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n\n")
	ranges := make([]Range, 0)

	for _, r := range strings.Split(parts[0], "\n") {
		rP := strings.Split(r, "-")
		min, _ := strconv.Atoi(rP[0])
		max, _ := strconv.Atoi(rP[1])
		ranges = append(ranges, Range{
			Min: min,
			Max: max,
		})
	}

	ingredients := make([]int, 0)
	for ingredient := range strings.SplitSeq(parts[1], "\n") {
		ing, _ := strconv.Atoi(ingredient)
		ingredients = append(ingredients, ing)
	}

	return ranges, ingredients
}

func solvePartOne(data string) int {
	ranges, ingredients := parse(data)

	sum := 0
	for _, i := range ingredients {
		for _, r := range ranges {
			if i >= r.Min && i <= r.Max {
				sum++
				break
			}
		}
	}
	return sum
}

func solvePartTwo(data string) int {
	ranges, _ := parse(data)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := merged[len(merged)-1]
		if r.Min <= last.Max {
			merged[len(merged)-1].Max = max(last.Max, r.Max)
		} else {
			merged = append(merged, r)
		}
	}

	sum := 0
	for _, r := range merged {
		sum += r.Max - r.Min + 1
	}
	return sum
}
