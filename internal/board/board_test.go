package board

import (
	"slices"
	"testing"
)

func TestGetNeighbors(t *testing.T) {
	// Board:
	//  0  1  2  3  4
	//  5  6  7  8  9
	// 10 11 12 13 14
	// 15 16 17 18 19
	// 20 21 22 23 24

	const size int = 5

	indices := []int{0, 3, 4, 10, 12, 19, 20, 22, 24}
	wants := [][]int{
		{1, 5, 6},                     // Top left corner (index: 0)
		{2, 4, 7, 8, 9},               // Top edge (index: 3)
		{3, 8, 9},                     // Top right corner (index: 4)
		{5, 6, 11, 15, 16},            // Left edge (index: 10)
		{6, 7, 8, 11, 13, 16, 17, 18}, // Middle (index: 12)
		{13, 14, 18, 23, 24},          // Right edge (index: 19)
		{15, 16, 21},                  // Bottom left corner (index: 20)
		{16, 17, 18, 21, 23},          // Bottom edge (index: 22)
		{18, 19, 23},                  // Bottom right corner (index: 24)
	}

	for i, index := range indices {
		got := getNeighbors(index, size)
		if !slices.Equal(got, wants[i]) {
			t.Errorf("Did not generate correct neighbors for index %d.\n\tExpected: %v\n\tGot: %v", index, wants[i], got)
		}
	}
}

func TestGetDimensions(t *testing.T) {
	valid := []rune{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	}
	got, _ := getDimensions(valid)
	want := 3

	if got != want {
		t.Errorf("Incorrect dimensions.\n\tExpected: %d\n\tGot:%d", want, got)
	}

	invalid := []rune{
		0, 1, 2,
		3, 4, 5,
		6, 7,
	}
	got, err := getDimensions(invalid)
	if err == nil {
		t.Error("Expected failure, got success")
	}
}
