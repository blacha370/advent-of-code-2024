package main

import (
	"log"
	"os"
	"strings"
)

func getInput() []string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(data), "\n")
}

func countDiagonal(rows []string) int {
	count := 0
	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i])-1; j++ {
			if rows[i][j] != 'A' {
				continue
			}
			if rows[i-1][j-1] == 'M' && rows[i+1][j+1] == 'S' || rows[i-1][j-1] == 'S' && rows[i+1][j+1] == 'M' {
				if rows[i-1][j+1] == 'M' && rows[i+1][j-1] == 'S' || rows[i-1][j+1] == 'S' && rows[i+1][j-1] == 'M' {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	input := getInput()
	log.Printf("Result: %v\n", countDiagonal(input))
}
