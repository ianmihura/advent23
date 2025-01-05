package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func reduce[T comparable](f func(a T, v T) T, input []T, initial T) T {
	trailing := initial
	for i := range input {
		trailing = f(trailing, input[i])
	}
	return trailing
}

// input example: ["?#?..#..??", "2,1,1"]
func calc_posible(line []string) int {
	mask := line[0]
	vector := line[1]
	fmt.Println(mask, vector)
	return 0
}

func run(data []string) int {
	var questions []int
	for _, d := range data {
		q := 0.0
		calc_posible(strings.Split(d, " "))
		for i := range d {
			if string(d[i]) == "?" {
				q++
			}
		}
		questions = append(questions, int(math.Pow(2.0, q)))
	}
	return reduce(func(a, v int) int {
		return a + v
	}, questions, 0)
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n")
	answer := run(fdata)
	fmt.Println(answer)
}
