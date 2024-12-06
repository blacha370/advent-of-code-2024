package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() [][]byte {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	result := [][]byte{}
	for _, row := range strings.Split(string(data), "\n") {
		result = append(result, []byte(row))
	}
	return result
}

func getPosition(input [][]byte) [2]int {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '^' {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}
func main() {
	input := getInput()
	position := getPosition(input)
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	direction := 0
	result := map[string]bool{fmt.Sprint(position): true}
	for {
		if position[0]+directions[direction][0] < 0 || position[1]+directions[direction][1] < 0 || position[0]+directions[direction][0] >= len(input) || position[1]+directions[direction][1] >= len(input[0]) {
			break
		}
		for input[position[0]+directions[direction][0]][position[1]+directions[direction][1]] == '#' {
			direction = (direction + 1) % 4
		}
		position[0] = position[0] + directions[direction][0]
		position[1] = position[1] + directions[direction][1]
		result[fmt.Sprint(position)] = true
	}
	log.Printf("Result: %v\n", len(result))
}
