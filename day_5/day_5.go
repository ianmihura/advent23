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
	seeds := get_seeds(fdata[0])
	mappings := get_mappings(fdata[2:])
	answer := run(seeds, mappings)
	fmt.Println(answer)
}

func run(seeds []int, mappings [][][]int) int {
	indexes := seeds
	for i := 0; i < 7; i++ {
		indexes = get_next_indexes(indexes, mappings[i])
	}
	return min(indexes...)
}

func min(list ...int) int {
	if len(list) == 1 {
		return list[0]
	} else if list[0] > list[1] {
		return min(list[1:]...)
	} else {
		return min(append(list[2:], list[0])...)
	}
}

func get_next_indexes(indexes []int, origin_map [][]int) []int {
	//
	// 	 50 98 2
	//           from 98 to 98+2 -1
	//     values map 50 to 50+2 -1
	//        => values between 98 and 98+2-1 map to 50 and 50+2-1
	//   52 50 48
	//           from 50 to 50+48 -1
	//     values map 52 to 52+48 -1
	//
	// 0: destination_range_start
	// 1: source_range_start
	// 2: range_length
	/*
		seed  soil
		0     0
		1     1

		14    14
		...   ...
		48    48
		49    49
		50    52
		51    53
		52    54
		53    55
		54    56
		55    57
		...   ...
		78    81
		...   ...
		96    98
		97    99
		98    50
		99    51
	*/
	var next_indexes []int
	for idx := range indexes {
		new_index := -1
		for i := 0; i < len(origin_map); i++ {
			destination_range_start := origin_map[i][0]
			source_range_start := origin_map[i][1]
			range_length := origin_map[i][2]
			if indexes[idx] > source_range_start && indexes[idx] < source_range_start+range_length {
				new_index = destination_range_start + indexes[idx] - source_range_start
				break
			}
		}
		if new_index > 0 {
			next_indexes = append(next_indexes, new_index)
		} else {
			next_indexes = append(next_indexes, indexes[idx])
		}
	}
	fmt.Println(next_indexes)
	return next_indexes
}

func get_seeds(fdata string) []int {
	var result []int
	data := strings.Split(fdata, " ")[1:]

	for i := 0; i < len(data); i++ {
		c, err := strconv.Atoi(data[i])
		if err == nil {
			result = append(result, c)
		}
	}
	return result
}

func get_mappings(fdata []string) [][][]int {
	var result [][][]int
	map_group := -1 // outer groupping (max: 7)
	map_line := 0

	for i := 0; i < len(fdata); i++ {
		// i: line of mapping

		line_data := strings.Split(fdata[i], " ")
		if len(line_data) == 3 {
			// these are instructions

			result[map_group] = append(result[map_group], []int{})
			for j := 0; j < len(line_data); j++ {
				// j: int in mapping line
				c, _ := strconv.Atoi(string(line_data[j]))
				result[map_group][map_line] = append(result[map_group][map_line], c)
			}
			map_line++

		} else if len(line_data) == 2 {
			// new group
			map_group++
			result = append(result, [][]int{})
			map_line = 0
		}
	}
	return result
}
