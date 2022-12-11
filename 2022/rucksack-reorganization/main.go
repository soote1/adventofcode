package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func sumPriorities(rucksacks []string) (int, int) {
	sum := 0
	groupPrioritySum := 0
	lowerMask := 96
	upperMask := 38
	var group []map[byte]bool
	for _, items := range rucksacks {
		aItems := make(map[byte]bool)
		bItems := make(map[byte]bool)
		allItems := make(map[byte]bool)
		compartmentBoundary := len(items) / 2
		for i := 0; i < len(items); i++ {
			if i < compartmentBoundary {
				aItems[items[i]] = true
			} else {
				bItems[items[i]] = true
			}
			allItems[items[i]] = true
		}

		for item := range aItems {
			if _, ok := bItems[item]; ok {
				if item >= 97 {
					sum += int(item) - lowerMask
				} else {
					sum += int(item) - upperMask
				}
			}
		}

		group = append(group, allItems)
		if len(group) == 3 {
			for item := range group[0] {
				_, bHasItem := group[1][item]
				_, cHasItem := group[2][item]

				if bHasItem && cHasItem {
					if item >= 97 {
						groupPrioritySum += int(item) - lowerMask
					} else {
						groupPrioritySum += int(item) - upperMask
					}
				}
			}
			group = []map[byte]bool{}
		}
	}
	return sum, groupPrioritySum
}

func main() {
	inputFileName := os.Args[1]

	bytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	content := string(bytes)
	answer1, answer2 := sumPriorities(strings.Split(content, "\n"))
	fmt.Println(answer1)
	fmt.Println(answer2)
}
