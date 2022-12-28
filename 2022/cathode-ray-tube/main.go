package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op    string
	value int
	cost  int
}

func getSignalStrengths(cycles map[int]bool, instructions []Instruction) []int {
	strengths := []int{}
	instructionCycle := 0
	programCycle := 0
	x := 1
	i := 0

	for i < len(instructions) {
		instructionCycle++
		programCycle++

		if _, ok := cycles[programCycle]; ok {
			strengths = append(strengths, programCycle*x)
		}

		if instructionCycle == instructions[i].cost {
			if instructions[i].op == "addx" {
				x += instructions[i].value
			}
			instructionCycle = 0
			i++
		}
	}

	return strengths
}

func draw(instructions []Instruction) []string {
	image := []string{}
	instructionCycle := 0
	programCycle := 0
	x := 1
	i := 0
	crtPos := 0
	rowLen := 39
	row := ""

	for i < len(instructions) {
		instructionCycle++
		programCycle++

		if crtPos > rowLen {
			image = append(image, row)
			crtPos = 0
			row = ""
		}

		// check if draw
		if x-1 <= crtPos && crtPos <= x+1 {
			row += "#"
		} else {
			row += "."
		}

		if instructionCycle == instructions[i].cost {
			if instructions[i].op == "addx" {
				x += instructions[i].value
			}
			instructionCycle = 0
			i++
		}

		crtPos++
	}

	image = append(image, row)

	return image
}

func parseIntructions(lines []string) []Instruction {
	instructions := []Instruction{}
	instruction := Instruction{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		instruction.op = parts[0]
		instruction.cost = 1
		if len(parts) == 2 {
			instruction.value, _ = strconv.Atoi(parts[1])
			instruction.cost = 2
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	instructions := parseIntructions(strings.Split(content, "\n"))
	cycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	strengths := getSignalStrengths(cycles, instructions)
	sum := 0
	for _, s := range strengths {
		sum += s
	}
	fmt.Printf("%v\n", sum)
	image := draw(instructions)
	for _, row := range image {
		fmt.Println(row)
	}
}
