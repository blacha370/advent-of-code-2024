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

func findGap(input []int, gapSize int) []int {
	result := []int{}
	for i, v := range input {
		if v == -1 {
			result = append(result, i)
		} else {
			result = []int{}
		}
		if len(result) == gapSize {
			return result
		}
	}
	return []int{}
}

func sortDisk(input []int) []int {
	tmp := input
	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] == -1 {
			continue
		}

		fileLen := 1
		for j := i - 1; j >= 0; j-- {
			if tmp[j] != tmp[i] {
				break
			}
			fileLen++
		}
		for j, v := range findGap(tmp[:i-fileLen+1], fileLen) {
			tmp[v], tmp[i-j] = tmp[i-j], tmp[v]
		}
		i -= fileLen - 1
	}
	return tmp
}

func main() {
	data := getInput()
	result := 0
	for i, val := range sortDisk(data) {
		if val == -1 {
			continue
		}
		result += i * val
	}
	log.Printf("Result: %v\n", result)
}
