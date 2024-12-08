package grid

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

func (v1 Coordinate) Add(v2 Coordinate) Coordinate {
	return Coordinate{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Coordinate) Subtract(v2 Coordinate) Coordinate {
	return Coordinate{v1.X - v2.X, v1.Y - v2.Y}
}

// x goes right, y goes bottom, listing vectors to corners from top-left clockwise:
var CornerVectors = []Coordinate{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}

// ... listing vectors to sides from top clockwise
var SideVectors = []Coordinate{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

var Vectors []Coordinate = append(SideVectors, CornerVectors...)

type Grid struct {
	values        [][]string
	Width, Height int
}

func NewGrid(w int, h int) *Grid {
	g := make([][]string, w)
	for x := 0; x < w; x++ {
		g[x] = make([]string, h)
	}
	return &Grid{values: g, Width: w, Height: h}
}

func GridFromLines(lines []string) *Grid {
	width, height := len(lines[0]), len(lines)
	var values [][]string
	for x := 0; x < width; x++ {
		values = append(values, make([]string, height))
		for y := 0; y < height; y++ {
			values[x][y] = string(lines[y][x])
		}
	}
	grid := &Grid{values: values}
	grid.Width = len(lines[0])
	grid.Height = len(lines)
	return grid
}

func (g Grid) GetLetterByCoordinate(c Coordinate) string {
	return g.values[c.X][c.Y]
}

func (g Grid) SetLetterByCoordinate(c Coordinate, letter string) {
	g.values[c.X][c.Y] = letter
}

func (g Grid) IsEdge(c Coordinate) bool {
	return c.X == 0 || c.Y == 0 || c.X == g.Width-1 || c.Y == g.Height-1
}

func (g Grid) IsValidCoord(c Coordinate) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < g.Width && c.Y < g.Height
}

func (g Grid) FindAll(letter string) []Coordinate {
	var occurences []Coordinate
	c := Coordinate{}
	for c.Y = range g.Height {
		for c.X = range g.Width {
			if g.GetLetterByCoordinate(c) == letter {
				occurences = append(occurences, c)
			}
		}
	}
	return occurences
}

func print2DArrayWithCoordinates(array [][]string) {
	// todo: generated, review and simplify
	// Determine the size of the array
	rows := len(array)
	if rows == 0 {
		fmt.Println("Empty array")
		return
	}
	cols := len(array[0])

	// Print column headers
	fmt.Print("    ") // Offset for row labels
	for col := 0; col < cols; col++ {
		fmt.Printf("%v ", col)
	}
	fmt.Println()

	// Print rows with coordinates
	for col := 0; col < cols; col++ {
		fmt.Printf("%3d ", col) // Row label
		for row := 0; row < rows; row++ {
			toPrint := array[row][col]
			if toPrint == "" {
				toPrint = " "
			}
			fmt.Printf("%v ", toPrint)
		}
		fmt.Println()
	}
}

func (g Grid) String() {
	print2DArrayWithCoordinates(g.values)
}
