package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1(fileName string) int {
	return solve(fileName, false)
}

func solve(fileName string, allowConcat bool) int {
	sum := 0
	lines, _ := util.ReadFileAsArray(fileName)
	for _, line := range lines {
		// format num1: num2 num3 ...
		split1 := strings.Split(line, ":")
		leftValue, _ := strconv.Atoi(split1[0])
		split2 := strings.Split(split1[1], " ")
		rightValues := make([]int, len(split2))
		for i, v := range split2 {
			rightValues[i], _ = strconv.Atoi(v)
		}
		if recursiveTestEquation(leftValue, 0, rightValues, allowConcat) {
			sum += leftValue
		}
	}
	return sum
}

func recursiveTestEquation(leftValue int, intermediateResult int, rightValues []int, allowConcat bool) bool {
	// if we have no valuzes left in rightValues, check if intermediateResult is equal to leftValue
	if len(rightValues) == 0 {
		return intermediateResult == leftValue
	}
	// if we have values left in rightValues, make the sum of intermediate with the first value in rightValues
	// and call the function recursively with the new intermediate value and the rest of the rightValues
	if recursiveTestEquation(leftValue, intermediateResult+rightValues[0], rightValues[1:], allowConcat) {
		return true
	}
	// if the previous call returned false, try to mulitply the first value in rightValues from the intermediate value
	// and call the function recursively with the new intermediate value and the rest of the rightValues
	if recursiveTestEquation(leftValue, intermediateResult*rightValues[0], rightValues[1:], allowConcat) {
		return true
	}
	if allowConcat {
		// check if we concatenate the first value in rightValues to the intermediate value
		// and call the function recursively with the new intermediate value and the rest of the rightValues
		concatVal, _ := strconv.Atoi(strconv.Itoa(intermediateResult) + strconv.Itoa(rightValues[0]))
		if recursiveTestEquation(leftValue, concatVal, rightValues[1:], allowConcat) {
			return true
		}
	}
	return false
}

func Part2(fileName string) int {
	return solve(fileName, true)
}

func main() {
	filename := "day07/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
