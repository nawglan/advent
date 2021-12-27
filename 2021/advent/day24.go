package main

import (
	"fmt"
	"math"
	"strings"
)

const DEBUG24 = false

type Program struct {
	registers map[string]int
	program   []string
}

func (p Program) execute(input modelNum) bool {
	inputIdx := 13
	p.registers = map[string]int{}
	p.registers["w"] = 0
	p.registers["x"] = 0
	p.registers["y"] = 0
	p.registers["z"] = 0

	if DEBUG24 {
		fmt.Printf("DEZ: Checking %s\n", input)
	}

	for _, line := range p.program {
		val := 0
		tokens := strings.Fields(line)
		switch tokens[0] {
		case "inp":
			val = input[inputIdx]
			if DEBUG24 {
				fmt.Printf("DEZ: %s = %d\n", tokens[1], val)
			}
			p.registers[tokens[1]] = val
			inputIdx--
		case "add":
			val = 0
			switch tokens[2] {
			case "w", "x", "y", "z":
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) + %s(%d) = ", tokens[1], tokens[1], p.registers[tokens[1]], tokens[2], p.registers[tokens[2]])
				}
				val = p.registers[tokens[2]]
			default:
				val = makeInt(tokens[2])
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) + %d = ", tokens[1], tokens[1], p.registers[tokens[1]], val)
				}
			}
			p.registers[tokens[1]] = p.registers[tokens[1]] + val
			if DEBUG24 {
				fmt.Printf("%d\n", p.registers[tokens[1]])
			}
		case "mul":
			val = 0
			switch tokens[2] {
			case "w", "x", "y", "z":
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) * %s(%d) = ", tokens[1], tokens[1], p.registers[tokens[1]], tokens[2], p.registers[tokens[2]])
				}
				val = p.registers[tokens[2]]
			default:
				val = makeInt(tokens[2])
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) * %d = ", tokens[1], tokens[1], p.registers[tokens[1]], val)
				}
			}
			p.registers[tokens[1]] = p.registers[tokens[1]] * val
			if DEBUG24 {
				fmt.Printf("%d\n", p.registers[tokens[1]])
			}
		case "div":
			val = 0
			switch tokens[2] {
			case "w", "x", "y", "z":
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) / %s(%d) = ", tokens[1], tokens[1], p.registers[tokens[1]], tokens[2], p.registers[tokens[2]])
				}
				val = p.registers[tokens[2]]
			default:
				val = makeInt(tokens[2])
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) / %d = ", tokens[1], tokens[1], p.registers[tokens[1]], val)
				}
			}
			p.registers[tokens[1]] = int(math.Floor(float64(p.registers[tokens[1]]) / float64(val)))
			if DEBUG24 {
				fmt.Printf("%d\n", p.registers[tokens[1]])
			}
		case "mod":
			val = 0
			switch tokens[2] {
			case "w", "x", "y", "z":
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) %% %s(%d) = ", tokens[1], tokens[1], p.registers[tokens[1]], tokens[2], p.registers[tokens[2]])
				}
				val = p.registers[tokens[2]]
			default:
				val = makeInt(tokens[2])
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) %% %d = ", tokens[1], tokens[1], p.registers[tokens[1]], val)
				}
			}
			p.registers[tokens[1]] = p.registers[tokens[1]] % val
			if DEBUG24 {
				fmt.Printf("%d\n", p.registers[tokens[1]])
			}
		case "eql":
			val = 0
			switch tokens[2] {
			case "w", "x", "y", "z":
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) == %s(%d) = ", tokens[1], tokens[1], p.registers[tokens[1]], tokens[2], p.registers[tokens[2]])
				}
				val = p.registers[tokens[2]]
			default:
				val = makeInt(tokens[2])
				if DEBUG24 {
					fmt.Printf("DEZ: %s = %s(%d) == %d = ", tokens[1], tokens[1], p.registers[tokens[1]], val)
				}
			}
			p.registers[tokens[1]] = 0
			if p.registers[tokens[1]] == val {
				p.registers[tokens[1]] = 1
			}
			if DEBUG24 {
				fmt.Printf("%d\n", p.registers[tokens[1]])
			}
		}
	}

	return p.registers["z"] == 0
}

type modelNum [14]int

func (n modelNum) nextSmallest() modelNum {
	var carry int
	var num modelNum

	num[0] = n[0] - 1
	if num[0] == 0 {
		num[0] = 9
		carry = 1
	}
	for idx := 1; idx < 14; idx++ {
		num[idx] = n[idx] - carry
		if num[idx] == 0 {
			num[idx] = 9
			carry = 1
		} else {
			carry = 0
		}
	}

	return num
}

func (n modelNum) nextLargest() modelNum {
	var carry int
	var num modelNum

	num[0] = n[0] + 1
	if num[0] == 10 {
		num[0] = 1
		carry = 1
	}
	for idx := 1; idx < 14; idx++ {
		num[idx] = n[idx] + carry
		if num[idx] == 10 {
			num[idx] = 1
			carry = 1
		} else {
			carry = 0
		}
	}

	return num
}

func (n modelNum) String() string {
	out := ""
	for i := 13; i >= 0; i-- {
		out += fmt.Sprintf("%d", n[i])
	}

	return out
}

func calcLargest(p Program, c chan string) {
	largestMN := modelNum{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

	for p.execute(largestMN) == false {
		largestMN = largestMN.nextSmallest()
		if fmt.Sprintf("%s", largestMN) == "99999999999999" {
			panic("wrapped")
		}
	}
	c <- fmt.Sprintf("Day 24 (part 1): Largest model number accepted is %s\n", largestMN)
}

func calcSmallest(p Program, c chan string) {
	smallestMN := modelNum{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for p.execute(smallestMN) == false {
		smallestMN = smallestMN.nextLargest()
		if fmt.Sprintf("%s", smallestMN) == "11111111111111" {
			panic("wrapped")
		}
	}
	c <- fmt.Sprintf("Day 24 (part 2): Smallest model number accepted is %s\n", smallestMN)
}

func day24(puzzle_data []string) {
	largest_chan := make(chan string)
	smallest_chan := make(chan string)

	go calcLargest(Program{program: puzzle_data}, largest_chan)
	go calcSmallest(Program{program: puzzle_data}, smallest_chan)

	fmt.Printf("%s%s", <-largest_chan, <-smallest_chan)
}
