package main

import (
	"fmt"
	"sort"
)

func day01(puzzle_data []string) {
	invoices := make([]int, len(puzzle_data))
	var A, B, C int
	for d, dataVal := range puzzle_data {
		invoices[d] = makeInt(dataVal)
		if A == B {
			for i := 0; i < d; i++ {
				if invoices[i]+invoices[d] == 2020 {
					A = invoices[i]
					B = invoices[d]
					break
				}
			}
		}
	}
	fmt.Printf("Day 01 (part 1): product of %d and %d is %d\n", A, B, A*B)
	A = 0
	B = 0
	sort.Ints(invoices)
	for i := 0; i < len(invoices)-2; i++ {
		for j := 1; j < len(invoices)-1; j++ {
			for k := len(invoices) - 1; k > 0; k-- {
				if invoices[i]+invoices[j]+invoices[k] == 2020 {
					A = invoices[i]
					B = invoices[j]
					C = invoices[k]
					break
				}
			}
			if A != B && B != C {
				break
			}
		}
		if A != B && B != C {
			break
		}
	}
	fmt.Printf("Day 01 (part 2): product of %d and %d and %d is %d\n", A, B, C, A*B*C)
}
