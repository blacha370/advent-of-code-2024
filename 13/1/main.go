package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() [][3][2]int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := [][3][2]int{}
	for _, game := range strings.Split(string(data), "\n\n") {
		gameData := [3][2]int{}
		operations := [][2]string{
			{"Button A: X+", ", Y+"},
			{"Button B: X+", ", Y+"},
			{"Prize: X=", ", Y="},
		}
		rows := strings.Split(game, "\n")
		for i, operation := range operations {
			row := strings.Split(strings.Replace(rows[i], operation[0], "", 1), operation[1])
			gameData[i] = [2]int{convertString(row[0]), convertString(row[1])}
		}
		result = append(result, gameData)
	}
	return result
}

func convertString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func countMinTokens(game [3][2]int) int {
	currentLocation := [2]int{0, 0}
	pressedButtons := [2]int{0, 0}
	for currentLocation[0] < game[2][0] && currentLocation[1] < game[2][1] {
		currentLocation[0] += game[1][0]
		currentLocation[1] += game[1][1]
		pressedButtons[1] += 1
	}
	for currentLocation[0] != game[2][0] || currentLocation[1] != game[2][1] {
		if currentLocation[0] > game[2][0] && currentLocation[1] > game[2][1] {
			if pressedButtons[1] == 0 {
				return 0
			}
			currentLocation[0] -= game[1][0]
			currentLocation[1] -= game[1][1]
			pressedButtons[1] -= 1
		} else {
			currentLocation[0] += game[0][0]
			currentLocation[1] += game[0][1]
			pressedButtons[0] += 1
		}
	}
	return pressedButtons[0]*3 + pressedButtons[1]
}

func main() {
	data := getInput()
	result := 0
	for _, game := range data {
		result += countMinTokens(game)
	}
	log.Printf("Result: %v\n", result)
}
