package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Programm struct {
	regA         int
	regB         int
	regC         int
	instructions []int
}

func (p *Programm) combo(i int) int {
	if i >= 0 && i <= 3 {
		return i
	} else if i == 4 {
		return p.regA
	} else if i == 5 {
		return p.regB
	} else if i == 6 {
		return p.regC
	} else {
		// error
		panic("Invalid register")
	}
}

func Part1(fileName string) string {
	programm := parseProgramm(fileName)
	step := 0
	output := make([]string, 0)
	for {
		if step >= len(programm.instructions) {
			break
		}
		instruction := programm.instructions[step]
		switch instruction {
		case 0:
			// adv -> division
			numerator := programm.regA
			denominator := math.Pow(2, float64(programm.combo(programm.instructions[step+1])))
			programm.regA = numerator / int(denominator)
			step += 2
		case 1:
			// bxl -> bitwise xor of B and liertale step+1
			programm.regB = programm.regB ^ programm.instructions[step+1]
			step += 2
		case 2:
			// bst combo mod 8 -> B
			programm.regB = programm.combo(programm.instructions[step+1]) % 8
			step += 2
		case 3:
			// jnz A == 0 -> do nothing
			if programm.regA == 0 {
				step += 2 // ??? or 1?
			} else {
				step = programm.instructions[step+1]
			}
		case 4:
			// bxc btwise xor of C and B stored in B
			programm.regB = programm.regB ^ programm.regC
			step += 2
		case 5:
			// out -> combo operand mod 8 will be outputted
			output = append(output, strconv.Itoa(programm.combo(programm.instructions[step+1])%8))
			step += 2
		case 6:
			// bdv like adv but stored in B
			numerator := programm.regA
			denominator := math.Pow(2, float64(programm.combo(programm.instructions[step+1])))
			programm.regB = numerator / int(denominator)
			step += 2
		case 7:
			// cdv like adv but stored in C
			numerator := programm.regA
			denominator := math.Pow(2, float64(programm.combo(programm.instructions[step+1])))
			programm.regC = numerator / int(denominator)
			step += 2
		}
	}

	return strings.Join(output, ",")
}

func parseProgramm(fileName string) *Programm {
	text, _ := util.ReadFileAsString(fileName)
	// line regex .*: (.*)
	pattern := `.*: (.*)`
	re, _ := regexp.Compile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)

	regA, _ := strconv.Atoi(matches[0][1])
	regB, _ := strconv.Atoi(matches[1][1])
	regC, _ := strconv.Atoi(matches[2][1])
	instructions := strings.Split(matches[3][1], ",")
	instructionsInt := make([]int, len(instructions))
	for i, v := range instructions {
		instructionsInt[i], _ = strconv.Atoi(v)
	}
	programm := &Programm{regA, regB, regC, instructionsInt}
	return programm
}

func Part2(fileName string) int {
	return 0
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
