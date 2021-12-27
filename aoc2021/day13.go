package main

import (
	"fmt"
	"strings"
	p "advent/util/parse"
)

func printPaper(paper map[string]int, maxX int, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, found := paper[fmt.Sprintf("%d,%d", x, y)]; found {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

var (
	maxX  int
	maxY  int
	folds = []string{}
	paper = map[string]int{}
)

func makePaper(puzzle_data []string) {
	for _, dataVal := range puzzle_data {
		if strings.Contains(dataVal, ",") {
			paper[dataVal] = 1 // store the mark

			coords := strings.Split(dataVal, ",")
			x := p.MakeInt(coords[0])
			y := p.MakeInt(coords[1])
			// store maxX and maxY
			if maxX < x {
				maxX = p.MakeInt(coords[0])
			}
			if maxY < y {
				maxY = p.MakeInt(coords[1])
			}
		}
		if strings.Contains(dataVal, "fold") {
			command := strings.Split(dataVal, " ")
			folds = append(folds, command[2])
		}
	}
}

func day13(puzzle_data []string) {
	makePaper(puzzle_data)

	for _, part := range []string{"part 1", "part 2"} {
		for f, fold := range folds {
			cmd := strings.Split(fold, "=")
			foldpos := p.MakeInt(cmd[1])

			// mark transposed marks
			for k := range paper {
				coords := strings.Split(k, ",")
				x := p.MakeInt(coords[0])
				y := p.MakeInt(coords[1])

				if cmd[0] == "y" {
					if y > foldpos {
						newY := foldpos - (y - foldpos)
						newKey := fmt.Sprintf("%d,%d", x, newY)
						paper[newKey] = 1
						delete(paper, k)
					}
				} else {
					if x > foldpos {
						newX := foldpos - (x - foldpos)
						newKey := fmt.Sprintf("%d,%d", newX, y)
						paper[newKey] = 1
						delete(paper, k)
					}
				}
			}
			// store new size of paper
			if cmd[0] == "y" {
				maxY = foldpos - 1
			} else {
				maxX = foldpos - 1
			}
			if part == "part 1" && f == 0 {
				fmt.Printf("Day 13 (part 1): There are %d marks on the paper\n", len(paper))
				// reset
				maxX = 0
				maxY = 0
				folds = []string{}
				makePaper(puzzle_data)
				break
			}
		}
	}
	fmt.Println("Day 13 (part 2): There code to activate the infrared thermal imaging camera:")
	printPaper(paper, maxX, maxY)
}
