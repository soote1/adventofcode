package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getNextDir(currentDir Position) Position {
	if currentDir.x == 0 && currentDir.y == -1 {
		return Position{x: 1, y: 0}
	}
	if currentDir.x == 1 && currentDir.y == 0 {
		return Position{x: 0, y: 1}
	}
	if currentDir.x == 0 && currentDir.y == 1 {
		return Position{x: -1, y: 0}
	}
	if currentDir.x == -1 && currentDir.y == 0 {
		return Position{x: 0, y: -1}
	}
	panic("invalid direction")
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	maze := []string{}
	scanner := bufio.NewScanner(file)
	currentPos := Position{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if x := strings.Index(line, "^"); x != -1 {
			currentPos.x = x
			currentPos.y = y
		}
		maze = append(maze, scanner.Text())
		y += 1
	}

	currentDir := Position{x: 0, y: -1}
	visited := make(map[Position]bool)
	for currentPos.x >= 0 && currentPos.x < len(maze[0]) && currentPos.y >= 0 && currentPos.y < len(maze) {
		if maze[currentPos.y][currentPos.x] == '#' {
			currentPos.x, currentPos.y = currentPos.x-currentDir.x, currentPos.y-currentDir.y
			currentDir = getNextDir(currentDir)
		} else {
			visited[currentPos] = true
			currentPos.x, currentPos.y = currentPos.x+currentDir.x, currentPos.y+currentDir.y
		}
	}
	fmt.Println(len(visited))
}
