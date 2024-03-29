package main

import (
	"fmt"
	"regexp"
	"strings"
)

func format(s string) string {
	r := s[1 : len(s)-1]
	r = strings.ReplaceAll(r, `\\`, `\`)
	r = strings.ReplaceAll(r, `\"`, `"`)

	re := regexp.MustCompile(`(\\x[a-f0-9][a-f0-9])`)
	m := re.FindAllStringSubmatch(r, -1)
	if m != nil {
		for _, match := range m {
			for i := 1; i < len(match); i++ {
				r = strings.ReplaceAll(r, match[i], "D")
			}
		}
	}

	return r
}

func reformat(s string) string {
	r := strings.ReplaceAll(s, `\`, `\\`)
	r = strings.ReplaceAll(r, `"`, `\"`)
	r = `"` + r + `"`

	return r
}

func day08(puzzle_data []string) {
	sum, sum2 := 0, 0
	for _, line := range puzzle_data {
		sum += len(line) - len(format(line))
		sum2 += len(reformat(line)) - len(line)
	}
	fmt.Printf("Day 8 (part 1): Sum of the difference is %d\n", sum)
	fmt.Printf("Day 8 (part 2): Sum of the difference is %d\n", sum2)
}
