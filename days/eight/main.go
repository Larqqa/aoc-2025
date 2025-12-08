package main

import (
	lib "aoc/2025"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayeight.txt")

	p1 := solvePartOne(data, 1000)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

func parse(data string) []lib.Point {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")
	points := make([]lib.Point, 0)
	for _, line := range parsed {
		p := strings.Split(line, ",")
		X, _ := strconv.Atoi(p[0])
		Y, _ := strconv.Atoi(p[1])
		Z, _ := strconv.Atoi(p[2])
		points = append(points, lib.Point{X: X, Y: Y, Z: Z})
	}
	return points
}

type Pair struct {
	first    lib.Point
	second   lib.Point
	distance float64
}

// calculate all distances of pairs of connecions
func getPairs(points []lib.Point) (map[float64]Pair, []float64) {
	pairs := make(map[float64]Pair)
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			if p1 == p2 {
				continue
			}
			dist := p1.Distance(p2)
			pairs[dist] = Pair{first: p1, second: p2, distance: dist}
		}
	}

	distances := make([]float64, 0, len(pairs))
	for k := range pairs {
		distances = append(distances, k)
	}
	sort.Float64s(distances)

	return pairs, distances
}

func solvePartOne(data string, conns int) int {
	points := parse(data)
	pairs, distances := getPairs(points)

	ufds := lib.NewUFDS[lib.Point]()
	for _, k := range distances[:conns] {
		p := pairs[k]
		ufds.Union(p.first, p.second)
	}

	// Find distances of all sets in ufds
	dists := make(map[lib.Point]int)
	for _, p := range ufds.Parent {
		dists[ufds.Find(p)]++
	}

	ds := make([]int, 0, len(dists))
	for _, v := range dists {
		ds = append(ds, v)
	}
	sort.Ints(ds)

	sum := 1
	for _, c := range ds[len(ds) - 3:] {
		sum *= c
	}
	return sum
}

func solvePartTwo(data string) int {
	points := parse(data)
	pairs, distances := getPairs(points)

	ufds := lib.NewUFDS[lib.Point]()
	unions := len(points)
	for _, k := range distances {
		p := pairs[k]

		// Doing unique union, reduce max count of unions
		if ufds.Find(p.first) != ufds.Find(p.second) {
			unions--
		}

		// All pairs combined to one circuit
		if unions == 1 {
			return p.first.X * p.second.X
		}

		ufds.Union(p.first, p.second)
	}

	return 0
}
