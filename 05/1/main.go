package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInput() []string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(data), "\n\n")
}

func formatInput(input []string) (map[string][]string, [][]string) {
	rules := map[string][]string{}
	for _, v := range strings.Split(input[0], "\n") {
		rule := strings.Split(v, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	rows := [][]string{}
	for _, v := range strings.Split(input[1], "\n") {
		rows = append(rows, strings.Split(v, ","))
	}
	return rules, rows
}

func checkRow(rules map[string][]string, row []string) bool {
	for i := 0; i < len(row); i++ {
		if _, ok := rules[row[i]]; !ok {
			continue
		}
		for j := 0; j < i; j++ {
			if slices.Contains(rules[row[i]], row[j]) {
				return false
			}
		}
	}
	return true
}

func main() {
	rules, rows := formatInput(getInput())
	result := 0
	for _, row := range rows {
		if checkRow(rules, row) {
			v, err := strconv.Atoi(row[int(len(row)/2)])
			if err != nil {
				log.Panic(err)
			}
			result += v
		}
	}
	log.Printf("Result: %v\n", result)
}
