package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    data, _ := os.ReadFile("./sample_input.txt")
	converted_data := strings.Fields(string(data))
	// answer := run(converted_data)
	answer := run_2(converted_data)
    fmt.Println("\n", answer)
}

func run(data []string) int {
	var trailing_sum int
	for i, data_line := range data {
		for j := 0; j < len(data_line); j++ {
			data_char := data_line[j]
			if is_int(data_char) {
				current_int, jn := get_full_number(data_line, j)
				if has_adjacent(data, i, j, jn) {
					trailing_sum += current_int
				}
				j = jn
			}
		}
	}
	return trailing_sum
}

func run_2(data []string) int {
	var trailing_sum int
	for i, data_line := range data {
		for j := 0; j < len(data_line); j++ {
			data_char := data_line[j]
			if string(data_char) == "*" {
				adjacent_product := get_adjacent_product(data, i, j)
				trailing_sum += adjacent_product
			}
		}
	}
	return trailing_sum
}

func is_int(c interface{}) bool {
	switch v := c.(type) {
	case byte:
		_, err := strconv.Atoi(string(v))
		return err == nil
	case rune:
		_, err := strconv.Atoi(string(v))
		return err == nil
	default:
		return false
	}
}

// Given a string and a starting index
// returns the full number and the ending index+1 in that string
func get_full_number(data string, i_0 int) (int, int) {
	i_n := i_0
	for len(data) > i_n && is_int(data[i_n]) {
		i_n++
	}
	full_number, _ := strconv.Atoi(data[i_0:i_n])

	// i_n is the index after the last digit
	return full_number, i_n
}

// Given a string and an index of any of its digits
// returns the full number
func get_full_number_backwards(data string, index int) int {
	// search for i_0
	i_0 := index
	for i_0 > 0 && is_int(data[i_0]) {
		i_0--
	}
	if string(data[i_0]) == "." || string(data[i_0]) == "*" {
		i_0++
	}
	full_number, _ := get_full_number(data, i_0)

	return full_number
}

func has_adjacent(data []string, line_i, j_0, j_n int) bool {
	if line_i == 0 {
		// don't check for negative index
		line_i = 1
	}
	
	if j_0 == 0 {
		// don't check for negative index
		j_0 = 1
	}

	for i := line_i-1; i <= line_i+1 && i < len(data); i++ {
		// loops at most 3 times

		for j := j_0-1; j <= j_n && j < len(data[line_i]); j++ {
			// loops at most j_n-j_0 times
			// that is, len(current_num)
			if string(data[i][j]) != "." && (is_int(data[i][j]) == false) {
				return true
			}
		}
	}
	return false
}

func get_adjacent_product(data []string, line_i, j_0 int) int {
	if line_i == 0 {
		// don't check for negative index
		line_i = 1
	}

	if j_0 == 0 {
		// don't check for negative index
		j_0 = 1
	}

	var factor_1, factor_2 int
	for i := line_i-1; i <= line_i+1 && i < len(data); i++ {
		// loops at most 3 times

		for j := j_0-1; j <= j_0+1 && j < len(data[line_i]); j++ {
			// loops at most 3 times
			if is_int(data[i][j]) {
				if factor_1 == 0 {
					factor_1 = get_full_number_backwards(data[i], j)
				} else if factor_2 == 0 {
					factor_2 = get_full_number_backwards(data[i], j)
					if factor_1 == factor_2 {
						factor_2 = 0
					}
				}
			}
		}
	}
	return factor_1 * factor_2
}
