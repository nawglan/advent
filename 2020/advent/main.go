package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	gridWidth  = 0
	gridLength = 0
)

func dumpGrid(grid []int) {
	for i := range grid {
		x := i % gridWidth
		y := i / gridWidth
		if y == gridLength-1 {
			break
		}
		if grid[i] == 0 {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
		if x == gridWidth-1 {
			fmt.Printf("\n")
		}
	}
}

func coords2pos(x int, y int) int {
	return x + y*gridWidth
}

func makeIntGrid(puzzle_data []string, width int, length int) []int {

	gridWidth = width
	gridLength = length

	grid := make([]int, width*length)

	for y, dataVal := range puzzle_data {
		for x, dataPoint := range dataVal {
			grid[coords2pos(x, y)] = makeInt(fmt.Sprintf("%c", dataPoint))
		}
	}

	return grid
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) []string {
	var string_data []string
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string_data = append(string_data, scanner.Text())
	}
	check(scanner.Err())

	return string_data
}

func stringContainsAll(input string, check string) bool {
	for _, checkLetter := range check {
		if !strings.ContainsRune(input, checkLetter) {
			return false
		}
	}
	return true
}

func makeInt(value string) int {
	integer, err := strconv.Atoi(value)
	check(err)
	return integer
}

func sumSlice(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return sum
}

func mulSlice(input []int) int {
	product := 1
	for _, val := range input {
		product *= val
	}
	return product
}

func neighbors(pos int) []int {
	posX := pos % gridWidth
	posY := pos / gridWidth
	neighborList := make([]int, 0)

	// north
	if posY > 0 {
		nY := posY - 1
		nX := posX
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// east
	if posX < gridWidth-1 {
		nY := posY
		nX := posX + 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// west
	if posX > 0 {
		nY := posY
		nX := posX - 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// south
	if posY < gridLength-1 {
		nY := posY + 1
		nX := posX
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	return neighborList
}

func neighborsDiag(pos int) []int {
	posX := pos % gridWidth
	posY := pos / gridWidth
	neighborList := make([]int, 0)

	// north-east
	if posX < gridWidth-1 && posY > 0 {
		nY := posY - 1
		nX := posX + 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// south-east
	if posX < gridWidth-1 && posY < gridLength-1 {
		nY := posY + 1
		nX := posX + 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// north-west
	if posX > 0 && posY > 0 {
		nY := posY - 1
		nX := posX - 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	// south-west
	if posX > 0 && posY < gridLength-1 {
		nY := posY + 1
		nX := posX - 1
		neighborList = append(neighborList, coords2pos(nX, nY))
	}

	return neighborList
}

func main() {
	filename := ""
	fmt.Println(filename)

	filename = "data/day01.txt"
	puzzle_data := readFile(filename)
	day01(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day02.txt"
	puzzle_data = readFile(filename)
	day02(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day03.txt"
	puzzle_data = readFile(filename)
	day03(puzzle_data)
	fmt.Println("-----------------------")
	filename = os.Args[1]
	//filename = "data/day04.txt"
	puzzle_data = readFile(filename)
	day04(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day05.txt"
	puzzle_data = readFile(filename)
	day05(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day06.txt"
	puzzle_data = readFile(filename)
	day06(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day07.txt"
	puzzle_data = readFile(filename)
	day07(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day08.txt"
	puzzle_data = readFile(filename)
	day08(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day09.txt"
	puzzle_data = readFile(filename)
	day09(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day10.txt"
	puzzle_data = readFile(filename)
	day10(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day11.txt"
	puzzle_data = readFile(filename)
	day11(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day12.txt"
	puzzle_data = readFile(filename)
	day12(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day13.txt"
	puzzle_data = readFile(filename)
	day13(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day14.txt"
	puzzle_data = readFile(filename)
	day14(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day15.txt"
	puzzle_data = readFile(filename)
	day15(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day16.txt"
	puzzle_data = readFile(filename)
	day16(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day17.txt"
	puzzle_data = readFile(filename)
	day17(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day18.txt"
	puzzle_data = readFile(filename)
	day18(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day19.txt"
	puzzle_data = readFile(filename)
	day19(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day20.txt"
	puzzle_data = readFile(filename)
	day20(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day21.txt"
	puzzle_data = readFile(filename)
	day21(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day22.txt"
	puzzle_data = readFile(filename)
	day22(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day23.txt"
	puzzle_data = readFile(filename)
	day23(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day24.txt"
	puzzle_data = readFile(filename)
	day24(puzzle_data)
	fmt.Println("-----------------------")
	filename = "data/day25.txt"
	puzzle_data = readFile(filename)
	day25(puzzle_data)
}
