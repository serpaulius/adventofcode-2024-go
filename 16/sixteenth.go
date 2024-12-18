package sixteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/labyrinth"
	"adventofcode/2024-go/util"
	"fmt"
)

func Run() {
	file := util.OpenFileOrPanicPlz("./16/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	g := grid.GridFromLines(input)
	res := labyrinth.FindLowestScorePath(g, labyrinth.LabyrinthOptions{ScoreForCorner: 1001, ScoreForSameDirection: 1})

	fmt.Println("16.1 price for reaching the end", res)
}
