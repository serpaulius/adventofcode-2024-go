package sixth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
)

func sumOfMiddleValues(lines []string) int64 {
	return 0
}

type Guard struct {
	position  grid.Coordinate
	direction grid.Coordinate
}

type Obstrucion struct {
	position grid.Coordinate
}

func Run() {
	file := util.OpenFileOrPanicPlz("./5/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	result1 := sumOfMiddleValues(lines)
	fmt.Println("6.1 - pages and rules", result1)

}
