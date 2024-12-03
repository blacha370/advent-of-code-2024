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

func cleanupInput(input string) string {
	result := input
	regexes := []string{
		"\n",                     // multiline
		"don't[(][)].*?do[(][)]", // everything between don't() and do()
		"don't[(][)].*",          // trailing don't()
	}

	for _, val := range regexes {
		re := regexp.MustCompile(val)
		result = re.ReplaceAllString(result, " ")
	}
	return result
}

func main() {
	input := cleanupInput(getInput())

	re := regexp.MustCompile("mul[(][0-9]{1,3},[0-9]{1,3}[)]")
	total := 0

	for _, val := range re.FindAllString(input, -1) {
		total += multiply(val)
	}
	log.Printf("Result: %v\n", total)
}
