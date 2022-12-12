package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func reArrangeStacks9000(stacks map[int][]string, moves [][]int) map[int][]string {
	for _, move := range moves {
		howMany := move[0]
		for i := 0; i < howMany; i++ {
			src := move[1]
			dest := move[2]
			stacks[dest] = append([]string{stacks[src][0]}, stacks[dest]...)
			stacks[src] = stacks[src][1:len(stacks[src])]
		}
	}

	return stacks
}

func reArrangeStacks9001(stacks map[int][]string, moves [][]int) map[int][]string {
	for _, move := range moves {
		howMany := move[0]
		src := move[1]
		dest := move[2]
		toBeAppended := []string{}

		for i := 0; i < howMany; i++ {
			toBeAppended = append(toBeAppended, stacks[src][0])
			stacks[src] = stacks[src][1:len(stacks[src])]
		}

		reversed := []string{}
		for i := 0; i < len(toBeAppended); i++ {
			reversed = append(reversed, toBeAppended[i])
		}

		stacks[dest] = append(reversed, stacks[dest]...)
	}

	return stacks
}

func parseStacks(stackLines []string) map[int][]string {
	stacks := make(map[int][]string)
	for _, line := range stackLines {
		for i, j := 1, 1; i < len(line); i, j = i+4, j+1 {
			item := string(line[i])
			if item != " " {
				// ignore number row
				if _, err := strconv.Atoi(item); err == nil {
					continue
				}
				stacks[j] = append(stacks[j], item)
			}
		}
	}

	return stacks
}

func parseMoves(moveLines []string) [][]int {
	moves := [][]int{}
	for _, line := range moveLines {
		if line == "" {
			continue
		}
		elements := strings.Split(line, " ")
		howMany, _ := strconv.Atoi(elements[1])
		from, _ := strconv.Atoi(elements[3])
		to, _ := strconv.Atoi(elements[5])
		moves = append(moves, []int{howMany, from, to})
	}

	return moves
}

func readFileContent(fileName string) ([]string, []string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	lines := strings.Split(content, "\n")
	stackLines := []string{}
	moveLines := []string{}
	state := "read-stacks"
	for _, line := range lines {
		if line == "" {
			state = "read-moves"
		}
		switch state {
		case "read-stacks":
			stackLines = append(stackLines, line)
		case "read-moves":
			moveLines = append(moveLines, line)
		}
	}

	return stackLines, moveLines
}

func main() {
	inputFileName := os.Args[1]
	stackLines, moveLines := readFileContent(inputFileName)
	stacks := parseStacks(stackLines)
	moves := parseMoves(moveLines)
	stacks = reArrangeStacks9001(stacks, moves)

	answer := ""
	for i := 1; i <= len(stacks); i++ {
		answer += stacks[i][0]
	}

	fmt.Print(answer)
}
