package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2022/utils"
)

// Converts "2-4,5-6" => [2,4],[5,6]
func toElfs(pair string) (elf1 []int, elf2 []int) {
	elfs := strings.Split(pair, ",")
	elf1, _ = utils.SliceAtoi(strings.Split(elfs[0], "-"))
	elf2, _ = utils.SliceAtoi(strings.Split(elfs[1], "-"))
	return elf1, elf2
}

func p1(pairs []string) (sum int) {
	for _, pair := range pairs {
		elf1, elf2 := toElfs(pair)
		// Check for full overlap
		if (elf2[1] >= elf1[1] && elf2[0] <= elf1[0]) ||
			(elf2[1] <= elf1[1] && elf2[0] >= elf1[0]) {
			sum += 1
		}
	}
	return sum
}

func p2(pairs []string) (sum int) {
	for _, pair := range pairs {
		elf1, elf2 := toElfs(pair)
		// Check for any overlap
		if (elf2[1] >= elf1[0] && elf2[0] <= elf1[1]) ||
			(elf2[1] <= elf1[0] && elf2[0] >= elf1[1]) {
			sum += 1
		}
	}
	return sum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
