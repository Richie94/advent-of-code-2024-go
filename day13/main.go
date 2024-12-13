package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	buttonA Button
	buttonB Button
	prize   util.Point
}

type Button struct {
	X    int
	Y    int
	cost int
}

func (game *Game) playNaive() int {
	// we have maximum hundred combinations, order doesnt matter
	cheapestCost := 100 * 100
	// Part 1
	// solve for i * aX + j * bX = pX and i * aY + j * bY = pY, i < 100, j < 100
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			X := i*game.buttonA.X + j*game.buttonB.X
			Y := i*game.buttonA.Y + j*game.buttonB.Y
			if X == game.prize.X && Y == game.prize.Y {
				cost := i*game.buttonA.cost + j*game.buttonB.cost
				if cost < cheapestCost {
					cheapestCost = cost
				}
			}
		}
	}
	if cheapestCost == 100*100 {
		return -1
	}
	return cheapestCost
}

func (game *Game) playMathProf() int {
	// optimize for cheaper b first
	bTop := game.prize.X*game.buttonA.Y - game.buttonA.X*game.prize.Y
	bBottom := game.buttonB.X*game.buttonA.Y - game.buttonA.X*game.buttonB.Y

	if bTop%bBottom != 0 {
		return -1
	}

	b := bTop / bBottom
	// now check for pricey a to match the prize
	aTop := game.prize.Y - b*game.buttonB.Y
	aBottom := game.buttonA.Y

	if aTop%aBottom != 0 {
		return -1
	}

	a := aTop / aBottom
	return a*game.buttonA.cost + b*game.buttonB.cost
}

func Part1(fileName string) int {
	games := parseGames(fileName)
	sum := 0
	for _, game := range games {
		cheapestCost := game.playNaive()
		if cheapestCost > 0 {
			sum += cheapestCost
		}
	}
	return sum
}

func parseGames(fileName string) []Game {
	text, _ := util.ReadFileAsString(fileName)
	games := make([]Game, 0)
	allPattern := "X(\\+|=)([0-9]+), Y(\\+|=)([0-9]+)"
	allRegex, _ := regexp.Compile(allPattern)
	for _, line := range strings.Split(text, "\n\n") {
		matches := allRegex.FindAllStringSubmatch(line, -1)
		aX, _ := strconv.Atoi(matches[0][2])
		aY, _ := strconv.Atoi(matches[0][4])
		buttonA := Button{X: aX, Y: aY, cost: 3}
		bX, _ := strconv.Atoi(matches[1][2])
		bY, _ := strconv.Atoi(matches[1][4])
		buttonB := Button{X: bX, Y: bY, cost: 1}
		pX, _ := strconv.Atoi(matches[2][2])
		pY, _ := strconv.Atoi(matches[2][4])
		prize := util.Point{X: pX, Y: pY}
		game := Game{buttonA: buttonA, buttonB: buttonB, prize: prize}
		games = append(games, game)
	}
	return games
}

func Part2(fileName string) int {
	games := parseGames(fileName)
	sum := 0
	for _, game := range games {
		// add 10000000000000 to the prize X and Y
		fixedPrice := util.Point{X: game.prize.X + 10000000000000, Y: game.prize.Y + 10000000000000}
		fixedGame := Game{buttonA: game.buttonA, buttonB: game.buttonB, prize: fixedPrice}
		cheapestCost := fixedGame.playMathProf()
		if cheapestCost > 0 {
			sum += cheapestCost
		}
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
