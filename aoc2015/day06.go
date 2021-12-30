package main

import (
	p "advent/util/parse"
	"fmt"
	"regexp"
)

func day06(puzzle_data []string) {
	grid := [1000][1000]bool{}

	regex := regexp.MustCompile(`^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

	for _, line := range puzzle_data {
		matches := regex.FindStringSubmatch(line)
		for y := p.MakeInt(matches[3]); y <= p.MakeInt(matches[5]); y++ {
			for x := p.MakeInt(matches[2]); x <= p.MakeInt(matches[4]); x++ {
				switch matches[1] {
				case "turn on":
					grid[x][y] = true
				case "turn off":
					grid[x][y] = false
				case "toggle":
					grid[x][y] = !grid[x][y]
				default:
					panic("Unknown command: " + matches[1])
				}
			}
		}
	}

	count := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if grid[x][y] {
				count++
			}
		}
	}

	fmt.Printf("Day 6 (part 1): %d lights are lit\n", count)

	grid2 := [1000][1000]int{}
	for _, line := range puzzle_data {
		matches := regex.FindStringSubmatch(line)
		for y := p.MakeInt(matches[3]); y <= p.MakeInt(matches[5]); y++ {
			for x := p.MakeInt(matches[2]); x <= p.MakeInt(matches[4]); x++ {
				switch matches[1] {
				case "turn on":
					grid2[x][y]++
				case "turn off":
					grid2[x][y]--
					if grid2[x][y] < 0 {
						grid2[x][y] = 0
					}
				case "toggle":
					grid2[x][y] += 2
				default:
					panic("Unknown command: " + matches[1])
				}
			}
		}
	}

	brightness := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			brightness += grid2[x][y]
		}
	}

	fmt.Printf("Day 6 (part 2): Total brightness is %d\n", brightness)
}
