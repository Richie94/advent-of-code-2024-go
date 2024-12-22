package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// Last PAD
//+---+---+---+
//| 7 | 8 | 9 |
//+---+---+---+
//| 4 | 5 | 6 |
//+---+---+---+
//| 1 | 2 | 3 |
//+---+---+---+
//    | 0 | A |
//    +---+---+
// All other pads
//    +---+---+
//    | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

var biDirectionalPad = map[util.Point]rune{
	util.Point{X: 1, Y: 0}: '^',
	util.Point{X: 2, Y: 0}: 'A',
	util.Point{X: 0, Y: 1}: '<',
	util.Point{X: 1, Y: 1}: 'v',
	util.Point{X: 2, Y: 1}: '>',
}

var oneDirectionalPad = map[util.Point]rune{
	util.Point{X: 0, Y: 0}: '7',
	util.Point{X: 1, Y: 0}: '8',
	util.Point{X: 2, Y: 0}: '9',
	util.Point{X: 0, Y: 1}: '4',
	util.Point{X: 1, Y: 1}: '5',
	util.Point{X: 2, Y: 1}: '6',
	util.Point{X: 0, Y: 2}: '1',
	util.Point{X: 1, Y: 2}: '2',
	util.Point{X: 2, Y: 2}: '3',
	util.Point{X: 1, Y: 3}: '0',
	util.Point{X: 2, Y: 3}: 'A',
}

func findShortestPathsFromPoints(start, end util.Point, pad map[util.Point]rune) []string {
	// find the shortest path from start to end
	// start with the start point
	current := start
	paths := make([]string, 0)
	if current != end {
		// go simple manhattan style
		rightPoint := util.Point{X: current.X + 1, Y: current.Y}
		_, rightExists := pad[rightPoint]
		leftPoint := util.Point{X: current.X - 1, Y: current.Y}
		_, leftExists := pad[leftPoint]
		upPoint := util.Point{X: current.X, Y: current.Y - 1}
		_, upExists := pad[upPoint]
		downPoint := util.Point{X: current.X, Y: current.Y + 1}
		_, downExists := pad[downPoint]
		if current.X < end.X && rightExists {
			recursiveShortestPaths := findShortestPathsFromPoints(rightPoint, end, pad)
			for _, path := range recursiveShortestPaths {
				paths = append(paths, ">"+path)
			}
		}
		if current.X > end.X && leftExists {
			recursiveShortestPaths := findShortestPathsFromPoints(leftPoint, end, pad)
			for _, path := range recursiveShortestPaths {
				paths = append(paths, "<"+path)
			}
		}
		if current.Y < end.Y && downExists {
			recursiveShortestPaths := findShortestPathsFromPoints(downPoint, end, pad)
			for _, path := range recursiveShortestPaths {
				paths = append(paths, "v"+path)
			}
		}
		if current.Y > end.Y && upExists {
			recursiveShortestPaths := findShortestPathsFromPoints(upPoint, end, pad)
			for _, path := range recursiveShortestPaths {
				paths = append(paths, "^"+path)
			}
		}
	} else {
		paths = append(paths, "")
	}
	return paths
}

func Part1(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	sum := 0
	for _, code := range lines {
		sequence := shortestSequence(code, 2)
		sum += sequenceToScore(sequence, code)
	}
	return sum
}

func findKeyByValue(pad map[util.Point]rune, value rune) (util.Point, bool) {
	for key, val := range pad {
		if val == value {
			return key, true
		}
	}
	return util.Point{}, false
}

