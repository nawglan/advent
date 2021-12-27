package main

import (
	"fmt"
	"math"
	"strings"
)

func day14(puzzle_data []string) {
	rules := map[string][]string{}
	var text string
	for _, dataVal := range puzzle_data {
		if len(dataVal) > 0 {
			if strings.Contains(dataVal, "->") {
				find_replace := strings.Fields(dataVal)
				rules[find_replace[0]] = append(rules[find_replace[0]], fmt.Sprintf("%c%s", find_replace[0][0], find_replace[2]))
				rules[find_replace[0]] = append(rules[find_replace[0]], fmt.Sprintf("%s%c", find_replace[2], find_replace[0][1]))
			} else {
				text = dataVal
			}
		}
	}

	first_char := fmt.Sprintf("%c", text[0])

	// split orig string into chunks
	chunk_counts := map[string]int{}
	for n := 0; n < len(text)-1; n++ {
		chunk_counts[text[n:n+2]]++
	}

	// process chunks
	for i := 0; i < 40; i++ {
		new_counts := map[string]int{}
		for chunk, chunk_count := range chunk_counts {
			for _, new_chunk := range rules[chunk] {
				new_counts[new_chunk] += chunk_count
			}
		}
		chunk_counts = new_counts

		if i == 9 || i == 39 {
			count := map[string]int{first_char: 1}
			for chunk, chunk_count := range chunk_counts {
				for c, char := range strings.Split(chunk, "") {
					if c != 0 { // only count last character in each chunk
						count[char] += chunk_count
					}
				}
			}

			min := math.MaxInt
			max := 0
			for _, v := range count {
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
			}
			part := "part 1"
			if i == 39 {
				part = "part 2"
			}
			fmt.Printf("Day 14 (%s): The max count is %d min count is %d, output is %d\n", part, max, min, max-min)
		}
	}
}
