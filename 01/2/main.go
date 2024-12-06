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

func formatInput(input []string) ([]int, map[int]int) {
	leftColumn := []int{}
	rightColumn := map[int]int{}
	for _, val := range input {
		values := strings.Split(val, "   ")
		leftColumn = append(leftColumn, convertString(values[0]))
		i := convertString(values[1])
		rightColumn[i] += i
	}

	return leftColumn, rightColumn
}

func main() {
	leftColumn, rightColumn := formatInput(getInput())
	result := 0
	for _, val := range leftColumn {
		if _, ok := rightColumn[val]; ok {
			result += rightColumn[val]
		}
	}
	log.Printf("Result: %v\n", result)
}
