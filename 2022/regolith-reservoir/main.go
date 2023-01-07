package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func parseInput(input []string) map[Position]bool {
	positions := make(map[Position]bool)

	for _, line := range input {
		points := strings.Split(line, "->")
		for _, p := range points {
			p = strings.TrimSpace(p)
			coordinates := strings.Split(p, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			positions[Position{x: x, y: y}] = true
		}
	}

	return positions
}

func loadInput(fileName string) []string {
	input := []string{}

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func main() {
	input := loadInput(os.Args[1])
	positions := parseInput(input)
	fmt.Println(positions)
}
