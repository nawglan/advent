package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	lookup = map[int]map[string]bool{}
	queue  = []int{}
)

func inQueue(check int) bool {
	for _, n := range queue {
		if n == check {
			return true
		}
	}

	return false
}

func isValid(check int) bool {
	if val, ok := lookup[check]; ok {
		for j := range val {
			vals := strings.Split(j, ",")
			if inQueue(makeInt(vals[0])) && inQueue(makeInt(vals[1])) {
				return true
			}
		}
	}

	return false
}

func addValue(val int) {
	queue = append(queue, val)

	if len(queue) > 1 {
		for _, k := range queue {
			if len(lookup[k+val]) == 0 {
				lookup[k+val] = make(map[string]bool, 0)
			}
			lookup[k+val][fmt.Sprintf("%d,%d", val, k)] = true
		}
	}
}

func day09(puzzle_data []string) {
	badValue := 0
	for i, dataVal := range puzzle_data {
		val := makeInt(dataVal)

		if i < 25 {
			addValue(val)
		} else {
			if isValid(val) {
				queue = queue[1:]
				addValue(val)
			} else {
				badValue = val
				break
			}
		}
	}

	list := []int{}

	sum := 0
	for i := len(puzzle_data)-1; i >= 0; i-- {
		val := makeInt(puzzle_data[i])
		list = append(list, val)
		if len(list) > 1 {
			sum = sumSlice(list)
			fmt.Printf("DEZ: sum (%d == %d), len is %d\n", badValue, sum, len(list))
			if sum == badValue {
				break
			} else {
				for sum > badValue && len(list) > 2 {
					list = list[1:]
					sum = sumSlice(list)
					fmt.Printf("DEZ: sum (%d == %d), len is %d\n", badValue, sum, len(list))
				}
			}
		}
	}
	sort.Ints(list)

	fmt.Printf("Day 09 (part 1) The first bad value is %d\n", badValue)
	fmt.Printf("Day 09 (part 2) The sum of %d and %d is %d\n", list[0], list[len(list)-1], list[0]+list[len(list)-1])
}
