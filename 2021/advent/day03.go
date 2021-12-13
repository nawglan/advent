package main

import (
	"fmt"
	"strconv"
)

func day03(puzzle_data []string) {
	// part 1
	var gammaRateBits string
	var epsilonRateBits string

	half := len(puzzle_data) / 2
	for colNum := range puzzle_data[0] {
		onesCount := 0
		for _, val := range puzzle_data {
			if val[colNum] == '1' {
				onesCount++
			}
		}
		if onesCount > half {
			gammaRateBits += "1"
			epsilonRateBits += "0"
		} else {
			gammaRateBits += "0"
			epsilonRateBits += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateBits, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBits, 2, 64)

	fmt.Printf("Day 3 (part 1): Gamma Rate = %d, Epsilon Rate = %d, Power Consumption = %d\n", gammaRate, epsilonRate, gammaRate*epsilonRate)

	// part 2
	// oxygen
	puzzle_data_copy := puzzle_data[:]
	for j := range puzzle_data_copy[0] {
		var ones []string
		var zeros []string
		if len(puzzle_data_copy) == 1 {
			break
		}
		for i := range puzzle_data_copy {
			if puzzle_data_copy[i][j] == '1' {
				ones = append(ones, puzzle_data_copy[i])
			} else {
				zeros = append(zeros, puzzle_data_copy[i])
			}
		}
		if len(ones) > len(zeros) || len(ones) == len(zeros) {
			puzzle_data_copy = ones
		} else {
			puzzle_data_copy = zeros
		}
	}
	oxygenRating, _ := strconv.ParseInt(puzzle_data_copy[0], 2, 64)

	// CO2
	puzzle_data_copy = puzzle_data[:]
	for j := range puzzle_data_copy[0] {
		var ones []string
		var zeros []string
		if len(puzzle_data_copy) == 1 {
			break
		}
		for i := range puzzle_data_copy {
			if puzzle_data_copy[i][j] == '1' {
				ones = append(ones, puzzle_data_copy[i])
			} else {
				zeros = append(zeros, puzzle_data_copy[i])
			}
		}
		if len(ones) < len(zeros) {
			puzzle_data_copy = ones
		} else {
			puzzle_data_copy = zeros
		}
	}
	co2Rating, _ := strconv.ParseInt(puzzle_data_copy[0], 2, 64)

	fmt.Printf("Day 3 (part 2): Oxygen Generator Rating = %d, CO2 Scrubber Rating = %d, Life Support Rating = %d\n", oxygenRating, co2Rating, oxygenRating*co2Rating)
}
