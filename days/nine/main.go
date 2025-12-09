package main

import (
	lib "aoc/2025"
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daynine.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)

	// too low 123179664
}

func parse(data string) []lib.Coord {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")
	coords := make([]lib.Coord, len(parsed))
	for i, line := range parsed {
		coords[i] = lib.CoordFromString(line)
	}

	return coords
}

func solvePartOne(data string) int {
	coords := parse(data)
	largest := 0
	for i, coord := range coords {
		for _, coord2 := range coords[i+1:] {
			area := coord.GetArea(coord2)
			if area > largest {
				largest = area
			}
		}
	}

	return largest
}

func solvePartTwo(data string) int {
	coords := parse(data)

	// cached points by the x and y values
	// x contains sorted y from min to max
	// y contains sorted x from min to max
	cacheX := make(map[int][]lib.Coord)
	cacheY := make(map[int][]lib.Coord)

	for _, coord := range coords {
		cacheX[coord.X] = append(cacheX[coord.X], coord)
		cacheY[coord.Y] = append(cacheY[coord.Y], coord)
	}

	for _, list := range cacheX {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Y < list[j].Y
		})
	}

	for _, list := range cacheY {
		sort.Slice(list, func(i, j int) bool {
			return list[i].X < list[j].X
		})
	}

	keysX := make([]int, 0, len(cacheX))
	for k := range cacheX {
		keysX = append(keysX, k)
	}
	sort.Ints(keysX)

	keysY := make([]int, 0, len(cacheY))
	for k := range cacheY {
		keysY = append(keysY, k)
	}
	sort.Ints(keysY)

	largest := 0
	for _, coord := range coords {
		for _, x := range cacheY[coord.Y] {
			if x.X == coord.X {
				continue
			}

			for _, y := range cacheX[coord.X] {
				if y.Y == coord.Y {
					continue
				}

				// both target coordinates always exists in the x and y cache,
				// but mignt not exist in the shape
				target := lib.Coord{
					X: x.X,
					Y: y.Y,
				}

				fmt.Println(coord, target)

				// Find the Y values for the target X
				value := cacheX[target.X]
				idx := sort.SearchInts(keysY, coord.Y)

				// if target Y is larger than coord Y
				// find next largest edge from coord Y

				if target.Y > coord.Y {
					for {
						idx++
						if idx >= len(keysY) {
							break
						}
						key := keysY[idx]
						edgeList := cacheY[key]

						min := edgeList[0].X
						max := edgeList[len(edgeList)-1].X

						if target.X >= min && target.X <= max {
							fmt.Println("InsideShape")
						}

						fmt.Println(coord, target, edgeList)
					}
				}

				// if target Y is smaller than edge Y
				// the shape is not valid

				// do the otherway if target Y smaller than coord Y

				// if target is outside of the shape
				if value[0].Y > target.Y || value[len(value)-1].Y < target.Y {
					fmt.Println("OutsideShape")
					continue
				}

				area := coord.GetArea(target)

				// fmt.Println(area, coord, target)

				if area > largest {
					largest = area
				}
			}
		}
	}

	return largest
}
