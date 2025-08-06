package board

import (
	"fmt"
	"math"
)

type Tile struct {
	Letter    rune
	Neighbors []int
}

func getDimensions(letters []rune) (int, error) {
	sqrt := math.Sqrt(float64(len(letters)))
	int_sqrt := int(sqrt)
	if float64(int_sqrt) == sqrt && int_sqrt*int_sqrt == len(letters) {
		return int_sqrt, nil
	} else {
		return -1, fmt.Errorf("Grid must be a square")
	}
}

func getNeighbors(index int, size int) []int {
	// Neighbor checking order
	// 1 2 3
	// 4 x 5
	// 6 7 8
	var neighbors []int

	// Directions
	var (
		topLeft     = index - size - 1
		top         = index - size
		topRight    = index - size + 1
		left        = index - 1
		right       = index + 1
		bottomLeft  = index + size - 1
		bottom      = index + size
		bottomRight = index + size + 1
	)

	// Edges
	var (
		topEdge    = index < size
		leftEdge   = index%size == 0
		rightEdge  = (index+1)%size == 0
		bottomEdge = index >= size*(size-1)
	)

	// Corners
	var (
		topLeftCorner     = 0
		topRightCorner    = size - 1
		bottomLeftCorner  = size * (size - 1)
		bottomRightCorner = size*size - 1
	)

	if topEdge {
		switch index {
		case topLeftCorner:
			neighbors = []int{right, bottom, bottomRight}
		case topRightCorner:
			neighbors = []int{left, bottomLeft, bottom}
		default:
			neighbors = []int{left, right, bottomLeft, bottom, bottomRight}
		}
	} else if leftEdge && index != bottomLeftCorner {
		neighbors = []int{top, topRight, right, bottom, bottomRight}
	} else if rightEdge && index != bottomRightCorner {
		neighbors = []int{topLeft, top, left, bottomLeft, bottom}
	} else if bottomEdge {
		switch index {
		case bottomLeftCorner:
			neighbors = []int{top, topRight, right}
		case bottomRightCorner:
			neighbors = []int{topLeft, top, left}
		default:
			neighbors = []int{topLeft, top, topRight, left, right}
		}
	} else {
		// interior
		neighbors = []int{topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight}
	}

	return neighbors
}

func PopulateBoard(letters []rune) ([]Tile, error) {
	gridLength, err := getDimensions(letters)
	if err != nil {
		return nil, err
	}

	grid := make([]Tile, len(letters))
	for i, letter := range letters {
		// TODO: Make sure to find a way to calculate size
		neighbors := getNeighbors(i, gridLength)
		grid[i] = Tile{
			Letter:    letter,
			Neighbors: neighbors,
		}
	}

	return grid, nil
}
