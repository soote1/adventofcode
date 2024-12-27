package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func findUniqueEndpoints(start Point, maze []string) map[Point]bool {
	endPoints := make(map[Point]bool)
	pending := []Point{start}
	for len(pending) != 0 {
		current := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if maze[current.y][current.x] == '9' {
			endPoints[current] = true
		}
		if current.x+1 < len(maze[0]) && int(maze[current.y][current.x+1]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x + 1, y: current.y})
		}
		if current.y+1 < len(maze) && int(maze[current.y+1][current.x]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x, y: current.y + 1})
		}
		if current.x-1 >= 0 && int(maze[current.y][current.x-1]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x - 1, y: current.y})
		}
		if current.y-1 >= 0 && int(maze[current.y-1][current.x]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x, y: current.y - 1})
		}
	}
	return endPoints
}

func findAllEndpoints(start Point, maze []string) []Point {
	endPoints := []Point{}
	pending := []Point{start}
	for len(pending) != 0 {
		current := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if maze[current.y][current.x] == '9' {
			endPoints = append(endPoints, current)
		}
		if current.x+1 < len(maze[0]) && int(maze[current.y][current.x+1]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x + 1, y: current.y})
		}
		if current.y+1 < len(maze) && int(maze[current.y+1][current.x]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x, y: current.y + 1})
		}
		if current.x-1 >= 0 && int(maze[current.y][current.x-1]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x - 1, y: current.y})
		}
		if current.y-1 >= 0 && int(maze[current.y-1][current.x]-'0')-int(maze[current.y][current.x]-'0') == 1 {
			pending = append(pending, Point{x: current.x, y: current.y - 1})
		}
	}
	return endPoints
}

func findStartPoints(maze []string) []Point {
	points := []Point{}
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == '0' {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	maze := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	startPoints := findStartPoints(maze)
	score := 0
	for i := range startPoints {
		endPoints := findAllEndpoints(startPoints[i], maze)
		score += len(endPoints)
	}
	fmt.Println(score)
}
