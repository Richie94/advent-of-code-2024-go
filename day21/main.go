package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strconv"
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

	for i := middleRobots + 1; i > 0; i-- {
		pad := biDirectionalPad
		if i == middleRobots+1 {
			pad = oneDirectionalPad
		}
		fmt.Println(code, i)
		// the code here is the sequence for the last robot
		detectedSequences := make([]string, 0)
		for _, robotCode := range sequences[i] {
			tmpSequences := make([]string, 0)
			if robotCode == "" {
				continue
			}
			for _, c := range robotCode {
				// find the shortest path from the current robot position to the code
				target, _ := findKeyByValue(pad, c)
				from := positions[i-1]
				robotSeqs := findShortestPathsFromPoints(from, target, pad)
				positions[i-1] = target

				if len(tmpSequences) == 0 {
					tmpSequences = make([]string, 1)
				}
				// multiply the sequences
				newSequences := make([]string, 0)
				for _, seq := range tmpSequences {
					for _, robotSeq := range robotSeqs {
						newSequences = append(newSequences, seq+robotSeq+"A")
					}
				}
				tmpSequences = newSequences
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
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
