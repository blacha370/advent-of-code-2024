package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() [][2]int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := [][2]int{}
	for _, row := range strings.Split(string(data), "\n") {
		values := strings.Split(row, ",")
		result = append(result, [2]int{convertString(values[0]), convertString(values[1])})
	}
	return result
}

func convertString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func findShortestWay(corrupted map[[2]int]bool, size [2]int) int {
	locations := map[[2]int]bool{{0, 0}: true}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; len(locations) != 0; i++ {
		newLocations := map[[2]int]bool{}
		for k := range locations {
			if k[0] == size[0] && k[1] == size[1] {
				return i
			}
			for _, direction := range directions {
				newLocation := [2]int{k[0] + direction[0], k[1] + direction[1]}
				if _, ok := corrupted[newLocation]; !ok && newLocation[0] >= 0 && newLocation[0] <= size[0] && newLocation[1] >= 0 && newLocation[1] <= size[1] {
					newLocations[[2]int{k[0] + direction[0], k[1] + direction[1]}] = true
				}
			}
		}
		locations = newLocations
	}
	return 0
}

func main() {
	locations := getInput()
	corrupted := map[[2]int]bool{}
	for i, location := range locations {
		if i == 1024 {
			break
		}
		corrupted[location] = true
	}

	result := findShortestWay(corrupted, [2]int{70, 70})

	log.Printf("Result: %v\n", result)
}
