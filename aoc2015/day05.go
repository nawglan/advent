package main

import (
	"fmt"
	"regexp"
	"strings"
)

func hasPairs(s string) bool {
	for _, a := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"} {
		for _, b := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"} {
			if strings.Count(s, fmt.Sprintf("%s%s", a, b)) > 1 {
				return true
			}
		}
	}

	return false
}

func hasRepeat(s string) bool {
	test := regexp.MustCompile(`a.a|b.b|c.c|d.d|e.e|f.f|g.g|h.h|i.i|j.j|k.k|l.l|m.m|n.n|o.o|p.p|q.q|r.r|s.s|t.t|u.u|v.v|w.w|x.x|y.y|z.z`)

	return test.MatchString(s)
}

func day05(puzzle_data []string) {
	count := 0
	alwaysGood := regexp.MustCompile(`ab|cd|pq|xy`)
	doubleChar := regexp.MustCompile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)
	for _, line := range puzzle_data {
		if alwaysGood.MatchString(line) {
			continue
		}
		vowels := strings.Count(line, "a") + strings.Count(line, "e") + strings.Count(line, "i") + strings.Count(line, "o") + strings.Count(line, "u")
		if vowels > 2 && doubleChar.MatchString(line) {
			count++
		}
	}
	fmt.Printf("Day 5 (part 1): Number of naughty words: %d\n", count)

	count = 0
	for _, line := range puzzle_data {
		if hasPairs(line) && hasRepeat(line) {
			count++
		}
	}
	fmt.Printf("Day 5 (part 2): Number of nice words: %d\n", count)
}
