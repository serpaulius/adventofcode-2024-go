package eighth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
)

func calcNodes(antenna grid.Coordinate, previousAntennas []grid.Coordinate, antinodes *grid.Grid) int64 {
	printedNew := 0
	for _, previousAntenna := range previousAntennas {
		antinode1 := previousAntenna.Add(previousAntenna.Subtract(antenna))
		antinode2 := antenna.Add(antenna.Subtract(previousAntenna))
		if antinodes.IsValidCoord(antinode1) && antinodes.GetLetterByCoordinate(antinode1) != "#" {
			antinodes.SetLetterByCoordinate(antinode1, "#")
			printedNew++
		}
		if antinodes.IsValidCoord(antinode2) && antinodes.GetLetterByCoordinate(antinode2) != "#" {
			antinodes.SetLetterByCoordinate(antinode2, "#")
			printedNew++
		}
	}
	return int64(printedNew)
}

func parseInput(g *grid.Grid) int64 {
	var antennaMap = map[string][]grid.Coordinate{}
	var coordinate = grid.Coordinate{}
	var antinodes = grid.NewGrid(g.Width, g.Height)
	var antiCount int64

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
}
