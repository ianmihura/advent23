package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    data, _ := os.ReadFile("./full_input.txt")
	sdata := strings.Split(string(data), "\n")
	desert_map, start_nodes := create(sdata[2:])
	answer := traverse_2(sdata[0], desert_map, start_nodes)
	fmt.Println(answer)
}

func create(data []string) (map[string]string, []string) {
	desert_map := make(map[string]string)
	var start_nodes []string

	for _, v := range data {
		fields := strings.Fields(v)
		name := fields[0]
		desert_map[name + "L"] = fields[2][1:4]
		desert_map[name + "R"] = fields[3][:3]
		if string(name[2]) == "A" {
			start_nodes = append(start_nodes, name)
		}
	}

	return desert_map, start_nodes
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

func traverse_2(instructions string, desert_map map[string]string, start_nodes []string) int {
	steps := 0
	current_nodes := start_nodes

	for {
		ins := steps % len(instructions)
		current_instruction := string(instructions[ins])
		steps++
		
		for i := 0; i < len(current_nodes); i++ {
			// First loop to save current_nodes
			current_nodes[i] = desert_map[current_nodes[i] + current_instruction]
			if string(current_nodes[i][2]) == "Z" && steps < 50000 {
				fmt.Println(i, steps)
			}
		}
	}

	return steps
}
