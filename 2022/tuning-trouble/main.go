package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func allDifferent(s []int) bool {
	set := make(map[int]bool)

	for _, item := range s {
		set[item] = true
	}

	return len(s) == len(set)
}

func getChunkLength(stream string, markerLength int) int {
	mark := []int{}

	for i, b := range stream {
		b := int(b)
		mark = append(mark, b)
		if len(mark) == markerLength {
			if allDifferent(mark) {
				return i + 1
			} else {
				mark = mark[1:]
			}
		}
	}

	return 0
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	packetLength := getChunkLength(content, 4)
	messageLength := getChunkLength(content, 14)
	fmt.Printf("%v\n", packetLength)
	fmt.Printf("%v\n", messageLength)
}
