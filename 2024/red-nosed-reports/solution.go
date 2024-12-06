package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func toInt(x string) int {
	n, err := strconv.Atoi(x)
	checkError(err)
	return n
}

func getDirection(left, right int) string {
	if left > right {
		return "-"
	}
	if left < right {
		return "+"
	}
	return "="
}

func isOutOfBounds(x int) bool {
	return x < -3 || x > 3 || x == 0
}

func isAllowed(from, to int, prevDir string) bool {
	direction := getDirection(from, to)
	isOutOfBounds := isOutOfBounds(from - to)
	return (prevDir == "" || direction == prevDir) && !isOutOfBounds
}

func isSafe(levels []int) bool {
	forbidden := 0
	i := 0
	dir := ""
	for i < len(levels)-1 {
		if isAllowed(levels[i], levels[i+1], dir) {
			dir = getDirection(levels[i], levels[i+1])
			i += 1
		} else {
			forbidden += 1
			if i == 0 {
				i += 1
				continue
			}
			if i+2 >= len(levels) {
				break
			}
			if forbidden > 1 || !isAllowed(levels[i], levels[i+2], dir) {
				return false
			}
			dir = getDirection(levels[i], levels[i+2])
			i += 2
		}
	}
	return forbidden < 2
}

func toIntSlice(levels []string) []int {
	var result []int
	for i := range len(levels) {
		result = append(result, toInt(levels[i]))
	}
	return result
}

func invert(slice []int) []int {
	var result []int
	i := len(slice) - 1
	for i >= 0 {
		result = append(result, slice[i])
		i -= 1
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	checkError(err)
	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		levels := toIntSlice(strings.Fields(scanner.Text()))
		if isSafe(levels) {
			safe += 1
		} else {
			if isSafe(invert(levels)) {
				safe += 1
			}
		}

	}
	fmt.Println(safe)
}
