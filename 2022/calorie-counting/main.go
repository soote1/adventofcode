package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MaxStack struct {
	capacity int
	elements *list.List
}

func NewMaxStack(capacity int) *MaxStack {
	return &MaxStack{capacity: capacity, elements: list.New()}
}

func (ms *MaxStack) Push(element int) {
	aux := []any{}

	for ms.elements.Back() != nil && ms.elements.Back().Value.(int) < element {
		back := ms.elements.Back().Value.(int)
		aux = append(aux, back)
		ms.elements.Remove(ms.elements.Back())
	}

	aux = append(aux, element)

	for i := len(aux) - 1; ms.elements.Len() < ms.capacity && i >= 0; i-- {
		ms.elements.PushBack(aux[i])
	}
}

func (ms *MaxStack) Pop() int {
	back := ms.elements.Back().Value.(int)
	ms.elements.Remove(ms.elements.Back())
	return back
}

func (ms *MaxStack) Peak() int {
	return ms.elements.Back().Value.(int)
}

func maxCalories(inventory []string) int {
	currentCount := 0
	maxValues := NewMaxStack(3)
	maxCaloriesSum := 0

	for _, item := range inventory {
		if item == "" {
			maxValues.Push(currentCount)
			currentCount = 0
		} else {
			calories, _ := strconv.Atoi(item)
			currentCount += calories
		}
	}

	for i := 0; i < 3; i++ {
		maxCaloriesSum += maxValues.Pop()
	}

	return maxCaloriesSum
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	answer := maxCalories(strings.Split(content, "\n"))
	fmt.Println(answer)
}
