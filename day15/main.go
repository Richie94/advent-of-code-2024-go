package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strings"
)

type Grid struct {
	xMax, yMax int
	robot      *util.Point
	grid       map[util.Point]string
}

func (grid *Grid) score() int {
	// calculate the score of the grid
	score := 0
	for y := 0; y < grid.yMax; y++ {
		for x := 0; x < grid.xMax; x++ {
			if grid.grid[util.Point{X: x, Y: y}] == "O" {
				score += y*100 + x
			} else if grid.grid[util.Point{X: x, Y: y}] == "[" {
				score += y*100 + x
			}
		}
	}
	return score
}

func (grid *Grid) debug() {
	for y := 0; y < grid.yMax; y++ {
		for x := 0; x < grid.xMax; x++ {
			if grid.robot.X == x && grid.robot.Y == y {
				fmt.Print("@")
			} else {
				if object, ok := grid.grid[util.Point{X: x, Y: y}]; ok {
					fmt.Print(object)
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func (grid *Grid) moveRobot(x, y int) {
	// check if the next position is  free, then just move
	nextPos := util.Point{X: grid.robot.X + x, Y: grid.robot.Y + y}
	object, ok := grid.grid[nextPos]
	if !ok {
		grid.robot.X += x
		grid.robot.Y += y
	} else {
		// if object is a wall (#) we cannot move
		if object == "#" {
			return
		} else if object == "O" {
			// it must be a box
			// check where the next free position in current direction is
			nextObj := "O"
			freePos := util.Point{X: nextPos.X, Y: nextPos.Y}
			for nextObj == "O" {
				freePos.X += x
				freePos.Y += y
				nextObj = grid.grid[freePos]
			}

			if nextObj == "#" {
				return
			} else {
				// move the box
				grid.grid[freePos] = "O"
				// remove nextPos from grid
				delete(grid.grid, nextPos)
				// move the robot
				grid.robot = &nextPos
			}
		} else {
			// must be a box [ or ]
			// we need to collect all related boxes wich will be pushed and check if they have all free space behind them
			// or if there is a wall
			boxPoints := make([]util.Point, 0)
			newBoxPoints := make([]util.Point, 0)
			if object == "[" {
				newBoxPoints = append(newBoxPoints, nextPos)
				newBoxPoints = append(newBoxPoints, util.Point{X: nextPos.X + 1, Y: nextPos.Y})
			} else {
				newBoxPoints = append(newBoxPoints, nextPos)
				newBoxPoints = append(newBoxPoints, util.Point{X: nextPos.X - 1, Y: nextPos.Y})
			}
			for len(newBoxPoints) > 0 {
				// get the first element
				boxPoint := newBoxPoints[0]
				boxPoints = append(boxPoints, boxPoint)
				newBoxPoints = newBoxPoints[1:]
				// check if in the direction behind this box is a wall or another box
				nextPoint := util.Point{X: boxPoint.X + x, Y: boxPoint.Y + y}
				nextObj, ok := grid.grid[nextPoint]
				if !ok {
					// is empty, so everything fine
				} else {
					// check what kind of object we have
					if nextObj == "#" {
						// dont do anything, cant push
						return
					} else if nextObj == "[" {
						// add this and the element to the right to the newBoxPoints
						if !slices.Contains(boxPoints, nextPoint) && !slices.Contains(newBoxPoints, nextPoint) {
							newBoxPoints = append(newBoxPoints, nextPoint)
						}

						pointRight := util.Point{X: nextPoint.X + 1, Y: nextPoint.Y}
						if !slices.Contains(boxPoints, pointRight) && !slices.Contains(newBoxPoints, pointRight) {
							newBoxPoints = append(newBoxPoints, pointRight)
						}
					} else if nextObj == "]" {
						// add this and the element to the left to the newBoxPoints
						if !slices.Contains(boxPoints, nextPoint) && !slices.Contains(newBoxPoints, nextPoint) {
							newBoxPoints = append(newBoxPoints, nextPoint)
						}

						pointLeft := util.Point{X: nextPoint.X - 1, Y: nextPoint.Y}
						if !slices.Contains(boxPoints, pointLeft) && !slices.Contains(newBoxPoints, pointLeft) {
							newBoxPoints = append(newBoxPoints, pointLeft)
						}
					}
				}
			}

			// if we didnt break so far, we can move all boxpoints one step to x y direction and move the robot too
			newGrid := make(map[util.Point]string)
			// add all boxpoints to the newGrid
			for _, boxPoint := range boxPoints {
				newGrid[util.Point{X: boxPoint.X + x, Y: boxPoint.Y + y}] = grid.grid[boxPoint]
			}
			// add all other points from the grid to the newGrid which are not in boxPoints
			for point, object := range grid.grid {
				if !slices.Contains(boxPoints, point) {
					newGrid[point] = object
				}
			}
			grid.grid = newGrid
			grid.robot.X += x
			grid.robot.Y += y
		}
	}
}

func Part1(fileName string) int {
	return solve(fileName, false)
}

func solve(fileName string, wide bool) int {
	text, _ := util.ReadFileAsString(fileName)
	textSplit := strings.Split(text, "\n\n")
	gridString := textSplit[0]
	moveString := textSplit[1]
	grid := parseGrid(gridString, wide)
	for _, move := range moveString {
		switch string(move) {
		case "^":
			grid.moveRobot(0, -1)
		case "v":
			grid.moveRobot(0, 1)
		case ">":
			grid.moveRobot(1, 0)
		case "<":
			grid.moveRobot(-1, 0)
		}
	}
	return grid.score()
}

func parseGrid(gridString string, wide bool) *Grid {
	grid := Grid{robot: &util.Point{}, grid: make(map[util.Point]string)}
	lines := strings.Split(gridString, "\n")
	grid.yMax = len(lines)
	grid.xMax = len(lines[0])
	if wide {
		grid.xMax *= 2
	}
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "@" {
				// set the robot position
				if wide {
					grid.robot = &util.Point{X: 2 * x, Y: y}
				} else {
					grid.robot = &util.Point{X: x, Y: y}
				}
			} else if string(char) == "#" || string(char) == "O" {
				if wide {
					if string(char) == "#" {
						grid.grid[util.Point{X: 2 * x, Y: y}] = "#"
						grid.grid[util.Point{X: 2*x + 1, Y: y}] = "#"
					} else {
						grid.grid[util.Point{X: 2 * x, Y: y}] = "["
						grid.grid[util.Point{X: 2*x + 1, Y: y}] = "]"
					}
				} else {
					grid.grid[util.Point{X: x, Y: y}] = string(char)
				}
			}
		}
	}
	return &grid
}

func Part2(fileName string) int {
	return solve(fileName, true)
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
