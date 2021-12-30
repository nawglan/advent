package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

var (
	contains = map[string]map[string]int{}
	seen     = map[string]bool{}
)

func can_hold_bag_count(bag_color string) int {

	count := 0

	for k, v := range contains {
		if _, ok := seen[k]; !ok {
			if _, found := v[bag_color]; found {
				seen[k] = true
				count += 1 + can_hold_bag_count(k)
			}
			continue
		}
	}

	return count
}

func bag_holds_how_many(bag_color string) int {
	count := 1

	if v, ok := contains[bag_color]; ok {
		if len(v) > 0 {
			for b, c := range v {
				count += (c * bag_holds_how_many(b))
			}
		}
	}

	return count
}

func day07(puzzle_data []string) {
	for _, dataVal := range puzzle_data {
		data := strings.Split(dataVal, " contain ")
		left_bag_color := strings.Replace(data[0], " bags", "", 1)
		contains[left_bag_color] = map[string]int{}
		if strings.Contains(dataVal, "no other bags.") {
			continue
		}
		right_bags := strings.Split(data[1], ",")
		for _, right_bag := range right_bags {
			right_bag_parts := strings.Fields(right_bag)
			right_bag_count := p.MakeInt(right_bag_parts[0])
			right_bag_color := strings.Join(right_bag_parts[1:len(right_bag_parts)-1], " ")

			contains[left_bag_color][right_bag_color] = right_bag_count
		}
	}

	count := can_hold_bag_count("shiny gold")

	fmt.Printf("Day 07 (part 1): %d bag colors can contain a shiny gold bag.\n", count)

	count = bag_holds_how_many("shiny gold") - 1

	fmt.Printf("Day 07 (part 2): %d bags fit inside a shiny gold bag.\n", count)
}
