package main

import (
	"fmt"
	"strings"
	p "advent/util/parse"
)

func day05(puzzle_data []string) {
	const gridSize int = 1000
	// part 1
	var grid [gridSize][gridSize]int

	for _, data_value := range puzzle_data {
		values := strings.Fields(data_value)
		pointA := strings.Split(values[0], ",")
		pointB := strings.Split(values[2], ",")

		xA := p.MakeInt(pointA[0])
		yA := p.MakeInt(pointA[1])
		xB := p.MakeInt(pointB[0])
		yB := p.MakeInt(pointB[1])

		// skip diagonal lines
		if xA == xB || yA == yB {
			grid[xB][yB]++
			for {
				grid[xA][yA]++
				if xA == xB { // vertical line
					if yA < yB {
						yA++
					} else {
						yA--
					}
				} else {
					if yA == yB { // horizontal line
						if xA < xB {
							xA++
						} else {
							xA--
						}
					}
				}
				if xA == xB && yA == yB {
					break
				}
			}
		}
	}
	count := 0
	for y := range grid {
		for _, val := range grid[y] {
			if val > 1 {
				count++
			}
		}
	}
	fmt.Printf("Day 5 (part 1): There are %d points crossed by more than one line\n", count)

	// part 2 - add 45 degree diagonal lines
	for _, data_val := range puzzle_data {
		values := strings.Fields(data_val)
		pointA := strings.Split(values[0], ",")
		pointB := strings.Split(values[2], ",")

		if pointA[0] != pointB[0] && pointA[1] != pointB[1] {
			xA := p.MakeInt(pointA[0])
			yA := p.MakeInt(pointA[1])
			xB := p.MakeInt(pointB[0])
			yB := p.MakeInt(pointB[1])

			grid[xB][yB]++
			for {
				grid[xA][yA]++
				if xA < xB {
					xA++
				} else {
					xA--
				}
				if yA < yB {
					yA++
				} else {
					yA--
				}
				if xA == xB && yA == yB {
					break
				}
			}
		}
	}

	count = 0
	for y := range grid {
		for _, val := range grid[y] {
			if val > 1 {
				count++
			}
		}
	}
	fmt.Printf("Day 5 (part 2): There are %d points crossed by more than one line\n", count)
}
