package main

import (
	"log"
	"os"
	"strings"
)

func getInput() [][]int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := [][]int{}
	for _, row := range strings.Split(string(data), "\n") {
		numbers := []int{}
		for _, char := range row {
			numbers = append(numbers, int(char)-48)
		}
		result = append(result, numbers)
	}
	return result
}

func countTrails(startingPos [2]int, input [][]int) int {
	result := map[[2]int]bool{}
	positions := [][2]int{}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, direction := range directions {
		if startingPos[0]+direction[0] < 0 || startingPos[0]+direction[0] >= len(input) || startingPos[1]+direction[1] < 0 || startingPos[1]+direction[1] >= len(input[0]) {
			continue
		}
		positions = append(positions, [2]int{startingPos[0] + direction[0], startingPos[1] + direction[1]})
	}

	for i := 1; len(positions) > 0 && i <= 9; i++ {
		tmpPositions := [][2]int{}
		for _, position := range positions {
			if input[position[0]][position[1]] != i {
				continue
			}
			if i == 9 {
				result[position] = true
				continue
			}
			for _, direction := range directions {
				if position[0]+direction[0] < 0 || position[0]+direction[0] >= len(input) || position[1]+direction[1] < 0 || position[1]+direction[1] >= len(input[0]) {
					continue
				}
				tmpPositions = append(tmpPositions, [2]int{position[0] + direction[0], position[1] + direction[1]})
			}
		}
		positions = tmpPositions
	}

	return len(result)
}

func main() {
	data := getInput()
	result := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 0 {
				result += countTrails([2]int{i, j}, data)
			}
		}
	}
	log.Printf("Result: %v\n", result)
}
