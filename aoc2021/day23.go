package main

import (
	"fmt"
)

// 012345678910
//#############
//#...........#
//###B#C#B#D###
//  #A#D#C#A#
//  #########

type Hallway struct {
	spot [11]string
}

type Room struct {
	name string    // A, B, C, D
	pos  int       // 2, 4, 6, 8 // under what spot in hallway
	spot [2]string // slot 0 is the top spot
}

func (r Room) isEmpty() bool {
	return r.spot[0] == "" && r.spot[1] == ""
}

func (r Room) contents() (string, string) {
	return r.spot[0], r.spot[1]
}

func (r Room) isFull() bool {
	return r.spot[0] == r.name && r.spot[1] == r.name
}

type Peg struct {
	name  string // A,  B,   C,    D
	cost  int    // 1, 10, 100, 1000
	steps int
	pos   [2][11]int // 0 == hallway, 1 == room, pos
}

var (
	rooms = map[string]Room{
		"A": {"A", 2, [2]string{}},
		"B": {"B", 4, [2]string{}},
		"C": {"C", 6, [2]string{}},
		"D": {"D", 8, [2]string{}},
	}
	pegs = map[string]Peg{
		"A1": {"A", 1, 0, [2][11]int{}},
	}
)

func day23(puzzle_data []string) {
	fmt.Println("day23")
	/*
		for _, dataVal := range puzzle_data {
		}
	*/
}
