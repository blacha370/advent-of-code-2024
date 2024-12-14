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

func moveRobot(robot [4]int, board [2]int, seconds int) [2]int {
	position := [2]int{(robot[0] + robot[2]*seconds) % board[0], (robot[1] + robot[3]*seconds) % board[1]}
	if position[0] < 0 {
		position[0] = board[0] + position[0]
	}
	if position[1] < 0 {
		position[1] = board[1] + position[1]
	}
	return position
}

func countConnected(positions map[[2]int]bool) int {
	visited := map[[2]int]bool{}
	result := 0
	for k := range positions {
		r := 0
		positionsToVisit := [][2]int{k}
		for len(positionsToVisit) > 0 {
			pos := positionsToVisit[0]
			positionsToVisit = positionsToVisit[1:]
			if _, ok := visited[pos]; ok {
				continue
			}
			if _, ok := positions[pos]; ok {
				r++
				positionsToVisit = append(positionsToVisit, [][2]int{
					{pos[0] - 1, pos[1]},
					{pos[0] + 1, pos[1]},
					{pos[0], pos[1] - 1},
					{pos[0], pos[1] + 1},
				}...)
				visited[pos] = true
			}
		}
		if r > result {
			result = r
		}
	}
	return result
}

func main() {
	data := getInput()
	board := [2]int{101, 103}
	for i := 1; i < 10000; i++ {
		positions := map[[2]int]bool{}
		for _, robot := range data {
			positions[moveRobot(robot, board, i)] = true
		}
		r := countConnected(positions)

		if r > 100 {
			b := [103][101]string{}
			for i := range b {
				for j := range b[i] {
					b[i][j] = " "
				}
			}
			for k := range positions {
				b[k[1]][k[0]] = "#"
			}
			for _, row := range b {
				log.Println(row)
			}
			log.Println("Please check if board contains easter egg")
			log.Printf("Current time: %v\n", i)
			break
		}
	}
}
