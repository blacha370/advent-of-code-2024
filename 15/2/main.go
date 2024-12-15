package main

import (
	"log"
	"os"
	"slices"
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
		updatedRow := []byte{}
		for j, v := range row {
			switch v {
			case '@':
				position = [2]int{i, j * 2}
				updatedRow = append(updatedRow, []byte("@.")...)
			case '#':
				updatedRow = append(updatedRow, []byte("##")...)
			case 'O':
				updatedRow = append(updatedRow, []byte("[]")...)
			case '.':
				updatedRow = append(updatedRow, []byte("..")...)
			}
		}
		board = append(board, updatedRow)
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
		if move[0] == 0 {
			i := 1
			for {
				if board[position[0]][position[1]+move[1]*i] == '#' {
					i = 0
					break
				}
				if board[position[0]][position[1]+move[1]*i] == '.' {
					break
				}
				i++
			}
			if i == 0 {
				continue
			}
			for j := i; j >= 2; j-- {
				board[position[0]][position[1]+move[1]*j] = board[position[0]][position[1]+move[1]*(j-1)]
			}
			board[position[0]][position[1]] = '.'
			board[position[0]][position[1]+move[1]] = '@'
			position[1] = position[1] + move[1]
		} else {
			fieldsToCheck := [][2]int{{position[0] + move[0], position[1]}}
			boxes := [][2]int{}
			i := 1
			for len(fieldsToCheck) > 0 {
				broken := false
				newBoxes := [][2]int{}
				for _, field := range fieldsToCheck {
					if board[field[0]][field[1]] == '#' {
						broken = true
						break
					}
					if board[field[0]][field[1]] == '[' {
						newBoxes = append(newBoxes, field)
						continue
					}
					if board[field[0]][field[1]] == ']' {
						if slices.Index(newBoxes, [2]int{field[0], field[1] - 1}) == -1 {
							newBoxes = append(newBoxes, [2]int{field[0], field[1] - 1})
						}
						continue
					}
				}
				if broken {
					i = 0
					break
				}
				i++
				fieldsToCheck = [][2]int{}
				if len(newBoxes) > 0 {
					newFieldsToCheck := [][2]int{}
					for _, box := range newBoxes {
						newFieldsToCheck = append(newFieldsToCheck, [][2]int{{box[0] + move[0], box[1]}, {box[0] + move[0], box[1] + 1}}...)
					}
					boxes = append(boxes, newBoxes...)
					fieldsToCheck = newFieldsToCheck
				}
			}
			if i == 0 {
				continue
			}

			for j := range boxes {
				box := boxes[len(boxes)-j-1]
				board[box[0]][box[1]], board[box[0]][box[1]+1] = '.', '.'
				board[box[0]+move[0]][box[1]], board[box[0]+move[0]][box[1]+1] = '[', ']'
			}
			board[position[0]][position[1]] = '.'
			board[position[0]+move[0]][position[1]] = '@'
			position[0] = position[0] + move[0]
		}
	}

	result := 0
	for i, row := range board {
		for j := range row {
			if row[j] == '[' {
				result += j + i*100
			}
		}
	}
	log.Printf("Result: %v\n", result)
}
