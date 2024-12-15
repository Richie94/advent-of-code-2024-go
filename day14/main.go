package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
)

type Robot struct {
	position util.Point
	velocity util.Point
}

func (robot *Robot) move(xMax, yMax int) {
	// moves and teleports the robot, will move with modulo
	robot.position.X = (robot.position.X + robot.velocity.X) % xMax
	robot.position.Y = (robot.position.Y + robot.velocity.Y) % yMax
	if robot.position.X < 0 {
		robot.position.X += xMax
	}
	if robot.position.Y < 0 {
		robot.position.Y += yMax
	}
}

func Part1(fileName string) int {
	robots := getRobots(fileName)
	// if we have below 15 robots its the test with xMax = 11 and yMax = 7
	// if above its 103 101
	xMax := 101
	yMax := 103
	if len(robots) < 15 {
		xMax = 11
		yMax = 7
	}
	// move each robot 100 times
	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			robot.move(xMax, yMax)
		}
	}
	// now detect the robots in each quadrant, the ones in the middle dont count
	return getSafetyFactor(robots, xMax, yMax)
}

func getSafetyFactor(robots []*Robot, xMax int, yMax int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.position.X < xMax/2 && robot.position.Y < yMax/2 {
			q1++
		} else if robot.position.X > xMax/2 && robot.position.Y < yMax/2 {
			q2++
		} else if robot.position.X < xMax/2 && robot.position.Y > yMax/2 {
			q3++
		} else if robot.position.X > xMax/2 && robot.position.Y > yMax/2 {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func getSafetyFactor16(robots []*Robot, xMax int, yMax int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	q5, q6, q7, q8 := 0, 0, 0, 0
	q9, q10, q11, q12 := 0, 0, 0, 0
	q13, q14, q15, q16 := 0, 0, 0, 0
	// divide the image in 16 parts, check for the maximum density in a part
	for _, robot := range robots {
		if robot.position.X < xMax/4 && robot.position.Y < yMax/4 {
			q1++
		} else if robot.position.X > xMax/4 && robot.position.X < xMax/2 && robot.position.Y < yMax/4 {
			q2++
		} else if robot.position.X > xMax/2 && robot.position.X < 3*xMax/4 && robot.position.Y < yMax/4 {
			q3++
		} else if robot.position.X > 3*xMax/4 && robot.position.Y < yMax/4 {
			q4++
		} else if robot.position.X < xMax/4 && robot.position.Y > yMax/4 && robot.position.Y < yMax/2 {
			q5++
		} else if robot.position.X > xMax/4 && robot.position.X < xMax/2 && robot.position.Y > yMax/4 && robot.position.Y < yMax/2 {
			q6++
		} else if robot.position.X > xMax/2 && robot.position.X < 3*xMax/4 && robot.position.Y > yMax/4 && robot.position.Y < yMax/2 {
			q7++
		} else if robot.position.X > 3*xMax/4 && robot.position.Y > yMax/4 && robot.position.Y < yMax/2 {
			q8++
		} else if robot.position.X < xMax/4 && robot.position.Y > yMax/2 && robot.position.Y < 3*yMax/4 {
			q9++
		} else if robot.position.X > xMax/4 && robot.position.X < xMax/2 && robot.position.Y > yMax/2 && robot.position.Y < 3*yMax/4 {
			q10++
		} else if robot.position.X > xMax/2 && robot.position.X < 3*xMax/4 && robot.position.Y > yMax/2 && robot.position.Y < 3*yMax/4 {
			q11++
		} else if robot.position.X > 3*xMax/4 && robot.position.Y > yMax/2 && robot.position.Y < 3*yMax/4 {
			q12++
		} else if robot.position.X < xMax/4 && robot.position.Y > 3*yMax/4 {
			q13++
		} else if robot.position.X > xMax/4 && robot.position.X < xMax/2 && robot.position.Y > 3*yMax/4 {
			q14++
		} else if robot.position.X > xMax/2 && robot.position.X < 3*xMax/4 && robot.position.Y > 3*yMax/4 {
			q15++
		} else if robot.position.X > 3*xMax/4 && robot.position.Y > 3*yMax/4 {
			q16++
		}
	}
	// pack the values into a list and get the maximum
	allList := []int{q1, q2, q3, q4, q5, q6, q7, q8, q9, q10, q11, q12, q13, q14, q15, q16}
	maxValue := 0
	for _, value := range allList {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func getRobots(fileName string) []*Robot {
	text, _ := util.ReadFileAsString(fileName)
	// line is in format p=x,y v=x,y
	pattern := `=(\d+),(\d+) v=(-?\d+),(-?\d+)`
	regex, _ := regexp.Compile(pattern)
	matches := regex.FindAllStringSubmatch(text, -1)
	robots := make([]*Robot, 0)
	for _, match := range matches {
		pX, _ := strconv.Atoi(match[1])
		pY, _ := strconv.Atoi(match[2])
		vX, _ := strconv.Atoi(match[3])
		vY, _ := strconv.Atoi(match[4])
		robot := &Robot{position: util.Point{X: pX, Y: pY}, velocity: util.Point{X: vX, Y: vY}}
		robots = append(robots, robot)
	}
	return robots
}

func Part2(fileName string) int {
	robots := getRobots(fileName)
	// if we have below 15 robots its the test with xMax = 11 and yMax = 7
	// if above its 103 101
	xMax := 101
	yMax := 103
	if len(robots) < 15 {
		xMax = 11
		yMax = 7
	}
	// get the top 20 lowest score
	highestDensity := 0
	for time := 0; time < 50000; time++ {
		for _, robot := range robots {
			robot.move(xMax, yMax)
		}

		density := getSafetyFactor16(robots, xMax, yMax)

		if density > highestDensity {
			highestDensity = density
			debug(time, density, yMax, xMax, robots)
		}
	}

	return 0
}

func debug(time int, safetyFactor int, yMax int, xMax int, robots []*Robot) {
	fmt.Println("Time:", time)
	fmt.Println("Safety Factor:", safetyFactor)
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			found := false
			for _, robot := range robots {
				if robot.position.X == x && robot.position.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	filename := "day03/data/input.txt"
	sum := Part1(filename)
	sum2 := Part2(filename)

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

}
