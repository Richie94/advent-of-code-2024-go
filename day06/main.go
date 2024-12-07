package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

type Point struct {
	X, Y int
}

// inBounds checks if a point is within the bounds of a grid.
func (p Point) inBounds(xBound, yBound int) bool {
	return p.X >= 0 && p.X < xBound && p.Y >= 0 && p.Y < yBound
}

// directions up, right, down, left als Array von Point
var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Part1(fileName string) int {
	obstructions, guardInit, xBound, yBound := parseFile(fileName)
	guardTraces, _ := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
	return len(util.Unique(guardTraces))
}

func parseFile(fileName string) ([]Point, Point, int, int) {
	lines, _ := util.ReadFileAsArray(fileName)
	var obstructions, guardPositions []Point
	xBound := len(lines[0])
	yBound := len(lines)
	obstructions, guardPositions = parseObstructionsAndGuardsPositions(lines)
	return obstructions, guardPositions[0], xBound, yBound
}

func parseObstructionsAndGuardsPositions(lines []string) ([]Point, []Point) {
	var obstructions, guardPositions []Point
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				obstructions = append(obstructions, Point{i, j})
			} else if char == '^' {
				guardPositions = append(guardPositions, Point{i, j})
			}
		}
	}
	return obstructions, guardPositions
}

func simulateGuardRunning(obstructions []Point, guardInit Point, xBound, yBound int) ([]Point, error) {
	direction := 0
	guardPositions := make([]Point, 0)
	guardPositions = append(guardPositions, guardInit)
	for guardPositions[len(guardPositions)-1].inBounds(xBound, yBound) {
		lastPosition := guardPositions[len(guardPositions)-1]
		nextPosition := Point{lastPosition.X + directions[direction].X, lastPosition.Y + directions[direction].Y}

		if slices.Contains(obstructions, nextPosition) {
			direction = (direction + 1) % len(directions)
		} else {
			foundAt := slices.Index(guardPositions, nextPosition)
			if foundAt > 0 && guardPositions[foundAt-1] == lastPosition {
				return guardPositions, fmt.Errorf("guard is running in circles")
			}
			guardPositions = append(guardPositions, nextPosition)
		}
	}
	// drop last as it out of bounds
	guardPositions = guardPositions[:len(guardPositions)-1]
	return guardPositions, nil
}

func Part2(fileName string) int {
	obstructions, guardInit, xBound, yBound := parseFile(fileName)
	sum := 0
	guardPositions, _ := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
	// drop first as we cant place something there
	guardPositions = guardPositions[1:]
	guardPositions = util.Unique(guardPositions)

	for _, pos := range guardPositions {
		obstructions = append(obstructions, pos)
		_, err := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
		if err != nil {
			sum++
		}
		obstructions = obstructions[:len(obstructions)-1] // remove the added obstruction
	}
	return sum
}

func main() {
	filename := "day06/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
