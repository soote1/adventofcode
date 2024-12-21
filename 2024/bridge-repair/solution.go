package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	value    int
	operator string
}

func toInt(x string) int {
	result, err := strconv.Atoi(x)
	checkError(err)
	return result
}

func buildOpTree(values []string) []Node {
	operators := map[string]string{"*": "+", "+": "||", "||": "*"}
	currentOperator := "*"
	nodes := []Node{{value: toInt(values[0])}}
	for i := range values {
		if i == 0 {
			continue
		}
		prevOperator := currentOperator
		j := int(math.Pow(float64(3), float64(i)))
		for range j {
			currentOperator := operators[prevOperator]
			node := Node{value: toInt(values[i]), operator: currentOperator}
			prevOperator = currentOperator
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func isValidOpTree(tree []Node, result int) bool {
	value := tree[0].value
	visited := make(map[int]bool)
	current := 0
	for {
		if _, ok := visited[current]; !ok && current >= 0 && current < len(tree) {
			// process
			visited[current] = true
			switch tree[current].operator {
			case "+":
				if value <= (math.MaxInt - tree[current].value) {
					value += tree[current].value
				} else {
					// backtrack
					parent := (current - 1) / 3
					current = parent
					continue
				}
			case "*":
				if value <= (math.MaxInt / tree[current].value) {
					value *= tree[current].value
				} else {
					// backtrack
					parent := (current - 1) / 3
					current = parent
					continue
				}
			case "||":
				joined, err := join(value, tree[current].value)
				if err != nil {
					// backtrack
					parent := (current - 1) / 3
					current = parent
					continue
				}
				value = joined
			}
		}
		left := (3 * current) + 1
		middle := (3 * current) + 2
		right := (3 * current) + 3
		if left >= len(tree) && right >= len(tree) && middle >= len(tree) {
			if value == result {
				return true
			}
			if current == 0 {
				break
			}
			// backtrack
			parent := (current - 1) / 3
			switch tree[current].operator {
			case "+":
				value -= tree[current].value
			case "*":
				value /= tree[current].value
			case "||":
				value = remove(value, tree[current].value)
			}
			current = parent
			continue
		}
		if _, ok := visited[left]; !ok {
			current = left
			continue
		}
		if _, ok := visited[middle]; !ok {
			current = middle
			continue
		}
		if _, ok := visited[right]; !ok {
			current = right
			continue
		}
		if current == 0 {
			break
		}
		// backtrack
		parent := (current - 1) / 3
		switch tree[current].operator {
		case "+":
			value -= tree[current].value
		case "*":
			value /= tree[current].value
		case "||":
			value = remove(value, tree[current].value)
		}
		current = parent
	}
	return false
}

func join(x, y int) (int, error) {
	joined := strconv.Itoa(x) + strconv.Itoa(y)
	z, err := strconv.Atoi(joined)
	if err != nil {
		return -1, err
	}
	return z, nil
}

func remove(x, y int) int {
	left := strconv.Itoa(x)
	right := strconv.Itoa(y)
	z := left[:len(left)-len(right)]
	result, err := strconv.Atoi(z)
	checkError(err)
	return result
}

func main() {
	totalCalibrationResult := 0
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colonPos := strings.Index(line, ":")
		result, err := strconv.Atoi(line[:colonPos])
		checkError(err)
		values := strings.Fields(line[colonPos+2:])
		tree := buildOpTree(values)
		if isValidOpTree(tree, result) {
			totalCalibrationResult += result
		}
	}
	fmt.Println(totalCalibrationResult)
}
