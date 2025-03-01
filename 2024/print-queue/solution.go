package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isCorrect(update []string, after map[string]map[string]bool, before map[string]map[string]bool) bool {
	i := 0
	for i < len(update) {
		j := i - 1
		for j >= 0 {
			if !before[update[i]][update[j]] {
				return false
			}
			j -= 1
		}
		j = i + 1
		for j < len(update) {
			if !after[update[i]][update[j]] {
				return false
			}
			j += 1
		}
		i += 1
	}
	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func sort(update []string, after map[string]map[string]bool) []string {
	i := 0
	for i < len(update) {
		j := i
		k := i + 1
		swaps := 0
		for k < len(update) {
			if !after[update[j]][update[k]] {
				update[j], update[k] = update[k], update[j]
				j = k
				swaps += 1
			}
			k += 1
		}
		if swaps == 0 {
			i += 1
		}
	}
	return update
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)
	mode := "scan_rule"
	after := make(map[string]map[string]bool)
	before := make(map[string]map[string]bool)
	updates := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mode = "scan_update"
			continue
		}
		switch mode {
		case "scan_rule":
			parts := strings.Split(line, "|")
			if _, ok := after[parts[0]]; !ok {
				after[parts[0]] = make(map[string]bool)
			}
			if _, ok := before[parts[1]]; !ok {
				before[parts[1]] = make(map[string]bool)
			}
			after[parts[0]][parts[1]] = true
			before[parts[1]][parts[0]] = true
		case "scan_update":
			parts := strings.Split(line, ",")
			updates = append(updates, parts)
		}
	}
	sum := 0
	sumFixed := 0
	for i := range updates {
		if isCorrect(updates[i], after, before) {
			x, err := strconv.Atoi(updates[i][len(updates[i])/2])
			checkError(err)
			sum += x
		} else {
			sorted := sort(updates[i], after)
			x, err := strconv.Atoi(sorted[len(sorted)/2])
			checkError(err)
			sumFixed += x
		}
	}
	fmt.Println(sum)
	fmt.Println(sumFixed)
}
