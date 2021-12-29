package main

import (
	"fmt"
	"strings"
)

func day01(puzzle_data []string) {
	fmt.Printf("Day 1 (part 2): Santa ends up on floor %d.\n", strings.Count(puzzle_data[0], "(")-strings.Count(puzzle_data[0], ")"))

	floor := 0
	i := 0
	for _, char := range puzzle_data[0] {
		i++
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			break
		}
	}
	fmt.Printf("Day 1 (part 2): Santa enters the basement at position %d.\n", i)
}
