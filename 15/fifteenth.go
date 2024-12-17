package fifteenth

import (
	"adventofcode/2024-go/grid"
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strings"
)

type MapObject interface {
	position() grid.Coordinate
}

type Robot struct {
	p grid.Coordinate
}

type Box struct {
	p grid.Coordinate
}

type Wall struct {
	p grid.Coordinate
}

func (r Robot) position() grid.Coordinate {
	return r.p
}
func (r Box) position() grid.Coordinate {
	return r.p
}
func (r Wall) position() grid.Coordinate {
	return r.p
}

type World struct {
	g       *grid.Grid
	objects [][]MapObject
	robot   *Robot
}

func (w *World) getObj(c grid.Coordinate) MapObject {
	if w.g.IsValidCoord(c) {
		return w.objects[c.X][c.Y]
	}
	return nil
}

func (w *World) sumOfGPS() int {
	var sum int
	for x := 0; x < w.g.Width; x++ {
		for y := 0; y < w.g.Height; y++ {
			obj := w.getObj(grid.Coordinate{X: x, Y: y})
			switch obj.(type) {
			case *Box:
				sum += (obj.position().Y)*100 + (obj.position().X)
			}
		}
	}
	return sum
}

func WorldFromLines(lines []string) (*World, []string) {
	var theMap *grid.Grid
	var moves []string
	for i, line := range lines {
		if line == "" {
			theMap = grid.GridFromLines(lines[:i])
			moves = lines[i+1:]
			moves = strings.Split(strings.Join(moves, ""), "")
			break
		}
	}

	world := &World{g: theMap, objects: make([][]MapObject, 0)}
	for x := 0; x < theMap.Width; x++ {
		world.objects = append(world.objects, make([]MapObject, 0))
		for y := 0; y < theMap.Height; y++ {
			coord := grid.Coordinate{X: x, Y: y}
			object := theMap.GetLetterByCoordinate(coord)
			var mapObject MapObject
			if object == "@" {
				world.robot = &Robot{p: coord}
				mapObject = world.robot
			}
			if object == "O" {
				mapObject = &Box{p: coord}
			}
			if object == "#" {
				mapObject = &Wall{p: coord}
			}
			if object == "." {
				mapObject = nil
			}
			world.objects[x] = append(world.objects[x], mapObject)
		}
	}
	return world, moves
}

func (w *World) moveRobot(directionStr string) bool {
	var direction grid.Coordinate
	switch directionStr {
	case "^":
		direction = grid.SideVectors[0]
	case ">":
		direction = (grid.SideVectors[1])
	case "v":
		direction = (grid.SideVectors[2])
	case "<":
		direction = (grid.SideVectors[3])
	}

	destination := w.robot.p.Add(direction)
	neighbor := w.getObj(destination)
	if neighbor == nil {
		w.objects[w.robot.p.X][w.robot.p.Y] = nil
		w.robot.p = destination
		w.objects[w.robot.p.X][w.robot.p.Y] = w.robot
		return true
	}
	// if neighbor was moved, move robot too
	if w.moveObject(direction, neighbor) {
		w.objects[w.robot.p.X][w.robot.p.Y] = nil
		w.robot.p = destination
		w.objects[w.robot.p.X][w.robot.p.Y] = w.robot
		return true
	}

	return false
}

// fixme: make it more "polymorphic"
func (w *World) moveObject(direction grid.Coordinate, o MapObject) bool {
	if o != nil {
		destination := o.position().Add(direction)
		neighbor := w.getObj(destination)
		switch object := o.(type) {
		case *Box:
			var wasMoved = false
			if neighbor != nil {
				wasMoved = w.moveObject(direction, neighbor)
				if !wasMoved {
					return false
				}
				neighbor = w.getObj(destination)
			}
			if neighbor == nil {
				// fixme: repeated a few times
				w.objects[object.p.X][object.p.Y] = nil
				object.p = destination
				w.objects[object.p.X][object.p.Y] = object
				return true
			}
		case *Wall:
			return false
		}
	}
	return true
}

func (w *World) print() {
	for x := 0; x < w.g.Width; x++ {
		for y := 0; y < w.g.Height; y++ {
			o := w.objects[x][y]
			var letter string
			switch o.(type) {
			case *Box:
				letter = "O"
			case *Wall:
				letter = "#"
			case *Robot:
				letter = "@"
			default:
				letter = "."
			}
			w.g.SetLetterByCoordinate(grid.Coordinate{X: x, Y: y}, letter)
		}
	}
	w.g.String()
}

func Run() {
	file := util.OpenFileOrPanicPlz("./15/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	w, moves := WorldFromLines(input)
	for _, m := range moves {
		log.Println("Move", m)
		w.moveRobot(m)
		// w.print()
	}
	fmt.Println("15.1 sum of GPS coords", w.sumOfGPS())
}
