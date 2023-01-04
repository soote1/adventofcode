package main

import (
	"bufio"
	"os"
    "fmt"
)

func diff(p1 any, p2 any) int {
    state := 0

    if x, ok := p1.(int); ok {
        if y, ok := p2.(int); ok {
            return x - y
        } else {
            p1 = []any{x}
        }
    } else {
        if y, ok := p2.(int); ok {
            p2 = []any{y}
        }
    }


    s1, ok := p1.([]any)
    if !ok {
        panic("can convert p1 to []any")
    }

    s2, ok := p2.([]any)
    if !ok {
        panic("can convert p2 to []any")
    }

    for i := 0; i < len(s1); i++ {
        if i == len(s2) {
            break
        }
        state = diff(s1[i], s2[i])
        if state < 0 || state > 0 {
            break
        }
    }

    if state == 0 {
        if len(s1) < len(s2) {
            state = -1
        }
        if len(s1) > len(s2) {
            state = 1
        }
    }

    return state
}

func collectSortedPacketIndices(packets []any) []int {
    sortedPacketIndices := []int{}
    for i, j := 1, 1; i < len(packets); i, j = i+2, j+1 {
        if diff(packets[i-1], packets[i]) < 0 {
            sortedPacketIndices = append(sortedPacketIndices, j)
        }
    }
    return sortedPacketIndices
}

func sortPackets(packets []any) []any {
    var swaps int
    var aux any
    for {
        swaps = 0
        for i := 0; i < len(packets)-1; i++ {
            if diff(packets[i], packets[i+1]) > 0 {
                aux = packets[i]
                packets[i] = packets[i+1]
                packets[i+1] = aux
                swaps++
            }
        }

        if swaps == 0 {
            break
        }
    }

    return packets
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
    sortedPacketIndices := collectSortedPacketIndices(packets)
    sum := 0
    for _, i := range sortedPacketIndices {
        sum += i
    }
    fmt.Println(sum)
    start := []any{[]any{2}}
    end := []any{[]any{6}}
    packets = append(packets, start)
    packets = append(packets, end)
    sortedPackets := sortPackets(packets)
    dividers := []int{}
    for i, packet := range sortedPackets {
        if diff(packet, start) == 0 || diff(packet, end) == 0 {
            dividers = append(dividers, i+1)
        }
    }
    fmt.Println(dividers[0]*dividers[1])
}
