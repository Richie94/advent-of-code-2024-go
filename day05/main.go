package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part1(fileName string) int {
	rules, updatesText := splitInputs(fileName)
	sum := 0
	for _, line := range updatesText {
		intLine, valid := buildUpdateAndCheckValidity(line, rules)

		if valid {
			// add the middle element of the line to the sum
			sum += intLine[len(intLine)/2]
		}
	}
	return sum
}

func Part2(fileName string) int {
	rules, updatesText := splitInputs(fileName)
	sum := 0
	for _, line := range updatesText {
		intLine, valid := buildUpdateAndCheckValidity(line, rules)

		if !valid {
			// fix the line
			sort.Slice(intLine, func(i, j int) bool {
				// compare if we have a match in the rules, if so we may swap
				for _, rule := range rules {
					if intLine[i] == rule[0] && intLine[j] == rule[1] {
						return true
					}
				}
				return false
			})

			// add the middle element of the line to the sum
			sum += intLine[len(intLine)/2]
		}
	}
	return sum
}

// buildUpdateAndCheckValidity converts a comma-separated string of numbers into a slice of integers and checks its validity.
// It verifies each number against rules and sets the validity to false if a rule's condition is violated. Returns the slice and validity.
func buildUpdateAndCheckValidity(line string, rules [][]int) ([]int, bool) {
	numStrings := strings.Split(line, ",")
	intLine := make([]int, 0)
	valid := true
	for _, num := range numStrings {
		n, _ := strconv.Atoi(num)
		// check if n occurs in rules on the left side
		for _, rule := range rules {
			if n == rule[0] {
				// there should not be the right number in the intlLine
				for _, num2 := range intLine {
					if num2 == rule[1] {
						valid = false
					}
				}
			}
		}
		intLine = append(intLine, n)
	}
	return intLine, valid
}

// splitInputs reads a file and splits its content into rules and updates text sections.
// It returns a slice of integer pairs extracted from rules and a slice of strings for updates.
func splitInputs(fileName string) ([][]int, []string) {
	text, _ := util.ReadFileAsString(fileName)
	// list of pairs of ints
	rules := make([][]int, 0)
	split := strings.Split(text, "\n\n")
	rulesText := strings.Split(split[0], "\n")
	updatesTexts := strings.Split(split[1], "\n")
	for _, line := range rulesText {
		numStrings := strings.Split(line, "|")
		a, _ := strconv.Atoi(numStrings[0])
		b, _ := strconv.Atoi(numStrings[1])
		rules = append(rules, []int{a, b})
	}
	return rules, updatesTexts
}

func main() {
	filename := "day05/data/test.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
