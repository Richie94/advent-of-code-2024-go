package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1(fileName string) int {
	return solve(fileName, 25)
}

func solve(fileName string, iterations int) int {
	numbers := parseText(fileName)
	// convert to map
	numbersMap := make(map[int]int)
	for _, num := range numbers {
		numbersMap[num]++
	}
	for i := 0; i < iterations; i++ {
		numbersMap = iterateNumbers(numbersMap)
	}
	sum := 0
	for _, amount := range numbersMap {
		sum += amount
	}
	return sum
}

func iterateNumbers(numbers map[int]int) map[int]int {
	result := make(map[int]int)
	for num, amount := range numbers {
		numAsString := strconv.Itoa(num)
		if num == 0 {
			// add 1 fpr result[1]
			result[1] += amount
		} else if len(numAsString)%2 == 0 {
			// split into two
			firstHalf := numAsString[:len(numAsString)/2]
			firstHalfAsInt, _ := strconv.Atoi(firstHalf)
			secondHalf := numAsString[len(numAsString)/2:]
			secondHalfAsInt, _ := strconv.Atoi(secondHalf)
			result[firstHalfAsInt] += amount
			result[secondHalfAsInt] += amount
		} else {
			result[num*2024] += amount
		}
	}
	return result
}

func parseText(fileName string) []int {
	text, _ := util.ReadFileAsString(fileName)
	// split text into list of int by space
	numbers := make([]int, 0)
	for _, s := range strings.Split(text, " ") {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}
	return numbers
}

func Part2(fileName string) int {
	return solve(fileName, 75)
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
