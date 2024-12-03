package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// readFileAsArray returns all lines of a file as a slice.
func readFileAsArray(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func main() {
	lines, err := readFileAsArray("day01/data/day01_input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var listLeft, listRight []int
	countMapRight := make(map[int]int)

	for _, line := range lines {
		split := strings.Fields(line) // nutzt Felder, um beliebige Leerzeichen zu handhaben

		if len(split) < 2 {
			fmt.Fprintf(os.Stderr, "warning: ignoring malformed line: %s\n", line)
			continue
		}

		atoi1, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping line due to conversion error: %s\n", err)
			continue
		}

		atoi2, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping line due to conversion error: %s\n", err)
			continue
		}

		listLeft = append(listLeft, atoi1)
		listRight = append(listRight, atoi2)

		countMapRight[atoi2]++
	}

	sort.Ints(listLeft)
	sort.Ints(listRight)

	var diffList []int
	for i, v := range listLeft {
		diff := listRight[i] - v
		if diff < 0 {
			diff = -diff
		}
		diffList = append(diffList, diff)
	}

	sum := 0
	for _, v := range diffList {
		sum += v
	}
	fmt.Println("Part 1:", sum)

	sumPart2 := 0
	for _, leftValue := range listLeft {
		sumPart2 += leftValue * countMapRight[leftValue]
	}
	fmt.Println("Part 2:", sumPart2)
}
