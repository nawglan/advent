package parse

import (
	"bufio"
	"os"
	"strconv"
)

func Bin2dec(bin string) int {
	mul := 1
	sum := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			sum += mul
		}
		mul *= 2
	}

	return sum
}

func MakeInt(value string) int {
	integer, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return integer
}


func ReadFile(fileName string) []string {
	var string_data []string
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string_data = append(string_data, scanner.Text())
	}
	if scanError := scanner.Err(); scanError != nil {
		panic(scanError)
	}

	return string_data
}

