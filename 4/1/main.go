package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func getInput() string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func countHorizontal(input string) int {
	re := regexp.MustCompile("(XMAS)")
	result := len(re.FindAllString(input, -1))
	re = regexp.MustCompile("(SAMX)")
	return result + len(re.FindAllString(input, -1))
}

func countVertical(input string) int {
	rows := strings.Split(input, "\n")
	count := 0
	for i, row := range rows {
		for j, c := range row {
			if c == 'X' {
				if i >= 3 && rows[i-1][j] == 'M' && rows[i-2][j] == 'A' && rows[i-3][j] == 'S' {
					count += 1
				}
				if len(rows)-i > 3 && rows[i+1][j] == 'M' && rows[i+2][j] == 'A' && rows[i+3][j] == 'S' {
					count += 1
				}
			}
		}
	}
	return count
}

func countDiagonal(input string) int {
	rows := strings.Split(input, "\n")
	count := 0
	for i, row := range rows {
		for j, c := range row {
			if c == 'X' {
				if i >= 3 && j >= 3 && rows[i-1][j-1] == 'M' && rows[i-2][j-2] == 'A' && rows[i-3][j-3] == 'S' {
					count += 1
				}
				if i >= 3 && len(row)-j > 3 && rows[i-1][j+1] == 'M' && rows[i-2][j+2] == 'A' && rows[i-3][j+3] == 'S' {
					count += 1
				}
				if len(rows)-i > 3 && j >= 3 && rows[i+1][j-1] == 'M' && rows[i+2][j-2] == 'A' && rows[i+3][j-3] == 'S' {
					count += 1
				}
				if len(rows)-i > 3 && len(row)-j > 3 && rows[i+1][j+1] == 'M' && rows[i+2][j+2] == 'A' && rows[i+3][j+3] == 'S' {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	input := getInput()
	result := countHorizontal(input)
	result += countVertical(input)
	result += countDiagonal(input)
	log.Printf("Result: %v\n", result)
}
