package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Item struct {
	value    *Node
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Node struct {
	index     int
	row       int
	column    int
	level     int
	elevation int
	label     rune
}

func canMove(a *Node, b *Node) bool {
	return b.elevation-a.elevation <= 1
}

func addNeighbor(pq *PriorityQueue, node *Node, dist int, dists *[]int) {
	node.level = dist
	heap.Push(pq, &Item{value: node, priority: dist})
	(*dists)[node.index] = dist
}

// use dijkstra algorithm to find shortest path from start to end
func minPathLen(matrix [][]*Node, start *Node, end *Node) int {
	var neighbor *Node
	var node *Node
	var minLen int = math.MaxInt
	var newDist int

	pending := &PriorityQueue{&Item{value: start}}
	heap.Init(pending)

	distances := make([]int, len(matrix)*len(matrix[0]))
	for i := 0; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}
	distances[start.index] = 0

	for len(*pending) > 0 {
		node = heap.Pop(pending).(*Item).value
		newDist = node.level + 1
		// add neighbors
		if node.column-1 >= 0 {
			neighbor = matrix[node.row][node.column-1]
			if canMove(node, neighbor) && newDist < distances[neighbor.index] {
				addNeighbor(pending, neighbor, newDist, &distances)
			}
		}
		if node.row-1 >= 0 {
			neighbor = matrix[node.row-1][node.column]
			if canMove(node, neighbor) && newDist < distances[neighbor.index] {
				addNeighbor(pending, neighbor, newDist, &distances)
			}
		}
		if node.column+1 < len(matrix[0]) {
			neighbor = matrix[node.row][node.column+1]
			if canMove(node, neighbor) && newDist < distances[neighbor.index] {
				addNeighbor(pending, neighbor, newDist, &distances)
			}
		}
		if node.row+1 < len(matrix) {
			neighbor = matrix[node.row+1][node.column]
			if canMove(node, neighbor) && newDist < distances[neighbor.index] {
				addNeighbor(pending, neighbor, newDist, &distances)
			}
		}

		if node == end {
			minLen = node.level
			break
		}
	}

	return minLen
}

func findStartEnd(matrix [][]*Node) (*Node, *Node) {
	var start *Node
	var end *Node

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j].label == 'S' {
				start = matrix[i][j]
			}

			if matrix[i][j].label == 'E' {
				end = matrix[i][j]
			}
		}

		if start != nil && end != nil {
			break
		}
	}

	return start, end
}

func findStartingPoints(matrix [][]*Node) []*Node {
	points := []*Node{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j].label == 'S' || matrix[i][j].label == 'a' {
				points = append(points, matrix[i][j])
			}
		}
	}
	return points
}

func shortestPathLen(matrix [][]*Node) int {
	start, end := findStartEnd(matrix)
	return minPathLen(matrix, start, end)
}

func parseInput(input []string) [][]*Node {
	matrix := [][]*Node{}
	minMask := 97
	maxMask := 122
	k := 0
	for i, line := range input {
		row := []*Node{}
		for j, char := range line {
			node := Node{}
			switch char {
			case 'S':
				node.elevation = minMask - minMask
			case 'E':
				node.elevation = maxMask - minMask
			default:
				node.elevation = int(char) - minMask
			}
			node.index = k
			node.row = i
			node.column = j
			node.label = char
			row = append(row, &node)
			k++
		}
		matrix = append(matrix, row)
	}
	return matrix
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
	//matrix := parseInput(input)
	//shortestPathLen := shortestPathLen(matrix)
	//fmt.Println(shortestPathLen)
	//shortestStartingPoint := shortestStartingPoint(parseInput(input))
	matrix := parseInput(input)
	length := 0
	minLength := math.MaxInt
	startingPoints := findStartingPoints(matrix)
	for _, start := range startingPoints {
		// I know, I shouldn't need to reload matrix and find end on every run
		// but my original implementation used pointers and this was the quickest
		// change to solve part 2 :P
		matrix := parseInput(input)
		_, end := findStartEnd(matrix)
		length = minPathLen(matrix, start, end)
		if length < minLength {
			minLength = length
		}
	}
	fmt.Println(minLength)
}
