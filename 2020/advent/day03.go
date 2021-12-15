package main

import (
	"fmt"
	"strings"
)

func day03(puzzle_data []string) {
	trees := map[string]bool{}
	// load up the trees
	width := len(puzzle_data[0]) - 1
	height := 0
	for y, dataVal := range puzzle_data {
		height = y
		for x, char := range strings.Split(dataVal, "") {
			if char == "#" {
				trees[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}
	}

	trees_hit_per_run := 1

	for run, slope := range []string{"3,1", "1,1", "5,1", "7,1", "1,2"} {
		y := 0
		x := 0
		slope_right := makeInt(slope[:1])
		slope_down := makeInt(slope[2:])
		trees_hit := map[string]bool{}
		for y <= height {
			x += slope_right
			if x > width {
				x -= width + 1
			}
			y += slope_down
			coords := fmt.Sprintf("%d,%d", x, y)
			if _, found := trees[coords]; found {
				trees_hit[coords] = true
			}
		}
		trees_hit_per_run *= len(trees_hit)
		if run == 0 {
			fmt.Printf("Day 03 (part 1): There were %d trees hit.\n", len(trees_hit))
		}
	}
	fmt.Printf("Day 03 (part 2): The product of the trees hit per run is %d.\n", trees_hit_per_run)
	/*
		for y := 0; y <= height; y++ {
			for x := 0; x <= width; x++ {
				coords := fmt.Sprintf("%d,%d", x, y)
				if _, tree := trees[coords]; tree {
					if _, hit := trees_hit[coords]; hit {
						fmt.Printf("X")
					} else {
						fmt.Printf("#")
					}
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	*/
}
