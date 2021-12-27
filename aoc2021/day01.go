package main

import (
	"fmt"
	p "advent/util/parse"
)

func day01(puzzle_data []string) {
	var puzzleVals []int
	for _, val := range puzzle_data {
		puzzleVals = append(puzzleVals, p.MakeInt(val))
	}

	// part 1
	increases := 0
	for i := 1; i < len(puzzleVals); i++ {
		if puzzleVals[i] > puzzleVals[i-1] {
			increases++
		}
	}
	fmt.Printf("Day 1 (part 1): There were %d increases in depth.\n", increases)

	// part 2
	increases = 0
	last_index := len(puzzleVals) - (len(puzzleVals) % 3)
	for i := 1; i < last_index; i++ {
		prev := puzzleVals[i-1 : i+2]
		next := puzzleVals[i : i+3]
		if sumSlice(prev) < sumSlice(next) {
			increases++
		}
	}
	fmt.Printf("Day 1 (part 2): There were %d increases in depth.\n", increases)
}
