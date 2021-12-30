package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

func day02(puzzle_data []string) {
	count_part1 := 0
	count_part2 := 0
	for _, dataVal := range puzzle_data {
		data := strings.Fields(dataVal)
		minmax := strings.Split(data[0], "-")
		need := data[1][0]
		contains := strings.Count(data[2], fmt.Sprintf("%c", need))
		if contains >= p.MakeInt(minmax[0]) && contains <= p.MakeInt(minmax[1]) {
			count_part1++
		}
		first_pos := p.MakeInt(minmax[0]) - 1
		second_pos := p.MakeInt(minmax[1]) - 1
		if (data[2][first_pos] == need || data[2][second_pos] == need) &&
			data[2][first_pos] != data[2][second_pos] {
			count_part2++
		}
	}
	fmt.Printf("Day 02 (part 1): There are %d valid passwords\n", count_part1)
	fmt.Printf("Day 02 (part 2): There are %d valid passwords\n", count_part2)
}
