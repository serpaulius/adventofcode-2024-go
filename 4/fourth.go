package fourth

import (
	grid "adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strings"
)

func countWordsForCoordinate(g *grid.Grid, word string, position grid.Coordinate) int {
	wordLength := len(word)
	count := 0
	for _, vect := range grid.Vectors {
		candidate := ""
		c := position
		for i := 0; i < wordLength; i++ {
			if g.IsValidCoord(c) {
				candidate += g.GetLetterByCoordinate(c)
			}
			c = c.Add(vect)
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
func ugliestCrossMASCheck(lines []string) int {
	g := grid.NewGrid(lines)
	aCoordinates := g.FindAll("A")

	xMasCount := 0
	for _, aCoord := range aCoordinates {
		if !g.IsEdge(aCoord) {
			var cornerPositions []grid.Coordinate
			for _, cornerVector := range grid.CornerVectors {
				cornerPositions = append(cornerPositions, aCoord.Add(cornerVector))
			}

			exactlyTwo := 0
			for i, corner := range cornerPositions {
				value := g.GetLetterByCoordinate(corner)
				oppositeIndex := (i + 2) % 4
				oppositeValue := g.GetLetterByCoordinate(cornerPositions[oppositeIndex])
				if strings.Compare(value, "M") == 0 && strings.Compare(oppositeValue, "S") == 0 {
					exactlyTwo += 1
				}
			}
			if exactlyTwo == 2 {
				xMasCount += 1
			}
		}
	}

	return xMasCount
}

func countWordsInLines(word string, lines []string) int {
	g := grid.NewGrid(lines)
	count := 0
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			count += countWordsForCoordinate(g, word, grid.Coordinate{X: x, Y: y})
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
