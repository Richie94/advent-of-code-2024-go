package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(fileName string) int {
	xMax := 70
	yMax := 70
	amount := 1024
	if strings.Contains(fileName, "test") {
		xMax = 6
		yMax = 6
		amount = 12
	}

	lines, _ := util.ReadFileAsArray(fileName)
	walls := make([]util.Point, 0)
	for _, line := range lines[:amount] {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[1])
		y, _ := strconv.Atoi(split[0])
		walls = append(walls, util.Point{X: x, Y: y})
	}
	// draw walls
	startPoint := util.Point{X: 0, Y: 0}
	endPoint := util.Point{X: xMax, Y: yMax}
	return getGscores(startPoint, endPoint, walls, xMax, yMax)[endPoint]
}

func debug(xMax int, yMax int, walls []util.Point, me util.Point) {
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			if slices.Contains(walls, util.Point{X: x, Y: y}) {
				fmt.Print("#")
			} else if me.X == x && me.Y == y {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getGscores(startPoint, endPoint util.Point, walls []util.Point, xMax, yMax int) map[util.Point]int {
	openList := make([]util.Point, 0)
	closedList := make([]util.Point, 0)
	openList = append(openList, startPoint)
	gScores := make(map[util.Point]int)
	gScores[startPoint] = 0
	fScores := make(map[util.Point]int)
	fScores[startPoint] = heuristic(startPoint, endPoint)
	for len(openList) > 0 {
		current := openList[0]
		for _, tile := range openList {
			if fScores[tile] < fScores[current] {
				current = tile
			}
		}
		//debug(xMax, yMax, walls, current)
		if current == endPoint {
			return gScores
		}
		openList = slices.DeleteFunc(openList, func(p util.Point) bool {
			return p == current
		})
		closedList = append(closedList, current)
		for _, neighbour := range getNeighbours(current, walls, xMax, yMax) {
			if slices.Contains(closedList, neighbour) {
				continue
			}
			tentativeGScore := gScores[current] + 1

			if !slices.Contains(openList, neighbour) {
				openList = append(openList, neighbour)
			} else if tentativeGScore >= gScores[neighbour] {
				continue
			}
			gScores[neighbour] = tentativeGScore
			fScores[neighbour] = gScores[neighbour] + heuristic(neighbour, endPoint)
		}
	}

	return gScores
}

func getNeighbours(point util.Point, walls []util.Point, xMax, yMax int) []util.Point {
	neighbours := make([]util.Point, 0)
	for _, direction := range util.Directions {
		neighbour := util.Point{X: point.X + direction.X, Y: point.Y + direction.Y}
		if slices.Contains(walls, neighbour) {
			continue
		}
		if neighbour.X < 0 || neighbour.X > xMax || neighbour.Y < 0 || neighbour.Y > yMax {
			continue
		}
		neighbours = append(neighbours, neighbour)
	}
	return neighbours
}

func heuristic(a util.Point, b util.Point) int {
	return util.AbsInt(a.X-b.X) + util.AbsInt(a.Y-b.Y)
}

func Part2(fileName string) string {
	xMax := 70
	yMax := 70

	lines, _ := util.ReadFileAsArray(fileName)
	walls := make([]util.Point, 0)
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[1])
		y, _ := strconv.Atoi(split[0])
		walls = append(walls, util.Point{X: x, Y: y})
	}
	// draw walls
	startPoint := util.Point{X: 0, Y: 0}
	endPoint := util.Point{X: xMax, Y: yMax}
	// test with 200 steps, if we didnt reach the end, we need to increase the amount of steps
	stepAmount := 1024
	decreaseMode := false
	for {
		score := getGscores(startPoint, endPoint, walls[:stepAmount], xMax, yMax)[endPoint]
		if score > 0 && !decreaseMode {
			// we need to go further
			stepAmount += 128
		} else if !decreaseMode {
			decreaseMode = true
			// now we decrease until we are above the score again
		} else {
			if score > 0 {
				answer := walls[stepAmount]
				return strconv.Itoa(answer.Y) + "," + strconv.Itoa(answer.X)
			} else {
				stepAmount -= 1
			}
		}

	}
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
