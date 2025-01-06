package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Position struct {
	row    int
	column int
}

func collectPositions(farm []string, start Position) []Position {
	positions := []Position{}
	rowLimit := len(farm)
	columnLimit := len(farm[0])
	pending := []Position{start}
	visited := make(map[Position]bool)
	target := farm[start.row][start.column]
	for len(pending) > 0 {
		current := pending[0]
		pending = pending[1:]
		if _, ok := visited[current]; !ok {
			positions = append(positions, current)
			visited[current] = true
			if current.column+1 < columnLimit && farm[current.row][current.column+1] == target {
				pending = append(pending, Position{row: current.row, column: current.column + 1})
			}
			if current.column-1 >= 0 && farm[current.row][current.column-1] == target {
				pending = append(pending, Position{row: current.row, column: current.column - 1})
			}
			if current.row+1 < rowLimit && farm[current.row+1][current.column] == target {
				pending = append(pending, Position{row: current.row + 1, column: current.column})
			}
			if current.row-1 >= 0 && farm[current.row-1][current.column] == target {
				pending = append(pending, Position{row: current.row - 1, column: current.column})
			}
		}
	}
	return positions
}

func createGardenPlots(farm []string) [][]Position {
	gardenPlots := [][]Position{}
	visited := make(map[Position]bool)
	for r := range farm {
		for c := range farm[r] {
			p := Position{row: r, column: c}
			if _, ok := visited[p]; !ok {
				visited[p] = true
				positions := collectPositions(farm, p)
				gardenPlots = append(gardenPlots, positions)
				for p := range positions {
					visited[positions[p]] = true
				}
			}
		}
	}
	return gardenPlots
}

func countWalls(segment []int) int {
	walls := 1
	sort.Slice(segment, func(i, j int) bool {
		return segment[i] < segment[j]
	})
	for i := 0; i < len(segment)-1; i++ {
		if (segment[i+1] - segment[i]) > 1 {
			walls += 1
		}
	}
	return walls
}

type Wall struct {
	id        int
	direction int
}

func calculatePrice(farm []string, gardenPlot []Position) int {
	area := 0
	perimeter := 0
	kind := farm[gardenPlot[0].row][gardenPlot[0].column]
	rowLimit := len(farm)
	columnLimit := len(farm[0])
	verticalWalls := make(map[Wall][]int)
	horizontalWalls := make(map[Wall][]int)
	for i := range gardenPlot {
		area += 1
		if gardenPlot[i].column+1 >= columnLimit || farm[gardenPlot[i].row][gardenPlot[i].column+1] != kind {
			w := Wall{id: gardenPlot[i].column + 1, direction: 1}
			verticalWalls[w] = append(verticalWalls[w], gardenPlot[i].row)
			perimeter += 1
		}
		if gardenPlot[i].column-1 < 0 || farm[gardenPlot[i].row][gardenPlot[i].column-1] != kind {
			w := Wall{id: gardenPlot[i].column - 1, direction: -1}
			verticalWalls[w] = append(verticalWalls[w], gardenPlot[i].row)
			perimeter += 1
		}
		if gardenPlot[i].row+1 >= rowLimit || farm[gardenPlot[i].row+1][gardenPlot[i].column] != kind {
			w := Wall{id: gardenPlot[i].row + 1, direction: 1}
			horizontalWalls[w] = append(horizontalWalls[w], gardenPlot[i].column)
			perimeter += 1
		}
		if gardenPlot[i].row-1 < 0 || farm[gardenPlot[i].row-1][gardenPlot[i].column] != kind {
			w := Wall{id: gardenPlot[i].row - 1, direction: -1}
			horizontalWalls[w] = append(horizontalWalls[w], gardenPlot[i].column)
			perimeter += 1
		}
	}
	sides := 0
	for _, segment := range verticalWalls {
		sides += countWalls(segment)
	}
	for _, segment := range horizontalWalls {
		sides += countWalls(segment)
	}
	return area * sides
}

func main() {
	f, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(f)
	farm := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		farm = append(farm, row)
	}
	gardenPlots := createGardenPlots(farm)
	total := 0
	for i := range gardenPlots {
		p := calculatePrice(farm, gardenPlots[i])
		total += p
	}
	fmt.Println(total)
}
