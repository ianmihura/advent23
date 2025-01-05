package main

import (
	"fmt"
	"os"
	"strings"
)

func count_stones(row []byte) int {
	count := 0
	for _, s := range row {
		if s == 'O' {
			count++
		}
	}
	return count
}

func get_weight(board [][]byte) int {
	weight := 0
	for r := 0; r < len(board); r++ {
		weight += count_stones(board[r]) * (len(board) - r)
	}
	return weight
}

func tilt_board(board [][]byte) [][]byte {
	t_board := board
	has_moved := true
	for has_moved {
		has_moved = false
		for r, row := range board {
			for s, stone := range row {
				if stone == 'O' && r > 0 && board[r-1][s] == '.' {
					has_moved = true
					t_board[r-1][s] = 'O'
					t_board[r][s] = '.'
				}
			}
		}
	}
	return t_board
}

func run(board [][]byte) int {
	t_board := tilt_board(board)
	return get_weight(t_board)
}

func format(data []string) [][]byte {
	var fdata [][]byte
	for i := range data {
		fdata = append(fdata, []byte{})
		for j := range data[i] {
			fdata[i] = append(fdata[i], data[i][j])
		}
	}
	return fdata
}

func main() {
	data, _ := os.ReadFile("./full_input.txt")
	fdata := format(strings.Split(string(data), "\n"))
	answer := run(fdata)
	fmt.Println(answer)
}
