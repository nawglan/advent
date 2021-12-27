package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func day07(puzzle_data []string) {
	crabs := make([]int, 1+strings.Count(puzzle_data[0], ","))
	for i, val := range strings.Split(puzzle_data[0], ",") {
		crabs[i] = makeInt(val)
	}
	sort.Ints(crabs)

	// part1
	median := crabs[len(crabs)/2]
	sum := 0
	for _, val := range crabs {
		sum += int(math.Abs(float64(median) - float64(val)))
	}
	fmt.Printf("Day 7 (part 1): %d fuel is required.\n", sum)

	// part2
	crabDiffs := make([]int, len(crabs))
	for i := range crabs {
		for _, val := range crabs {
			diff := int(math.Abs(float64(val) - float64(crabs[i])))
			crabDiffs[i] += (diff * (diff + 1)) / 2
		}
	}
	min := math.MaxInt
	for _, val := range crabDiffs {
		if val < min {
			min = val
		}
	}
	fmt.Printf("Day 7 (part 2): %d fuel is required.\n", min)
}
