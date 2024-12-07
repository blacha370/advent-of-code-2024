package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() ([]int, [][]int) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	results := []int{}
	values := [][]int{}
	for _, row := range strings.Split(string(data), "\n") {
		r := strings.Split(row, ": ")
		results = append(results, convertString(r[0]))
		rowValues := []int{}
		for _, value := range strings.Split(r[1], " ") {
			rowValues = append(rowValues, convertString(value))
		}
		values = append(values, rowValues)
	}
	return results, values
}

func convertString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func checkRow(result int, values []int) bool {
	currentValues := []int{values[0]}
	for i := 1; i < len(values); i++ {
		tmpValues := []int{}
		for _, val := range currentValues {
			if tmpValue := val + values[i]; tmpValue <= result {
				tmpValues = append(tmpValues, tmpValue)
			}
			if tmpValue := val * values[i]; tmpValue <= result {
				tmpValues = append(tmpValues, tmpValue)
			}
		}
		currentValues = tmpValues
	}
	for _, val := range currentValues {
		if val == result {
			return true
		}
	}
	return false
}

func main() {
	results, values := getInput()
	result := 0
	for i := range results {
		if checkRow(results[i], values[i]) {
			result += results[i]
		}
	}
	log.Printf("Result: %v\n", result)
}
