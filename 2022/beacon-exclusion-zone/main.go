package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

type ExclusionZone struct {
	rows map[int][]Point
}

func calculateExclusionZone(sensors []Point, beacons []Point) *ExclusionZone {
	exclusionZone := ExclusionZone{}
	for i := 0; i < len(sensors); i++ {
		// TODO: calculate distance to closest beacon
		// TODO: Save corresponding range for upper rows (from sensor coordinates)
		// TODO: Save corresponding range for lower rows (from sensor coordinates)
		// TODO: Save corresponding range for current row (from sensor coordinates)
	}
	return &exclusionZone
}

func countExcludedCells(ez *ExclusionZone, beacons []Point, row int) int {
	count := 0
	// TODO: calculate min Point
	// TODO: calculate max Point
	// TODO: count = max-min
	// TODO: remove beacons in row from count
	return count
}

func collectNumber(offset *int, data string) int {
	buf := ""
	for *offset < len(data) && data[*offset] != ',' && data[*offset] != ':' {
		buf += string(data[*offset])
		*offset++
	}
	number, err := strconv.Atoi(buf)
	if err != nil {
		panic(err)
	}
	return number
}

func parseInput(input []string) ([]Point, []Point) {
	var offset int
	sensors := []Point{}
	beacons := []Point{}
	for _, line := range input {
		sensor := Point{}
		beacon := Point{}
		offset = 12
		sensor.x = collectNumber(&offset, line)
		offset += 4
		sensor.y = collectNumber(&offset, line)
		offset += 25
		beacon.x = collectNumber(&offset, line)
		offset += 4
		beacon.y = collectNumber(&offset, line)
		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)

	}
	return sensors, beacons
}

func loadInput(fileName string) []string {
	input := []string{}

	f, err := os.Open(fileName)
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
	sensors, beacons := parseInput(input)
	fmt.Println(sensors)
	fmt.Println(beacons)
}
