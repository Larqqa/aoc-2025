package lib

import (
	"math"
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

// Chunk a string into pieces by given chunk size
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
			cell := grid.Cells[GetIndexOfCoord(Coord{X: x, Y: y}, grid.Width)]
			if r, ok := any(cell).(rune); ok {
				print(string(r))
			} else {
				print(cell)
			}
			if _, ok := any(grid.Cells[0]).(int); ok {
				print(" ")
			}
		}
		println()
	}
}

// Get a single number at a specific index from an integer
// Returns -1 if the index is out of bounds
func GetNumberAtIndex(val int, index int) int {
	length := int(math.Log10(float64(val)))
	if length < index || index < 0 {
		return -1
	}

	mod := int(math.Pow(10, float64(length-index)))

	return (val / mod) % 10
}

type Point struct {
	X int
	Y int
	Z int
}

func (p Point) Distance(other Point) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	dz := float64(p.Z - other.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type UFDS[T comparable] struct {
	Parent map[T]T
	Rank   map[T]int
}

func (u *UFDS[T]) CountSets() int {
	roots := make(map[T]bool)
	for node := range u.Parent {
		roots[u.Find(node)] = true
	}
	return len(roots)
}

func NewUFDS[T comparable]() *UFDS[T] {
	return &UFDS[T]{
		Parent: make(map[T]T),
		Rank:   make(map[T]int),
	}
}

func (u *UFDS[T]) Find(x T) T {
	if _, exists := u.Parent[x]; !exists {
		u.Parent[x] = x
		u.Rank[x] = 1
	}
	if u.Parent[x] == x {
		return x
	}
	u.Parent[x] = u.Find(u.Parent[x])
	return u.Parent[x]
}

func (u *UFDS[T]) Union(x, y T) {
	rootX := u.Find(x)
	rootY := u.Find(y)

	if rootX == rootY {
		return
	}

	if u.Rank[rootX] < u.Rank[rootY] {
		u.Parent[rootX] = rootY
	} else if u.Rank[rootX] > u.Rank[rootY] {
		u.Parent[rootY] = rootX
	} else {
		u.Parent[rootY] = u.Parent[rootX]
		u.Rank[rootX]++
	}
}
