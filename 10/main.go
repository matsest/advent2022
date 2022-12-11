package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2022/utils"
)

func p1(lines []string) (sum int) {

	x, cycle := 1, 1
	nextCycleToCheck := 20
	queue := []int{}
	remainingWorkingCycles := 1

	for _, l := range lines {
		//fmt.Println("\ncycle ", cycle, " x is ", x, "remaining working cycles: ", remainingWorkingCycles)
		//fmt.Println("queue: ", queue)
		//fmt.Println(l)
		parts := strings.Split(l, " ")

		if cycle == nextCycleToCheck {
			fmt.Println("cycle * x: ", cycle*x, cycle, x)
			sum += cycle * x
			nextCycleToCheck += 40
		}

		if parts[0] == "noop" {
			queue = append(queue, 0)
		} else if parts[0] == "addx" {
			num, _ := strconv.Atoi(parts[1])
			queue = append(queue, num)
			//fmt.Println("adding to queue: ", num, "queue length ", len(queue))
		}

		if remainingWorkingCycles == 0 {
			if queue[0] != 0 {
				x += queue[0]
				//fmt.Println("adding ", queue[0], "to x")
				queue = queue[1:]
				remainingWorkingCycles = 1
			} else {
				queue = queue[1:]
			}
		} else if remainingWorkingCycles > 0 {
			remainingWorkingCycles -= 1
		}
		cycle += 1
	}

	// Finish the queue
	for len(queue) > 0 {
		//fmt.Println("\ncycle ", cycle, " x is ", x)
		if cycle == nextCycleToCheck {
			fmt.Println("cycle * x: ", cycle*x, cycle, x)
			sum += cycle * x
			nextCycleToCheck += 40
		}

		if remainingWorkingCycles == 0 {
			if queue[0] != 0 {
				x += queue[0]
				queue = queue[1:]
				remainingWorkingCycles = 1
			} else {
				queue = queue[1:]
			}
		} else if remainingWorkingCycles > 0 {
			remainingWorkingCycles -= 1
		}
		cycle += 1
	}
	return sum
}

func p2(lines []string) {

	x, cycle := 1, 1
	nextCycleToCheck := 40
	queue := []int{}
	remainingWorkingCycles := 1

	img := [6][40]string{}
	row, col := 0, 0

	for _, l := range lines {

		//fmt.Println("\ncycle ", cycle, " x is ", x, "remaining working cycles: ", remainingWorkingCycles)
		//fmt.Println("queue: ", queue)
		//fmt.Println(l)
		parts := strings.Split(l, " ")

		if x-1 <= col && col <= x+1 {
			img[row][col] = "#"
		} else {
			img[row][col] = "."
		}
		col += 1

		if cycle == nextCycleToCheck {
			//fmt.Println("cycle * x: ", cycle*x, cycle, x)
			row += 1
			col = 0
			nextCycleToCheck += 40
		}

		if parts[0] == "noop" {
			queue = append(queue, 0)
		} else if parts[0] == "addx" {
			num, _ := strconv.Atoi(parts[1])
			queue = append(queue, num)
			//fmt.Println("adding to queue: ", num, "queue length ", len(queue))
		}

		if remainingWorkingCycles == 0 {
			if queue[0] != 0 {
				x += queue[0]
				//fmt.Println("adding ", queue[0], "to x")

				queue = queue[1:]
				remainingWorkingCycles = 1
			} else {
				queue = queue[1:]
			}
		} else if remainingWorkingCycles > 0 {
			remainingWorkingCycles -= 1
		}
		cycle += 1
	}

	// Finish the queue
	for cycle <= 240 {
		//fmt.Println("\ncycle ", cycle, " x is ", x)

		if x-1 <= col && col <= x+1 {
			img[row][col] = "#"
		} else {
			img[row][col] = "."
		}
		col += 1

		if cycle == nextCycleToCheck {
			//fmt.Println("cycle * x: ", cycle*x, cycle, x)
			row += 1
			col = 0
			nextCycleToCheck += 40
		}

		if remainingWorkingCycles == 0 {
			if queue[0] != 0 {
				x += queue[0]

				queue = queue[1:]
				remainingWorkingCycles = 1
			} else {
				queue = queue[1:]
			}
		} else if remainingWorkingCycles > 0 {
			remainingWorkingCycles -= 1
		}
		cycle += 1
	}
	for _, row := range img {
		fmt.Println(row)
	}
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// P1
	fmt.Println(p1(lines))
	p2(lines)
}
