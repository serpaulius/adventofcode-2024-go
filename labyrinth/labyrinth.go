package labyrinth

import (
	"adventofcode/2024-go/grid"
	"math"
	"slices"
)

type Location struct {
	position  grid.Coordinate
	direction int
}

func traverseMazeBFS(start Location, g *grid.Grid, visited [][][]int64, end grid.Coordinate, o LabyrinthOptions) int64 {
	queue := make([]Location, 0)
	queue = append(queue, Location{position: start.position, direction: start.direction})
	visited[start.position.X][start.position.Y][start.direction] = 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < len(grid.SideVectors); i++ {
			// rotate from current to other directions
			nextDirectionIndex := (current.direction + i) % 4
			nextDirection := grid.SideVectors[nextDirectionIndex]
			nextCell := current.position.Add(nextDirection)
			var cost int64

			// todo: review
			if nextDirectionIndex == current.direction {
				cost = o.ScoreForSameDirection
			} else {
				cost = o.ScoreForCorner
			}

			if !g.IsValidCoord(nextCell) {
				continue
			}
			if letter := g.GetLetterByCoordinate(nextCell); letter == "#" {
				continue
			}

			if visited[nextCell.X][nextCell.Y][nextDirectionIndex] > visited[current.position.X][current.position.Y][current.direction]+cost {
				visited[nextCell.X][nextCell.Y][nextDirectionIndex] = visited[current.position.X][current.position.Y][current.direction] + cost
				if nextDirectionIndex == current.direction {
					queue = slices.Insert(queue, 0, Location{position: nextCell, direction: nextDirectionIndex})
				}
				queue = append(queue, Location{position: nextCell, direction: nextDirectionIndex})
			}
		}
	}
	return slices.Min(visited[end.X][end.Y])
}

type LabyrinthOptions struct {
	ScoreForCorner        int64
	ScoreForSameDirection int64
}

func FindLowestScorePath(g *grid.Grid, o LabyrinthOptions) int64 {
	start := g.FindAll("S")
	end := g.FindAll("E")
	visited := initVisited(g.Width, g.Height)

	// starts facing east
	deer := Location{position: start[0], direction: 1}
	res := traverseMazeBFS(deer, g, visited, end[0], o)

	return res
}

func initVisited(w, h int) [][][]int64 {
	var visited [][][]int64
	for i := 0; i < w; i++ {
		visited = append(visited, make([][]int64, h))
		for j := 0; j < h; j++ {
			visited[i][j] = make([]int64, 4)
			for k := 0; k < 4; k++ {
				visited[i][j][k] = math.MaxInt64
			}
		}
	}
	return visited
}
