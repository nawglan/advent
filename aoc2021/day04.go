package main

import (
	p "advent/util/parse"
	"fmt"
	"strings"
)

type cardCell struct {
	value int
	found bool
}

type bingoCard struct {
	cells [5][5]cardCell
	found bool
}

func day04(puzzle_data []string) {
	// first line is numbers to draw
	bingoCards := make([]bingoCard, (len(puzzle_data)-1)/6)

	cardNum := 0
	for i := 2; i < len(puzzle_data); i += 6 {
		var newCard bingoCard
		for rowNum := range newCard.cells {
			for colNum, val := range strings.Fields(puzzle_data[i+rowNum]) {
				newCard.cells[rowNum][colNum].value = p.MakeInt(val)
			}
		}
		bingoCards[cardNum] = newCard
		cardNum++
	}

	var messages []string
	part1 := true
	for {
		found := false
		for _, number := range strings.Split(puzzle_data[0], ",") {
			num := p.MakeInt(number)
			for c := range bingoCards {
				if !bingoCards[c].found {
					// mark card`
					for y := range bingoCards[c].cells {
						for x := range bingoCards[c].cells[y] {
							if bingoCards[c].cells[x][y].value == num {
								bingoCards[c].cells[x][y].found = true
							}
						}
					}
					// check row`
					for i := 0; i < 5; i++ {
						bingoCards[c].found = bingoCards[c].cells[i][0].found &&
							bingoCards[c].cells[i][1].found &&
							bingoCards[c].cells[i][2].found &&
							bingoCards[c].cells[i][3].found &&
							bingoCards[c].cells[i][4].found
						// check col`
						if !bingoCards[c].found {
							bingoCards[c].found = bingoCards[c].cells[0][i].found &&
								bingoCards[c].cells[1][i].found &&
								bingoCards[c].cells[2][i].found &&
								bingoCards[c].cells[3][i].found &&
								bingoCards[c].cells[4][i].found
						}
						if bingoCards[c].found {
							break
						}
					}
					if bingoCards[c].found {
						sum := 0
						for _, row := range bingoCards[c].cells {
							for _, cell := range row {
								if !cell.found {
									sum += cell.value
								}
							}
						}
						messages = append(messages, fmt.Sprintf("Sum = %d, Last Number = %s, product = %d\n", sum, number, sum*p.MakeInt(number)))
						found = true
						if part1 {
							break
						}
					}
				}
			}
			if found && part1 {
				break
			}
		}
		if part1 {
			fmt.Print("Day 4 (part 1):" + messages[len(messages)-1])
			part1 = !part1
		} else {
			fmt.Print("Day 4 (part 2):" + messages[len(messages)-1])
			break
		}
	}
}
