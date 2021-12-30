package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

func day04(puzzle_data []string) {
	i := 0
	zeros := regexp.MustCompile(`^00000.*`)
	for {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", puzzle_data[0], i))
		if zeros.MatchString(fmt.Sprintf("%x", h.Sum(nil))) {
			break
		}
		i++
	}
	fmt.Printf("Day 4 (part 1): %d\n", i)

	// start from where part 1 left off
	sixzeros := regexp.MustCompile(`^000000.*`)
	for {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", puzzle_data[0], i))
		if sixzeros.MatchString(fmt.Sprintf("%x", h.Sum(nil))) {
			break
		}
		i++
	}
	fmt.Printf("Day 4 (part 2): %d\n", i)
}
