package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	kind string
	data string
}

type Test struct {
	evaluation  string
	value       string
	trueAction  Action
	falseAction Action
}

type WorryLevelChange struct {
	op    string
	value string
}

type Monkey struct {
	items            []int
	worryLevelChange WorryLevelChange
	test             Test
}

func parseInput(input []string) []*Monkey {
	monkeys := []*Monkey{}
	i := -1
	for _, line := range input {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Monkey") {
			monkey := &Monkey{}
			monkeys = append(monkeys, monkey)
			i++
		}
		if strings.HasPrefix(line, "Starting") {
			parts := strings.Split(line, ": ")
			items := parts[1]
			for _, item := range strings.Split(items, ", ") {
				item, _ := strconv.Atoi(item)
				monkeys[i].items = append(monkeys[i].items, item)
			}
		}
		if strings.HasPrefix(line, "Operation") {
			parts := strings.Split(line, ": ")
			operation := parts[1]
			operationParts := strings.Split(operation, " ")
			monkeys[i].worryLevelChange.op = operationParts[3]
			monkeys[i].worryLevelChange.value = operationParts[4]
		}
		if strings.HasPrefix(line, "Test") {
			parts := strings.Split(line, ": ")
			monkeys[i].test.evaluation = "divisible by"
			monkeys[i].test.value = strings.Split(parts[1], " ")[2]
		}
		if strings.HasPrefix(line, "If true") {
			parts := strings.Split(line, ": ")
			action := strings.Split(parts[1], " ")
			monkeys[i].test.trueAction.kind = strings.Join(action[:2], " ")
			monkeys[i].test.trueAction.data = action[3]
		}
		if strings.HasPrefix(line, "If false") {
			parts := strings.Split(line, ": ")
			action := strings.Split(parts[1], " ")
			monkeys[i].test.falseAction.kind = strings.Join(action[:2], " ")
			monkeys[i].test.falseAction.data = action[3]
		}
	}
	return monkeys
}

func loadInput(fileName string) []string {
	input := []string{}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func changeWorryLevel(worryLevel int, operation string, factor int) int {
	result := 0

	switch operation {
	case "*":
		result = worryLevel * factor
	case "+":
		result = worryLevel + factor
	case "⌊/⌋":
		result = int(math.Floor(float64(worryLevel) / float64(factor)))
	}

	return result
}

func testWorryLevel(worryLevel int, op string, value int) bool {
	switch op {
	case "divisible by":
		return worryLevel%value == 0
	}
	panic("unsupported test operation")
}

func runSimulation(rounds int, monkeys []*Monkey) {
	worryLevel := 0
	op := ""
	value := 0
	m := 0
	inspectedItems := make(map[int]int)
	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			for len(monkey.items) > 0 {
				worryLevel = monkey.items[0]
				monkey.items = monkey.items[1:]
				inspectedItems[j]++
				// increase worry worry level
				op = monkey.worryLevelChange.op
				if monkey.worryLevelChange.value == "old" {
					value = worryLevel
					worryLevel = changeWorryLevel(worryLevel, op, value)
				} else {
					value, _ = strconv.Atoi(monkey.worryLevelChange.value)
					worryLevel = changeWorryLevel(worryLevel, op, value)
				}
				// reduce worry level
				worryLevel = changeWorryLevel(worryLevel, "⌊/⌋", 3)
				// test worry level
				op = monkey.test.evaluation
				value, _ = strconv.Atoi(monkey.test.value)
				if testWorryLevel(worryLevel, op, value) {
					if monkey.test.trueAction.kind == "throw to" {
						m, _ = strconv.Atoi(monkey.test.trueAction.data)
						monkeys[m].items = append(monkeys[m].items, worryLevel)
					}
				} else {
					if monkey.test.falseAction.kind == "throw to" {
						m, _ = strconv.Atoi(monkey.test.falseAction.data)
						monkeys[m].items = append(monkeys[m].items, worryLevel)
					}
				}
			}
		}
	}

	fmt.Println(inspectedItems)
}

func main() {
	input := loadInput(os.Args[1])
	monkeys := parseInput(input)
	runSimulation(20, monkeys)
}
