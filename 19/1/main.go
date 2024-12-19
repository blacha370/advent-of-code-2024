package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func getInput() ([]string, string) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	rows := strings.Split(string(data), "\n")
	regex := fmt.Sprintf("^((%s)){1,}$", strings.ReplaceAll(rows[0], ", ", ")|("))
	return rows[2:], regex
}

func main() {
	rows, regex := getInput()
	re, err := regexp.Compile(regex)

	if err != nil {
		log.Fatal(err)
	}

	result := 0
	for _, row := range rows {
		if re.MatchString(row) {
			result += 1
		}
	}

	log.Printf("Result: %v\n", result)
}
