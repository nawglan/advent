package main

import (
	"fmt"
	"sort"
)

type person struct {
	name       string
	preference []string
	happiness  map[string]int
}

func day13(puzzle_data []string) {

	people := map[string]*person{}

	for _, line := range puzzle_data {
		who := ""
		what := ""
		target := ""
		value := 0

		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &who, &what, &value, &target)
		target = target[:len(target)-1] // strip off the period
		if what == "lose" {
			value *= -1
		}
		if val, ok := people[who]; ok {
			val.preference = append(val.preference, target)
			val.happiness[target] = value
			sort.SliceStable(val.preference, func(i, j int) bool { return val.happiness[val.preference[i]] > val.happiness[val.preference[j]] })
		} else {
			people[who] = &person{name: who, preference: []string{target}, happiness: map[string]int{target: value}}
		}
	}
	for _, p := range people {
		fmt.Printf("DEZ: %+v\n", *p)
	}
	fmt.Printf("Day 13 (part 1): \n")
	fmt.Printf("Day 13 (part 2): \n")
}
