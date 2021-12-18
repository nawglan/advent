package main

import (
	"fmt"
	"strings"
)

func runPrg(prg [][]string) (int, bool) {
	atEnd := false

	ptr := 0
	acc := 0
	seen := map[int]bool{}

	for {
		if _, found := seen[ptr]; !found {
			seen[ptr] = true
			val := makeInt(prg[ptr][1])
			switch prg[ptr][0] {
			case "acc":
				acc += val
				ptr++
			case "jmp":
				ptr += val
			case "nop":
				ptr++
			}
		} else {
			break
		}
		if ptr >= len(prg) {
			atEnd = true
			break
		}
	}

	return acc, atEnd
}

func day08(puzzle_data []string) {

	prg := make([][]string, len(puzzle_data))
	for i, dataVal := range puzzle_data {
		prg[i] = strings.Fields(dataVal)
	}

	acc, _ := runPrg(prg)

	fmt.Printf("Day 08 (part 1): The value in acc is %d\n", acc)

	ok := false
	for i := 0; i < len(prg); i++ {
		if prg[i][0] == "jmp" {
			prg[i][0] = "nop"
			if acc, ok = runPrg(prg); ok {
				break
			} else {
				prg[i][0] = "jmp"
			}
		} else {
			if prg[i][0] == "nop" {
				prg[i][0] = "jmp"
				if acc, ok = runPrg(prg); ok {
					break
				} else {
					prg[i][0] = "nop"
				}
			}
		}
	}

	fmt.Printf("Day 08 (part 2): The value in acc is %d\n", acc)
}
