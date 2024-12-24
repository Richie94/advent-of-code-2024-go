package main

import (
	"advent-of-code-2024/util"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	operation string
	value1    string
	value2    string
	target    string
}

func Part1(fileName string) int {
	knownElements, operations := parseOperations(fileName)

	solveOperations(operations, knownElements)
	// take all z elements, make binary out of it and sum it up
	sum := extractBinaryValue(knownElements, "z")
	return sum
}

func extractBinaryValue(knownElements map[string]int, prefix string) int {
	sum := 0
	for key, value := range knownElements {
		if strings.HasPrefix(key, prefix) && value == 1 {
			index, _ := strconv.Atoi(key[1:])
			sum += int(math.Pow(2, float64(index)))
		}
	}
	return sum
}

func solveOperations(operations []Operation, knownElements map[string]int) {
	for len(operations) > 0 {
		operation := operations[0]
		solved := false
		if v1, okV1 := knownElements[operation.value1]; okV1 {
			if v2, okV2 := knownElements[operation.value2]; okV2 {
				switch operation.operation {
				case "AND":
					knownElements[operation.target] = v1 & v2
				case "OR":
					knownElements[operation.target] = v1 | v2
				case "XOR":
					knownElements[operation.target] = v1 ^ v2
				}
				solved = true
			}
		}
		if solved {
			operations = operations[1:]
		} else {
			operations = append(operations[1:], operations[0])
		}
	}
}

func parseOperations(fileName string) (map[string]int, []Operation) {
	text, _ := util.ReadFileAsString(fileName)
	split := strings.Split(text, "\n\n")
	knownElements := make(map[string]int)
	for _, line := range strings.Split(split[0], "\n") {
		elements := strings.Split(line, ": ")
		knownElements[elements[0]], _ = strconv.Atoi(elements[1])
	}
	operations := make([]Operation, 0)
	rules := strings.Split(split[1], "\n")
	pattern := "(.*) (XOR|OR|AND) (.*) -> (.*)"
	regex := regexp.MustCompile(pattern)
	for _, rule := range rules {
		matches := regex.FindAllStringSubmatch(rule, -1)[0]
		operations = append(operations, Operation{operation: matches[2], value1: matches[1], value2: matches[3], target: matches[4]})
	}
	return knownElements, operations
}

func Part2(fileName string) string {
	_, operations := parseOperations(fileName)
	// validate output
	swapCandidates := make([]string, 0)
	for _, op := range operations {
		isValid := true
		// if its a z output, it has to be XOR of two values
		if strings.HasPrefix(op.target, "z") && op.operation != "XOR" && op.target != "z45" {
			isValid = false
		}
		// each AND must go to an OR
		if op.operation == "AND" && !strings.HasPrefix(op.value1, "y") {
			found := false
			for _, op2 := range operations {
				if (op2.value1 == op.target || op2.value2 == op.target) && op2.operation == "OR" {
					found = true
				}
			}
			if !found {
				isValid = false
			}
		}
		// each XOR must go to xor or AND
		if op.operation == "XOR" && !strings.HasPrefix(op.target, "z") {
			found := false
			for _, op2 := range operations {
				if (op2.value1 == op.target || op2.value2 == op.target) && (op2.operation == "XOR" || op2.operation == "AND") {
					found = true
				}
			}
			if !found {
				isValid = false
			}
		}

		if !isValid {
			swapCandidates = append(swapCandidates, op.target)
		}
	}

	slices.Sort(swapCandidates)

	return strings.Join(swapCandidates, ",")
}
