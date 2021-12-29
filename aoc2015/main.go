package main

import (
	p "advent/util/parse"
	"fmt"
	"os"
)

func main() {
	filename := ""
	fmt.Println(filename)

	//filename = "data/day01.txt"
	filename = "data/day01.txt"
	puzzle_data := p.ReadFile(filename)
	day01(puzzle_data)
	fmt.Println("-----------------------")
	filename = os.Args[1]
	puzzle_data = p.ReadFile(filename)
	day02(puzzle_data)
	/*
		fmt.Println("-----------------------")
		filename = "data/day03.txt"
		puzzle_data = p.ReadFile(filename)
		day03(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day04.txt"
		puzzle_data = p.ReadFile(filename)
		day04(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day05.txt"
		puzzle_data = p.ReadFile(filename)
		day05(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day06.txt"
		puzzle_data = p.ReadFile(filename)
		day06(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day07.txt"
		puzzle_data = p.ReadFile(filename)
		day07(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day08.txt"
		puzzle_data = p.ReadFile(filename)
		day08(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day09.txt"
		puzzle_data = p.ReadFile(filename)
		day09(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day10.txt"
		puzzle_data = p.ReadFile(filename)
		day10(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day11.txt"
		puzzle_data = p.ReadFile(filename)
		day11(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day12.txt"
		puzzle_data = p.ReadFile(filename)
		day12(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day13.txt"
		puzzle_data = p.ReadFile(filename)
		day13(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day14.txt"
		puzzle_data = p.ReadFile(filename)
		day14(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day15.txt"
		puzzle_data = p.ReadFile(filename)
		day15(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day16.txt"
		puzzle_data = p.ReadFile(filename)
		day16(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day17.txt"
		puzzle_data = p.ReadFile(filename)
		day17(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day18.txt"
		puzzle_data = p.ReadFile(filename)
		day18(puzzle_data)
		fmt.Println("-----------------------")
		filename = os.Args[1]
		puzzle_data = p.ReadFile(filename)
		day19(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day20.txt"
		puzzle_data = p.ReadFile(filename)
		day20(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day21.txt"
		puzzle_data = p.ReadFile(filename)
		day21(puzzle_data)
		fmt.Println("-----------------------")
		filename = os.Args[1]
		puzzle_data = p.ReadFile(filename)
		day22(puzzle_data)
		fmt.Println("-----------------------")
		filename = os.Args[1]
		puzzle_data = p.ReadFile(filename)
		day23(puzzle_data)
		fmt.Println("-----------------------")
		filename = os.Args[1]
		puzzle_data = p.ReadFile(filename)
		day24(puzzle_data)
		fmt.Println("-----------------------")
		filename = "data/day25.txt"
		puzzle_data = p.ReadFile(filename)
		day25(puzzle_data)
	*/
}