func shortestSequence(code string, middleRobots int) string {
	positions := make([]util.Point, middleRobots+1)
	sequences := make([][]string, middleRobots+2)
	sequences[middleRobots+1] = []string{code}

	// the first one are on the
	for i := 0; i < middleRobots; i++ {
		positions[i] = util.Point{X: 2, Y: 0}
	}
	positions[middleRobots] = util.Point{X: 2, Y: 3}

	cache := make(map[string][]string)

	for i := middleRobots + 1; i > 0; i-- {
		pad := biDirectionalPad
		if i == middleRobots+1 {
			pad = oneDirectionalPad
		}
		fmt.Println(code, i, len(sequences[i]))
		// the code here is the sequence for the last robot
		detectedSequences := make([]string, 0)
		for _, robotCode := range sequences[i] {
			tmpSequences := make([]string, 0)
			if robotCode == "" {
				continue
			}
			for runeIdx, targetRune := range robotCode {
				// check if we have the sequence in cache
				if sequences, exists := cache[robotCode[:runeIdx]]; exists {
					tmpSequences = sequences
					break
				}
				// find the shortest path from the current robot position to the code
				target, _ := findKeyByValue(pad, targetRune)
				from := positions[i-1]
				robotSeqs := findShortestPathsFromPoints(from, target, pad)
				// optimization: if we have multiple paths and some have same char in a row, get rid of the others
				skipSequences := make([]string, 0)
				if len(robotSeqs) > 1 {
					minCharSwitches := len(robotSeqs[0])
					charSwitchMap := make(map[string]int)
					for _, seq := range robotSeqs {
						charSwitches := 0
						for i := 1; i < len(seq); i++ {
							if seq[i] != seq[i-1] {
								charSwitches++
							}
						}
						charSwitchMap[seq] = charSwitches
						if charSwitches < minCharSwitches {
							minCharSwitches = charSwitches
						}
					}
					for _, seq := range robotSeqs {
						if charSwitchMap[seq] > minCharSwitches {
							skipSequences = append(skipSequences, seq)
						}
					}
					// prefer starting with up or right, as they are closer to the letter A
					// check if any starts with right Or Up, if it does, remove all others
					startsWithRightOrUp := false
					startsWithRightOrUpMap := make(map[string]bool)
					for _, seq := range robotSeqs {
						if seq == "" {
							continue
						}
						startsWithRightOrUpMap[seq] = seq[0] == '>' || seq[0] == '^'
						if seq[0] == '>' || seq[0] == '^' {
							startsWithRightOrUp = true
						}
					}

					if startsWithRightOrUp {
						for seq, startsWith := range startsWithRightOrUpMap {
							if !startsWith {
								skipSequences = append(skipSequences, seq)
							}
						}
					}
				}
				positions[i-1] = target

				if len(tmpSequences) == 0 {
					tmpSequences = make([]string, 1)
				}
				// multiply the sequences
				newSequences := make([]string, 0)
				for _, seq := range tmpSequences {
					for _, robotSeq := range robotSeqs {
						if slices.Contains(skipSequences, robotSeq) {
							continue
						}
						newSequences = append(newSequences, seq+robotSeq+"A")
					}
				}
				tmpSequences = newSequences
				cache[robotCode[:runeIdx]] = tmpSequences
			}

			detectedSequences = append(detectedSequences, tmpSequences...)
		}
		// remove all sequences with are longer than the shortest one
		minSeqLen := len(detectedSequences[0])
		for _, seq := range detectedSequences {
			if len(seq) < minSeqLen {
				minSeqLen = len(seq)
			}
		}
		slices.DeleteFunc(detectedSequences, func(seq string) bool {
			return len(seq) > minSeqLen
		})

		sequences[i-1] = detectedSequences
	}

	minSeqLen := len(sequences[0][0])
	minSeq := sequences[0][0]
	for _, seq := range sequences[0] {
		if len(seq) < minSeqLen && len(seq) > 0 {
			minSeqLen = len(seq)
			minSeq = seq
		}
	}
	return minSeq
}

func sequenceToScore(sequence, code string) int {
	codeString := ""
	for _, c := range code {
		// if code is digit, append to codeString
		if c >= '0' && c <= '9' {
			codeString += string(c)
		}
	}
	codeNum, _ := strconv.Atoi(codeString)
	return len(sequence) * codeNum
}

func Part2(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	sum := 0
	for _, code := range lines {
		sequence := shortestSequence(code, 25)
		sum += sequenceToScore(sequence, code)
	}
	return sum
}

