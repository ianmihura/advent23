package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n")
	answer := run(fdata)
	fmt.Println(answer)

}

func run(fdata []string) int {
	var trailing int
	for _, v := range fdata {
		var current_seq [][]int
		current_line := convert_line(strings.Split(v, " "))
		current_seq = append(current_seq, current_line)
		var step int
		for {
			differences := get_differences(current_seq[step])
			current_seq = append(current_seq, differences)
			if is_final(differences) {
				break
			} else {
				step++
			}
		}
		trailing += get_next_number_2(current_seq)
	}
	return trailing
}

func get_next_number(current_seq [][]int) int {
	var next_numbers []int
	for i := len(current_seq) - 1; i >= 0; i-- {
		next_numbers = append(next_numbers, current_seq[i][len(current_seq[i])-1])
	}
	fmt.Println(next_numbers)
	return reduce_add(next_numbers)
}

func get_next_number_2(current_seq [][]int) int {
	var next_numbers []int
	for i := 0; i < len(current_seq); i++ {
		next_numbers = append(next_numbers, current_seq[i][0])
	}
	fmt.Println(next_numbers)
	return reduce_add_sub(next_numbers)
}

func reduce_add(list []int) int {
	var value int
	for i := range list {
		value += list[i]
	}
	return value
}

func reduce_add_sub(list []int) int {
	var value int
	for i := range list {
		if i%2 == 0 {
			value += list[i]
		} else {
			value -= list[i]
		}
	}
	return value
}

func is_final(line []int) bool {
	for k := range line {
		if line[k] != 0 {
			return false
		}
	}
	return true
}

func get_differences(line []int) []int {
	var differences []int
	for k := range line {
		if k != 0 {
			differences = append(differences, line[k]-line[k-1])
		}
	}
	return differences
}

func convert_line(line []string) []int {
	var fline []int
	for k := range line {
		line_k, _ := strconv.Atoi(line[k])
		fline = append(fline, line_k)
	}
	return fline
}
