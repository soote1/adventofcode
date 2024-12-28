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
	left, err := strconv.Atoi(y[0 : len(y)/2])
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

func blink(stones []int, count int) []int {
	transformed := []int{}
	cache := make(map[int][]int)
	for range count {
		for i := range len(stones) {
			t, ok := cache[stones[i]]
			if !ok {
				t = transform(stones[i])
			}
			transformed = append(transformed, t...)
			cache[stones[i]] = t
		}
		stones = []int{}
		stones = append(stones, transformed...)
		transformed = []int{}

	}
	return stones
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
	result := blink(stones, 75)
	fmt.Println(len(result))
}
