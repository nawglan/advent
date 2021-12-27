package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

func day06(puzzle_data []string) {
	maxDays := [2]int{80, 256}
	for _, max := range maxDays {
		counts := make([]int, 9)
		for _, val := range strings.Split(puzzle_data[0], ",") {
			counts[p.MakeInt(val)]++
		}

		for numDays := 0; numDays < max; numDays++ {
			reproduced := counts[0]
			counts = counts[1:]
			counts[6] += reproduced
			counts = append(counts, reproduced)
		}

		sum := 0
		for _, val := range counts {
			sum += val
		}

		if max == 80 {
			fmt.Printf("Day 6 (part 1): There are %d fish in the school\n", sum)
		} else {
			fmt.Printf("Day 6 (part 2): There are %d fish in the school\n", sum)
		}
	}
}
