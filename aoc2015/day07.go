package main

import (
	"fmt"
	"regexp"
)

func day07(puzzle_data []string) {
	not := regexp.MustCompile(`^NOT (\c+) -> (\c+)`)
	or := regexp.MustCompile(`^(\c+) OR (\c+) -> (\c+)`)
	applycurrent := regexp.MustCompile(`^(\d+) -> (\c+)`)
	and := regexp.MustCompile(`^(\c+) AND (\c+) -> (\c+)`)
	rshift := regexp.MustCompile(`^(\c+) RSHIFT (\d+) -> (\c+)`)
	lshift := regexp.MustCompile(`^(\c+) LSHIFT (\d+) -> (\c+)`)

	for _, line := range puzzle_data {
		for key, re := range map[string]*regexp.Regexp{"not": not, "or": or, "applycurrent": applycurrent, "and": and, "rshift": rshift, "lshift": lshift} {
			m := re.FindStringSubmatch(line)
			if len(m) > 0 {
				switch key {
				}
			}
		}
	}

	fmt.Printf("Day 7 (part 1): \n")
}
