package main

import (
	"log"
	"os"
	"strings"
)

func getInput() ([][]byte, [][2]int, [2]int) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	d := strings.Split(string(data), "\n\n")
	board := [][]byte{}
	position := [2]int{-1, -1}
	for i, row := range strings.Split(d[0], "\n") {
		if index := strings.Index(row, "@"); index != -1 {
			position = [2]int{i, index}
		}
		board = append(board, []byte(row))
	}
	moves := [][2]int{}
	for _, v := range d[1] {
		switch v {
		case '^':
			moves = append(moves, [2]int{-1, 0})
		case 'v':
			moves = append(moves, [2]int{1, 0})
		case '>':
			moves = append(moves, [2]int{0, 1})
		case '<':
			moves = append(moves, [2]int{0, -1})
		}
	}
	return board, moves, position
}

func main() {
	board, moves, position := getInput()
	for _, move := range moves {
		i := 1
		for {
			if board[position[0]+move[0]*i][position[1]+move[1]*i] == '#' {
				i = 0
				break
			}
			if board[position[0]+move[0]*i][position[1]+move[1]*i] == '.' {
				break
			}
			i++
		}
		if i == 0 {
			continue
		}
		board[position[0]][position[1]] = '.'
		board[position[0]+move[0]][position[1]+move[1]] = '@'
		position[0], position[1] = position[0]+move[0], position[1]+move[1]
		if i > 1 {
			board[position[0]+move[0]*(i-1)][position[1]+move[1]*(i-1)] = 'O'
		}
	}

	result := 0
	for i, row := range board {
		for j := range row {
			if row[j] == 'O' {
				result += j + i*100
			}
		}
	}
	log.Printf("Result: %v\n", result)
}
