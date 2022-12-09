package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2022/utils"
)

func p1(trees [][]int) int {

	visibleTrees := []int{}

	// Go through interior
	for i := 0; i < len(trees); i++ {
		//fmt.Println("\nrow ", i)
		for j := 0; j < len(trees[i][:]); j++ {
			//fmt.Println("\ncurrent: ", trees[i][j], "(", i, j, ")")

			// Add exterior
			if i == 0 || i == len(trees)-1 {
				visibleTrees = append(visibleTrees, trees[i][j])
				continue
			}
			if j == 0 || j == len(trees[i][:])-1 {
				visibleTrees = append(visibleTrees, trees[i][j])
				continue
			}

			// Check right / left
			row := trees[i]
			isVisible := false
			leftCount, rightCount := 0, 0
			for k := range row {
				// left
				if k < j && row[k] < trees[i][j] {
					leftCount += 1
				}
				// right
				if k > j && row[k] < trees[i][j] {
					rightCount += 1
				}
			}
			if leftCount == len(row[:j]) {
				isVisible = true
				//fmt.Println(trees[i][j], "is visible FROM LEFT, position ", i, j)
			}
			if rightCount == len(row[j+1:]) {
				isVisible = true
				//fmt.Println(trees[i][j], "is visible FROM RIGHt, position ", i, j)
			}

			// check top/bottom
			var col []int
			for k := range trees {
				col = append(col, trees[k][j])
			}
			topCount := 0
			bottomCount := 0
			for k := range col {
				// top
				if k < i && col[k] < trees[i][j] {
					//fmt.Println("over", col[k], i, k)
					topCount += 1
				}
				// bottom
				if k > i && col[k] < trees[i][j] {
					bottomCount += 1
				}
			}
			if topCount == len(col[:i]) {
				isVisible = true
				//fmt.Println(trees[i][j], "is visible FROM TOP, position ", i, j)
			}
			if bottomCount == len(col[i+1:]) {
				isVisible = true
				//fmt.Println(trees[i][j], "is visible FROM BOTTOM, position ", i, j)
			}

			// Add tree if visible
			if isVisible {
				visibleTrees = append(visibleTrees, trees[i][j])
			}
		}

	}
	//fmt.Println(visibleTrees)

	return len(visibleTrees)
}

func p2(trees [][]int) int {

	maxScenic := 0
	for i := 0; i < len(trees); i++ {
		//fmt.Println("\n row", i)
		for j:=0; j<len(trees[i]); j++ {
			// check
			//fmt.Println(trees[i][j])
			row := trees[i]

			// count right
			rightCount := 0
			latestRight := j
			for k := j+1; k < len(row)-1; k++ {
				if row[k] < trees[i][j] {
					rightCount += 1
					latestRight = k
				} else {
					break
				}
			}
			if j < len(row)-1 && latestRight != len(row)-1 {
				rightCount += 1
			}

			// count left
			leftCount := 0
			latestLeft := j
			for k := j-1; k >= 0; k-- {
				if row[k] < trees[i][j] {
					leftCount += 1
					latestLeft = k
				} else {
					break
				}
			}
			if j > 0 && latestLeft != 0  {
				leftCount += 1
			}

			var col []int
			for k := range trees {
				col = append(col, trees[k][j])
			}

			// count top
			topCount := 0
			latestTop := i
			for k := i-1; k >= 0; k-- {
				if col[k] < trees[i][j] {
					topCount += 1
					latestTop = k
				} else {
					break
				}
			}
			if i > 0 && latestTop != 0  {
				topCount += 1
			}

			// count bottom
			bottomCount := 0
			latestBottom := j
			for k := i+1; k < len(col)-1; k++ {
				if col[k] < trees[i][j] {
					bottomCount += 1
					latestBottom = k
				} else {
					break
				}
			}
			if i < len(col)-1 && latestBottom != len(col)-1 {
				bottomCount += 1
			}

			//fmt.Println("right count ", rightCount)
			//fmt.Println("left count ", leftCount)
			//fmt.Println("top count ", topCount )
			//fmt.Println("bottom count ", bottomCount)
			scenic := rightCount*leftCount*topCount*bottomCount
			//fmt.Println("scenic ", scenic)
			if scenic > maxScenic {
				maxScenic = scenic
			}

		}
	}
    return maxScenic
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Parse input
	var trees [][]int
	for _, v := range lines {
		row := strings.Split(v, "")
		rowInt, _ := utils.SliceAtoi(row)
		trees = append(trees, rowInt)
	}

	fmt.Println(p1(trees))
	fmt.Println(p2(trees))
}