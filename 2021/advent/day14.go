package main

import (
	"fmt"
	"math"
	"strings"
)

func day14(puzzle_data []string) {
	rules := map[string]string{}
	var text string
	for _, dataVal := range puzzle_data {
		if len(dataVal) > 0 {
			if strings.Contains(dataVal, "->") {
				find_replace := strings.Fields(dataVal)
				rules[find_replace[0]] = fmt.Sprintf("%c%s", find_replace[0][0], find_replace[2])
			} else {
				text = dataVal
			}
		}
	}

	for _, itr := range []int{10, 40} {
		for i := 0; i < itr; i++ {
			chunks := []string{}
			for n := 0; n < len(text)-1; n++ {
				chunk := text[n : n+2]
				if _, found := rules[chunk]; found {
					chunk = rules[chunk]
				}
				chunks = append(chunks, chunk)
			}
			chunks = append(chunks, fmt.Sprintf("%c", text[len(text)-1]))
			text = strings.Join(chunks, "")
		}
		count := map[string]int{}
		for _, char := range strings.Split(text, "") {
			count[char]++
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
		if itr == 40 {
			part = "part 2"
		}
		fmt.Printf("Day 14 (%s): The max count is %d min count is %d, output is %d\n", part, max, min, max-min)
	}
}
