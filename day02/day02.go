package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrayUnsafeAt(numbers []int) int {
	// check if its mostly ascending or descending
	ascCount := 0
	descCount := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			ascCount++
		} else {
			descCount++
		}
	}
	safeArray := -1
	for i := 0; i < len(numbers)-1; i++ {
		safe := safeCompare(numbers[i], numbers[i+1], ascCount > descCount)
		if !safe {
			safeArray = i
			break
		}
	}
	return safeArray
}

func safeCompare(a int, b int, asc bool) bool {
	if asc {
		if b-a < 1 || b-a > 3 {
			return false
		}
	}
	if !asc {
		if a-b < 1 || a-b > 3 {
			return false
		}
	}
	return true
}

func main() {
	lines, err := util.ReadFileAsArray("day02/data/day02_input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	safeSum1 := 0
	safeSum2 := 0

	// iterate over all lines
	for idx, line := range lines {
		fmt.Println("Line", idx, ":", line)
		// parse line to int array
		var numbersStrings = strings.Fields(line)
		var numbers = make([]int, len(numbersStrings))
		for i, numberString := range numbersStrings {
			numbers[i], _ = strconv.Atoi(numberString)
		}
		unsafeAt := arrayUnsafeAt(numbers)
		if unsafeAt == -1 {
			safeSum1++
			safeSum2++
		} else {
			// check if the array without the unsafe element is safe
			copySlice := make([]int, len(numbers))
			copy(copySlice, numbers)
			newSlice1 := append(copySlice[:unsafeAt], copySlice[unsafeAt+1:]...)
			fmt.Println("New slice:", newSlice1)
			unsafeAt2 := arrayUnsafeAt(newSlice1)
			if unsafeAt2 == -1 {
				safeSum2++
			} else {
				// check if its okay if we drop the unsafeAt+1 element
				copySlice := make([]int, len(numbers))
				copy(copySlice, numbers)
				newSlice2 := append(copySlice[:unsafeAt+1], copySlice[unsafeAt+2:]...)
				fmt.Println("New slice:", newSlice2)
				unsafeAt2 := arrayUnsafeAt(newSlice2)
				if unsafeAt2 == -1 {
					safeSum2++
				} else {
					fmt.Println("\tUnsafe at", unsafeAt, "and", unsafeAt2)
				}
			}
		}
	}

	fmt.Println("Part 1:", safeSum1)
	fmt.Println("Part 2:", safeSum2)

}
