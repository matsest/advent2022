package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/matsest/advent2022/utils"
)

func p1(elves [][]int) (maxSum int) {
	maxSum = 0
	for i := range elves {
		elfSum := 0
		for j := range elves[i] {
			elfSum += elves[i][j]
		}
		if elfSum > maxSum {
			maxSum = elfSum
		}
	}
	return maxSum
}

func p2(elves [][]int) int {
	var sums []int
	for i := range elves {
		elfSum := 0
		for j := range elves[i] {
			elfSum += elves[i][j]
		}
		sums = append(sums, elfSum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	sum := 0
	for i := 0; i < 3; i++ {
		sum += sums[i]
	}
	return sum
}

func main() {
	// Read input
	lines, _ := utils.ReadLines("input.txt")
	var elves [][]int
	var elf []int
	current := 0
	for i := range lines {
		if lines[i] == "" {
			current += 1
			elves = append(elves, elf)
			elf = []int{}
		} else {
			num, _ := strconv.Atoi(lines[i])
			elf = append(elf, num)

			if i == len(lines)-1 {
				elves = append(elves, elf)
			}
		}
	}

	fmt.Println(p1(elves))
	fmt.Println(p2(elves))
}
