package sixth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
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

func moveGuardOut(g *grid.Grid) (int, error) {
	guard := prepareGuard(g)
	visitedCoordinates := map[grid.Coordinate]int64{guard.position: 1}
	obstaclesHit := map[grid.Coordinate]int64{}

	for {
		nextStep := guard.position.Add(guard.direction)
		if !g.IsValidCoord(nextStep) {
			break
		}

		isLoop := obstaclesHit[nextStep] > 1
		if isLoop {
			return 0, fmt.Errorf("guard is stuck")
		}

		isObstruction := g.GetLetterByCoordinate(nextStep) == "#"
		if isObstruction {
			obstaclesHit[nextStep]++
			guard.direction = rotateVectorClockwise(guard.direction)
		}
		if !isObstruction {
			guard.position = nextStep
			visitedCoordinates[guard.position]++
		}
	}
	return len(visitedCoordinates), nil
}

func findPositionsForLoopObstacle(g *grid.Grid) int {
	loopPositions := 0
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			coord := grid.Coordinate{X: x, Y: y}
			object := g.GetLetterByCoordinate(coord)
			if object == "." {
				g.SetLetterByCoordinate(coord, "#")
				_, err := moveGuardOut(g)
				if err != nil {
					loopPositions++
				}
				g.SetLetterByCoordinate(coord, ".")
			}
		}
	}
	return loopPositions
}

func Run() {
	file := util.OpenFileOrPanicPlz("./6/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	g := grid.GridFromLines(lines)
	positionsVisited, _ := moveGuardOut(g)
	fmt.Println("6.1 - guard positions visited", positionsVisited)

	// todo: very inefficient, optimize
	loopPositions := findPositionsForLoopObstacle(g)
	fmt.Println("6.2 - obstacles for looping a guard", loopPositions)
}
