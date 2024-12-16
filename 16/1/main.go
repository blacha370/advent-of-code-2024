package main

import (
	"log"
	"os"
	"strings"
)

func getInput() ([][]byte, [2]int) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	board := [][]byte{}
	start := [2]int{-1, -1}
	for i, v := range strings.Split(string(data), "\n") {
		board = append(board, []byte(v))
		if index := strings.Index(v, "S"); index != -1 {
			start[0], start[1] = i, index
		}
	}
	return board, start
}

func filterOutDuplicates(paths [][4]int) [][4]int {
	result := [][4]int{}
	for i := range paths {
		if paths[i][3] == -1 {
			continue
		}
		for j := i + 1; j < len(paths); j++ {
			if paths[i][0] == paths[j][0] && paths[i][1] == paths[j][1] {
				if paths[i][3] > paths[j][3] && paths[j][3] != -1 {
					paths[i] = paths[j]
					paths[j][3] = -1
				} else if paths[i][3] < paths[j][3] {
					paths[j][3] = -1
				}
			}
		}
		result = append(result, paths[i])
	}
	return result
}

func findBestPath(board [][]byte, start [2]int) int {
	directions := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	paths := [][4]int{{start[0], start[1], 0, 0}}
	lowestScore := -1
	for len(paths) > 0 {
		newPaths := [][4]int{}
		for _, path := range paths {
			if lowestScore != -1 && path[3] > lowestScore {
				continue
			}
			if board[path[0]][path[1]] == 'E' {
				lowestScore = path[3]
				continue
			}
			if board[path[0]+directions[path[2]][0]][path[1]+directions[path[2]][1]] != '#' {
				newPaths = append(newPaths, [4]int{path[0] + directions[path[2]][0], path[1] + directions[path[2]][1], path[2], path[3] + 1})
			}
			left := path[2] - 1
			if left < 0 {
				left = 3
			}
			if board[path[0]+directions[left][0]][path[1]+directions[left][1]] != '#' {
				newPaths = append(newPaths, [4]int{path[0] + directions[left][0], path[1] + directions[left][1], left, path[3] + 1001})
			}
			right := path[2] + 1
			if right > 3 {
				right = 0
			}
			if board[path[0]+directions[right][0]][path[1]+directions[right][1]] != '#' {
				newPaths = append(newPaths, [4]int{path[0] + directions[right][0], path[1] + directions[right][1], right, path[3] + 1001})
			}
		}
		paths = filterOutDuplicates(newPaths)
	}
	return lowestScore
}

func main() {
	log.Printf("Result: %v\n", findBestPath(getInput()))
}
