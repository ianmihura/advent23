package main

import (
	"fmt"
	"os"
	"strings"
)

// Remainder of a / b
// ex: 10 / 3 = 3r1
func get_div_remainder(a, b int) int {
	r := a
	for r >= b {
		r -= b
	}
	return r
}

func get_next_hash_value(data string) int {
	result := 0
	for d := range data {
		result += int(data[d])
		result *= 17
		result = get_div_remainder(result, 256)
	}
	return result
}

func run(data []string) int {
	value := 0
	for d := range data {
		value += get_next_hash_value(data[d])
		fmt.Println(data[d], value)
	}
	return value
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), ",")
	answer := run(fdata)
	fmt.Println(answer)
}
