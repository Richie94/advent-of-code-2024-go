package main

import (
	"advent-of-code-2024/util"
	"fmt"
)

type Point struct {
	X, Y int
}

// directions up, right, down, left als Array von Point
var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Part1(fileName string) int {
	obstructions, guardInit, xBound, yBound := parseFile(fileName)
	guardTraces, _ := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
	return len(removeDuplicates(guardTraces)) - 2
}

func simulateGuardRunning(obstructions []Point, guardInit Point, xBound, yBound int) ([]Point, error) {
	direction := 0
	guardPositions := make([]Point, 1)
	guardPositions = append(guardPositions, guardInit)
	for inBounds(guardPositions[len(guardPositions)-1], xBound, yBound) {
		lastPosition := guardPositions[len(guardPositions)-1]
		nextPosition := Point{lastPosition.X + directions[direction].X, lastPosition.Y + directions[direction].Y}

		if contains(obstructions, nextPosition) {
			direction = (direction + 1) % len(directions)
		} else {
			foundAt, err := containsAt(guardPositions, nextPosition)
			if err == nil && foundAt > 0 && guardPositions[foundAt-1] == lastPosition {
				return guardPositions, fmt.Errorf("guard is running in circles")
			}
			guardPositions = append(guardPositions, nextPosition)
		}
	}
	return guardPositions, nil
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

func inBounds(guard Point, xBound, yBound int) bool {
	return guard.X >= 0 && guard.X < xBound && guard.Y >= 0 && guard.Y < yBound
}

func contains(s []Point, e Point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsAt(s []Point, e Point) (int, error) {
	for i, a := range s {
		if a == e {
			return i, nil
		}
	}
	return -1, fmt.Errorf("element not found")
}

func removeDuplicates(elements []Point) []Point {
	seen := make(map[string]bool)
	var result []Point

	for _, element := range elements {
		key := sliceToString(element)
		if !seen[key] {
			seen[key] = true
			result = append(result, element)
		}
	}
	return result
}

// sliceToString konvertiert ein Point in eine string Darstellung, die als Schlüssel in einer Map verwendet werden kann.
func sliceToString(p Point) string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func Part2(fileName string) int {
	obstructions, guardInit, xBound, yBound := parseFile(fileName)
	sum := 0
	guardPositions, _ := simulateGuardRunning(obstructions, guardInit, xBound, yBound)
	guardPositions = removeDuplicates(guardPositions)

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
