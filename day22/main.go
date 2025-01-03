package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

func Part1(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	sum := 0
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		sum += generate(num, 2000)
	}
	return sum
}

func generate(secretNumber, rounds int) int {
	for i := 0; i < rounds; i++ {
		secretNumber = apply(secretNumber)
	}
	return secretNumber
}

func apply(secretNumber int) int {
	result := secretNumber * 64
	secretNumber ^= result
	secretNumber %= 16777216

	result = secretNumber / 32
	secretNumber ^= result
	secretNumber %= 16777216

	result = secretNumber * 2048
	secretNumber ^= result
	secretNumber %= 16777216
	return secretNumber
}

func getSequence(number, rounds int) map[string]int {
	sequence := make([]int, rounds)
	changeMap := make(map[string]int)
	sequence[0] = number % 10
	for i := 1; i < rounds; i++ {
		number = apply(number)
		remainder := number % 10
		sequence[i] = remainder
		if i > 3 {
			changeSequence := ""
			for j := 4; j > 0; j-- {
				changeSequence += strconv.Itoa(sequence[i-j+1] - sequence[i-j])
			}
			if _, ok := changeMap[changeSequence]; !ok {
				changeMap[changeSequence] = remainder
			}
		}
	}
	return changeMap

}

func Part2(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	prices := make(map[string]int, 0)
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		changeMap := getSequence(num, 2000)
		for key, val := range changeMap {
			if _, ok := prices[key]; !ok {
				prices[key] = val
			} else {
				prices[key] += val
			}
		}
	}
	bestPrice := 0
	for _, val := range prices {
		if val > bestPrice {
			bestPrice = val
		}
	}
	return bestPrice
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
