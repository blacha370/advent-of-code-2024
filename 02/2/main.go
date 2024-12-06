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

func createRowVariations(row []int) [][]int {
	result := [][]int{}
	for range row {
		result = append(result, []int{})
	}

	for i := range row {
		for j := range result {
			if i == j {
				continue
			}
			result[j] = append(result[j], row[i])
		}
	}

	return result
}

func main() {
	rows := formatInput(getInput())
	safeCount := 0
	for _, row := range rows {
		if isRowSafe(row) {
			safeCount += 1
		} else {
			log.Println(row)
			tmpRows := createRowVariations(row)
			for _, tmpRow := range tmpRows {
				if isRowSafe(tmpRow) {
					safeCount += 1
					break
				}
			}
		}
	}
	log.Println(safeCount)
}
