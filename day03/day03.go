package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func Part1(fileName string) int {
	lines, err := util.ReadFileAsArray(fileName)
	if err != nil {
		panic(err)
	}

	pattern := `mul\((\d+),(\d+)\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) >= 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			}
		}
	}
	return sum
}

func Part2(fileName string) int {
	lines, err := util.ReadFileAsArray(fileName)
	if err != nil {
		panic(err)
	}

	pattern := `(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	sum := 0
	active := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				active = true
			}
			if match[0] == "don't()" {
				active = false
			}
			if len(match) >= 4 && active {
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				sum += num1 * num2
			}
		}
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
