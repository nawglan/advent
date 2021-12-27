package main

import (
	"fmt"
	"sort"
	"strings"
	p "advent/util/parse"
)

func day08(puzzle_data []string) {
	// part 1
	count := 0
	for _, dataLine := range puzzle_data {
		outputWires := strings.Split(dataLine, " ")[11:]

		for _, val := range outputWires {
			if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
				count++
			}
		}
	}
	fmt.Printf("Day 8 (part 1): There are %d 1, 4, 7, and 8s\n", count)

	/* part 2
	0 => a b c	 e f g
	1 =>		 c		 f
	2 => a	 c d e	 g
	3 => a	 c d	 f g
	4 =>	 b c d	 f
	5 => a b	 d	 f g
	6 => a b	 d e f g
	7 => a	 c		 f
	8 => a b c d e f g
	9 => a b c d	 f g
	*/

	sum := 0
	for _, dataLine := range puzzle_data {
		inputWires := strings.Split(dataLine, " ")[0:10]
		outputWires := strings.Split(dataLine, " ")[11:]

		digits := make(map[string]string)
		for {
			for _, input := range inputWires {
				chars := strings.Split(input, "")
				sort.Strings(chars)
				letters := strings.Join(chars, "")
				switch len(letters) {
				case 2: // digit 1 =>		 c		 f
					digits["1"] = letters
					digits[letters] = "1"
				case 3: // digit 7 => a	 c		 f
					digits["7"] = letters
					digits[letters] = "7"
				case 4: // digit 4
					digits["4"] = letters
					digits[letters] = "4"
				case 5: // digits 2, 3, 5
					_, foundSeven := digits["7"]
					_, foundEG := digits["EG"]
					if foundSeven && foundEG {
						if stringContainsAll(letters, digits["7"]) {
							digits[letters] = "3"
							digits["3"] = letters
						} else { // 2 or 5
							if stringContainsAll(letters, digits["EG"]) {
								digits["2"] = letters
								digits[letters] = "2"
							} else {
								digits["5"] = letters
								digits[letters] = "5"
							}
						}
					}
				case 6: // digits 0, 6, 9
					_, foundFour := digits["4"]
					_, foundOne := digits["1"]
					if foundFour && foundOne {
						if stringContainsAll(letters, digits["4"]) {
							digits["9"] = letters
							digits[letters] = "9"
						} else {
							if stringContainsAll(letters, digits["1"]) {
								digits["0"] = letters
								digits[letters] = "0"
							} else {
								digits["6"] = letters
								digits[letters] = "6"
							}
						}
					}
				case 7: // digit 8
					digits["8"] = letters
					digits[letters] = "8"
					_, foundFour := digits["4"]
					_, foundSeven := digits["7"]
					_, foundEG := digits["EG"]
					if foundFour && foundSeven && !foundEG {
						// if the letters are not a part of the number 4 or the number 7, then it is part of segments EG
						for _, char := range strings.Split(letters, "") {
							if strings.Index(digits["4"]+digits["7"], char) == -1 {
								digits["EG"] += char
							}
						}
					}
				}
			}
			if len(digits) == 21 {
				break
			}
		}
		outputNum := ""
		for _, output := range outputWires {
			chars := strings.Split(output, "")
			sort.Strings(chars)
			letters := strings.Join(chars, "")
			outputNum += digits[letters]
		}
		sum += p.MakeInt(outputNum)
	}
	fmt.Printf("Day 8 (part 2): The sum of all the output numbers is %d\n", sum)
}
