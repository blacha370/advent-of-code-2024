package main

import (
	"log"
	"os"
)

func getInput() []int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := []int{}
	for i := 0; i < len(data); i++ {
		for j := 0; j <= int(data[i])-49; j++ {
			if i%2 == 1 {
				result = append(result, -1)
			} else {
				result = append(result, i/2)
			}
		}
	}
	return result
}

func sortDisk(input []int) []int {
	tmp := input
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != -1 {
			continue
		}
		for j := len(tmp) - 1; j > i; j-- {
			if tmp[j] == -1 {
				continue
			}
			tmp[i], tmp[j] = tmp[j], -1
			break

		}
	}
	return tmp
}

func main() {
	data := getInput()
	result := 0
	for i, val := range sortDisk(data) {
		if val == -1 {
			break
		}
		result += i * val
	}
	log.Printf("Result: %v\n", result)
}
