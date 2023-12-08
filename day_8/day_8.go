package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    data, _ := os.ReadFile("./full_input.txt")
	sdata := strings.Split(string(data), "\n")
	desert_map := create(sdata[2:])
	answer := traverse(sdata[0], desert_map)
	fmt.Println(answer)
}

func create(data []string) map[string]string {
	desert_map := make(map[string]string)
	for _, v := range data {
		fields := strings.Fields(v)
		name := fields[0]
		desert_map[name + "L"] = fields[2][1:4]
		desert_map[name + "R"] = fields[3][:3]
	}
	return desert_map
}

func traverse(instructions string, desert_map map[string]string) int {
	steps := 0
	current_node := "AAA"

	for current_node != "ZZZ" {
		i := steps % len(instructions)
		current_instruction := string(instructions[i])
		current_node = desert_map[current_node + current_instruction]
		steps++
	}

	return steps
}
