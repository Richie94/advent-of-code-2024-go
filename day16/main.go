package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

type Maze struct {
	freeTiles []util.Point
	start     util.Point
	end       util.Point
}

func Part1(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	maze := parseMaze(lines)
	// fine the way from start to end with A*
	startMove := Move{point: maze.start, direction: util.Point{X: 0, Y: 1}}
	gscores := getGscores(startMove, maze)
	score := 10000000
	for m, value := range gscores {
		if m.point == maze.end && value < score {
			score = value
		}
	}
	return score
}

func getGscores(startMove Move, maze *Maze) map[Move]int {
	openList := make([]Move, 0)
	closedList := make([]Move, 0)
	openList = append(openList, startMove)
	gScores := make(map[Move]int)
	gScores[startMove] = 0
	fScores := make(map[Move]int)
	fScores[startMove] = heuristic(startMove, maze.end)
	for len(openList) > 0 {
		current := openList[0]
		for _, tile := range openList {
			if fScores[tile] < fScores[current] {
				current = tile
			}
		}
		if current.point == maze.end {
			return gScores
		}
		openList = slices.DeleteFunc(openList, func(p Move) bool {
			return p == current
		})
		closedList = append(closedList, current)
		for _, moveWithCost := range getNeighbours(current, maze) {
			move := Move{point: moveWithCost.point, direction: moveWithCost.direction}
			if slices.Contains(closedList, move) {
				continue
			}
			tentativeGScore := gScores[current] + moveWithCost.cost

			if !slices.Contains(openList, move) {
				openList = append(openList, move)
			} else if tentativeGScore >= gScores[move] {
				continue
			}
			gScores[move] = tentativeGScore
			fScores[move] = gScores[move] + heuristic(move, maze.end)
		}
	}

	return gScores
}

type Move struct {
	point     util.Point
	direction util.Point
}

type MoveWithCost struct {
	point     util.Point
	direction util.Point
	cost      int
}

func getNeighbours(move Move, maze *Maze) []MoveWithCost {
	neighbours := make([]MoveWithCost, 0)
	neighbour := util.Point{X: move.point.X + move.direction.X, Y: move.point.Y + move.direction.Y}
	if slices.Contains(maze.freeTiles, neighbour) {
		neighbours = append(neighbours, MoveWithCost{point: neighbour, direction: move.direction, cost: 1})
	}
	// also add the moves for turning right or left
	left := util.TurnLeft(move.direction)
	if slices.Contains(maze.freeTiles, util.Point{X: move.point.X + left.X, Y: move.point.Y + left.Y}) {
		neighbours = append(neighbours, MoveWithCost{point: move.point, direction: left, cost: 1000})
	}
	right := util.TurnRight(move.direction)
	if slices.Contains(maze.freeTiles, util.Point{X: move.point.X + right.X, Y: move.point.Y + right.Y}) {
		neighbours = append(neighbours, MoveWithCost{point: move.point, direction: right, cost: 1000})
	}
	return neighbours
}

func heuristic(a Move, b util.Point) int {
	walkPart := util.AbsInt(a.point.X-b.X) + util.AbsInt(a.point.Y-b.Y)
	turnPart := 0
	if util.AbsInt(a.point.X-b.X) > 0 && util.AbsInt(a.point.Y-b.Y) > 0 {
		turnPart++
	}
	return walkPart + 1000*turnPart
}

func parseMaze(lines []string) *Maze {
	maze := &Maze{}
	for i, line := range lines {
		for j, char := range line {
			if char == '.' {
				maze.freeTiles = append(maze.freeTiles, util.Point{X: i, Y: j})
			} else if char == 'S' {
				maze.start = util.Point{X: i, Y: j}
				maze.freeTiles = append(maze.freeTiles, util.Point{X: i, Y: j})
			} else if char == 'E' {
				maze.end = util.Point{X: i, Y: j}
				maze.freeTiles = append(maze.freeTiles, util.Point{X: i, Y: j})
			}
		}
	}
	return maze

}

func Part2(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	maze := parseMaze(lines)
	startMove := Move{point: maze.start, direction: util.Point{X: 0, Y: 1}}
	gscoresA := getGscores(startMove, maze)

	shortestPathScore := 10000000
	shortestDir := util.Point{X: 0, Y: 1}
	for m, value := range gscoresA {
		if m.point == maze.end && value < shortestPathScore {
			shortestPathScore = value
			shortestDir = m.direction
		}
	}

	mazeReverse := &Maze{freeTiles: maze.freeTiles, start: maze.end, end: maze.start}
	startMoveReverse := Move{point: maze.end, direction: util.TurnLeft(util.TurnLeft(shortestDir))}
	gscoresB := getGscores(startMoveReverse, mazeReverse)

	pointsOnPath := make([]util.Point, 0)
	for m, value := range gscoresA {
		for m2, value2 := range gscoresB {
			if m.point == m2.point {
				s := value + value2
				if shortestPathScore == s && !slices.Contains(pointsOnPath, m.point) {
					pointsOnPath = append(pointsOnPath, m.point)
				}
			}
		}
	}
	// debug
	xMax := len(lines)
	yMax := len(lines[0])
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if slices.Contains(pointsOnPath, util.Point{X: x, Y: y}) {
				fmt.Print("X")
			} else if slices.Contains(maze.freeTiles, util.Point{X: x, Y: y}) {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	return len(pointsOnPath)
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
