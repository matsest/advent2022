package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2022/utils"
)

// findPriority converts a-z to 1-26 and A-Z to 27-52
func findPriority(letter string) (pri int) {
	r := []rune(letter)[0]
	if r > 96 {
		r -= 96
	} else {
		r -= 38
	}
	return int(r)
}

func p1(rucksacks []string) (sum int) {
	matches := map[string]int{}
	// Divide each rucksack into two compartments
	for _, rucksack := range rucksacks {
		compartment1 := rucksack[0 : len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]
		inRucksack := map[string]int{}

		// Find matching entries across compartments - only add one match per rucksack
		for _, r := range compartment1 {
			var letter string = string(r)
			_, exists := inRucksack[letter]

			if strings.Contains(compartment2, letter) && !exists {
				inRucksack[letter] += 1
			}
		}
		// Add to matches
		for k, v := range inRucksack {
			matches[k] += v
		}
	}

	// Summarize total matches across rucksacks
	for l, i := range matches {
		sum += findPriority(l) * i
	}

	return sum
}

func p2(rucksacks []string) (sum int) {
	matches := map[string]int{}

	// Divide rucksacks into groups of threee
	for i := 0; i < len(rucksacks); i += 3 {
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]
		inGroup := map[string]int{}

		// Check that there is a match across all rucksacks in group
		for _, r := range r1 {
			var letter string = string(r)
			_, exists := inGroup[letter]
			if (strings.Contains(r3, letter) && strings.Contains(r2, letter)) && !exists {
				inGroup[letter] += 1
			}
		}
		// Add to matches
		for k, v := range inGroup {
			matches[k] += v
		}
	}
	// Summarize total matches across rucksacks
	for l, i := range matches {
		sum += findPriority(l) * i
	}

	return sum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
