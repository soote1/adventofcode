package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func canMove(p Position, r *map[Position]bool, s *map[Position]bool) bool {
	isSand, _ := (*s)[p]
	isRock, _ := (*r)[p]
	return !isSand && !isRock
}

func generateSand(rocks *map[Position]bool, limit Position) *map[Position]bool {
	var moves int
	var nextMove string
	var position Position
	sand := make(map[Position]bool)
	for {
		nextMove = "down"
		position.x = 500
		position.y = 0
		moves = 0
		for nextMove != "" {
            if position.y >= limit.y {
                moves = 0
                break
            }
			switch nextMove {
			case "down":
				position.y++
				if !canMove(position, rocks, &sand) {
					nextMove = "down-left"
					position.y--
				} else {
					moves++
					nextMove = "down"
				}
			case "down-left":
				position.y++
				position.x--
				if !canMove(position, rocks, &sand) {
					nextMove = "down-right"
					position.y--
					position.x++
				} else {
					moves++
					nextMove = "down"
				}
			case "down-right":
				position.y++
				position.x++
				if !canMove(position, rocks, &sand) {
					position.y--
					position.x--
					sand[position] = true
					nextMove = ""
				} else {
					moves++
					nextMove = "down"
				}
			}
		}
		if moves == 0 {
			break
		}
	}
	return &sand
}

func generateLine(p1 Position, p2 Position) []Position {
	i := 0
	line := []Position{p1, p2}
	position := Position{}

	if p1.x == p2.x {
		i = p1.y
	} else if p1.y == p2.y {
		i = p1.x
	} else {
		panic("Can't draw line")
	}

	for {
		if p1.x == p2.x {
			if p1.y < p2.y {
				i++
			}
			if p1.y > p2.y {
				i--
			}
			if i == p2.y {
				break
			}
			position.x = p1.x
			position.y = i
		} else {
			if p1.x < p2.x {
				i++
			}
			if p1.x > p2.x {
				i--
			}
			if i == p2.x {
				break
			}
			position.x = i
			position.y = p1.y
		}
		line = append(line, position)
	}

	return line
}

func parseInput(input []string) (map[Position]bool, Position) {
	start := Position{}
	end := Position{}
	positions := make(map[Position]bool)
    limit := Position{}

	for _, line := range input {
		if line == "" {
			continue
		}
		points := strings.Split(line, "->")
		for i, p := range points {
			p = strings.TrimSpace(p)
			coordinates := strings.Split(p, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			end.x = x
			end.y = y
            if end.y > limit.y {
                limit = end
            }
			if i >= 1 {
				line := generateLine(start, end)
				for _, p := range line {
					positions[p] = true
				}
			}
			start = end
		}
	}

	return positions, limit
}

func loadInput(fileName string) []string {
	input := []string{}

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func main() {
	input := loadInput(os.Args[1])
	positions, limit := parseInput(input)
	fmt.Println(positions)
	sand := generateSand(&positions, limit)
	fmt.Println(len(*sand))
}
