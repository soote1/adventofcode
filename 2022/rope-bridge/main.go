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
	knots []*Point
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

func getTailPositions(movements []Movement, rope *Rope) map[Point]bool {
	head := rope.knots[0]
	tail := rope.knots[len(rope.knots)-1]
	subHead := &Point{}
	subTail := &Point{}
	positions := make(map[Point]bool)
	positions[*tail] = true

	for _, movement := range movements {
		for i := 0; i < movement.steps; i++ {
			switch movement.direction {
			case "L":
				head.x--
			case "U":
				head.y++
			case "R":
				head.x++
			case "D":
				head.y--
			}

			for j := 0; j < len(rope.knots)-1; j++ {
				subHead = rope.knots[j]
				subTail = rope.knots[j+1]
				if shouldMoveTail(subHead, subTail) {
					if subHead.x == subTail.x && subHead.y > subTail.y {
						subTail.y++
					} else if subHead.x == subTail.x && subHead.y < subTail.y {
						subTail.y--
					} else if subHead.y == subTail.y && subHead.x > subTail.x {
						subTail.x++
					} else if subHead.y == subTail.y && subHead.x < subTail.x {
						subTail.x--
					} else if subHead.y > subTail.y && subHead.x > subTail.x {
						subTail.x++
						subTail.y++
					} else if subHead.y > subTail.y && subHead.x < subTail.x {
						subTail.x--
						subTail.y++
					} else if subHead.y < subTail.y && subHead.x > subTail.x {
						subTail.x++
						subTail.y--
					} else if subHead.y < subTail.y && subHead.x < subTail.x {
						subTail.x--
						subTail.y--
					}
				}
			}

			if _, ok := positions[*tail]; !ok {
				positions[*tail] = true
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
	rope := Rope{}
	for i := 0; i < 10; i++ {
		rope.knots = append(rope.knots, &Point{})
	}
	positions := getTailPositions(movements, &rope)
	fmt.Printf("%v\n", len(positions))
}
