package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2022/utils"
)

// Maps entries to hands
var m = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissor",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissor",
}

// Maps hands to values
var values = map[string]int{
	"Rock":    1,
	"Paper":   2,
	"Scissor": 3,
}

type Hand struct {
	P1     string
	P2     string
	Winner string
}

func (hand *Hand) setWinner() {
	p1 := m[hand.P1]
	p2 := m[hand.P2]

	if p1 == p2 {
		hand.Winner = "Draw"
		return
	} else {
		hand.Winner = "2" // Default
	}

	switch p1 {
	case "Rock":
		if p2 == "Scissor" {
			hand.Winner = "1"
		}
	case "Paper":
		if p2 == "Rock" {
			hand.Winner = "1"
		}
	case "Scissor":
		if p2 == "Paper" {
			hand.Winner = "1"
		}
	}
}

func calculatePoints(hand Hand) (res int) {
	draw := 3
	win := 6

	switch hand.Winner {
	case "2":
		res = win + values[m[hand.P2]]
	case "Draw":
		res = draw + values[m[hand.P2]]
	case "1":
		res = values[m[hand.P2]]
	}
	return res
}

func p1(hands []Hand) (sum int) {
	for i := range hands {
		hand := &hands[i]
		hand.setWinner()
		sum += calculatePoints(*hand)
	}
	return sum
}

func p2(hands []Hand) (sum int) {
	vals := []string{"A", "B", "C", "A"}

	for i := range hands {
		hand := &hands[i]

		switch hand.P2 {
		case "X": // Lose and choose losing hand
			hand.Winner = "1"
			for k, v := range vals {
				if v == hand.P1 && k != 0 {
					hand.P2 = vals[k-1]
					break
				}
			}
		case "Y": // Draw and choose same hand
			hand.Winner = "Draw"
			hand.P2 = hand.P1
		case "Z": // Win and choose winning hand
			hand.Winner = "2"
			for k, v := range vals {
				if v == hand.P1 {
					hand.P2 = vals[k+1]
					break
				}
			}
		}
		sum += calculatePoints(*hand)
	}
	return sum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	var hands []Hand
	for i := range lines {
		parts := strings.Split(lines[i], " ")
		hand := Hand{P1: parts[0], P2: parts[1], Winner: "unknown"}
		hands = append(hands, hand)
	}

	fmt.Println(p1(hands))
	fmt.Println(p2(hands))
}
