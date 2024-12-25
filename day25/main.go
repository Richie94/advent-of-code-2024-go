package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strings"
)

func Part1(fileName string) int {
	text, _ := util.ReadFileAsString(fileName)
	keys := make([][]int, 0)
	locks := make([][]int, 0)
	for _, part := range strings.Split(text, "\n\n") {
		// if first line is all #, then its a lock
		lines := strings.Split(part, "\n")
		isLock := strings.Contains(lines[0], "#")
		// count the amount of # in each column
		counts := make([]int, len(lines[0]))
		// init with -1
		for i := range counts {
			counts[i] = -1
		}
		for _, line := range lines {
			for i, c := range line {
				if c == '#' {
					counts[i]++
				}
			}
		}
		if isLock {
			locks = append(locks, counts)
		} else {
			keys = append(keys, counts)
		}
	}

	// find all matching keys and locks
	matchSum := 0
	for _, key := range keys {
		for _, lock := range locks {
			match := true
			// if they dont add up to 6, they cant be a match
			for idx, _ := range key {
				if key[idx]+lock[idx] >= 6 {
					match = false
					break
				}
			}
			if match {
				matchSum++
			}
		}
	}
	return matchSum
}

func Part2(fileName string) int {
	return 0
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
