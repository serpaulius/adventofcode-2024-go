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

func findGuard(g *grid.Grid) *Guard {
	guardPosition := g.FindAll("^")
	if len(guardPosition) != 1 {
		panic(fmt.Errorf("should be a single guard, instead there are %v", guardPosition))
	}
	return &Guard{position: guardPosition[0], direction: grid.SideVectors[0]}
}

type arr2d [][]int64

var visitedCoordinates arr2d
var obstaclesHit arr2d

func moveGuardOut(g *grid.Grid, guard Guard) (int, bool) {
	visitedCoordinates = make(arr2d, g.Width)
	obstaclesHit = make(arr2d, g.Width)
	for i := 0; i < g.Width; i++ {
		visitedCoordinates[i] = make([]int64, g.Height)
		obstaclesHit[i] = make([]int64, g.Height)
	}
	visitedCoordinates[guard.position.X][guard.position.Y] = 1
	visitedSum := 1

	var nextStep grid.Coordinate
	for {
		nextStep = guard.position.Add(guard.direction)
		if !g.IsValidCoord(nextStep) {
			break
		}

		if obstaclesHit[nextStep.X][nextStep.Y] > 1 {
			return 0, false
		}

		isObstruction := g.GetLetterByCoordinate(nextStep) == "#"
		if isObstruction {
			obstaclesHit[nextStep.X][nextStep.Y]++
			guard.direction = guard.direction.NextSideVector()
		}
		if !isObstruction {
			guard.position = nextStep
			// count first times separately
			if visitedCoordinates[guard.position.X][guard.position.Y] == 0 {
				visitedSum++
			}
			visitedCoordinates[guard.position.X][guard.position.Y]++
		}
	}
	return visitedSum, true
}

func findPositionsForLoopObstacle(g *grid.Grid, guard Guard) int {
	loopPositions := 0
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			coord := grid.Coordinate{X: x, Y: y}
			object := g.GetLetterByCoordinate(coord)
			if object == "." {
				g.SetLetterByCoordinate(coord, "#")
				_, err := moveGuardOut(g, guard)
				if !err {
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
	guard := findGuard(g)
	positionsVisited, _ := moveGuardOut(g, *guard)
	fmt.Println("6.1 - guard positions visited", positionsVisited)

	loopPositions := findPositionsForLoopObstacle(g, *guard)
	fmt.Println("6.2 - obstacles for looping a guard", loopPositions)
}
