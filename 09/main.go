package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2022/utils"
)

// info about a coordinate
type coordinate struct {
	visitedByHeadCount int
	visitedByTailCount int
}

// {2,3}:{visitedHEad, visitedTail}
type area map[[2]int]coordinate

// direction
type direction struct {
	orientation string
	length      int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(lines []string) (directions []direction) {
	for _, line := range lines {
		arr := strings.Split(line, " ")
		orientation, l := arr[0], arr[1]
		length, _ := strconv.Atoi(l)
		directions = append(directions, direction{orientation, length})
	}
	return directions
}

func p1(directions []direction) int {

	pos := [2]int{0, 0}
	tailPos := [2]int{0, 0}
	area := make(area)
	area[pos] = coordinate{1, 1}

	for _, dir := range directions {
		pos, tailPos = move(dir, pos, tailPos, area)
	}

	//fmt.Println(area)
	visitedByTail := 0
	for _, v := range area {
		if v.visitedByTailCount > 0 {
			//fmt.Println(c, v)
			visitedByTail += 1
		}
	}
	return visitedByTail
}

func move(dir direction, pos, tailPos [2]int, a area) ([2]int, [2]int) {
	orientation := dir.orientation
	length := dir.length
	//fmt.Println("\n", dir, pos, tailPos)

	var currPos [2]int
	switch orientation {
	case "R":
		for k := pos[0] + 1; k <= pos[0]+length; k++ {
			currPos = [2]int{k, pos[1]}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			if Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] - 1, currPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	case "L":
		for k := pos[0] - 1; k >= pos[0]-length; k-- {
			currPos = [2]int{k, pos[1]}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			if Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] + 1, currPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
			}
			//fmt.Println("visiting", currPos, tailPos)
		}

	case "U":
		for k := pos[1] + 1; k <= pos[1]+length; k++ {
			currPos = [2]int{pos[0], k}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			if Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{currPos[0], currPos[1] - 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	case "D":
		for k := pos[1] - 1; k >= pos[1]-length; k-- {
			currPos = [2]int{pos[0], k}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			if Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{currPos[0], currPos[1] + 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	}
	return currPos, tailPos
}

func main() {
	lines, _ := utils.ReadLines("test.txt")
	directions := parseInput(lines)
	//fmt.Println("directions: ", directions)

	fmt.Println(p1(directions))
}
