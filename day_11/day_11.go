package main

import (
	"fmt"
	"os"
	"strings"
)

type Gallaxies struct {
	x []int
	y []int
}

func (g *Gallaxies) append(x, y int) {
	g.x = append(g.x, x)
	g.y = append(g.y, y)
}

func make_range[T any](length int, symbol T) []T {
	var r []T
	for i := 0; i < length; i++ {
		r = append(r, symbol)
	}
	return r
}

func str_make_range(length int, symbol string) string {
	var r string
	for i := 0; i < length; i++ {
		r += symbol
	}
	return r
}

func replace_at_index[T any](input []T, replace T, index, offset int) []T {
	return append(input[:index], append([]T{replace}, input[index+offset:]...)...)
}

func add_at_index(input string, replace string, index int) string {
	return input[:index] + replace + input[index:]
}

func expand_universe(universe []string, gallaxies Gallaxies) []string {
	range_x := make_range(len(universe), 1)
	range_y := make_range(len(universe), 1)
	for i := 0; i < len(gallaxies.x); i++ {
		// create binary maps of what row and column we have to expand
		range_x = replace_at_index(range_x, 0, gallaxies.x[i], 1)
		range_y = replace_at_index(range_y, 0, gallaxies.y[i], 1)
	}
	for i := len(range_x) - 1; i > 0; i-- {
		if range_x[i] == 1 {
			for j := range universe {
				// add one char to every line at this index
				one_char := strings.Repeat(".", 1_000_000)
				// one_char := str_make_range(1_000_000, ".")
				universe[j] = add_at_index(universe[j], one_char, i)
				// TODO for every galaxy bigger than i, add 1M
			}
		}
	}
	for i := len(range_x) - 1; i > 0; i-- {
		if range_y[i] == 1 {
			new_line := strings.Repeat(".", len(universe[0]))
			new_lines := []string{}
			for m := 0; m < 1_000_000; m++ {
				new_lines = append(new_lines, new_line)
			}
			// add one full line to list of str at this index
			universe = append(universe[:i], append(new_lines, universe[i:]...)...)
			// TODO for every galaxy bigger than i, add 1M
		}
	}
	return universe
}

func find_all_gallaxies(universe []string) Gallaxies {
	gallaxies := Gallaxies{}
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			if string(universe[y][x]) == "#" {
				gallaxies.append(x, y)
			}
		}
	}
	return gallaxies
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func get_sum_distances(universe []string, gallaxies Gallaxies) int {
	trailing := 0
	for i := 0; i < len(gallaxies.y); i++ {
		for j := i; j < len(gallaxies.x); j++ {
			trailing += abs(gallaxies.x[i]-gallaxies.x[j]) + abs(gallaxies.y[i]-gallaxies.y[j])
		}
	}
	return trailing
}

func run(universe []string) int {
	gallaxies := find_all_gallaxies(universe)
	universe = expand_universe(universe, gallaxies)
	fmt.Println("finished expanding")
	gallaxies = find_all_gallaxies(universe)
	fmt.Println("finished finding second time")
	distances := get_sum_distances(universe, gallaxies)
	return distances
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n")
	answer := run(fdata)
	fmt.Println(answer)
}
