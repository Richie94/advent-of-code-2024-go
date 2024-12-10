package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

type DataPoint struct {
	fileId int // if isFree, fileId = -1
}

func Part1(fileName string) int {
	text, _ := util.ReadFileAsString(fileName)
	dataPoints := generateDatapoints(text)
	compactDatapointsSingle(dataPoints)
	sum := calcSum(dataPoints)
	return sum
}

func calcSum(dataPoints []DataPoint) int {
	sum := 0
	for idx, dp := range dataPoints {
		if dp.fileId != -1 {
			sum += idx * dp.fileId
		}
	}
	return sum
}

func compactDatapointsSingle(dataPoints []DataPoint) {
	lastEmpty := 0
	// iterate in reverse through the datapoints
	for i := len(dataPoints) - 1; i >= 0; i-- {
		// if we have a file id, we can start the process
		if dataPoints[i].fileId != -1 {
			// find the first empty block
			for j := lastEmpty; j < i; j++ {
				if dataPoints[j].fileId == -1 {
					// swap the two elements
					dataPoints[i].fileId, dataPoints[j].fileId = dataPoints[j].fileId, dataPoints[i].fileId
					break
				}
			}
		}
	}
}

func compactDatapoints(dataPoints []DataPoint) {
	// here we try to move only complete blocks
	// iterate in reverse through the datapoints
	amountFilled := 0
	lastId := dataPoints[len(dataPoints)-1].fileId
	for i := len(dataPoints) - 1; i >= 0; i-- {
		// if we have a file id, we can start the process
		if dataPoints[i].fileId == lastId {
			amountFilled++
		} else {
			// we read the complete block, now we try to merge it
			amountFree := 0
			matched := false
			for j := 0; j <= i; j++ {
				if amountFree == amountFilled {
					// we found a complete block
					// swap the two blocks
					for k := 0; k < amountFilled; k++ {
						dataPoints[i+k+1].fileId, dataPoints[j-k-1].fileId = dataPoints[j-k-1].fileId, dataPoints[i+k+1].fileId
					}
					matched = true
					amountFilled = 0
					// last id is the next not -1 value behind i
					for k := i; k >= 0; k-- {
						if dataPoints[k].fileId != -1 {
							if k == i {
								amountFilled++
							}
							lastId = dataPoints[k].fileId
							break
						}
					}
					break

				}
				if dataPoints[j].fileId == -1 {
					amountFree++
				} else {
					amountFree = 0
				}
			}

			if !matched {
				// we did not match, we let that block there and check for the next fileId
				amountFilled = 0
				for k := i; k >= 0; k-- {
					if dataPoints[k].fileId != -1 {
						if k == i {
							amountFilled++
						}
						lastId = dataPoints[k].fileId
						break
					}
				}
			}
		}
	}

}

func generateDatapoints(text string) []DataPoint {
	dataPoints := make([]DataPoint, 0)
	fileId := -1
	for idx, char := range text {
		num, _ := strconv.Atoi(string(char))
		isMissing := false
		if idx%2 == 0 {
			// input
			fileId++
		} else {
			// empty block
			isMissing = true
		}
		for i := 0; i < num; i++ {
			if !isMissing {
				dataPoints = append(dataPoints, DataPoint{fileId: fileId})
			} else {
				dataPoints = append(dataPoints, DataPoint{fileId: -1})
			}
		}
	}
	return dataPoints
}

func Part2(fileName string) int {
	text, _ := util.ReadFileAsString(fileName)
	dataPoints := generateDatapoints(text)
	compactDatapoints(dataPoints)
	// write to file, join the dataPoints
	fmt.Println(dataPoints)
	sum := calcSum(dataPoints)
	return sum
}

func main() {
	filename := "day09/data/test.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
