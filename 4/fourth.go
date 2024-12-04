package fourth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strings"
)

func transpose(lines []string) []string {
	transposed := make([]string, len(lines[0]))
	for _, line := range lines {
		for j, strChar := range line {
			transposed[j] = transposed[j] + string(strChar)
		}
	}
	fmt.Println(transposed)
	return transposed
}

type Vector struct {
	x int
	y int
}

var vectors []Vector = []Vector{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

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

func (g Grid) getLetter(cx, cy int) string {
	return string(g.lines[cy][cx])
}

func (g Grid) countWordsForCoordinate(word string, x, y int) int {
	wordLength := len(word)
	count := 0
	for _, vect := range vectors {
		candidate := ""
		cx, cy := x, y
		for i := 0; i < wordLength; i++ {
			if cx >= 0 && cy >= 0 && cx < g.width && cy < g.height {
				candidate += g.getLetter(cx, cy)
			}
			cx, cy = cx+vect.x, cy+vect.y
		}
		if strings.Compare(candidate, word) == 0 {
			log.Println(Vector{x, y}, vect, candidate)
			count++
		}
	}
	log.Println("count for", Vector{x, y}, "is", count)
	return count
}

func countWords(lines []string) int {
	grid := NewGrid(lines)
	count := 0
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			count += grid.countWordsForCoordinate("XMAS", x, y)
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

	result := countWords(lines)
	fmt.Println("4.1 - XMAS count", result)

}
