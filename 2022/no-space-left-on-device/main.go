package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	program  string
	argument string
	output   []string
}

type FileSystemNode struct {
	nodeType string
	name     string
	size     int
	children map[string]*FileSystemNode
	parent   *FileSystemNode
}

func parseOutput(output string) *FileSystemNode {
	var node FileSystemNode
	parts := strings.Split(output, " ")
	switch parts[0] {
	case "dir":
		node.nodeType = parts[0]
	default:
		node.nodeType = "file"
		node.size, _ = strconv.Atoi(parts[0])
	}
	node.name = parts[1]
	node.children = make(map[string]*FileSystemNode)

	return &node
}

func buildFileSystem(commands []Command) *FileSystemNode {
	var rootNode *FileSystemNode
	var currentNode *FileSystemNode

	for i := len(commands) - 1; i >= 0; i-- {
		switch commands[i].program {
		case "cd":
			switch commands[i].argument {
			case "..":
				currentNode = currentNode.parent
			case "/":
				if rootNode == nil {
					rootNode = &FileSystemNode{
						name:     commands[i].argument,
						nodeType: "dir",
						children: make(map[string]*FileSystemNode),
					}
				}
				currentNode = rootNode
			default:
				newNode := &FileSystemNode{
					name:     commands[i].argument,
					nodeType: "dir",
					parent:   currentNode,
					children: make(map[string]*FileSystemNode),
				}
				currentNode.children[newNode.name] = newNode
				currentNode = newNode
			}
		case "ls":
			for _, output := range commands[i].output {
				if output == "" {
					continue
				}
				node := parseOutput(output)
				currentNode.children[node.name] = node
			}
		}
	}

	return rootNode
}

func sumDirSizes(fs *FileSystemNode, maxAllowedSize int) int {
	visited := make(map[*FileSystemNode]bool)
	pending := []*FileSystemNode{fs}
	currentNode := &FileSystemNode{}
	sum := 0

	for len(pending) > 0 {
		currentNode = pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if _, ok := visited[currentNode]; !ok {
			for _, node := range currentNode.children {
				pending = append(pending, node)
			}
			visited[currentNode] = true
			if currentNode.nodeType == "dir" && currentNode.size <= maxAllowedSize {
				sum += currentNode.size
			}
		}
	}
	return sum
}

func setDirSize(node *FileSystemNode) {
	size := 0
	for _, children := range node.children {
		if children.nodeType == "dir" {
			setDirSize(children)
			size += children.size
		} else {
			size += children.size
		}
	}

	node.size = size
}

func parseInput(input []string) []Command {
	var commands []Command
	var command Command
	var output []string

	for i := len(input) - 1; i >= 0; i-- {
		parts := strings.Split(input[i], " ")
		if parts[0] == "$" {
			if len(parts) == 3 {
				command = Command{
					program:  parts[1],
					argument: parts[2],
					output:   output,
				}
			} else {
				command = Command{program: parts[1], output: output}
			}
			output = output[len(output):]
			commands = append(commands, command)
		} else {
			output = append(output, input[i])
		}
	}
	return commands
}

func getDirToDelete(totalSpaceAvailable int, target int, fs *FileSystemNode) *FileSystemNode {
	spaceNeeded := target - (totalSpaceAvailable - fs.size)
	pending := []*FileSystemNode{fs}
	var currentNode *FileSystemNode
	nodeToDelete := fs

	for len(pending) > 0 {
		currentNode = pending[0]
		pending = pending[1:]

		size := currentNode.size
		if size >= spaceNeeded && size < nodeToDelete.size && currentNode.nodeType == "dir" {
			nodeToDelete = currentNode
		}

		for _, children := range currentNode.children {
			pending = append(pending, children)
		}
	}

	return nodeToDelete
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	input := strings.Split(string(bytes), "\n")
	commands := parseInput(input)
	fs := buildFileSystem(commands)
	setDirSize(fs)
	// sum := sumDirSizes(fs, 100000)
	answer := getDirToDelete(70000000, 30000000, fs)
	fmt.Println(answer.size)
}
