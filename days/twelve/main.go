package main

import (
	lib "aoc/2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/daytwelve.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)
}

type region struct {
	grid       lib.Grid[rune]
	quantities []int
}

type present struct {
	grid         lib.Grid[rune]
	orientations []lib.Grid[rune]
	area         int
}

func getAllDistinctOrientations(grid lib.Grid[rune]) []lib.Grid[rune] {
	var orientations []lib.Grid[rune]
	seen := make(map[string]bool)

	current := grid.Copy()
	for range 4 {
		key := string(current.Cells)
		if !seen[key] {
			orientations = append(orientations, current.Copy())
			seen[key] = true
		}

		{
			flip := current.Copy()
			flip.Flip(lib.Up)
			key := string(flip.Cells)
			if !seen[key] {
				orientations = append(orientations, flip)
				seen[key] = true
			}
		}

		current.Rotate(lib.Right)
	}

	return orientations
}

func parse(data string) ([]present, []region) {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n\n")

	presentsStr := parsed[:len(parsed)-1]
	presents := make([]present, len(presentsStr))
	for i, pStr := range presentsStr {
		lines := strings.Split(pStr, "\n")
		grid := lib.NewGrid[rune](len(lines)-1, len(lines[1]))
		for y, line := range lines[1:] {
			for x, ch := range line {
				grid.Set(lib.NewCoord(x, y), ch)
			}
		}

		area := 0
		for _, cell := range grid.Cells {
			if cell == '#' {
				area++
			}
		}
		presents[i] = present{
			grid:         grid,
			orientations: getAllDistinctOrientations(grid),
			area:         area,
		}
	}

	regionsStr := strings.Split(parsed[len(parsed)-1], "\n")
	regions := make([]region, len(regionsStr))
	for i, regionStr := range regionsStr {
		s := strings.Split(regionStr, ": ")

		gridStr := strings.Split(s[0], "x")
		width, _ := strconv.Atoi(gridStr[0])
		height, _ := strconv.Atoi(gridStr[1])
		grid := lib.NewGrid[rune](height, width)
		for y := range height {
			for x := range width {
				grid.Set(lib.NewCoord(x, y), '.')
			}
		}

		quantitiesStr := strings.Split(s[1], " ")
		quantities := make([]int, len(presentsStr))
		for i, qStr := range quantitiesStr {
			q, _ := strconv.Atoi(qStr)
			quantities[i] = q
		}

		regions[i] = region{
			grid:       grid,
			quantities: quantities,
		}
	}

	return presents, regions
}

// Check if tile can be placed into target at offset
func canPlaceTile(target lib.Grid[rune], tile lib.Grid[rune], offset lib.Coord) bool {
	// Check if this is the first tile to add
	isEmpty := true
	for _, cell := range target.Cells {
		if cell == '#' {
			isEmpty = false
			break
		}
	}

	// First tile can be placed anywhere
	if isEmpty {
		return true
	}

	for y := 0; y < tile.Height; y++ {
		for x := 0; x < tile.Width; x++ {
			aX := x + offset.X
			aY := y + offset.Y

			// Skip empty cells of tile
			bCell := tile.Get(lib.NewCoord(x, y))
			if bCell != '#' {
				continue
			}

			// Directly collides with other tile
			aCell := target.Get(lib.NewCoord(aX, aY))
			if aCell == '#' {
				return false
			}
		}
	}

	// Minimize available placement opportunities by requiring at least one neighbor
	// Check neighboring tiles at top and left edge since we iterate always starting from top left
	neighbors := []lib.Coord{}

	for i := range tile.Height {
		neighbors = append(neighbors, lib.NewCoord(offset.X-i, offset.Y))
	}
	for i := range tile.Width {
		neighbors = append(neighbors, lib.NewCoord(offset.X, offset.Y-i))
	}

	hasNeighbor := false
	for _, neighbor := range neighbors {
		if neighbor.X >= 0 && neighbor.X < target.Width && neighbor.Y >= 0 && neighbor.Y < target.Height {
			if target.Get(neighbor) == '#' {
				hasNeighbor = true
				break
			}
		}
	}

	return hasNeighbor
}

func placeTile(a *lib.Grid[rune], b lib.Grid[rune], offset lib.Coord, tile rune) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Get(lib.NewCoord(x, y)) == '#' {
				a.Set(lib.NewCoord(x+offset.X, y+offset.Y), tile)
			}
		}
	}
}

func getRegionKey(region lib.Grid[rune], index int) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(index))
	sb.WriteString(":")
	sb.WriteString(string(region.Cells))
	return sb.String()
}

func canPlaceAllPresents(region lib.Grid[rune], presents []present, index int, memo map[string]bool) bool {
	if index >= len(presents) {
		return true
	}

	key := getRegionKey(region, index)
	if result, exists := memo[key]; exists {
		return result
	}

	present := presents[index]

	for _, orientation := range present.orientations {
		for y := 0; y <= region.Height-orientation.Height; y++ {
			for x := 0; x <= region.Width-orientation.Width; x++ {
				offset := lib.NewCoord(x, y)

				if !canPlaceTile(region, orientation, offset) {
					continue
				}

				placeTile(&region, orientation, offset, '#')

				if canPlaceAllPresents(region, presents, index+1, memo) {
					return true
				}

				// Backtrack
				placeTile(&region, orientation, offset, '.')
			}
		}
	}

	memo[key] = false
	return false
}

func solvePartOne(data string) int {
	presents, regions := parse(data)

	// Max region size is 50x50
	// Max shapes is 385
	// 50x50: 73 58 64 63 64 63

	// Prune impossible regions by total areas
	for i := len(regions) - 1; i >= 0; i-- {
		reg := regions[i]
		totalArea := reg.grid.Width * reg.grid.Height
		presentArea := 0
		for j, q := range reg.quantities {
			presentArea += presents[j].area * q
		}

		if totalArea < presentArea {
			regions = append(regions[:i], regions[i+1:]...)
		}
	}
	// Presents can fit in 541 regions

	total := 0
	for _, reg := range regions {
		// Get all presents for region
		var presentList []present
		for i, q := range reg.quantities {
			for range q {
				presentList = append(presentList, presents[i])
			}
		}

		// Sort presents by area
		for i := 0; i < len(presentList)-1; i++ {
			for j := i + 1; j < len(presentList); j++ {
				if presentList[j].area > presentList[i].area {
					presentList[i], presentList[j] = presentList[j], presentList[i]
				}
			}
		}

		if canPlaceAllPresents(reg.grid, presentList, 0, make(map[string]bool)) {
			total++
			// reg.grid.Print()
			// fmt.Println()
		}
	}

	return total
}
