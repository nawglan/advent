package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

func move(start string, d rune) string {
	coords := strings.Split(start, ",")
	x, y := p.MakeInt(coords[0]), p.MakeInt(coords[1])
	switch d {
	case '>':
		x++
	case '<':
		x--
	case '^':
		y--
	case 'v':
		y++
	}

	idx := fmt.Sprintf("%d,%d", x, y)
	return idx
}

func day03(puzzle_data []string) {
	santa := "0,0"

	houses := map[string]int{}

	// deliver a present to the first house
	houses[santa]++
	for _, char := range puzzle_data[0] {
		santa = move(santa, char)
		houses[santa]++
	}

	fmt.Printf("Day 3 (part 1): Number of houses that received at least one present: %d\n", len(houses))

	houses = map[string]int{}
	santa = "0,0"
	robot_santa := "0,0"
	// deliver a present to the first house
	houses[santa]++
	houses[robot_santa]++
	santa_turn := true
	for _, char := range puzzle_data[0] {
		if santa_turn {
			santa = move(santa, char)
			houses[santa]++
			santa_turn = false
		} else {
			robot_santa = move(robot_santa, char)
			houses[robot_santa]++
			santa_turn = true
		}
	}

	fmt.Printf("Day 3 (part 2): Number of houses that received at least one present: %d\n", len(houses))
}
