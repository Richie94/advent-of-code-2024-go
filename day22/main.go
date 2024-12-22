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
	changeMaps := make([]map[string]int, 0)
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		changeMap := getSequence(num, 2000)
		changeMaps = append(changeMaps, changeMap)
	}
	// check the changeMaps for highest value over all changeMaps
	allKeys := make(map[string]bool)
	for _, changeMap := range changeMaps {
		for key := range changeMap {
			allKeys[key] = true
		}
	}
	bestPrice := 0
	for key := range allKeys {
		price := 0
		for _, changeMap := range changeMaps {
			if val, ok := changeMap[key]; ok {
				price += val
			}
		}
		if price > bestPrice {
			bestPrice = price
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
