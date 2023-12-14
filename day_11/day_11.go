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
			// add one char to every line at this index (hard)
			for j := range universe {
				universe[j] = add_at_index(universe[j], ".", i)
			}
		}
	}
	for i := len(range_x) - 1; i > 0; i-- {
		if range_y[i] == 1 {
			new_line := str_make_range(len(universe[0]), ".")
			universe = replace_at_index(universe, new_line, i, 0)
		}
	}
	return universe
}

func find_all_gallaxies(universe []string) Gallaxies {
	gallaxies := Gallaxies{}
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe); x++ {
			if string(universe[y][x]) == "#" {
				gallaxies.append(x, y)
			}
		}
	}
	return gallaxies
}

func run(universe []string) Gallaxies {
	gallaxies := find_all_gallaxies(universe)
	expand_universe(universe, gallaxies)
	return gallaxies
}

func main() {
	data, _ := os.ReadFile("./sample_input.txt")
	fdata := strings.Split(string(data), "\n")
	answer := run(fdata)
	fmt.Println(answer)
}
