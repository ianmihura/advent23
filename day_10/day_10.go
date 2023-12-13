package main

import (
	"fmt"
	"os"
	"strings"
)

// directions {x, y} == {j, i}
var TILE_DIRECTION = map[string][]int{
	"F": {1, 1},
	"7": {-1, 1},
	"L": {1, -1},
	"J": {-1, -1},
	"-": {1, 0},
	"|": {0, 1},
}

type Path struct {
	x []int
	y []int
}

func (p *Path) append(x, y int) {
	p.x = append(p.x, x)
	p.y = append(p.y, y)
}

func get_farthest_dist(p Path) int {
	return len(p.x) / 2
}

func find_starting_pos(chart []string) (int, int) {
	for y := 0; y < len(chart); y++ {
		for x := 0; x < len(chart[y]); x++ {
			if string(chart[y][x]) == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func get_current_tile(chart []string, p Path) (string, int, int) {
	p_x := p.x[len(p.x)-1]
	p_y := p.y[len(p.y)-1]
	return string(chart[p_y][p_x]), p_x, p_y
}

func get_previous_tile_pos(chart []string, p Path) (int, int) {
	p_x := p.x[len(p.x)-2]
	p_y := p.y[len(p.y)-2]
	return p_x, p_y
}

func get_next_tile_pos(chart []string, p Path) (int, int) {
	current_tile, cur_x, cur_y := get_current_tile(chart, p)
	prev_x, prev_y := get_previous_tile_pos(chart, p)

	tile_direction := TILE_DIRECTION[current_tile]
	if tile_direction[0] == 0 {
		if prev_y == cur_y+1 {
			// going: up
			return cur_x, cur_y - 1
		} else {
			// going: down
			return cur_x, cur_y + 1
		}
	} else if tile_direction[1] == 0 {
		if prev_x == cur_x+1 {
			// going: backwards
			return cur_x - 1, cur_y
		} else {
			// going: forward
			return cur_x + 1, cur_y
		}
	} else {
		if prev_x == cur_x {
			// going: up or down
			return cur_x + tile_direction[0], cur_y
		} else {
			// going: back or forward
			return cur_x, cur_y + tile_direction[1]
		}
	}
}

func run(chart []string) int {
	var x, y = find_starting_pos(chart)
	path := Path{
		x: []int{x},
		y: []int{y},
	}
	path.append(x+1, y)

	for {
		next_x, next_y := get_next_tile_pos(chart, path)

		if string(chart[next_y][next_x]) == "S" {
			break
		} else {
			path.append(next_x, next_y)
		}
	}
	fmt.Println("final path:", path)

	return get_farthest_dist(path)
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n")
	answer := run(fdata)
	fmt.Println(answer)
}
