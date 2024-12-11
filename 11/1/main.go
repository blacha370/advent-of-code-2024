package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput() []int {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result := []int{}
	for _, value := range strings.Split(string(data), " ") {
		result = append(result, convertString(value))
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

func processStone(stone int, times int, cache map[[2]int]int) int {
	if val, ok := cache[[2]int{stone, times}]; ok {
		return val
	}
	if times == 0 {
		cache[[2]int{stone, times}] = 1
		return 1
	}
	if stone == 0 {
		result := processStone(1, times-1, cache)
		cache[[2]int{stone, times}] = result
		return result
	}
	stoneString := fmt.Sprint(stone)
	if len(stoneString)%2 == 0 {
		re := regexp.MustCompile("^0{1,}$")
		left, right := stoneString[:len(stoneString)/2], re.ReplaceAllString(stoneString[len(stoneString)/2:], "0")
		result := processStone(convertString(left), times-1, cache) + processStone(convertString(right), times-1, cache)
		cache[[2]int{stone, times}] = result
		return result
	}
	result := processStone(stone*2024, times-1, cache)
	cache[[2]int{stone, times}] = result
	return result
}

func main() {
	data := getInput()
	result := 0
	cache := map[[2]int]int{}
	for _, stone := range data {
		result += processStone(stone, 25, cache)
	}
	log.Printf("Result: %v\n", result)
}
