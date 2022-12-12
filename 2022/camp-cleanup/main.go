package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func overlaps(aLeft, aRight, bLeft, bRight int) bool {
	return aLeft <= bLeft && aRight >= bRight ||
		bLeft <= aLeft && bRight >= aRight ||
		aLeft <= bLeft && aRight <= bRight && aRight >= bLeft ||
		aLeft >= bLeft && aLeft <= bRight && aRight >= bRight
}

func countOverlaps(pairs []string) int {
	overlapCount := 0

	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		a := strings.Split(pair, ",")[0]
		aLeft, _ := strconv.Atoi(strings.Split(a, "-")[0])
		aRight, _ := strconv.Atoi(strings.Split(a, "-")[1])

		b := strings.Split(pair, ",")[1]
		bLeft, _ := strconv.Atoi(strings.Split(b, "-")[0])
		bRight, _ := strconv.Atoi(strings.Split(b, "-")[1])

		if overlaps(aLeft, aRight, bLeft, bRight) {
			overlapCount++
		}
	}

	return overlapCount
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	answer := countOverlaps(strings.Split(content, "\n"))
	fmt.Println(answer)
}
