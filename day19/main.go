package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"sort"
	"strings"
)

func Part1(fileName string) int {
	return solve(fileName, true)
}

func solve(fileName string, part1 bool) int {
	text, _ := util.ReadFileAsString(fileName)
	// line 1 are all the elements splitted by komma
	split := strings.Split(text, "\n\n")
	elements := strings.Split(split[0], ", ")
	// sort elements by size
	sort.Slice(elements, func(i, j int) bool {
		return len(elements[i]) > len(elements[j])
	})
	sum := 0
	for _, line := range strings.Split(split[1], "\n") {
		rejectedCache := make(map[string]bool)
		acceptedCache := make(map[string]int)
		resolutions := solveLine(line, elements, rejectedCache, acceptedCache)
		if part1 {
			if resolutions > 0 {
				sum++
			}
		} else {
			sum += resolutions
		}

	}
	return sum
}

func solveLine(line string, elements []string, rejectedCache map[string]bool, acceptedCache map[string]int) int {
	if rejectedCache[line] {
		return 0
	}
	if acceptedCache[line] > 0 {
		return acceptedCache[line]
	}
	if line == "" {
		return 1
	}

	solvedBy := 0
	for _, element := range elements {
		// if line starts with, try this recursive
		if strings.HasPrefix(line, element) {
			solvedBy += solveLine(line[len(element):], elements, rejectedCache, acceptedCache)
		}
	}

	if solvedBy > 0 {
		acceptedCache[line] = solvedBy
		return solvedBy
	}

	rejectedCache[line] = true

	return 0
}

func Part2(fileName string) int {
	return solve(fileName, false)
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
