package main

import (
	lib "aoc/2025"
	"fmt"
	"slices"
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

type Node struct {
	first    lib.Point
	second   lib.Point
	distance float64
}

func solvePartOne(data string, conns int) int {
	points := parse(data)

	// calculate all distances of pairs of connecions
	distances := make(map[float64]Node)
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			if p1 == p2 {
				continue
			}
			dist := p1.Distance(p2)
			distances[dist] = Node{first: p1, second: p2, distance: dist}
		}
	}

	// Get keys, sort them, take n smallest connections
	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	pairs := keys[:conns]

	connections := make([][]lib.Point, 0)
	for _, k := range pairs {
		node := distances[k]

		p1i := -1
		p2i := -1

		for i := 0; i < len(connections); i++ {
			conn := connections[i]
			if slices.Contains(conn, node.first) {
				p1i = i
			}
			if slices.Contains(conn, node.second) {
				p2i = i
			}
		}

		// add new connection
		if p1i < 0 && p2i < 0 {
			connections = append(connections, []lib.Point{node.first, node.second})
			continue
		}

		// connected,skip
		if p1i == p2i {
			continue
		}

		// connect 2 to 1
		if 0 <= p1i && p2i < 0 {
			connections[p1i] = append(connections[p1i], node.second)
			continue
		}

		// connect 1 to 2
		if 0 <= p2i && p1i < 0 {
			connections[p2i] = append(connections[p2i], node.first)
			continue
		}

		// merge connections
		if 0 <= p1i && 0 <= p2i && p1i != p2i {
			connections[p1i] = append(connections[p1i], connections[p2i]...)
			connections = append(connections[:p2i], connections[p2i+1:]...)
			continue
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return len(connections[i]) > len(connections[j])
	})

	sum := 1
	for _, c := range connections[:3] {
		sum *= len(c)
	}
	return sum
}

func solvePartTwo(data string) int {
	sum := 0
	return sum
}
