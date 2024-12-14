package eighth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
)

var oncePerDirection = true
var includeAllAntennas = false

func tryAddingNewAntinode(where grid.Coordinate, antinodes *grid.Grid) int {
	if antinodes.GetLetterByCoordinate(where) != "#" {
		antinodes.SetLetterByCoordinate(where, "#")
		return 1
	}
	return 0
}

func addAntinodesInDirection(antenna grid.Coordinate, direction grid.Coordinate, antinodes *grid.Grid) int {
	currPosition := antenna.Add(direction)
	var sumAdded = 0
	for antinodes.IsValidCoord(currPosition) {
		sumAdded += tryAddingNewAntinode(currPosition, antinodes)
		currPosition = currPosition.Add(direction)
		if oncePerDirection {
			break
		}
	}
	return sumAdded
}

func calcNodes(antenna grid.Coordinate, previousAntennas []grid.Coordinate, antinodes *grid.Grid) int {
	newAntinodesAdded := 0
	for _, previousAntenna := range previousAntennas {
		vector1 := previousAntenna.Subtract(antenna)
		vector2 := antenna.Subtract(previousAntenna)
		if includeAllAntennas {
			// find shortest natural number vector for going through every coordinate on the way including in between antennas
			vectorGcd := util.Gcd(vector1.X, vector1.Y)
			vector1 = grid.Coordinate{X: vector1.X / vectorGcd, Y: vector1.Y / vectorGcd}
			vector2 = grid.Coordinate{X: vector2.X / vectorGcd, Y: vector2.Y / vectorGcd}

		}
		newAntinodesAdded += addAntinodesInDirection(previousAntenna, vector1, antinodes)
		newAntinodesAdded += addAntinodesInDirection(antenna, vector2, antinodes)

		// antinodes.String()
	}
	return newAntinodesAdded
}

func parseInput(g *grid.Grid) int {
	var antennaMap = map[string][]grid.Coordinate{}
	var coordinate = grid.Coordinate{}
	var antinodes = grid.NewGrid(g.Width, g.Height)
	var antiCount int

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			coordinate.X = x
			coordinate.Y = y
			letter := g.GetLetterByCoordinate(coordinate)
			if letter != "." && letter != "#" {
				if len(antennaMap[letter]) > 0 {
					antiCount += calcNodes(coordinate, antennaMap[letter], antinodes)
				}
				antennaMap[letter] = append(antennaMap[letter], coordinate)
			}
		}
	}
	return antiCount
}

func Run() {
	file := util.OpenFileOrPanicPlz("./8/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	g := grid.GridFromLines(lines)
	antinodeCount := parseInput(g)
	fmt.Println("8.1 - antinodes for antennas in line", antinodeCount)

	// fixme: not nice
	oncePerDirection = false
	includeAllAntennas = true
	antinodeCount2 := parseInput(g)
	fmt.Println("8.2 - antinodes with antennas till end", antinodeCount2)
	// uniqueLocationsInBounds()
}
