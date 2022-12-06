package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2022/utils"
)

func findMarker(buffer string, n int) (marker int) {

	nLast := []string{}
	allChars := strings.Split(buffer, "")

	for i := 0; i < len(buffer); i++ {
		nLastMap := map[string]int{}
		nLast = allChars[i : i+n]
		for _, v := range nLast {
			nLastMap[v] += 1
		}

		// Check if we have n consecutive unique values
		if len(nLastMap) == n {
			return i + n
		}
	}
	return marker
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	buffer := lines[0]

	fmt.Println(findMarker(buffer, 4))
	fmt.Println(findMarker(buffer, 14))
}
