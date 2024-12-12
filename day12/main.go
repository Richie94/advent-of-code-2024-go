package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

type Area struct {
	points []util.Point
	color  string
}

// adjacent should check if the point is adjacent to any point in the area
func (area *Area) adjacent(point util.Point) bool {
	for _, p := range area.points {
		for _, direction := range util.Directions {
			if p.X+direction.X == point.X && p.Y+direction.Y == point.Y {
				return true
			}
		}
	}
	return false
}

func (area *Area) addPoint(point util.Point) {
	area.points = append(area.points, point)
}

func (area *Area) perimeter() int {
	// we have to iterate over all points and check if it has a neighbour inside our points
	perimeter := 0
	for _, p := range area.points {
		for _, direction := range util.Directions {
			neighbour := util.Point{X: p.X + direction.X, Y: p.Y + direction.Y}
			if !slices.Contains(area.points, neighbour) {
				perimeter++
			}
		}
	}
	return perimeter
}

func (area *Area) sides() int {
	// get all adjacent points, use their direction as the point "color"
	edgeAreas := make([]*Area, 0)
	for _, p := range area.points {
		for _, direction := range util.Directions {
			neighbour := util.Point{X: p.X + direction.X, Y: p.Y + direction.Y}
			if !slices.Contains(area.points, neighbour) {
				color := util.DirectionsToStringMap[direction]
				// check if we can merge it with an edge area
				edgeAreas = getAreas(edgeAreas, color, neighbour)
			}
		}
	}
	return len(edgeAreas)
}

func Part1(fileName string) int {
	areas := calculateAreas(fileName)
	sum := 0
	for _, area := range areas {
		amount := len(area.points)
		perimeter := area.perimeter()
		sum += amount * perimeter
	}
	return sum
}

func calculateAreas(fileName string) []*Area {
	lines, _ := util.ReadFileAsArray(fileName)
	areas := make([]*Area, 0)
	for i, line := range lines {
		for j, char := range line {
			color := string(char)
			// check if the point is adjacent to any area
			point := util.Point{X: i, Y: j}
			areas = getAreas(areas, color, point)
		}
	}
	return areas
}

func getAreas(areas []*Area, color string, point util.Point) []*Area {
	adjacentAreas := make([]*Area, 0)
	for _, area := range areas {
		if area.color == color && area.adjacent(point) {
			adjacentAreas = append(adjacentAreas, area)
		}
	}
	// if no adjacent area is found, create a new one
	if len(adjacentAreas) == 0 {
		areas = append(areas, &Area{points: []util.Point{point}, color: color})
	} else if len(adjacentAreas) > 1 {
		// merge the areas
		mergedArea := &Area{points: []util.Point{point}, color: color}
		for _, area := range adjacentAreas {
			mergedArea.points = append(mergedArea.points, area.points...)
			areas = slices.DeleteFunc(areas, func(other *Area) bool {
				return other == area
			})
		}
		areas = append(areas, mergedArea)
	} else {
		// add the point to the adjacent area
		adjacentAreas[0].addPoint(point)
	}
	return areas
}

func Part2(fileName string) int {
	areas := calculateAreas(fileName)
	sum := 0
	for _, area := range areas {
		amount := len(area.points)
		sides := area.sides()
		sum += amount * sides
	}
	return sum
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
