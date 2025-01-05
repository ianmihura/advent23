package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./sample_input.txt")
	fmt.Println(run(strings.Fields(string(data))))
}

func convert(data []byte) []string {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var converted []string
	for _, s := range strings.Fields(string(data)) {
		for k, v := range numbers {
			s = strings.Replace(s, k, v, -1)
		}
		converted = append(converted, s)
	}
	return converted
}

func run(data []string) int {
	var result int
	var current string
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, v := range data {
		for i, vv := range v {
			value, err := strconv.Atoi(string(vv))

			if err == nil {
				// Is a number
				current += strconv.Itoa(value)
			} else {
				for ii, n := range numbers {
					if strings.HasPrefix(v[i:], n) {
						current += fmt.Sprint(ii + 1)
					}
				}
			}
		}

		if current != "" {
			first_last_int := string(current[0]) + string(current[len(current)-1])
			converted_current, _ := strconv.Atoi(first_last_int)
			fmt.Println(converted_current)
			result += converted_current
			current = ""
		}
	}
	return result
}
