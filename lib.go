package lib

import (
	"os"
	"path/filepath"
	"runtime"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// Get the adjacent directory to lib, which should be root
func getRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// Read a file starting from the root directory
func ReadFile(path string) string {
	root := getRootDir()
	osPath := filepath.Join(root, path)
	dat, err := os.ReadFile(osPath)
	CheckError(err)

	return string(dat)
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := min(i+chunkSize, len(s))
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

type Grid[T any] struct {
	Width  int
	Height int
	Cells  []T
}

type Coord struct {
	X int
	Y int
}

func GetIndexOfCoord(coord Coord, width int) int {
	return coord.Y*width + coord.X
}

func GetCoordOfIndex(index int, width int) Coord {
	return Coord{
		X: index % width,
		Y: index / width,
	}
}

var AdjacencyMatrix = []Coord{
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

var FullAdjacencyMatrix = []Coord{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func Print2DGrid[T any](grid Grid[T]) {
	for y := 0; y < grid.Height; y++ {
		for x := range grid.Width {
			print(grid.Cells[GetIndexOfCoord(Coord{X: x, Y: y}, grid.Width)])
			if _, ok := any(grid.Cells[0]).(int); ok {
				print(" ")
			}
		}
		println()
	}
}
