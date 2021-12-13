package main

import (
	"fmt"
)

var (
	octoEnergy []int
)

func increaseEnergy(grid []int) ([]int, int) {
	flashed := make(map[int]bool, 0)

	for i := range grid {
		grid[i]++
		if grid[i] > 9 {
			flashed[i] = true
		}
	}

	seen := make(map[int]bool, 0)
	for {
		for k := range flashed {
			seen[k] = true // can only flash once
			delete(flashed, k)
			for _, n := range append(neighbors(k), neighborsDiag(k)...) {
				if _, found := seen[n]; !found {
					grid[n]++
					if grid[n] > 9 {
						flashed[n] = true
					}
				}
			}
		}
		if len(flashed) == 0 {
			break
		}
	}

	for k := range seen {
		grid[k] = 0
	}

	return grid, len(seen)
}

func day11(puzzle_data []string) {
	// part 1
	sum := 0
	octoEnergy = makeIntGrid(puzzle_data, len(puzzle_data[0]), len(puzzle_data))

	for step := 0; step < 100; step++ {
		numFlashed := 0
		octoEnergy, numFlashed = increaseEnergy(octoEnergy)
		sum += numFlashed
	}

	fmt.Printf("Day 11 (part 1): There were %d flashes after 100 steps.\n", sum)

	// part 2
	octoEnergy = makeIntGrid(puzzle_data, len(puzzle_data[0]), len(puzzle_data))

	step := 1
	for {
		numFlashed := 0
		octoEnergy, numFlashed = increaseEnergy(octoEnergy)
		if numFlashed == len(octoEnergy) {
			break
		}
		step++
	}

	fmt.Printf("Day 11 (part 2): All of the Octopuses flashed on step %d\n", step)
}
