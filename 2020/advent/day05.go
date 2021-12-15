package main

import (
	"fmt"
	"sort"
	"strings"
)

func day05(puzzle_data []string) {
	fmt.Println("day05")
	max_seatid := 0
	seatids := []int{}
	for _, dataVal := range puzzle_data {
		min := 0
		max := 127
		row_template := dataVal[0:7]
		seat_template := dataVal[7:]
		for _, char := range strings.Split(row_template, "") {
			half := (min + max) / 2
			switch char {
			case "F":
				max = half
			case "B":
				min = half + 1
			}
		}
		row := 0
		if row_template[len(row_template)-1] == 'F' {
			row = min
		} else {
			row = max
		}
		min = 0
		max = 7
		for _, char := range strings.Split(seat_template, "") {
			half := (min + max) / 2
			switch char {
			case "L":
				max = half
			case "R":
				min = half + 1
			}
		}
		seat := 0
		if seat_template[len(seat_template)-1] == 'L' {
			seat = min
		} else {
			seat = max
		}
		seatid := row*8 + seat
		seatids = append(seatids, seatid)
		if seatid > max_seatid {
			max_seatid = seatid
		}
	}
	fmt.Printf("Day 05 (part 1): The highest seat id is %d\n", max_seatid)
	sort.Ints(seatids)
	for i := 0; i < len(seatids)-1; i++ {
		if seatids[i]+1 != seatids[i+1] {
			fmt.Printf("Day 05 (part 2): My seat id is %d\n", seatids[i]+1)
			break
		}
	}
}
