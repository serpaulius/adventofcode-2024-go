package sixth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
)

type Guard struct {
	position  grid.Coordinate
	direction grid.Coordinate
}

func prepareGuard(g *grid.Grid) Guard {
	guardPosition := g.FindAll("^")
	if len(guardPosition) != 1 {
		panic(fmt.Errorf("should be a single guard, instead there are %v", guardPosition))
	}
	return Guard{position: guardPosition[0], direction: grid.SideVectors[0]}
}

func rotateVectorClockwise(c grid.Coordinate) grid.Coordinate {
	for i, side := range grid.SideVectors {
		if c == side {
			nextIndex := (i + 1) % 4
			return grid.SideVectors[nextIndex]
		}
	}
	panic("no side matched, check vector")
}

func moveGuardOut(lines []string) int {
	g := grid.GridFromLines(lines)
	guard := prepareGuard(g)

	g.SetLetterByCoordinate(guard.position, ".")

	visitedCoordinates := map[grid.Coordinate]int64{guard.position: 1}

	for {
		nextStep := guard.position.Add(guard.direction)
		if !g.IsValidCoord(nextStep) {
			break
		}

		isObstruction := g.GetLetterByCoordinate(nextStep) == "#"
		if isObstruction {
			guard.direction = rotateVectorClockwise(guard.direction)
		}
		if !isObstruction {
			guard.position = nextStep
			visitedCoordinates[guard.position]++
		}
	}

	log.Println(len(visitedCoordinates), visitedCoordinates)
	return len(visitedCoordinates)
}

func Run() {
	file := util.OpenFileOrPanicPlz("./6/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	positionsVisited := moveGuardOut(lines)
	fmt.Println("6.1 - pages and rules", positionsVisited)

}
