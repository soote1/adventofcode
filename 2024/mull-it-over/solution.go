package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Instruction struct {
	action string
	values []int
}

func main() {
	input, err := os.ReadFile("input.txt")
	data := string(input)
	check(err)
	state := "search"
	i := 0
	j := 0
	numbers_collected := 0
	x, y := 0, 0
	instructions := []Instruction{}
	products := []int{}
	for i < len(data) {
		switch state {
		case "search":
			x, y = 0, 0
			numbers_collected = 0
			if data[i] == 'm' {
				state = "collect_mul"
				j = i
			} else if data[i] == 'd' {
				state = "collect_do_or_dont"
				j = i
			} else {
				i += 1
			}
		case "collect_do_or_dont":
			target := "do()"
			p := j
			valid := true
			for q := range target {
				if p == len(data) {
					break
				}
				if target[q] != data[p] {
					valid = false
					break
				}
				p += 1
			}
			if valid {
				state = "search"
				instructions = append(instructions, Instruction{action: "do"})
				i += len(target)
			} else {
				target = "don't()"
				valid = true
				p = j
				for q := range target {
					if p == len(data) {
						break
					}
					if target[q] != data[p] {
						valid = false
						break
					}
					p += 1
				}
				if valid {
					state = "search"
					instructions = append(instructions, Instruction{action: "don't"})
					i += len(target)
				} else {
					state = "search"
					i += 1
				}
			}
		case "collect_mul":
			target := "mul("
			p := j
			valid := true
			for q := range target {
				if p == len(data) {
					break
				}
				if target[q] != data[p] {
					valid = false
					break
				}
				p += 1
			}
			if valid {
				state = "collect_number"
				j = p
			} else {
				state = "search"
				i += 1
			}
		case "collect_number":
			p := j
			length := 0
			number := []byte{}
			valid := true
			for {
				if p == len(data) {
					break
				}
				if data[p] >= '0' && data[p] <= '9' {
					if length == 3 {
						valid = false
						break
					}
					number = append(number, data[p])
					length += 1
					p += 1
				} else {
					break
				}
			}
			if valid {
				if numbers_collected == 0 {
					state = "collect_comma"
					j = p
					numbers_collected += 1
					x, _ = strconv.Atoi(string(number))
				} else if numbers_collected == 1 {
					state = "collect_end"
					j = p
					numbers_collected += 1
					y, _ = strconv.Atoi(string(number))
				} else {
					panic("invalid state")
				}
			} else {
				state = "search"
				i += 1
			}
		case "collect_comma":
			if data[j] == ',' {
				state = "collect_number"
				j += 1
			} else {
				state = "search"
				i += 1
			}
		case "collect_end":
			if data[j] == ')' {
				instructions = append(instructions, Instruction{action: "mul", values: []int{x, y}})
				products = append(products, x*y)

			}
			state = "search"
			i += 1
		}
	}
	sum := 0
	do_mul := true
	for i := range instructions {
		switch instructions[i].action {
		case "do":
			do_mul = true
		case "don't":
			do_mul = false
		case "mul":
			if do_mul {
				sum += instructions[i].values[0] * instructions[i].values[1]
			}
		}
	}
	fmt.Println(sum)
}
