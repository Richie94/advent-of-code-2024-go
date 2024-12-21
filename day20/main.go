package main

import (
	"advent-of-code-2024/util"
	"slices"
	"strings"
)

func Part1(fileName string) int {
	testMode := strings.Contains(fileName, "test")
	minSave := 100
	if testMode {
		minSave = 20
	}
	return solve(fileName, 2, minSave)
}

func solve(fileName string, cheatDistance, minSave int) int {
	lines, _ := util.ReadFileAsArray(fileName)
	xMax := len(lines)
	yMax := len(lines[0])
	start := util.Point{X: 0, Y: 0}
	end := util.Point{X: xMax, Y: yMax}
	blocks := make([]util.Point, 0)
	for x, line := range lines {
		for y, char := range line {
			if char == '#' {
				blocks = append(blocks, util.Point{X: x, Y: y})
			} else if char == 'S' {
				start = util.Point{X: x, Y: y}
			} else if char == 'E' {
				end = util.Point{X: x, Y: y}
			}
		}
	}
	path := getShortestPath(start, end, blocks, xMax, yMax)
	// check for every point in the path if there is a cheat possible,
	// e.g. if there is a point after it in manhatten distance of cheatDistance
	sum := 0
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		for j := i + minSave; j < len(path); j++ {
			next := path[j]
			manDist := manhattan(current, next)

			if manDist <= cheatDistance {
				// check whats the difference between i and j (normal distance) vs what we take now with manhattan cheat
				cheat := j - i - manDist
				if minSave <= cheat {
					sum++
				}
			}
		}
	}
	return sum
}

func getShortestPath(startPoint, endPoint util.Point, walls []util.Point, xMax, yMax int) []util.Point {
	// there is only one path
	path := make([]util.Point, 0)
	path = append(path, startPoint)
	for {
		current := path[len(path)-1]
		if current == endPoint {
			return path
		}
		// get next neighbour who is not before me
		neighbours := getNeighbours(current, walls, xMax, yMax)
		for _, neighbour := range neighbours {
			if len(path) > 1 && path[len(path)-2] == neighbour {
				continue
			}
			path = append(path, neighbour)
			break
		}
	}
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

func manhattan(a, b util.Point) int {
	return util.AbsInt(a.X-b.X) + util.AbsInt(a.Y-b.Y)
}

func Part2(fileName string) int {
	testMode := strings.Contains(fileName, "test")
	minSave := 100
	if testMode {
		minSave = 50
	}
	return solve(fileName, 20, minSave)
}
