package main

import (
	"log"
	"os"
	"strings"
)

func getInput() (map[rune][][2]int, [2]int) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	rows := strings.Split(string(data), "\n")
	size := [2]int{len(rows), len(rows[0])}
	result := map[rune][][2]int{}
	for i, row := range rows {
		for j, val := range row {
			if val != '.' {
				result[val] = append(result[val], [2]int{i, j})
			}
		}
	}
	return result, size
}

func main() {
	antenas, size := getInput()
	positions := map[[2]int]bool{}
	for _, val := range antenas {
		for i := 0; i < len(val); i++ {
			positions[val[i]] = true
			for j := i + 1; j < len(val); j++ {
				offset := [2]int{val[i][0] - val[j][0], val[i][1] - val[j][1]}
				for k := 1; true; k++ {
					position := [2]int{val[i][0] + offset[0]*k, val[i][1] + offset[1]*k}
					if position[0] < 0 || position[1] < 0 || position[1] >= size[1] {
						break
					}
					positions[position] = true
				}
				for k := 1; true; k++ {
					position := [2]int{val[i][0] - offset[0]*k, val[i][1] - offset[1]*k}
					if position[0] >= size[0] || position[1] < 0 || position[1] >= size[1] {
						break
					}
					positions[position] = true
				}
			}
		}
	}
	log.Printf("Result: %v\n", len(positions))
}
