package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countPoints(rounds []string) int {
	totalPoints := 0
	values := map[string]int{"A": 0, "B": 1, "C": 2}

	for _, round := range rounds {
		if round == "" {
			continue
		}
		round := strings.Split(round, " ")
		a, _ := values[round[0]]
		result := round[1]
		points := 0

		switch result {
		case "X": // loose
			a = a - 1
			if a < 0 {
				a = 2
			}
			points = a + 1
		case "Y": // draw
			points = a + 1
			points += 3
		case "Z": // win
			points = ((a + 1) % 3) + 1
			points += 6
		}

		totalPoints += points
	}

	return totalPoints
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	answer := countPoints(strings.Split(content, "\n"))
	fmt.Println(answer)
}
