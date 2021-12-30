package main

import (
	p "advent/util/parse"
	"fmt"
	"regexp"
)

func day12(puzzle_data []string) {
	re := regexp.MustCompile(`(-?\d+)`)
	m := re.FindAllStringSubmatch(puzzle_data[0], -1)
	sum := 0
	for _, match := range m {
		for i := 1; i < len(match); i++ {
			sum += p.MakeInt(match[i])
		}
	}
	fmt.Printf("Day 12 (part 1): sum of all the numbers is %d\n", sum)
	fmt.Printf("Day 12 (part 2): \n")
}