func main() {
	input, _ := util.ReadFileAsArray("day21/data/test.txt")
	numericalMap := make(map[string]util.Point)
	numericalMap["A"] = util.Point{2, 0}
	numericalMap["0"] = util.Point{1, 0}
	numericalMap["1"] = util.Point{0, 1}
	numericalMap["2"] = util.Point{1, 1}
	numericalMap["3"] = util.Point{2, 1}
	numericalMap["4"] = util.Point{0, 2}
	numericalMap["5"] = util.Point{1, 2}
	numericalMap["6"] = util.Point{2, 2}
	numericalMap["7"] = util.Point{0, 3}
	numericalMap["8"] = util.Point{1, 3}
	numericalMap["9"] = util.Point{2, 3}

	directionalMap := make(map[string]util.Point)
	directionalMap["A"] = util.Point{2, 1}
	directionalMap["^"] = util.Point{1, 1}
	directionalMap["<"] = util.Point{0, 0}
	directionalMap["v"] = util.Point{1, 0}
	directionalMap[">"] = util.Point{2, 0}

	//fmt.Println("answer to part 1: ", getSequence(input, numericalMap, directionalMap, 2))
	fmt.Println("answer to part 1: ", getSequence(input, numericalMap, directionalMap, 25))
}

/*
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//	   | 0 | A |
//	   +---+---+
*/
func getPressesForNumericPad(input []string, start string, numericalMap map[string]util.Point) []string {
	current := numericalMap[start]
	output := []string{}

	for _, char := range input {
		dest := numericalMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < util.AbsInt(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < util.AbsInt(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if current.Y == 0 && dest.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if current.X == 0 && dest.Y == 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX >= 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}

		current = dest
		output = append(output, "A")
	}
	return output
}

/*
//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+
*/
func getPressesForDirectionalPad(input []string, start string, directionlMap map[string]util.Point) []string {
	current := directionlMap[start]
	output := []string{}

	for _, char := range input {
		dest := directionlMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < util.AbsInt(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < util.AbsInt(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if current.X == 0 && dest.Y == 1 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if current.Y == 1 && dest.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX >= 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}
		current = dest
		output = append(output, "A")
	}
	return output
}

func getSequence(input []string, numericalMap, directionalMap map[string]util.Point, robots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, line := range input {
		row := strings.Split(line, "")
		seq1 := getPressesForNumericPad(row, "A", numericalMap)
		num := getCountAfterRobots(seq1, robots, 1, cache, directionalMap)
		codeString := ""
		for _, c := range line {
			// if code is digit, append to codeString
			if c >= '0' && c <= '9' {
				codeString += string(c)
			}
		}
		codeNum, _ := strconv.Atoi(codeString)
		count += codeNum * num
	}
	return count
}

func getCountAfterRobots(input []string, maxRobots int, robot int, cache map[string][]int, directionalMap map[string]util.Point) int {
	if val, ok := cache[strings.Join(input, "")]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[strings.Join(input, "")] = make([]int, maxRobots)
	}

	seq := getPressesForDirectionalPad(input, "A", directionalMap)
	cache[strings.Join(input, "")][0] = len(seq)

	if robot == maxRobots {
		return len(seq)
	}

	splitSeq := getIndividualSteps(seq)

	count := 0
	for _, s := range splitSeq {
		c := getCountAfterRobots(s, maxRobots, robot+1, cache, directionalMap)
		if _, ok := cache[strings.Join(s, "")]; !ok {
			cache[strings.Join(s, "")] = make([]int, maxRobots)
		}
		cache[strings.Join(s, "")][0] = c
		count += c
	}

	cache[strings.Join(input, "")][robot-1] = count
	return count
}

func getIndividualSteps(input []string) [][]string {
	output := [][]string{}
	current := []string{}
	for _, char := range input {
		current = append(current, char)

		if char == "A" {
			output = append(output, current)
			current = []string{}
		}
	}
	return output
}
