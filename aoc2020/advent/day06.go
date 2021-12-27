package main

import (
	"fmt"
	"strings"
)

func day06(puzzle_data []string) {
	questions := map[string]int{}
	counts_part_1 := []int{}
	counts_part_2 := []int{}
	group_count := 0
	for _, dataVal := range puzzle_data {
		if len(dataVal) != 0 {
			group_count++
			for _, char := range strings.Split(dataVal, "") {
				questions[char]++
			}
		} else {
			counts_part_1 = append(counts_part_1, len(questions))
			all_yes := 0
			for _, q_count := range questions {
				if q_count == group_count {
					all_yes++
				}
			}
			counts_part_2 = append(counts_part_2, all_yes)
			questions = map[string]int{}
			group_count = 0
		}
	}
	// add the last group
	counts_part_1 = append(counts_part_1, len(questions))
	all_yes := 0
	for _, q_count := range questions {
		if q_count == group_count {
			all_yes++
		}
	}
	counts_part_2 = append(counts_part_2, all_yes)

	fmt.Printf("Day 06 (part 1): The sum is %d\n", sumSlice(counts_part_1))
	fmt.Printf("Day 06 (part 2): The sum is %d\n", sumSlice(counts_part_2))
}
