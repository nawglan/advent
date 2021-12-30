package main

import (
	p "advent/util/parse"
	"fmt"
	"math"
	"sort"
)

func day10(puzzle_data []string) {
	adaptors := []int{0}
	for _, dataVal := range puzzle_data {
		adaptors = append(adaptors, p.MakeInt(dataVal))
	}

	sort.Ints(adaptors)
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

	one := 0
	three := 0
	diff := adaptors[0]
	for i := 1; i < len(adaptors); i++ {
		diff = int(math.Abs(float64(adaptors[i-1] - adaptors[i])))
		switch diff {
		case 1:
			one++
		case 3:
			three++
		}
	}

	fmt.Printf("Day 10 (part 1): The number of 1-jolt (%d) mul by 3-jolt (%d) is %d\n", one, three, one*three)

	count := 0
	fmt.Printf("Day 10 (part 2): The number of ways to connect the adaptors is %d\n", count)
}
