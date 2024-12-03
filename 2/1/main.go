package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(data), "\n")
}

func convertString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func formatInput(input []string) [][]int {
	result := [][]int{}
	for _, row := range input {
		values := []int{}
		for _, val := range strings.Split(row, " ") {
			values = append(values, convertString(val))
		}
		result = append(result, values)
	}
	return result
}

func isOk(x, y int, increasing bool) bool {
	if increasing {
		if x-y >= -3 && x-y <= -1 {
			return true
		}
		return false
	}

	if x-y <= 3 && x-y >= 1 {
		return true
	}
	return false
}

func isRowSafe(row []int) bool {
	increasing := row[0] < row[1]
	for i := 0; i < len(row)-1; i++ {
		if !isOk(row[i], row[i+1], increasing) {
			return false
		}
	}
	return true
}

func main() {
	values := formatInput(getInput())
	safeCount := 0
	for _, row := range values {
		if isRowSafe(row) {
			safeCount += 1
		}
	}
	log.Println(safeCount)
}
