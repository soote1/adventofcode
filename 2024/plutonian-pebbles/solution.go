package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func splitNumber(x int) (int, int) {
	y := strconv.Itoa(x)
	left, err := strconv.Atoi(y[:len(y)/2])
	checkError(err)
	right, err := strconv.Atoi(y[len(y)/2:])
	checkError(err)
	return left, right
}

func isEvenLength(x int) bool {
	return len(strconv.Itoa(x))%2 == 0
}

func transform(x int) []int {
	result := []int{}
	if x == 0 {
		result = append(result, 1)
	} else if isEvenLength(x) {
		y, z := splitNumber(x)
		result = append(result, y)
		result = append(result, z)
	} else {
		result = append(result, x*2024)
	}
	return result
}

func solve(stones []int, limit int) int {
	transformed := make(map[int]map[int]int)
	transformed[0] = make(map[int]int)
	for i := range stones {
		transformed[0][stones[i]] = 1
	}
	for i := range limit {
		for s, c := range transformed[i] {
			x := transform(s)
			for j := range x {
				if _, ok := transformed[i+1]; !ok {
					transformed[i+1] = make(map[int]int)
				}
				if _, ok := transformed[i+1][x[j]]; ok {
					transformed[i+1][x[j]] += c
				} else {
					transformed[i+1][x[j]] = c
				}
			}
		}
	}
	total := 0
	for _, c := range transformed[limit] {
		total += c
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	stones := []int{}
	for i := range input {
		s, err := strconv.Atoi(input[i])
		checkError(err)
		stones = append(stones, s)
	}
	answer := solve(stones, 75)
	fmt.Println(answer)
}
