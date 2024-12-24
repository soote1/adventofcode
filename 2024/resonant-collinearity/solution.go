package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unicode"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Point struct {
	x float64
	y float64
}

type Line struct {
	a Point
	b Point
}

func buildLines(grid []string) []Line {
	lines := []Line{}
	prevPoints := make(map[byte][]Point)
	for y := range grid {
		for x := range grid[y] {
			if unicode.IsLetter(rune(grid[y][x])) || unicode.IsDigit(rune(grid[y][x])) {
				if points, ok := prevPoints[grid[y][x]]; ok {
					for i := range points {
						lines = append(lines, Line{a: points[i], b: Point{float64(x), float64(y)}})
					}
				}
				prevPoints[grid[y][x]] = append(prevPoints[grid[y][x]], Point{float64(x), float64(y)})
			}
		}
	}
	return lines
}

func getDistance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func getMagnitude(p Point) float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func getUnitVector(p Point) Point {
	magnitude := getMagnitude(p)
	return Point{x: p.x / magnitude, y: p.y / magnitude}
}

func getDifferenceVector(a, b Point) Point {
	return Point{x: b.x - a.x, y: b.y - a.y}
}

func getPointAtDistance(a, b Point, distance float64) Point {
	diffVector := getDifferenceVector(a, b)
	unitVector := getUnitVector(diffVector)
	return Point{x: math.Round((unitVector.x * distance) + a.x), y: math.Round((unitVector.y * distance) + a.y)}
}

func findAntinodes(line Line) (Point, Point) {
	distance := getDistance(line.a, line.b)
	p1 := getPointAtDistance(line.a, line.b, 2*distance)
	p2 := getPointAtDistance(line.b, line.a, 2*distance)
	return p1, p2
}

func isWithinBounds(grid []string, p Point) bool {
	return p.y < float64(len(grid)) && p.y >= 0 && p.x < float64(len(grid[0])) && p.x >= 0
}

func getAntinodes(grid []string, line Line) []Point {
	points := []Point{line.a, line.b}
	distance := getDistance(line.a, line.b)
	i := 2.0
	for {
		p := getPointAtDistance(line.a, line.b, distance*i)
		points = append(points, p)
		if !isWithinBounds(grid, p) {
			break
		}
		i += 1
	}
	i = 2.0
	for {
		p := getPointAtDistance(line.b, line.a, distance*i)
		points = append(points, p)
		if !isWithinBounds(grid, p) {
			break
		}
		i += 1
	}
	return points
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	grid := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	lines := buildLines(grid)
	antinodes := make(map[Point]bool)
	for i := range lines {
		nodes := getAntinodes(grid, lines[i])
		for j := range nodes {
			if isWithinBounds(grid, nodes[j]) {
				antinodes[nodes[j]] = true
			}
		}
	}
	fmt.Println(len(antinodes))
}
