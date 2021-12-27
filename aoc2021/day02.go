package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

func day02(puzzle_data []string) {
	// part 1
	h := 0
	d := 0
	for _, dataVal := range puzzle_data {
		cmd := strings.Split(dataVal, " ")
		val := p.MakeInt(cmd[1])
		if cmd[0] == "forward" {
			h += val
		}
		if cmd[0] == "down" {
			d += val
		}
		if cmd[0] == "up" {
			d -= val
		}
	}
	fmt.Printf("Day 2 (part 1): Depth = %d, Horizontal Position = %d, product = %d\n", d, h, d*h)

	// part 2
	h = 0
	aim := 0
	d = 0
	for _, dataVal := range puzzle_data {
		cmd := strings.Fields(dataVal)
		val := p.MakeInt(cmd[1])
		if cmd[0] == "forward" {
			h += val
			if aim > 0 {
				d += aim * val
			}
		}
		if cmd[0] == "down" {
			aim += val
		}
		if cmd[0] == "up" {
			aim -= val
		}
	}
	fmt.Printf("Day 2 (part 2): Aim = %d, Depth = %d, Horizontal Position = %d, product = %d\n", aim, d, h, d*h)
}
