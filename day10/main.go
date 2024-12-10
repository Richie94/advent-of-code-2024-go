package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

func Part1(fileName string) int {
	sum := solveTrails(fileName, true)
	return sum
}

func solveTrails(fileName string, onlyOnce bool) int {
	trailMap := generateTrailMap(fileName)
	// for every zero, try to find all paths to a 9
	scores := make([]int, 0)
	for startPoint, value := range trailMap {
		if value != 0 {
			continue
		}
		// we can use a breadth first search for this
		// we have to keep track of the visited nodes
		visited := make(map[util.Point]int)
		// we have to keep track of the nodes we have to visit
		toVisit := make([]util.Point, 0)
		toVisit = append(toVisit, startPoint)
		foundNines := 0
		for len(toVisit) > 0 {
			// get the first element of the toVisit list
			current := toVisit[0]
			if trailMap[current] == 9 {
				foundNines++
			}
			visited[current] = trailMap[current]
			toVisit = toVisit[1:]
			// check for all directions
			for _, direction := range []util.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				next := util.Point{X: current.X + direction.X, Y: current.Y + direction.Y}
				// check if the next point is in the trail map and therefore in bounds
				if _, ok := trailMap[next]; !ok {
					continue
				}
				// check if the point is exactly +1 to the current value
				if trailMap[next] != trailMap[current]+1 {
					continue
				}
				// check if we have not visited this node yet
				if _, ok := visited[next]; ok && onlyOnce {
					continue
				}
				// add the point to the toVisit list
				toVisit = append(toVisit, next)
			}
		}
		scores = append(scores, foundNines)
	}

	sum := 0
	for _, score := range scores {
		sum += score
	}
	return sum
}

func generateTrailMap(fileName string) map[util.Point]int {
	// construct an [][]int from the input file
	lines, _ := util.ReadFileAsArray(fileName)
	trailMap := make(map[util.Point]int)
	// iterate over the lines
	for i, line := range lines {
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			trailMap[util.Point{X: i, Y: j}] = value
		}
	}
	return trailMap
}

func Part2(fileName string) int {
	return solveTrails(fileName, false)
}

func main() {
	filename := "day10/data/test.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
