package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

// directions up, right, down, left als Array von Point
var directions = []util.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Part1(fileName string) int {
	obstructions, guardInit, xBound, yBound := parseFile(fileName)
	guardTraces, _ := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
	return len(util.Unique(guardTraces))
}

func parseFile(fileName string) ([]util.Point, util.Point, int, int) {
	lines, _ := util.ReadFileAsArray(fileName)
	var obstructions, guardPositions []util.Point
	xBound := len(lines[0])
	yBound := len(lines)
	obstructions, guardPositions = parseObstructionsAndGuardsPositions(lines)
	return obstructions, guardPositions[0], xBound, yBound
}

func parseObstructionsAndGuardsPositions(lines []string) ([]util.Point, []util.Point) {
	var obstructions, guardPositions []util.Point
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				obstructions = append(obstructions, util.Point{X: i, Y: j})
			} else if char == '^' {
				guardPositions = append(guardPositions, util.Point{X: i, Y: j})
			}
		}
	}
	return obstructions, guardPositions
}

func simulateGuardRunning(obstructions []util.Point, guardInit util.Point, xBound, yBound int) ([]util.Point, error) {
	direction := 0
	guardPositions := make([]util.Point, 0)
	guardPositions = append(guardPositions, guardInit)
	for guardPositions[len(guardPositions)-1].InBounds(xBound, yBound) {
		lastPosition := guardPositions[len(guardPositions)-1]
		nextPosition := util.Point{X: lastPosition.X + directions[direction].X, Y: lastPosition.Y + directions[direction].Y}

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
