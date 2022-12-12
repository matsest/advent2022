package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/matsest/advent2022/utils"
)

func (m Monkey) Operation(old int) (int, error) {
	if m.OperationWithSelf {
		if m.Operator == "plus" {
			return old + old, nil
		} else if m.Operator == "multiply" {
			return old * old, nil
		}
	}

	if m.Operator == "plus" {
		return old + m.OperationConst, nil

	} else if m.Operator == "multiply" {
		return old * m.OperationConst, nil

	} else {
		return 0, fmt.Errorf("invalid operator for monkey %v", m.Id)
	}
}

func (m Monkey) Test(num int) (bool, error) {
	if m.DivisionConst <= 0 {
		return false, fmt.Errorf("invalid division const for monkey %v", m.Id)
	}

	if num%m.DivisionConst == 0 {
		return true, nil
	}
	return false, nil
}

type Monkey struct {
	Id                int
	Items             []int
	InspectCount      int
	Operator          string
	OperationConst    int
	OperationWithSelf bool
	DivisionConst     int
	TestTrueMonkeyint int
	TestFalseMonkeyId int
}

func createMonkeys(lines []string) []Monkey {
	entryPos := 0
	monkeys := []Monkey{}
	monkey := new(Monkey)

	for _, line := range lines {
		//fmt.Println(line)

		switch entryPos {
		case 0: // set id
			re, _ := regexp.Compile(`\d`)
			parts := re.FindString(line)
			//fmt.Println("match: ", parts)
			num, _ := strconv.Atoi(parts)
			monkey.Id = num

		case 1: // add items
			re, _ := regexp.Compile(`\d+`)
			parts := re.FindAllString(line, 10)
			//fmt.Println("match: ", parts)
			nums, _ := utils.SliceAtoi(parts)
			monkey.Items = nums
		case 2: // Add operation
			parts := strings.Split(line, " ")
			if parts[6] == "+" {
				monkey.Operator = "plus"
			} else if parts[6] == "*" {
				monkey.Operator = "multiply"
			}
			re, _ := regexp.Compile(`old`)
			parts = re.FindAllString(line, 2)
			if len(parts) == 2 {
				monkey.OperationWithSelf = true
				monkey.OperationConst = 0
			} else {
				monkey.OperationWithSelf = false
				re, _ := regexp.Compile(`\d+`)
				parts := re.FindString(line)
				num, _ := strconv.Atoi(parts)
				monkey.OperationConst = num
			}
		case 3: // division
			re, _ := regexp.Compile(`\d+`)
			parts := re.FindString(line)
			num, _ := strconv.Atoi(parts)
			monkey.DivisionConst = num
		case 4: // throw to if true
			re, _ := regexp.Compile(`\d+`)
			parts := re.FindString(line)
			num, _ := strconv.Atoi(parts)
			monkey.TestTrueMonkeyint = num
		case 5: // throw to if false
			re, _ := regexp.Compile(`\d+`)
			parts := re.FindString(line)
			num, _ := strconv.Atoi(parts)
			monkey.TestFalseMonkeyId = num
		}
		entryPos += 1
		if entryPos == 7 {
			entryPos = 0
			monkeys = append(monkeys, *monkey)
		}
	}
	// add latest monkey
	monkeys = append(monkeys, *monkey)
	return monkeys
}

func roundDivision(i int) int {
	return int(math.Floor(float64(i) / 3))
}

func p1(monkeys []Monkey, divide bool, count int) int {

	round := 1

	newMod := 1
	for _, m := range monkeys {
		newMod *= m.DivisionConst
	}

	for round <= count {
		//fmt.Println("\n round ", round)
		// Play
		for _, m := range monkeys {
			//fmt.Println("\nmonkey ", m.Id)
			for _, i := range m.Items {
				//fmt.Println("inspecting ", i)
				i, err := m.Operation(i)
				if err != nil {
					fmt.Println("ERROR operating")
					return 0
				}
				// divide by three and round
				var toGive int
				if divide {
					toGive = roundDivision(i)
				} else {
					toGive = i%newMod
				}
				test, err := m.Test(toGive)
				if err != nil {
					fmt.Println("ERROR divising")
					return 0
				}
				if test {
					monkeys[m.TestTrueMonkeyint].Items = append(monkeys[m.TestTrueMonkeyint].Items, toGive)
				} else {
					monkeys[m.TestFalseMonkeyId].Items = append(monkeys[m.TestFalseMonkeyId].Items, toGive)
				}
			}
			monkeys[m.Id].InspectCount += len(m.Items)
			monkeys[m.Id].Items = []int{}
		}
		round += 1
	}

	inspectionCounts := []int{}
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectCount)
	}
	//fmt.Println(inspectionCounts)
	sort.Ints(inspectionCounts)
	return inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2]
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	monkeys := createMonkeys(lines)
	fmt.Println(p1(monkeys, true, 20))

	newMonkeys := createMonkeys(lines)
	fmt.Println(p1(newMonkeys, false, 10000))
}
