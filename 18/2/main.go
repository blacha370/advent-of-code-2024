package main

import (
	"log"
	"os"
	"slices"
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

func findShortestWay(corrupted map[[2]int]bool, size [2]int) [][2]int {
	locations := map[[2]int][][2]int{{0, 0}: {}}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; len(locations) != 0; i++ {
		newLocations := map[[2]int][][2]int{}
		for k, v := range locations {
			if k[0] == size[0] && k[1] == size[1] {
				return v
			}
			for _, direction := range directions {
				newLocation := [2]int{k[0] + direction[0], k[1] + direction[1]}
				if slices.Index(v, newLocation) != -1 {
					continue
				}
				if _, ok := corrupted[newLocation]; ok {
					continue
				}
				if newLocation[0] >= 0 && newLocation[0] <= size[0] && newLocation[1] >= 0 && newLocation[1] <= size[1] {
					newLocations[newLocation] = append(v, k)
				}
			}
		}
		locations = newLocations
	}
	return [][2]int{}
}

func main() {
	boardSize := [2]int{70, 70}
	startBytes := 1024

	locations := getInput()
	corrupted := map[[2]int]bool{}
	for i, location := range locations {
		if i == startBytes {
			break
		}
		corrupted[location] = true
	}

	path := findShortestWay(corrupted, boardSize)

	var result [2]int
	for i := startBytes; i < len(locations); i++ {
		corrupted[locations[i]] = true
		if slices.Index(path, locations[i]) != -1 {
			path = findShortestWay(corrupted, boardSize)
			if len(path) == 0 {
				result = locations[i]
				break
			}
		}
	}

	log.Printf("Result: \"%v,%v\"\n", result[0], result[1])
}
