package main

import (
	"fmt"
	"sort"
)

func day10(puzzle_data []string) {
	errorValues := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	expectedValues := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	autocompleteValues := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	// part 1
	syntaxScore := 0
	var autocompleteScores []int
	for _, dataVal := range puzzle_data {
		var stack []rune
		for _, dataRune := range dataVal {
			// opening value, push onto stack
			if _, ok := expectedValues[dataRune]; ok {
				stack = append(stack, dataRune)
			} else { // closing value, check if it matches the top of the stack
				check := stack[len(stack)-1]
				stack = stack[:len(stack)-1] // pop off the top
				if _, ok := expectedValues[check]; ok && dataRune != expectedValues[check] {
					if value, ok := errorValues[dataRune]; ok {
						syntaxScore += value
						break
					}
				}
			}
		}
		if len(stack) > 0 {
			autocompleteScore := 0
			for len(stack) > 0 {
				stackVal := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				autocompleteScore *= 5
				need, ok := expectedValues[stackVal]
				if ok {
					score, found := autocompleteValues[need]
					if found {
						autocompleteScore += score
					}
				}
			}
			autocompleteScores = append(autocompleteScores, autocompleteScore)
		}
	}
	sort.Ints(autocompleteScores)

	fmt.Printf("Day 10 (part 1): The syntax error score is %d\n", syntaxScore)
	fmt.Printf("Day 10 (part 2): The autocomplete score is %d\n", autocompleteScores[len(autocompleteScores)/2])
}
