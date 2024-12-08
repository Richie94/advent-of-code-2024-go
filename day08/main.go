package main

import (
	"advent-of-code-2024/util"
	"fmt"
)

func Part1(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	// we collect a map of string to array of points
	xMax := len(lines)
	yMax := len(lines[0])
	antennaMap := calculateAntennaMap(lines)
	sum := 0
	// for every antenna, we iterate over all other antennas get the antinodes
	antiNodes := findAntiNodes(antennaMap, xMax, yMax, true)
	sum += len(util.Unique(antiNodes))
	return sum
}

func calculateAntennaMap(lines []string) map[int32][]util.Point {
	antennaMap := make(map[int32][]util.Point)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				// we found an antenna, add it to the map
				antennaMap[char] = append(antennaMap[char], util.Point{X: i, Y: j})
			}
		}
	}
	return antennaMap
}

func findAntiNodes(antennaMap map[int32][]util.Point, xMax int, yMax int, cutoff bool) []util.Point {
	antiNodes := make([]util.Point, 0)
	for _, points := range antennaMap {
		for i, p1 := range points {
			if !cutoff {
				// add the point itself if we have that part 2 multiplier
				antiNodes = append(antiNodes, p1)
			}
			for j, p2 := range points {
				if i <= j {
					continue
				}
				// there is a line between those points,
				//which we double in each direction to get the positions of the antinodes
				vector := util.Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
				// direct to p1
				antiNodeInPoint1Dir := util.Point{X: p1.X - vector.X, Y: p1.Y - vector.Y}
				for antiNodeInPoint1Dir.InBounds(xMax, yMax) {
					antiNodes = append(antiNodes, antiNodeInPoint1Dir)
					if cutoff {
						break
					}
					antiNodeInPoint1Dir = util.Point{X: antiNodeInPoint1Dir.X - vector.X, Y: antiNodeInPoint1Dir.Y - vector.Y}
				}

				antiNodeInPoint2Dir := util.Point{X: p2.X + vector.X, Y: p2.Y + vector.Y}
				for antiNodeInPoint2Dir.InBounds(xMax, yMax) {
					antiNodes = append(antiNodes, antiNodeInPoint2Dir)
					if cutoff {
						break
					}
					antiNodeInPoint2Dir = util.Point{X: antiNodeInPoint2Dir.X + vector.X, Y: antiNodeInPoint2Dir.Y + vector.Y}
				}
			}
		}
	}
	return antiNodes
}

func Part2(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	// we collect a map of string to array of points
	xMax := len(lines)
	yMax := len(lines[0])
	antennaMap := calculateAntennaMap(lines)
	sum := 0
	// for every antenna, we iterate over all other antennas get the antinodes
	antiNodes := findAntiNodes(antennaMap, xMax, yMax, false)
	sum += len(util.Unique(antiNodes))
	return sum
}

func main() {
	filename := "day08/data/test.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
