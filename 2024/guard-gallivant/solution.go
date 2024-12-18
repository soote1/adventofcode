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

type Step struct {
	position  Position
	direction Position
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

func findExit(start Position, direction Position, maze []string) []Step {
	currentPos := start
	currentDir := direction
	path := []Step{}
	visited := make(map[Step]bool)
	for isInsideBounds(currentPos, maze) {
		step := Step{position: currentPos, direction: currentDir}
		if _, ok := visited[step]; ok {
			return []Step{}
		}
		if maze[currentPos.y][currentPos.x] == '#' {
			currentPos.x, currentPos.y = currentPos.x-currentDir.x, currentPos.y-currentDir.y
			currentDir = getNextDir(currentDir)
		} else {
			visited[step] = true
			path = append(path, step)
			currentPos.x, currentPos.y = currentPos.x+currentDir.x, currentPos.y+currentDir.y
		}
	}
	if len(path) == 0 {
		return []Step{{position: start, direction: direction}}
	}
	return path
}

func getNextPos(current Position, direction Position) Position {
	return Position{x: current.x + direction.x, y: current.y + direction.y}
}

func isInsideBounds(position Position, maze []string) bool {
	return position.x >= 0 && position.x < len(maze[0]) && position.y >= 0 && position.y < len(maze)
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	maze := []string{}
	scanner := bufio.NewScanner(file)
	initialPosition := Position{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if x := strings.Index(line, "^"); x != -1 {
			initialPosition.x = x
			initialPosition.y = y
		}
		maze = append(maze, scanner.Text())
		y += 1
	}

	initialDirection := Position{x: 0, y: -1}
	path := findExit(initialPosition, initialDirection, maze)
	loopPositions := make(map[Position]bool)
	for i := range path {
		nextPos := getNextPos(path[i].position, path[i].direction)
		if _, ok := loopPositions[nextPos]; ok {
			continue
		}
		if isInsideBounds(nextPos, maze) && maze[nextPos.y][nextPos.x] != '#' && maze[nextPos.y][nextPos.x] != '^' {
			row := []rune(maze[nextPos.y])
			prevValue := row[nextPos.x]
			row[nextPos.x] = '#'
			maze[nextPos.y] = string(row)
			subPath := findExit(initialPosition, initialDirection, maze)
			if len(subPath) == 0 {
				loopPositions[nextPos] = true
			}
			row[nextPos.x] = prevValue
			maze[nextPos.y] = string(row)
		}

	}

	fmt.Println("solution", len(loopPositions))
}
