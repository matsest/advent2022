package main

import (
	"fmt"
	"regexp"

	"github.com/matsest/advent2022/utils"
)

// Reads stacks from input
func getStacks(lines []string) [][]string {
	// Find out how many stacks to initialize based on 4 characters length
	numbersLength := (len(lines[0]) + 1) / 4
	arrs := make([][]string, numbersLength)

	for _, v := range lines {
		// Lazy match to stop at the first non-interesting line..
		stop, _ := regexp.Match(`^(\s+(\d\s+)+)\d`, []byte(v))
		if stop {
			break
		}

		// Find all letters - assume 4 character length
		tmp := []string{}
		for i := 1; i < len(v); i += 4 {
			current := string(v[i])
			tmp = append(tmp, current)
		}

		// Flip index to add to columns instead of rows
		for k := range tmp {
			if tmp[k] != " " {
				arrs[k] = append(arrs[k], tmp[k])
			}
		}
	}
	return arrs
}

// Reads instructions from input
func getInstructions(lines []string) (instructions [][]int) {
	for _, v := range lines {
		// Find lines matching instruction pattern
		re := regexp.MustCompile(`^move \d{1,2} from \d to \d$`)
		match := re.MatchString(v)
		if match {
			// Find integers
			re := regexp.MustCompile(`\d{1,2}`)
			numsString := re.FindAllString(v, -1)
			nums, _ := utils.SliceAtoi(numsString)
			count := nums[0]

			// Remove 1 due to 1-indexing in input
			fromIndex := nums[1] - 1
			toIndex := nums[2] - 1

			// Add to instructions
			current := []int{count, fromIndex, toIndex}
			instructions = append(instructions, current)
		}
	}
	return instructions
}

// Helper function to copy 2D slice
func clone(arr [][]string) (res [][]string) {
	res = make([][]string, len(arr))
	for i := range arr {
		res[i] = append([]string{}, arr[i]...)
	}
	return
}

func doInstructions(stacks [][]string, instructions [][]int, reverse bool) (ans [][]string) {
	ans = clone(stacks)

	for _, match := range instructions {
		count := match[0]
		from := match[1]
		to := match[2]

		// Pop from the FROM stack
		tmpFrom := make([]string, len(ans[from]))
		copy(tmpFrom, ans[from])
		front, _ := tmpFrom[:count], tmpFrom[count:]

		// Reverse the front crates before adding
		if reverse {
			for i, j := 0, len(front)-1; i < j; i, j = i+1, j-1 {
				front[i], front[j] = front[j], front[i]
			}
		}

		// Add front crates to TO stack
		ans[to] = append(front, ans[to]...)

		// Clean up the FROM stack
		ans[from] = ans[from][count:]
	}
	return ans
}

func getTopCrates(stacks [][]string) (ans string) {
	for i := range stacks {
		ans += (stacks[i][0])
	}
	return ans
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Find stacks from input
	stacks := getStacks(lines)

	// Find instructions from input
	instructions := getInstructions(lines)

	// Part 1
	arrs1 := doInstructions(stacks, instructions, true)
	fmt.Println(getTopCrates(arrs1))

	// Part 2
	arrs2 := doInstructions(stacks, instructions, false)
	fmt.Println(getTopCrates(arrs2))
}
