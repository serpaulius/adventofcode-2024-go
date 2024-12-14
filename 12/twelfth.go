package twelfth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"log"
)

/*
......
.AAAA.
.BBCD.
.BBCC.
.EEEC.
......
everywhere where A touches something, is a perimeter part
so there are 4 A's with every A touching .'s and B's anc, etc

fixme: I will need to go around all regions, unfortunately
*/
var touchesPerLetter = make(map[string]int64)
var letters = make(map[string]int64)

func traverseLetters(g *grid.Grid) {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			currCoord := grid.Coordinate{X: x, Y: y}
			currLetter := g.GetLetterByCoordinate(currCoord)
			letters[currLetter]++
			for _, v := range grid.SideVectors {
				toCheck := currCoord.Add(v)
				isValid := g.IsValidCoord(toCheck)
				if isValid && g.GetLetterByCoordinate(toCheck) != currLetter || !isValid {
					// this is perimieter, mate
					touchesPerLetter[currLetter]++
				}
			}
		}
	}
}

func Run() {
	file := util.OpenFileOrPanicPlz("./12/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	g := grid.GridFromLines(input)
	traverseLetters(g)
	log.Println(letters)
	log.Println(touchesPerLetter)
	var sum int64
	for letter, count := range letters {
		sum += count * touchesPerLetter[letter]
	}
	// fixme:
	//fmt.Println("12.1 perimeters", sum)

}
