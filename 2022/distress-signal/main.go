package main

import (
	"bufio"
	"fmt"
	"os"
)

func countSortedPackets(packets []any) {
    for _, packet := range packets {
        fmt.Println(packet)
    }
}

func buildInt(runes []rune) int {
    factor := 1
    number := 0
    for i := len(runes)-1; i >= 0; i-- {
        number += int(runes[i]-'0')*factor
        factor *= 10
    }
    return number
}

func parseInput(input []string) []any {
    packets := []any{}

    for _, line := range input {
        if line == "" {
            continue
        }

        number := []rune{}
        var data []any
        prev := [][]any{}
        integer := 0

        for _, char := range line {
            switch char {
            case '[':
                if data != nil {
                    prev = append(prev, data)
                    data = []any{}
                } else {
                    data = []any{}
                }
            case ']':
                if len(number) > 0 {
                    integer = buildInt(number)
                    data = append(data, integer)
                    number = []rune{}
                }
                if len(prev) > 0 {
                    prev[len(prev)-1] = append(prev[len(prev)-1], data)
                    data = prev[len(prev)-1]
                    prev = prev[:len(prev)-1]
                }
            case ',':
                if len(number) > 0 {
                    integer = buildInt(number)
                    data = append(data, integer)
                }
                number = []rune{}
            default:
                number = append(number, char)
            }
        }
        packets = append(packets, data)
    }

    return packets
}

func loadInput(filename string) []string {
	input := []string{}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func main() {
    input := loadInput(os.Args[1])
    packets := parseInput(input)
    countSortedPackets(packets)
}
