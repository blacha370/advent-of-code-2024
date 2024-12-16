package main

import (
	"log"
	"os"
	"slices"
	"strings"
)

type Path struct {
	Fields [][3]int
	Score  int
}

func (p *Path) isLoop() bool {
	for i := range p.Fields {
		for j := i + 1; j < len(p.Fields); j++ {
			if p.Fields[i][0] == p.Fields[j][0] && p.Fields[i][1] == p.Fields[j][1] && p.Fields[i][2] == p.Fields[j][2] {
				return true
			}
		}
	}
	return false
}

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

func removeDuplicates(paths []Path) []Path {
	pathsMap := map[[3]int][]int{}
	for i, p := range paths {
		if p.isLoop() {
			continue
		}
		field := p.Fields[len(p.Fields)-1]
		pathsMap[field] = append(pathsMap[field], i)
	}
	newPaths := []Path{}
	for k, v := range pathsMap {
		newPath := Path{
			Score: -1,
		}
		for _, i := range v {
			if newPath.Score == -1 || paths[i].Score < newPath.Score {
				newPath.Score = paths[i].Score
			}
		}
		fields := map[[3]int]bool{}
		for _, i := range v {
			p := paths[i]
			if p.Score != newPath.Score {
				continue
			}
			for _, field := range p.Fields[:len(p.Fields)-1] {
				fields[field] = true
			}
		}
		for k1 := range fields {
			newPath.Fields = append(newPath.Fields, k1)
		}
		newPath.Fields = append(newPath.Fields, [3]int{k[0], k[1], k[2]})
		newPaths = append(newPaths, newPath)
	}
	return newPaths
}

func countUniqueFields(fields [][3]int) int {
	result := map[[2]int]bool{}
	for _, v := range fields {
		result[[2]int{v[0], v[1]}] = true
	}
	return len(result)
}

func findBestPath(board [][]byte, start [2]int) int {
	directions := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	paths := []Path{
		{
			Fields: [][3]int{{start[0], start[1], 0}},
			Score:  0,
		},
	}
	successfulPaths := [][3]int{}

	// lowestScore from previous part
	lowestScore := 105508
	for len(paths) > 0 {
		newPaths := []Path{}
		for _, path := range paths {
			if path.Score > lowestScore {
				continue
			}
			field := path.Fields[len(path.Fields)-1]
			if board[field[0]][field[1]] == 'E' {
				successfulPaths = append(successfulPaths, path.Fields...)
				continue
			}
			if board[field[0]+directions[field[2]][0]][field[1]+directions[field[2]][1]] != '#' {
				newPaths = append(newPaths, Path{
					Fields: append(slices.Clone(path.Fields), [3]int{field[0] + directions[field[2]][0], field[1] + directions[field[2]][1], field[2]}),
					Score:  path.Score + 1,
				})

			}
			left := field[2] - 1
			if left < 0 {
				left = 3
			}
			if board[field[0]+directions[left][0]][field[1]+directions[left][1]] != '#' {
				newPaths = append(newPaths, Path{
					Fields: append(slices.Clone(path.Fields), [3]int{field[0] + directions[left][0], field[1] + directions[left][1], left}),
					Score:  path.Score + 1001,
				})
			}
			right := field[2] + 1
			if right > 3 {
				right = 0
			}
			if board[field[0]+directions[right][0]][field[1]+directions[right][1]] != '#' {
				newPaths = append(newPaths, Path{
					Fields: append(slices.Clone(path.Fields), [3]int{field[0] + directions[right][0], field[1] + directions[right][1], right}),
					Score:  path.Score + 1001,
				})
			}
		}
		paths = removeDuplicates(newPaths)
	}
	return countUniqueFields(successfulPaths)
}

func main() {
	log.Printf("Result: %v\n", findBestPath(getInput()))
}
