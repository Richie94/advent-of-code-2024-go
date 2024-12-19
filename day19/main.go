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
		cache := make(map[string]int)
		resolutions := solveLine(line, elements, cache)
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

func solveLine(line string, elements []string, cache map[string]int) int {
	// return cache hits
	if entry, exists := cache[line]; exists {
		return entry
	}

	if line == "" {
		return 1
	}

	solvedBy := 0
	for _, element := range elements {
		// if line starts with, try this recursive
		if strings.HasPrefix(line, element) {
			solvedBy += solveLine(line[len(element):], elements, cache)
		}
	}

	cache[line] = solvedBy

	return solvedBy
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
