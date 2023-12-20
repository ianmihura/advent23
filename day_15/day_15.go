package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_hash_value(data string) int {
	result := 0
	for d := range data {
		result = ((result + int(data[d])) * 17) % 256
	}
	return result
}

func get_lens_index(box [][]string, lens_name string) int {
	for l, lens := range box {
		if lens[0] == lens_name {
			return l
		}
	}
	return -1
}

func run(data []string) int {
	boxes := make([][][]string, 256)
	for _, d := range data {

		if d[len(d)-1] == '-' {
			lens_name := strings.Split(d, "-")[0]
			box := get_hash_value(lens_name)

			lens_index := get_lens_index(boxes[box], lens_name)
			if lens_index > -1 {
				boxes[box] = append(boxes[box][:lens_index], boxes[box][lens_index+1:]...)
			}

		} else {
			lens_name := strings.Split(d, "=")[0]
			lens_power := strings.Split(d, "=")[1]
			box := get_hash_value(lens_name)

			lens_index := get_lens_index(boxes[box], lens_name)
			if lens_index > -1 {
				boxes[box][lens_index][1] = lens_power
			} else {
				boxes[box] = append(boxes[box], []string{lens_name, lens_power})
			}
		}
	}

	value := 0
	for b, box := range boxes {
		for l, lens := range box {
			lens_power, _ := strconv.Atoi(lens[1])
			value += (b + 1) * (l + 1) * lens_power
		}
	}

	return value
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := strings.Split(string(data), ",")
	answer := run(fdata)
	fmt.Println(answer)
}
