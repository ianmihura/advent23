package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    data, _ := os.ReadFile("./sample_input.txt")
    fmt.Println(run_2(strings.Fields(string(data))))
}

func run_1(data []string) (int) {
	var current_game, trailing_sum int
	for k, v := range data {
		current_int, _ := strconv.Atoi(v)

		if v == "Game" {
			trailing_sum += current_game
			current_game, _ = strconv.Atoi(data[k+1][0:len(data[k+1])-1])
			
		} else if current_int > 0 {
			is_valid_red := strings.HasPrefix(data[k+1], "red") && current_int <= 12
			is_valid_green := strings.HasPrefix(data[k+1], "green") && current_int <= 13
			is_valid_blue := strings.HasPrefix(data[k+1], "blue") && current_int <= 14
			if !(is_valid_red || is_valid_green || is_valid_blue) {
				current_game = 0
			}
			fmt.Println(trailing_sum, current_game)
		}
	}
	trailing_sum += current_game
	return trailing_sum
}

func run_2(data []string) (int) {
	var red, green, blue, trailing_sum int
	for k, v := range data {
		current_int, _ := strconv.Atoi(v)

		if v == "Game" {
			fmt.Println(data[k+1], red*green*blue)
			trailing_sum += (red * green * blue)
			red, green, blue = 0, 0, 0
			
		} else if current_int > 0 {
			if strings.HasPrefix(data[k+1], "red") {
				red = max(red, current_int)
			} else if strings.HasPrefix(data[k+1], "green") {
				green = max(green, current_int)
			} else if strings.HasPrefix(data[k+1], "blue") {
				blue = max(blue, current_int)
			}
			fmt.Println(trailing_sum, red, green, blue)
		}
	}
	// Add last trailing game
	trailing_sum += (red * green * blue)

	return trailing_sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}