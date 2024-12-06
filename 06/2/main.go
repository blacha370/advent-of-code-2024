package main

import (
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

func getAllPossiblePositions(input [][]byte, startingPosition [2]int) map[[2]int]bool {
	position := [3]int{startingPosition[0], startingPosition[1], 0}
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	result := map[[2]int]bool{}
	for {
		if position[0]+directions[position[2]][0] < 0 || position[1]+directions[position[2]][1] < 0 || position[0]+directions[position[2]][0] >= len(input) || position[1]+directions[position[2]][1] >= len(input[0]) {
			break
		}
		for input[position[0]+directions[position[2]][0]][position[1]+directions[position[2]][1]] == '#' {
			position[2] = (position[2] + 1) % 4
		}
		position[0] = position[0] + directions[position[2]][0]
		position[1] = position[1] + directions[position[2]][1]
		if position[0] == startingPosition[0] && position[1] == startingPosition[1] {
			continue
		}
		result[[2]int{position[0], position[1]}] = true
	}
	return result
}

func isLoop(input [][]byte, startingPosition [2]int) bool {
	position := [3]int{startingPosition[0], startingPosition[1], 0}
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	fields := map[[3]int]bool{position: true}
	for {
		if position[0]+directions[position[2]][0] < 0 || position[1]+directions[position[2]][1] < 0 || position[0]+directions[position[2]][0] >= len(input) || position[1]+directions[position[2]][1] >= len(input[0]) {
			break
		}
		for input[position[0]+directions[position[2]][0]][position[1]+directions[position[2]][1]] == '#' {
			position[2] = (position[2] + 1) % 4
		}
		position[0] = position[0] + directions[position[2]][0]
		position[1] = position[1] + directions[position[2]][1]
		if _, ok := fields[position]; ok {
			return true
		}
		fields[position] = true
	}
	return false
}

func main() {
	input := getInput()
	position := getPosition(input)
	possiblePositions := getAllPossiblePositions(input, position)
	result := 0
	for k := range possiblePositions {
		input := getInput()
		input[k[0]][k[1]] = '#'
		if isLoop(input, position) {
			result += 1
		}
	}
	log.Printf("Result: %v\n", result)
}
