package fourteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"regexp"
)

type Robot struct {
	position grid.Coordinate
	velocity grid.Coordinate
}

func parseInput(lines []string) []Robot {
	var robots []Robot
	coordFinder, _ := regexp.Compile(`-?\d{1,}`)
	for i := 0; i < len(lines); i++ {
		arr := coordFinder.FindAllString(lines[i], 4)
		robot := Robot{position: grid.CoordinateFromArray(arr[:2]), velocity: grid.CoordinateFromArray(arr[2:])}
		robots = append(robots, robot)
	}
	return robots
}

func safetyFactor(g *grid.Grid, robots []Robot) int {
	quadrantX := (g.Width - 1) / 2
	quadrantY := (g.Height - 1) / 2
	var q1, q2, q3, q4 int
	for _, robot := range robots {
		x := robot.position.X
		y := robot.position.Y
		if x < quadrantX {
			if y < quadrantY {
				q1++
			}
			if y > quadrantY {
				q3++
			}
		}
		if x > quadrantX {
			if y < quadrantY {
				q2++
			}
			if y > quadrantY {
				q4++
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func Run() {
	file := util.OpenFileOrPanicPlz("./14/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	const SECONDS = 100
	const WIDTH = 101  //11
	const HEIGHT = 103 //7
	robots := parseInput(input)
	gstart := grid.NewGrid(WIDTH, HEIGHT)

	for _, robot := range robots {
		gstart.SetLetterByCoordinate(robot.position, "X")
	}
	// gstart.String()
	g := grid.NewGrid(WIDTH, HEIGHT)

	for i, robot := range robots {
		newPos := robot.position
		for i := 0; i < SECONDS; i++ {
			newPos = newPos.Add(robot.velocity)
		}

		robots[i].position = g.WrapAroundEdge(newPos)
		g.SetLetterByCoordinate(robot.position, "X")
	}
	// g.String()
	log.Println(robots)
	sf := safetyFactor(g, robots)
	fmt.Println("14.2 robots", sf)
}
