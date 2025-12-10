package lib

import (
	"testing"
)

func TestChunkString(t *testing.T) {
	{
		s := "abcdefghij"
		chunked := ChunkString(s, 3)
		expected := []string{"abc", "def", "ghi", "j"}
		if len(chunked) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(chunked))
		}
		for i, v := range expected {
			if chunked[i] != v {
				t.Errorf("At index %d, expected %s, got %s", i, v, chunked[i])
			}
		}
	}
	{
		s := "hello"
		chunked := ChunkString(s, 2)
		expected := []string{"he", "ll", "o"}
		if len(chunked) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(chunked))
		}
		for i, v := range expected {
			if chunked[i] != v {
				t.Errorf("At index %d, expected %s, got %s", i, v, chunked[i])
			}
		}
	}
}

func TestNumberAtIndex(t *testing.T) {
	{
		v := GetNumberAtIndex(12345, 0)
		if v != 1 {
			t.Errorf("Failed at index 0, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, 1)
		if v != 2 {
			t.Errorf("Failed at index 1, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, 2)
		if v != 3 {
			t.Errorf("Failed at index 2, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, 3)
		if v != 4 {
			t.Errorf("Failed at index 3, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, 4)
		if v != 5 {
			t.Errorf("Failed at index 4, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, 400)
		if v != -1 {
			t.Errorf("Failed at index 400, got %d", v)
		}
	}
	{
		v := GetNumberAtIndex(12345, -100)
		if v != -1 {
			t.Errorf("Failed at index -100, got %d", v)
		}
	}
}

func TestEdgesIntersect(t *testing.T) {
	e1 := Edge{A: Coord{X: 11, Y: 1}, B: Coord{X: 2, Y: 3}}

	e2 := Edge{A: Coord{X: 7, Y: 3}, B: Coord{X: 7, Y: 1}}
	if !e1.Intersects(e2) {
		t.Errorf("Expected edges to intersect")
	}

	e3 := Edge{A: Coord{X: 0, Y: 1}, B: Coord{X: 4, Y: 1}}
	if e1.Intersects(e3) {
		t.Errorf("Expected edges not to intersect")
	}
}
