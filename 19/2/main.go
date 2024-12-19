package main

import (
	"log"
	"os"
	"strings"
)

func getInput() ([]string, []string) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	rows := strings.Split(string(data), "\n")
	stripes := strings.Split(rows[0], ", ")
	return rows[2:], stripes
}

func findAllPossibleWays(c map[string]int, towel string, stripes []string) int {
	if v, ok := c[towel]; ok {
		return v
	}
	result := 0
	for _, stripe := range stripes {
		if strings.HasPrefix(towel, stripe) {
			if len(towel) == len(stripe) {
				result += 1
				continue
			}
			result += findAllPossibleWays(c, strings.TrimPrefix(towel, stripe), stripes)
		}
	}
	c[towel] = result
	return result
}

func main() {
	rows, stripes := getInput()
	result := 0
	c := map[string]int{}
	for _, row := range rows {
		result += findAllPossibleWays(c, row, stripes)
	}
	log.Printf("Result: %v\n", result)
}
