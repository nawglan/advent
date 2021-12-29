package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

type Box struct {
	l, w, h int
}

func (b Box) paper_needed() int {
	min := b.l * b.w
	if b.w*b.h < min {
		min = b.w * b.h
	}
	if b.h*b.l < min {
		min = b.h * b.l
	}
	need := min + 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l

	return need
}

func (b Box) ribbon_needed() int {
	min := 2*b.l + 2*b.w
	if 2*b.w+2*b.h < min {
		min = 2*b.w + 2*b.h
	}
	if 2*b.h+2*b.l < min {
		min = 2*b.h + 2*b.l
	}
	need := min + b.l*b.w*b.h

	return need
}

func day02(puzzle_data []string) {
	paper_sqft := 0
	ribbon_ft := 0
	for _, line := range puzzle_data {
		sides := strings.Split(line, "x")
		box := Box{
			l: p.MakeInt(sides[0]),
			w: p.MakeInt(sides[1]),
			h: p.MakeInt(sides[2]),
		}
		paper_sqft += box.paper_needed()
		ribbon_ft += box.ribbon_needed()
	}
	fmt.Printf("Day 2 (part 1): Total square feet of wrapping paper needed is %d\n", paper_sqft)
	fmt.Printf("Day 2 (part 2): Total feet of ribbon needed is %d\n", ribbon_ft)
}
