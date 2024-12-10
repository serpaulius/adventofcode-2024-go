package tenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"strconv"
)

func traverseTrails(g *grid.Grid, coordinate grid.Coordinate, currValue int64, visited *grid.Grid) int64 {
	if currValue == 9 {
		if visited == nil {
			return 1
		} else if visited.GetLetterByCoordinate(coordinate) != "X" {
			visited.SetLetterByCoordinate(coordinate, "X")
			return 1
		}
	}

	var peaksFound int64
	for _, vector := range grid.SideVectors {
		nextCoordinate := coordinate.Add(vector)
		if g.IsValidCoord(nextCoordinate) {
			nextVal, _ := strconv.ParseInt(g.GetLetterByCoordinate(nextCoordinate), 10, 64)
			if nextVal == currValue+1 {
				peaksFound += traverseTrails(g, nextCoordinate, nextVal, visited)
			}
		}
	}
	return peaksFound
}

func parseData(g *grid.Grid) int64 {
	zeros := g.FindAll("0")
	var visitedPeaksPerTrailhead int64
	for _, zero := range zeros {
		visitedCoords := grid.NewGrid(g.Width, g.Height)
		visited := traverseTrails(g, zero, 0, visitedCoords)
		visitedPeaksPerTrailhead += visited
	}
	return visitedPeaksPerTrailhead
}

func trailsPossiblePerPeak(g *grid.Grid) int64 {
	zeros := g.FindAll("0")
	var trailsPerPeak int64
	for _, zero := range zeros {
		visited := traverseTrails(g, zero, 0, nil)
		trailsPerPeak += visited
	}
	return trailsPerPeak
}

//take a map
//find trailheads
//for each position
//  find bigger position
//		for each bigger position
//			find bigger position
//			...
//			until 9 reached
//			add to result
//			go back til last place with another bigger position
// e.g. recursion

func Run() {
	file := util.OpenFileOrPanicPlz("./10/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	g := grid.GridFromLines(lines)
	result1 := parseData(g)
	fmt.Println("10.1 - peaks per trail", result1)
	result2 := trailsPossiblePerPeak(g)
	fmt.Println("10.2 - trails per peak", result2)

}
