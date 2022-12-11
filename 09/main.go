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

	// Move
	for _, v := range directions {
		pos, tailPos = move(v, pos, tailPos, area)
	}

	//fmt.Println(area)
	//fmt.Println(len(area))
	tot := 0
	for _, v := range area {
		if v.visitedByTailCount > 0 {
			//fmt.Println(c, v)
			tot += 1
		}
	}
	return tot
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
			currPos[0] = k
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			// Move to right
			if tailPos[1] == currPos[1] && Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] - 1, tailPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving to right", tailPos)
			}
			// Move diagonal
			if tailPos[1] != currPos[1] && Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] - 1, currPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving diag", tailPos)
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	case "L":
		for k := pos[0] - 1; k >= pos[0]-length; k-- {
			currPos = [2]int{k, pos[1]}
			currPos[0] = k
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount + 1), (a[currPos].visitedByTailCount)}
			// Check tail
			// Move to left
			if tailPos[1] == currPos[1] && Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] + 1, tailPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving to left", tailPos)
			}
			// Move diagonal
			if tailPos[1] != currPos[1] && Abs(currPos[0]-tailPos[0]) > 1 {
				tailPos = [2]int{currPos[0] + 1, currPos[1]}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving diag", tailPos)
			}
			//fmt.Println("visiting", currPos, tailPos)
		}

	case "U":
		for k := pos[1] + 1; k <= pos[1]+length; k++ {
			currPos = [2]int{pos[0], k}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount+1), (a[currPos].visitedByTailCount)}
			// Check tail
			// Move up
			if tailPos[0] == currPos[0] && Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{currPos[0], tailPos[1] + 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving up", tailPos)
			}
			// Move diagonal
			if tailPos[0] != currPos[0] && Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{currPos[0], currPos[1] - 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving diag", tailPos)
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	case "D":
		for k := pos[1] - 1; k >= pos[1]-length; k-- {
			currPos = [2]int{pos[0], k}
			a[currPos] = coordinate{(a[currPos].visitedByHeadCount+1), (a[currPos].visitedByTailCount)}
			// Check tail
			// Move up
			if tailPos[0] == currPos[0] && Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{tailPos[0], tailPos[1] - 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving down", tailPos)
			}
			// Move diagonal
			if tailPos[0] != currPos[0] && Abs(currPos[1]-tailPos[1]) > 1 {
				tailPos = [2]int{currPos[0], currPos[1] + 1}
				a[tailPos] = coordinate{(a[tailPos].visitedByHeadCount), (a[tailPos].visitedByTailCount + 1)}
				//fmt.Println("tail moving diag", tailPos)
			}
			//fmt.Println("visiting", currPos, tailPos)
		}
	}
	return currPos, tailPos
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	directions := parseInput(lines)
	//fmt.Println("directions: ", directions)

	fmt.Println(p1(directions))

}