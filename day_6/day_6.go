package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	time_data := strings.Fields(strings.Split(string(data), "\n")[0])[1:]
	dist_data := strings.Fields(strings.Split(string(data), "\n")[1])[1:]
	answer := run(time_data, dist_data)
	fmt.Println(answer)
}

func run(time_data, dist_data []string) int {
	var race_duration, record_dist string
	for i := 0; i < len(time_data); i++ {
		race_duration += time_data[i]
		record_dist += dist_data[i]
	}
	int_race_duration, _ := strconv.Atoi(race_duration)
	int_record_dist, _ := strconv.Atoi(record_dist)
	return get_best_for_race(int_race_duration, int_record_dist)
}

func get_best_for_race(race_duration, record_dist int) int {
	best := 0

	for i := 0; i < race_duration; i++ {
		if i*(race_duration-i) > record_dist {
			best++
		}
	}
	return best
}
