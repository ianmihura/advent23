package main2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func experiment() {
	messages := make(chan int, 1_000_000_000)
	done := make(chan bool, 1)

	go func(d chan bool) {
		for i := 0; i < 100_000_000; i++ {
			messages <- i
		}
		d <- true
	}(done)

	<-done

	fmt.Println(len(messages))
	fmt.Println(<-messages)
	fmt.Println(len(messages))
	fmt.Println(<-messages)
	fmt.Println(len(messages))
}

func main2() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), "\n")
	seeds := get_seeds(fdata[0])
	mappings := get_mappings(fdata[2:])
	answer := run(seeds, mappings)
	fmt.Println(answer)
}

func run(seeds []int, mappings [][][]int) int {
	M := 100_000_000
	indexes := make(chan int, 5*M)
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(i + 1)

		go get_next_indexes(seeds[i*M:(i+1)*M], mappings[0], indexes, wg, i)
	}
	wg.Wait()
	return <-indexes
}

// indexes := seeds
// for i := 0; i < 7; i++ {
// 	indexes = get_next_indexes(indexes, mappings[i])
// }
// return min(indexes...)

func get_next_indexes(seeds []int, origin_map [][]int, indexes chan int, wg sync.WaitGroup, order int) {
	fmt.Println("starting", len(seeds))
	defer wg.Done()
	for idx := range seeds {
		// fmt.Println(len(indexes), order)
		new_index := -1
		for i := 0; i < len(origin_map); i++ {
			destination_range_start := origin_map[i][0]
			source_range_start := origin_map[i][1]
			range_length := origin_map[i][2]
			if seeds[idx] > source_range_start && seeds[idx] < source_range_start+range_length {
				new_index = destination_range_start + seeds[idx] - source_range_start
				break
			}
		}
		if new_index > 0 {
			indexes <- new_index
		} else {
			indexes <- seeds[idx]
		}
	}
	fmt.Println("done", order)
}

func get_seeds(fdata string) []int {
	result := make([]int, 500_000_000)
	data := strings.Split(fdata, " ")[1:]
	current_range_start := -1

	for i := 0; i < len(data); i++ {
		c, _ := strconv.Atoi(data[i])
		if i%2 == 0 {
			current_range_start = c
		} else {
			result = append(result, new_range(current_range_start, c)...)
		}
	}
	return result
}

func new_range(start, end int) []int {
	var r []int
	for i := start; i < end; i++ {
		r = append(r, i)
	}
	return r
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

func min(list ...int) int {
	if len(list) == 1 {
		return list[0]
	} else if list[0] > list[1] {
		return min(list[1:]...)
	} else {
		return min(append(list[2:], list[0])...)
	}
}
