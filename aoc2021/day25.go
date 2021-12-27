package main

import (
	"fmt"
)

var (
	eastFlock  = map[int]bool{}
	southFlock = map[int]bool{}
)

func dump() {
	for y := 0; y < gridLength; y++ {
		fmt.Printf("DEZ:")
		for x := 0; x < gridWidth; x++ {
			idx := coords2pos(x, y)
			if eastFlock[idx] {
				fmt.Printf(">")
			} else {
				if southFlock[idx] {
					fmt.Printf("v")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println()
	}
}

func day25(puzzle_data []string) {
	gridWidth = len(puzzle_data[0])
	gridLength = len(puzzle_data)

	for y, dataVal := range puzzle_data {
		for x, char := range dataVal {
			switch char {
			case 'v':
				southFlock[coords2pos(x, y)] = true
			case '>':
				eastFlock[coords2pos(x, y)] = true
			}
		}
	}
	//dump()
	step := 0

	movedEast := false
	movedSouth := false
	for {
		step++
		//fmt.Printf("DEZ: ---------- %d -----------\n", step)
		//fmt.Printf("DEZ: ---------- EAST -----------\n")
		eastFlock, movedEast = move(eastFlock, true)
		//dump()
		//fmt.Printf("DEZ: ---------- SOUTH -----------\n")
		southFlock, movedSouth = move(southFlock, false)
		//dump()
		if !movedEast && !movedSouth {
			break
		}
	}

	fmt.Printf("Day 25 (part 1): The first step that no sea cucumbers can move is: %d\n", step)
}

func move(flock map[int]bool, east bool) (map[int]bool, bool) {
	newPos := make(map[int]bool, len(flock))

	moved := false

	for i := range flock {
		x := i % gridWidth
		y := i / gridWidth
		newIdx := 0
		if east {
			if x == gridWidth-1 {
				x = 0
			} else {
				x++
			}
			newIdx = coords2pos(x, y)
		} else {
			if y == gridLength-1 {
				y = 0
			} else {
				y++
			}
			newIdx = coords2pos(x, y)
		}
		if !eastFlock[newIdx] && !southFlock[newIdx] {
			moved = true
			newPos[newIdx] = true
		} else {
			newPos[i] = true
		}
	}

	return newPos, moved
}
