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
			if i == 0 || i == len(trees) - 1 {
				visibleTrees = append(visibleTrees, trees[i][j])
				continue
			}
			if j == 0 || j == len(trees[i][:]) - 1 {
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

}