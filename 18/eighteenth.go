package eighteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/labyrinth"
	"adventofcode/2024-go/util"
	"fmt"
	"strconv"
	"strings"
)

func preParseData(input []string) *grid.Grid {
	g := grid.NewGrid(71, 71)
	simulated := 0
	for _, str := range input {
		str := strings.Split(str, ",")
		x, _ := strconv.ParseInt(str[0], 10, 64)
		y, _ := strconv.ParseInt(str[1], 10, 64)
		g.SetLetterByCoordinate(grid.Coordinate{X: int(x), Y: int(y)}, "#")
		simulated++
		if simulated == 1024 {
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
	g := preParseData(input)
	res := labyrinth.FindLowestScorePath(g, labyrinth.LabyrinthOptions{ScoreForCorner: 1, ScoreForSameDirection: 1})

	fmt.Println("18.1 labyrinth again", res)
}
