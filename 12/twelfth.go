package twelfth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
)

type Region struct {
	letter    string
	curr      grid.Coordinate
	perimeter int64
	area      int64
}

var regions []Region = make([]Region, 0)

// var regions = map[grid.Coordinate][]grid.Coordinate
// take a grid
// visit a cell
// check sides
//   visit a cell
//   check sides
//     repeat while nowhere to go

func flood(region *Region, g *grid.Grid, visited *grid.Grid) {
	curr := region.curr
	visited.SetLetterByCoordinate(region.curr, "X")
	region.area++
	for _, side := range grid.SideVectors {
		candidate := curr.Add(side)
		// is candidate out of bounds?
		isValid := g.IsValidCoord(candidate)
		if !isValid {
			region.perimeter++
			continue
		}

		candidateLetter := g.GetLetterByCoordinate(candidate)
		// is candidate end of region?
		if candidateLetter != region.letter {
			region.perimeter++
			continue
		}
		isVisited := visited.GetLetterByCoordinate(candidate) != ""
		// is candidate visited?
		if isVisited {
			continue
		}
		// is candidate same letter?
		if candidateLetter == region.letter {
			region.curr = candidate
			flood(region, g, visited)
		}
	}
}

func traverseLetters(g *grid.Grid, visited *grid.Grid) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			currCoord := grid.Coordinate{X: x, Y: y}
			if visited.GetLetterByCoordinate(currCoord) == "" {
				region := Region{letter: g.GetLetterByCoordinate(currCoord), curr: currCoord}
				flood(&region, g, visited)
				log.Println("one done", region)
				regions = append(regions, region)
			}
		}
	}
}

func Run() {
	file := util.OpenFileOrPanicPlz("./12/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	g := grid.GridFromLines(input)
	visited := grid.NewGrid(g.Width, g.Height)
	traverseLetters(g, visited)
	var sum int64
	for _, region := range regions {
		log.Println(region.area, region.perimeter, region.area*region.perimeter)
		sum += region.area * region.perimeter
	}
	fmt.Println("12.1 price of regions", sum)
}
