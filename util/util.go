package util

import (
	"bufio"
	"fmt"
	"os"
)

// readFileAsArray returns all lines of a file as a slice.
func ReadFileAsArray(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func ReadFileAsString(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var text string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	// remove last newline
	text = text[:len(text)-1]

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return text, nil
}

func Unique[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type Point struct {
	X, Y int
}

// inBounds checks if a point is within the bounds of a grid.
func (p Point) InBounds(xBound, yBound int) bool {
	return p.X >= 0 && p.X < xBound && p.Y >= 0 && p.Y < yBound
}
