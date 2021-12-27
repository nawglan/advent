package main

import (
	"fmt"
	"strings"
)

func traverse_part_one(g map[string][]string, step string, prev string) []string {
	if step == "end" {
		return []string{prev + step}
	}
	routes := []string{}
	for _, next := range g[step] {
		if strings.ToLower(next) == next && strings.Contains(prev, fmt.Sprintf(",%s,", next)) {
			continue
		} else {
			routes = append(routes, traverse_part_one(g, next, prev+step+",")...)
		}
	}
	return routes
}

func traverse_part_two(g map[string][]string, step string, prev string, seen int) []string {
	if step == "end" {
		return []string{prev + step}
	}
	routes := []string{}
	for _, next := range g[step] {
		if strings.ToLower(next) == next && strings.Contains(prev, ","+next+",") {
			if seen == 1 || next == "end" || next == "start" {
				continue
			} else {
				routes = append(routes, traverse_part_two(g, next, prev+step+",", 1)...)
			}
		} else {
			routes = append(routes, traverse_part_two(g, next, prev+step+",", seen)...)
		}
	}
	return routes
}

func day12(puzzle_data []string) {
	g := map[string][]string{}

	for _, dataVal := range puzzle_data {
		caves := strings.Split(dataVal, "-")
		g[caves[0]] = append(g[caves[0]], caves[1])
		g[caves[1]] = append(g[caves[1]], caves[0])
	}

	// part 1
	routes := traverse_part_one(g, "start", ",")

	fmt.Printf("Day 12 (part 1): There are %d paths through the caves.\n", len(routes))

	routes = traverse_part_two(g, "start", ",", 0)

	// part 2
	fmt.Printf("Day 12 (part 2): There are %d paths through the caves.\n", len(routes))
}
