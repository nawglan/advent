package main

import (
	"fmt"
	"sort"
)

var (
	heightMap []int
)

func calcBasinSize(lowPoint int) int {
	sum := 1
	heightMap[lowPoint] = 9
	for _, check := range neighbors(lowPoint) {
		if heightMap[check] != 9 {
			sum += calcBasinSize(check)
		}
	}
	return sum
}

func day09(puzzle_data []string) {
	// part 1
	heightMap = makeIntGrid(puzzle_data, len(puzzle_data[0]), len(puzzle_data))

	sum := 0
	lowPoints := make([]int, 0)
	for i := range heightMap {
		found := true
		for _, n := range neighbors(i) {
			if heightMap[n] <= heightMap[i] {
				found = false
			}
		}
		if found {
			sum += heightMap[i] + 1
			lowPoints = append(lowPoints, i)
		}
	}
	fmt.Printf("Day 9 (part 1): Sum of the risk levels is %d\n", sum)

	// part 2
	for l, lowPoint := range lowPoints {
		lowPoints[l] = calcBasinSize(lowPoint)
	}
	sort.Ints(lowPoints)
	fmt.Printf("Day 9 (part 2): Product of the top three basin sizes is %d\n", mulSlice(lowPoints[len(lowPoints)-3:]))
}
