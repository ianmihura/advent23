package main

import (
	"fmt"
	"os"
	"strings"
)

func is_equal_col(shape []string, a, b int) bool {
	for i := 0; i < len(shape); i++ {
		if shape[i][a] != shape[i][b] {
			return false
		}
	}
	return true
}

// cols to left
// 100 times rows above
func find_reflection(s []string) int {
	shape := s[:len(s)-1]

	// seek horizontal lines
	for i := 1; i < len(shape); i++ {
		is_row := false
		for j := i; j > 0; j-- {
			// odd number secuence: j + 2n+1
			n_index := j + 2*(i-j) + 1
			if n_index > len(shape) {
				continue
			}
			// if keeps being false, this is a reflection line
			// fmt.Println("horizontal:", i, is_row, shape[j-1], shape[n_index-1])
			is_row = is_row || (shape[j-1] != shape[n_index-1])
			if is_row {
				break
			}
		}
		// fmt.Println()
		if !is_row {
			// fmt.Println("horizontal:", i)
			return i * 100
		}
	}

	// seek vertical
	for i := 1; i < len(shape[0]); i++ {
		is_col := false
		for j := i; j > 0; j-- {
			// odd number secuence; j + 2n+1
			n_index := j + 2*(i-j) + 1
			if n_index > len(shape[0]) {
				continue
			}
			// if keeps being false, this is a reflection line
			// fmt.Println("vertical:", i, is_col, j, n_index)
			is_col = is_col || !is_equal_col(shape, j-1, n_index-1)
			if is_col {
				break
			}
		}
		// fmt.Println()
		if !is_col {
			// fmt.Println("vertical:", i)
			return i
		}
	}

	fmt.Println(shape)
	return 0
}

func run(data []string) int {
	trailing := 0
	for d := range data {
		trailing += find_reflection(strings.Split(data[d], "\n"))
	}

	return trailing
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n\n")
	answer := run(fdata)
	fmt.Println(answer)
}

// 23991 too low
