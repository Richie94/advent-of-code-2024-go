package main

import (
	"advent-of-code-2024/util"
	"fmt"
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
	// idea: run A* algorithm, reconstruct the path
	path := getShortestPath(start, end, blocks, xMax, yMax)
	// check for every point in the path if there is a cheat possible, e.g. if there is a point after it with x+/-2 or y+/-2
	sum := 0
	cheatMap := make(map[int]int)
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		for j := i + 1; j < len(path); j++ {
			next := path[j]
			manDist := manhattan(current, next)

			if manDist <= cheatDistance {
				// check whats the difference between i and j
				cheat := j - i - manDist
				cheatMap[cheat] = cheatMap[cheat] + 1
				if minSave <= cheat {
					sum++
				}
			}
		}
	}
	fmt.Println(cheatMap)
	return sum
}

func getShortestPath(startPoint, endPoint util.Point, walls []util.Point, xMax, yMax int) []util.Point {
	openList := make([]util.Point, 0)
	closedList := make([]util.Point, 0)
	openList = append(openList, startPoint)
	gScores := make(map[util.Point]int)
	gScores[startPoint] = 0
	fScores := make(map[util.Point]int)
	fScores[startPoint] = heuristic(startPoint, endPoint)
	parents := make(map[util.Point]util.Point)
	for len(openList) > 0 {
		current := openList[0]
		for _, tile := range openList {
			if fScores[tile] < fScores[current] {
				current = tile
			}
		}
		if current == endPoint {
			// reconstruct the path
			path := make([]util.Point, 0)
			for current != startPoint {
				path = append(path, current)
				current = parents[current]
			}
			path = append(path, startPoint)
			// now reverse the list
			slices.Reverse(path)
			return path
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
			parents[neighbour] = current
			gScores[neighbour] = tentativeGScore
			fScores[neighbour] = gScores[neighbour] + heuristic(neighbour, endPoint)
		}
	}

	return []util.Point{}
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
	return manhattan(a, b)
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
