package main

import (
	"log"
	"os"
	"sort"
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

func formatInput(input []string) [2][]int {
	result := [2][]int{{}, {}}
	for _, val := range input {
		values := strings.Split(val, "   ")
		result[0] = append(result[0], convertString(values[0]))
		result[1] = append(result[1], convertString(values[1]))
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func main() {
	values := formatInput(getInput())
	sort.Sort(sort.IntSlice(values[0]))
	sort.Sort(sort.IntSlice(values[1]))

	result := 0
	for i := range values[0] {
		result += abs(values[0][i] - values[1][i])
	}
	log.Printf("Result: %v\n", result)
}
