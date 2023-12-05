package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	converted_data := strings.Fields(string(data))
	answer := run(converted_data)
    fmt.Println("\n", answer)
}

type int_list struct {
	list []int
}

func (list *int_list) contains(item int) bool {
	for _, v := range list.list {
		if item == v {
			return true
		}
	}
	return false
}

func (list *int_list) append(item int) []int {
    list.list = append(list.list, item)
    return list.list
}

func run(data []string) int {
	var result int
	is_reading_lottery := false
	var lottery int_list
	var trailing int

	for _, v := range data {
		int_v, err := strconv.Atoi(string(v))

		if err == nil {
			if is_reading_lottery {
				lottery.append(int_v)
			} else {
				if lottery.contains(int_v) {
					if trailing == 0 {
						trailing = 1
					} else {
						trailing <<= 1
					}
				}
			}

		} else if v == "|" {
			// Starting reading my numbers
			is_reading_lottery = false

		} else if v == "Card" {
			// New card
			result += trailing
			trailing = 0
			lottery.list = nil
			is_reading_lottery = true
		}
	}
	return result
}
