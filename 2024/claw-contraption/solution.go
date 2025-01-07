package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Button struct {
	cost int
	x    int
	y    int
}

type Prize struct {
	x int
	y int
}

type Game struct {
	a     Button
	b     Button
	prize Prize
}

func parseButtonLine(line string) Button {
	lineItems := strings.Fields(line)
	plusIndex := strings.Index(lineItems[2], "+")
	commaIndex := strings.Index(lineItems[2], ",")
	x, err := strconv.Atoi(lineItems[2][plusIndex+1 : commaIndex])
	checkError(err)
	plusIndex = strings.Index(lineItems[3], "+")
	y, err := strconv.Atoi(lineItems[3][plusIndex+1:])
	checkError(err)
	cost := 0
	if lineItems[1][0] == 'A' {
		cost = 3
	} else {
		cost = 1
	}
	return Button{x: x, y: y, cost: cost}
}

func parsePrizeLine(line string) Prize {
	lineItems := strings.Fields(line)
	equalIndex := strings.Index(lineItems[1], "=")
	commaIndex := strings.Index(lineItems[1], ",")
	x, err := strconv.Atoi(lineItems[1][equalIndex+1 : commaIndex])
	checkError(err)
	equalIndex = strings.Index(lineItems[2], "=")
	y, err := strconv.Atoi(lineItems[2][equalIndex+1:])
	checkError(err)
	return Prize{x: x, y: y}
}

func findOptimalCostToWin(combinations [][]int, game Game) int {
	cost := 0
	for i := range combinations {
		aAmount, bAmount := combinations[i][0], combinations[i][1]
		x := (game.a.x * aAmount) + (game.b.x * bAmount)
		y := (game.a.y * aAmount) + (game.b.y * bAmount)
		if x == game.prize.x && y == game.prize.y {
			c := game.a.cost*aAmount + game.b.cost*bAmount
			if cost == 0 || c < cost {
				cost = c
			}
		}
	}
	return cost
}

func main() {
	f, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(f)
	games := []Game{}
	for scanner.Scan() {
		aButton := parseButtonLine(scanner.Text())
		scanner.Scan()
		bButton := parseButtonLine(scanner.Text())
		scanner.Scan()
		prize := parsePrizeLine(scanner.Text())
		games = append(games, Game{a: aButton, b: bButton, prize: prize})
		scanner.Scan()
	}
	combinations := [][]int{}
	for i := range 100 {
		for j := range 100 {
			combinations = append(combinations, []int{i, j})
		}
	}
	total := 0
	for i := range games {
		c := findOptimalCostToWin(combinations, games[i])
		if c != 0 {
			total += c
		}
	}
	fmt.Println(total)
}
