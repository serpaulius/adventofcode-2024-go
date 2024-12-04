package fourth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func (v1 Coordinate) add(v2 Coordinate) Coordinate {
	return Coordinate{v1.x + v2.x, v1.y + v2.y}
}

// x goes right, y goes bottom, listing vectors to corners from top-left clockwise:
var cornerVectors = []Coordinate{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
var vectors []Coordinate = append([]Coordinate{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, cornerVectors...)

type Grid struct {
	lines         []string
	width, height int
}

func NewGrid(lines []string) *Grid {
	grid := &Grid{lines: lines}
	grid.width = len(lines[0])
	grid.height = len(lines)
	return grid
}

func (g Grid) getLetterByCoordinate(c Coordinate) string {
	return string(g.lines[c.y][c.x])
}

func (g Grid) findAll(letter string) []Coordinate {
	var occurences []Coordinate
	c := Coordinate{}
	for c.y = range g.height {
		for c.x = range g.width {
			if strings.Compare(g.getLetterByCoordinate(c), letter) == 0 {
				occurences = append(occurences, c)
			}
		}
	}
	return occurences
}

func (g Grid) countWordsForCoordinate(word string, position Coordinate) int {
	wordLength := len(word)
	count := 0
	for _, vect := range vectors {
		candidate := ""
		c := position
		for i := 0; i < wordLength; i++ {
			if c.x >= 0 && c.y >= 0 && c.x < g.width && c.y < g.height {
				candidate += g.getLetterByCoordinate(c)
			}
			c = c.add(vect)
		}
		if strings.Compare(candidate, word) == 0 {
			log.Println(position, vect, candidate)
			count++
		}
	}
	log.Println("count for", position, "is", count)
	return count
}

// find all A coordinates and check diagonal vectors for M and S in the opposite
// 4,4 6,6 4,6 6,4
// 0.1  M.M  S.M  ...
// ...  .A.  .S.  ...
// 2.3  S.S  S.M  ...
//
// NOT these!
// M.S  S.M
// .A.  .A.
// S.M  M.S

func (g Grid) isEdge(c Coordinate) bool {
	return c.x < 1 || c.y < 1 || c.x >= g.width-1 || c.y >= g.height-1
}

func ugliestCrossMASCheck(lines []string) int {
	grid := NewGrid(lines)
	foundLetters := grid.findAll("A")

	xMasCount := 0
	for _, letter := range foundLetters {
		if grid.isEdge(letter) {
			continue
		}
		var cornerPositions []Coordinate
		for _, cornerVector := range cornerVectors {
			cornerPositions = append(cornerPositions, letter.add(cornerVector))
		}

		exactlyTwo := 0
		for i, corner := range cornerPositions {
			oppositeIndex := (i + 2) % 4
			if strings.Compare(grid.getLetterByCoordinate(corner), "M") == 0 {
				if strings.Compare(grid.getLetterByCoordinate(cornerPositions[oppositeIndex]), "S") == 0 {
					exactlyTwo += 1
				}
			}
		}
		if exactlyTwo == 2 {
			xMasCount += 1
		}
	}

	return xMasCount
}

func countWordsInLines(word string, lines []string) int {
	grid := NewGrid(lines)
	count := 0
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			count += grid.countWordsForCoordinate(word, Coordinate{x, y})
		}
	}
	log.Println("Final result", count)
	return count
}

func Run() {
	file := util.OpenFileOrPanicPlz("./4/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	result := countWordsInLines("XMAS", lines)
	fmt.Println("4.1 - XMAS count", result)

	result2 := ugliestCrossMASCheck(lines)
	fmt.Println("4.2 - X-MAS count", result2)

}
