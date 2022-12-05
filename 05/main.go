package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/matsest/advent2022/utils"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Find out how many arrays to initialize based on 4 characters
	numbersLength := (len(lines[0]) + 1) / 4
	//fmt.Println(numbersLength)
	arrs := make([][]string, numbersLength)
	//fmt.Println(arrs)

	for _, v := range lines {
		// Dirty match to find the first non-interesting line..
		toStop, _ := regexp.Match(`^(\s+(\d\s+)+)\d`, []byte(v))
		if toStop {
			break
		}

		// Find all letters - quick and dirty based on 4 characters per
		tmp := []string{}
		for i := 1; i < len(v); i += 4 {
			current := string(v[i])
			tmp = append(tmp, current)

		}
		// Flip index to add to columns instead of rows
		//fmt.Println("tmp: ", tmp)
		for k := range tmp {
			//fmt.Println("tmp[k]:", tmp[k])
			if tmp[k] != " " {
				arrs[k] = append(arrs[k], tmp[k])
			}
		}
	}

	fmt.Println("initial arrs: ", arrs)

	// Find instructions
	//fmt.Println("MATCHES:")
	var matches [][]int
	for _, v := range lines {
		re := regexp.MustCompile(`^move \d{1,2} from \d to \d$`)
		match := re.MatchString(v)
		if match {
			//fmt.Printf("%q\n", re.FindAll([]byte(v), -1))
			re := regexp.MustCompile(`\d{1,2}`)
			nums := re.FindAllString(v, -1)
			count, _ := strconv.Atoi(nums[0])
			fromIndex, _ := strconv.Atoi(nums[1])
			fromIndex -= 1
			toIndex, _ := strconv.Atoi(nums[2])
			toIndex -= 1
			//fmt.Println(count, fromIndex, toIndex)
			current := []int{count, fromIndex, toIndex}
			matches = append(matches, current)
		}
	}
	//fmt.Println(matches)

	// Final output
	//fmt.Println("initial arr: ", arrs)
	//for i := range arrs {
	//	fmt.Println(arrs[i])
	//}

	//fmt.Println("beginning altering arr")
	for _, match := range matches {
		count := match[0]
		from := match[1]
		to := match[2]

		//fmt.Println("moving value ", arrs[from][0:count], "from", from, "to ", to, " (count ", count, ")")

		// pop from front of from
		tmpFrom := make([]string, len(arrs[from]))
		copy(tmpFrom, arrs[from])
		front, _ := tmpFrom[:count], tmpFrom[count:]
		//fmt.Println("front, prev:", front, prev)

		// Reverse front stack
		// comment this out for part 2
		for i, j := 0, len(front)-1; i < j; i, j = i+1, j-1 {
			front[i], front[j] = front[j], front[i]
		}

		// Add to front of to
		arrs[to] = append(front, arrs[to]...)

		// Remove from back
		arrs[from] = arrs[from][count:]

		//fmt.Println("changed arrs: ", arrs)
		//for i := range arrs {
		//	fmt.Println(arrs[i])
		//}

	}

	// Final output
	fmt.Println("arrs: ", arrs)
	//for i := range arrs {
	//	fmt.Println(arrs[i])
	//}

	// Ans
	ans := ""
	for i := range arrs {
		ans += (arrs[i][0])
	}
	fmt.Println("ANSWER", ans)
}
