package fourteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"
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

func safetyFactor(width int, height int, robots []Robot) int {
	quadrantX := (width - 1) / 2
	quadrantY := (height - 1) / 2
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

func increaseNumber(str string) string {
	if str == "" || str == "." {
		return "1"
	}
	i, _ := strconv.ParseInt(str, 10, 64)
	return strconv.FormatInt(i+1, 10)
}

func saveImage(g *grid.Grid, second int) {
	width, height := g.Width, g.Height
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelValue := 0
			letter := g.GetLetterByCoordinate(grid.Coordinate{X: x, Y: y})
			if letter != "" {
				pixelValue = 255
			}

			img.Set(x, y, color.RGBA{R: uint8(pixelValue), G: uint8(pixelValue), B: uint8(pixelValue), A: uint8(255)})
		}
	}

	file, err := os.Create("./images/" + strconv.FormatInt(int64(second), 10) + ".png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

func Run() {
	file := util.OpenFileOrPanicPlz("./14/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	const TRY_2ND_PART = false
	var seconds = 100
	const WIDTH = 101
	const HEIGHT = 103
	robots := parseInput(input)
	robots2 := parseInput(input)
	g := grid.NewGrid(WIDTH, HEIGHT)

	for i, robot := range robots {
		newPos := robot.position
		for i := 0; i < seconds; i++ {
			newPos = newPos.Add(robot.velocity)
		}
		robots[i].position = g.WrapAroundEdge(newPos)
	}
	log.Println(robots)
	sf := safetyFactor(WIDTH, HEIGHT, robots)
	fmt.Println("14.1 robots", sf)

	if TRY_2ND_PART {
		seconds = 10000
		for i := 0; i < seconds; i++ {
			g = grid.NewGrid(WIDTH, HEIGHT)
			for i, robot := range robots2 {
				newPos := robot.position
				newPos = newPos.Add(robot.velocity)
				newPos = g.WrapAroundEdge(newPos)
				robots2[i].position = newPos
				g.SetLetterByCoordinate(robots2[i].position, increaseNumber(g.GetLetterByCoordinate(newPos)))
			}
			saveImage(g, i)
		}
	}
	fmt.Println("14.2 img index", 6586)
}
