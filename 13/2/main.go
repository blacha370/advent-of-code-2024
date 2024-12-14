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
		gameData[2][0] += 10000000000000
		gameData[2][1] += 10000000000000
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
	x := game[2][0] / game[1][0]
	currentLocation := [2]int{game[1][0] * x, game[1][1] * x}
	pressedButtons := [2]int{0, x}

	for currentLocation[0] != game[2][0] || currentLocation[1] != game[2][1] {
		if currentLocation[0] > game[2][0] || currentLocation[1] > game[2][1] {
			offset := (currentLocation[0] - game[2][0]) / game[1][0]
			if offset < (currentLocation[1]-game[2][1])/game[1][1] {
				offset = (currentLocation[1] - game[2][1]) / game[1][1]
			}
			offset += 1
			currentLocation[0] -= game[1][0] * offset
			currentLocation[1] -= game[1][1] * offset
			pressedButtons[1] -= offset

			if pressedButtons[1] < 0 {
				return 0
			}
			if currentLocation[0]-game[2][0]+game[1][0] == 0 && currentLocation[1]-game[2][1]+game[1][1] == 0 {
				pressedButtons[1] += 1
				break
			}
		} else {
			offset := (game[2][0] - currentLocation[0]) / game[0][0]
			if offset < (game[2][1]-currentLocation[1])/game[0][1] {
				offset = (game[2][1] - currentLocation[1]) / game[0][1]
			}
			offset += 1
			currentLocation[0] += game[0][0] * offset
			currentLocation[1] += game[0][1] * offset
			pressedButtons[0] += offset
			if currentLocation[0]-game[2][0]-game[0][0] == 0 && currentLocation[1]-game[2][1]-game[0][1] == 0 {
				pressedButtons[0] -= 1
				break
			}
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
