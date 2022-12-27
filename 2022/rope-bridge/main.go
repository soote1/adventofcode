package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	steps     int
}

type Point struct {
	x int
	y int
}

type Rope struct {
	knots []Point
}

func euclideanDistance(a *Point, b *Point) float64 {
	ax := float64(a.x)
	ay := float64(a.y)
	bx := float64(b.x)
	by := float64(b.y)

	return math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
}

func shouldMoveTail(h *Point, t *Point) bool {
	distance := math.Floor(euclideanDistance(h, t))
	return distance > 1.0
}

func getTailPositions(movements []Movement) map[Point]bool {
	head := Point{}
	tail := Point{}
	positions := make(map[Point]bool)
	positions[tail] = true

	for _, movement := range movements {
		for i := 0; i < movement.steps; i++ {
			switch movement.direction {
			case "L":
				head.x--
				if shouldMoveTail(&head, &tail) {
					tail.x = head.x + 1
					tail.y = head.y
				}
			case "U":
				head.y++
				if shouldMoveTail(&head, &tail) {
					tail.x = head.x
					tail.y = head.y - 1
				}
			case "R":
				head.x++
				if shouldMoveTail(&head, &tail) {
					tail.x = head.x - 1
					tail.y = head.y
				}
			case "D":
				head.y--
				if shouldMoveTail(&head, &tail) {
					tail.x = head.x
					tail.y = head.y + 1
				}
			}

			if _, ok := positions[tail]; !ok {
				positions[tail] = true
			}
		}
	}

	return positions
}

func parseMovements(lines []string) []Movement {
	movements := []Movement{}
	parts := []string{}
	s := 0
	d := ""

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts = strings.Split(line, " ")
		s, _ = strconv.Atoi(parts[1])
		d = parts[0]
		movements = append(movements, Movement{direction: d, steps: s})
	}

	return movements
}

func main() {
	inputFileName := os.Args[1]
	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}
	content := string(bytes)

	movements := parseMovements(strings.Split(content, "\n"))
	positions := getTailPositions(movements)
	fmt.Printf("%v\n", len(positions))
}
