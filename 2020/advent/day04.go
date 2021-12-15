package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func day04(puzzle_data []string) {
	fields := []string{
		"byr:",
		"iyr:",
		"eyr:",
		"hgt:",
		"hcl:",
		"ecl:",
		"pid:",
	}

	seen := map[string]bool{}
	part1_valid := 0
	part2_valid := 0
	valid_fields := 0
	for _, dataVal := range puzzle_data {
		if len(dataVal) > 0 {
			for _, data_field := range strings.Fields(dataVal) {
				for _, field := range fields {
					if strings.Contains(data_field, field) {
						// part 1
						seen[field] = true

						// part 2 validate the field
						data := strings.Replace(data_field, field, "", 1)

						switch field {
						case "byr:":
							if len(data) == 4 {
								if t, err := strconv.Atoi(data); err == nil {
									if t >= 1920 && t <= 2002 {
										valid_fields++
									}
								}
							}
						case "iyr:":
							if len(data) == 4 {
								if t, err := strconv.Atoi(data); err == nil {
									if t >= 2010 && t <= 2020 {
										valid_fields++
									}
								}
							}
						case "eyr:":
							if len(data) == 4 {
								if t, err := strconv.Atoi(data); err == nil {
									if t >= 2020 && t <= 2030 {
										valid_fields++
									}
								}
							}
						case "hgt:":
							if strings.Contains(data, "cm") {
								data = strings.Replace(data, "cm", "", 1)
								height := makeInt(data)
								if height >= 150 && height <= 193 {
									valid_fields++
								}
							} else {
								if strings.Contains(data, "in") {
									data = strings.Replace(data, "in", "", 1)
									height := makeInt(data)
									if height >= 59 && height <= 76 {
										valid_fields++
									}
								}
							}
						case "hcl:":
							if data[0:1] == "#" && len(data) == 7 {
								data = strings.ToLower(data)
								if _, err := hex.DecodeString(data[1:]); err == nil {
									valid_fields++
								}
							}
						case "ecl:":
							switch data {
							case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
								valid_fields++
							}
						case "pid:":
							if len(data) == 9 {
								if _, err := strconv.Atoi(data); err == nil {
									valid_fields++
								}
							}
						}
					}
				}
			}
		} else {
			if len(seen) == len(fields) {
				part1_valid++
				if valid_fields == len(fields) {
					part2_valid++
				}
			}
			// reset
			for _, field := range fields {
				delete(seen, field)
			}
			valid_fields = 0
		}
	}
	// was the last record valid?
	if len(seen) == len(fields) {
		part1_valid++
		if valid_fields == len(fields) {
			part2_valid++
		}
	}
	fmt.Printf("Day 04 (part 1): There were %d valid passports.\n", part1_valid)
	fmt.Printf("Day 04 (part 2): There were %d valid passports.\n", part2_valid)
}
