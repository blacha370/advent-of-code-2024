package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() [][4]int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := [][4]int{}
	for _, row := range strings.Split(string(data), "\n") {
		r := strings.Split(row, " v=")
		p := strings.Split(r[0], ",")
		v := strings.Split(r[1], ",")
		result = append(result, [4]int{
			convertString(p[0][2:]),
			convertString(p[1]),
			convertString(v[0]),
			convertString(v[1]),
		})
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

func moveRobot(robot [4]int, board [2]int, seconds int) int {
	position := [2]int{(robot[0] + robot[2]*seconds) % board[0], (robot[1] + robot[3]*seconds) % board[1]}
	if position[0] < 0 {
		position[0] = board[0] + position[0]
	}
	if position[1] < 0 {
		position[1] = board[1] + position[1]
	}
	if position[0] == board[0]/2 || position[1] == board[1]/2 {
		return 4
	}
	if position[0] < board[0]/2 {
		if position[1] < board[1]/2 {
			return 0
		}
		return 2
	}
	if position[1] < board[1]/2 {
		return 1
	}
	return 3
}

func main() {
	data := getInput()
	quadrants := [5]int{0, 0, 0, 0, 0}
	board := [2]int{101, 103}
	seconds := 100
	for _, robot := range data {
		quadrants[moveRobot(robot, board, seconds)] += 1
	}
	result := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	log.Printf("Result: %v\n", result)
}
