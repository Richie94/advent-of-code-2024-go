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
	openList := make([]Move, 0)
	closedList := make([]Move, 0)
	startMove := Move{point: maze.start, direction: util.Point{X: 0, Y: 1}}
	openList = append(openList, startMove)
	parents := make(map[util.Point]util.Point)
	gScores := make(map[Move]int)
	gScores[startMove] = 0
	fScores := make(map[Move]int)
	fScores[startMove] = heustic(maze.start, maze.end)
	for len(openList) > 0 {
		current := openList[0]
		for _, tile := range openList {
			if fScores[tile] < fScores[current] {
				current = tile
			}
		}
		//debug(maze, current)
		if current.point == maze.end {
			// reconstruct path
			path := make([]util.Point, 0)
			path = append(path, current.point)
			currentPoint := current.point
			for {
				currentPoint = parents[currentPoint]
				path = append(path, currentPoint)
				if currentPoint == maze.start {
					break
				}
			}
			return gScores[current]
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
			// check how many 90 degree turns we had to make

			tentativeGScore := gScores[current] + moveWithCost.cost

			if !slices.Contains(openList, move) {
				openList = append(openList, move)
			} else if tentativeGScore >= gScores[move] {
				continue
			}
			if moveWithCost.point != current.point {
				parents[moveWithCost.point] = current.point
			}
			gScores[move] = tentativeGScore
			fScores[move] = gScores[move] + heustic(moveWithCost.point, maze.end)
		}
	}

	return 0
}

func debug(maze *Maze, point Move) {
	xMax := 0
	yMax := 0
	for _, tile := range maze.freeTiles {
		if tile.X+1 > xMax {
			xMax = tile.X + 1
		}
		if tile.Y+1 > yMax {
			yMax = tile.Y + 1
		}
	}
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			if point.point.X == i && point.point.Y == j {
				dirString := util.DirectionsToStringMap[point.direction]
				switch dirString {
				case "up":
					fmt.Print("^")
				case "down":
					fmt.Print("v")
				case "left":
					fmt.Print("<")
				case "right":
					fmt.Print(">")
				}
			} else if slices.Contains(maze.freeTiles, util.Point{X: i, Y: j}) {
				fmt.Print(".")
			} else {
				fmt.Print("#")

			}
		}
		fmt.Println()
	}
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
	neighbours = append(neighbours, MoveWithCost{point: move.point, direction: util.TurnLeft(move.direction), cost: 1000})
	neighbours = append(neighbours, MoveWithCost{point: move.point, direction: util.TurnRight(move.direction), cost: 1000})
	return neighbours
}

func heustic(a, b util.Point) int {
	walkPart := util.AbsInt(a.X-b.X) + util.AbsInt(a.Y-b.Y)
	turnPart := 0
	if util.AbsInt(a.X-b.X) > 0 {
		turnPart++
	}
	if util.AbsInt(a.Y-b.Y) > 0 {
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
	return 0
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
