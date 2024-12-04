package main

import (
	"advent-of-code-2024/util"
	"fmt"
)

type Matrix struct {
	Rows int
	Cols int
	data []string
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		Rows: rows,
		Cols: cols,
		data: make([]string, rows*cols),
	}
}

func (m *Matrix) At(row, col int) string {
	return m.data[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, v string) {
	m.data[row*m.Cols+col] = v
}

func Part1(fileName string) int {
	matrix := getMatrix(fileName)
	// iterate over the matrix
	sum := 0
	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < matrix.Cols; j++ {
			// check the word XMAS in all directions
			// check horizontal left to right
			if j+3 < matrix.Cols {
				if matrix.At(i, j) == "X" && matrix.At(i, j+1) == "M" && matrix.At(i, j+2) == "A" && matrix.At(i, j+3) == "S" {
					sum++
				}
			}
			// check horizontal right to left
			if j-3 >= 0 {
				if matrix.At(i, j) == "X" && matrix.At(i, j-1) == "M" && matrix.At(i, j-2) == "A" && matrix.At(i, j-3) == "S" {
					sum++
				}
			}
			// check vertical top to bottom
			if i+3 < matrix.Rows {
				if matrix.At(i, j) == "X" && matrix.At(i+1, j) == "M" && matrix.At(i+2, j) == "A" && matrix.At(i+3, j) == "S" {
					sum++
				}
			}

			// check vertical bottom to top
			if i-3 >= 0 {
				if matrix.At(i, j) == "X" && matrix.At(i-1, j) == "M" && matrix.At(i-2, j) == "A" && matrix.At(i-3, j) == "S" {
					sum++
				}
			}

			// check diagonal top left to bottom right
			if i+3 < matrix.Rows && j+3 < matrix.Cols {
				if matrix.At(i, j) == "X" && matrix.At(i+1, j+1) == "M" && matrix.At(i+2, j+2) == "A" && matrix.At(i+3, j+3) == "S" {
					sum++
				}
			}
			// check diagonal bottom right to top left
			if i-3 >= 0 && j-3 >= 0 {
				if matrix.At(i, j) == "X" && matrix.At(i-1, j-1) == "M" && matrix.At(i-2, j-2) == "A" && matrix.At(i-3, j-3) == "S" {
					sum++
				}
			}
			// check diagonal top right to bottom left
			if i+3 < matrix.Rows && j-3 >= 0 {
				if matrix.At(i, j) == "X" && matrix.At(i+1, j-1) == "M" && matrix.At(i+2, j-2) == "A" && matrix.At(i+3, j-3) == "S" {
					sum++
				}
			}
			// check diagonal bottom left to top right
			if i-3 >= 0 && j+3 < matrix.Cols {
				if matrix.At(i, j) == "X" && matrix.At(i-1, j+1) == "M" && matrix.At(i-2, j+2) == "A" && matrix.At(i-3, j+3) == "S" {
					sum++
				}
			}
		}
	}

	return sum
}

func getMatrix(fileName string) *Matrix {
	lines, _ := util.ReadFileAsArray(fileName)
	rowAmount := len(lines)
	colAmount := len(lines[0])
	matrix := NewMatrix(rowAmount, colAmount)
	for i, line := range lines {
		for j, char := range line {
			matrix.Set(i, j, string(char))
		}
	}
	return matrix
}

func Part2(fileName string) int {
	matrix := getMatrix(fileName)
	// iterate over the matrix
	sum := 0
	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < matrix.Cols; j++ {
			// we check for having an A here
			matches := 0
			if matrix.At(i, j) == "A" {
				// now we check for the word MAS twice
				// check diagonal top left to bottom right and vice versa
				if i+1 < matrix.Rows && j+1 < matrix.Cols && i-1 >= 0 && j-1 >= 0 {
					if matrix.At(i-1, j-1) == "M" && matrix.At(i+1, j+1) == "S" {
						matches++
					}
					if matrix.At(i+1, j+1) == "M" && matrix.At(i-1, j-1) == "S" {
						matches++
					}
				}
				// check diagonal top right to bottom left and vice versa
				if i+1 < matrix.Rows && j-1 >= 0 && i-1 >= 0 && j+1 < matrix.Cols {
					if matrix.At(i-1, j+1) == "M" && matrix.At(i+1, j-1) == "S" {
						matches++
					}
					if matrix.At(i+1, j-1) == "M" && matrix.At(i-1, j+1) == "S" {
						matches++
					}
				}

				if matches > 1 {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	filename := "day04/data/test.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
