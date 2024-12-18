package eighteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/labyrinth"
	"adventofcode/2024-go/util"
	"fmt"
	"strconv"
	"strings"
)

func parseAndAddCorruption(line string, g *grid.Grid) {
	str := strings.Split(line, ",")
	x, _ := strconv.ParseInt(str[0], 10, 64)
	y, _ := strconv.ParseInt(str[1], 10, 64)
	g.SetLetterByCoordinate(grid.Coordinate{X: int(x), Y: int(y)}, "#")

}

func preParseData(input []string, amount int) *grid.Grid {
	g := grid.NewGrid(71, 71)
	simulated := 0
	for _, str := range input {
		parseAndAddCorruption(str, g)
		simulated++
		if simulated == amount {
			break
		}
	}
	g.SetLetterByCoordinate(grid.Coordinate{X: 0, Y: 0}, "S")
	g.SetLetterByCoordinate(grid.Coordinate{X: g.Width - 1, Y: g.Height - 1}, "E")

	return g
}

func Run() {
	file := util.OpenFileOrPanicPlz("./18/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	amount := 1024
	g := preParseData(input, amount)

	opt := labyrinth.LabyrinthOptions{ScoreForCorner: 1, ScoreForSameDirection: 1}
	res := labyrinth.FindLowestScorePath(g, opt)

	fmt.Println("18.1 labyrinth again", res)

	firstToCorrupt := ""
	for i := amount + 1; i < len(input); i++ {
		parseAndAddCorruption(input[i], g)
		res := labyrinth.FindLowestScorePath(g, opt)

		// stop if not reachable (approx. more steps needed than coordinates exist)
		if res > 71*71 {
			firstToCorrupt = input[i]
			break
		}
	}

	fmt.Println("18.2 first to corrupt", firstToCorrupt)
}
