package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput() string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func convertString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func multiply(s string) int {
	values := strings.Split(s[4:len(s)-1], ",")
	return convertString(values[0]) * convertString(values[1])
}

func main() {
	input := getInput()
	re := regexp.MustCompile("mul[(][0-9]{1,3},[0-9]{1,3}[)]")
	total := 0
	for _, val := range re.FindAllString(input, -1) {
		total += multiply(val)
	}
	log.Printf("Result: %v\n", total)
}
