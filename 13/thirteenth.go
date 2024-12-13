package thirteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Machine struct {
	a     grid.Coordinate
	b     grid.Coordinate
	prize grid.Coordinate
}

func justMakeCoord(str []string) grid.Coordinate {
	x, _ := strconv.ParseInt(str[0], 10, 64)
	y, _ := strconv.ParseInt(str[1], 10, 64)
	return grid.Coordinate{X: int(x), Y: int(y)}
}

func parseInput(lines []string) []Machine {
	var machines []Machine
	coordFinder, _ := regexp.Compile(`\d{1,}`)
	linesPerMachine := 4
	for i := 0; i <= len(lines)/linesPerMachine; i++ {
		currIndex := i * linesPerMachine
		a := coordFinder.FindAllString(lines[currIndex], 2)
		b := coordFinder.FindAllString(lines[currIndex+1], 2)
		prize := coordFinder.FindAllString(lines[currIndex+2], 2)
		machine := Machine{a: justMakeCoord(a), b: justMakeCoord(b), prize: justMakeCoord(prize)}
		machines = append(machines, machine)
	}
	return machines
}

const priceA int64 = 3
const priceB int64 = 1

// for linear system: x=nA*​ax​+nB*​by; ​y=nA*​ay​+nB*​by
// determinant: det=ax​*by​−ay​*bx​
// count of a: nA​=(x*by​−y*bx)/det​
// count of b: nB​=(y*ax​−x*a​y)/detå
// if det != 0 and both nA nB are natural numbers, it is reachable
// (thanks gipitee for equations)
func cheapify(machine Machine) int64 {
	det := machine.a.X*machine.b.Y - machine.a.Y*machine.b.X
	na := float64(machine.prize.X*machine.b.Y-machine.prize.Y*machine.b.X) / float64(det)
	nb := float64(machine.prize.X*machine.a.Y-machine.prize.Y*machine.a.X) / float64(det)
	if det != 0 && na == float64(int64(na)) && nb == float64(int64(nb)) {
		log.Println("will reach", machine, det, na, nb)
		return util.Abs(int64(na))*priceA + util.Abs(int64(nb))*priceB
	}
	return 0
}

func Run() {
	file := util.OpenFileOrPanicPlz("./13/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	machines := parseInput(input)
	var sum int64
	for _, machine := range machines {
		sum += cheapify(machine)
	}

	fmt.Println("13.2 buttons to press", sum)
}
