package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, err := util.ReadFileAsArray("day03/data/input.txt")
	if err != nil {
		panic(err)
	}

	pattern := `(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`
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
	sum2 := 0
	active := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fmt.Println(match)
			if match[0] == "do()" {
				active = true
			}
			if match[0] == "don't()" {
				active = false
			}
			if len(match) >= 4 && active {
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				sum2 += num1 * num2
			}
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
