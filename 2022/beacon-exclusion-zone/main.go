package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Sensor struct {
	location      Point
	closestBeacon Beacon
}

type Beacon struct {
	location Point
}

type Segment struct {
	start Point
	end   Point
}

type ExclusionZone struct {
	rows map[int][]Segment
}

func calculateExclusionZone(sensors []Sensor) *ExclusionZone {
	var d int
	exclusionZone := ExclusionZone{rows: make(map[int][]Segment)}
	for i := 0; i < len(sensors); i++ {
		s := sensors[i]
		xdiff := int(math.Abs(float64(s.location.x - s.closestBeacon.location.x)))
		ydiff := int(math.Abs(float64(s.location.y - s.closestBeacon.location.y)))
		d = xdiff + ydiff
		for x, y := d-1, s.location.y-1; x > 0; y, x = y-1, x-1 {
			segment := Segment{
				start: Point{x: s.location.x - x},
				end:   Point{x: s.location.x + x},
			}
			exclusionZone.rows[y] = append(exclusionZone.rows[y], segment)
		}
		for x, y := d-1, s.location.y+1; x > 0; x, y = x-1, y+1 {
			segment := Segment{
				start: Point{x: s.location.x - x},
				end:   Point{x: s.location.x + x},
			}
			exclusionZone.rows[y] = append(exclusionZone.rows[y], segment)
		}
		segment := Segment{
			start: Point{x: s.location.x - d},
			end:   Point{x: s.location.x + d},
		}
		exclusionZone.rows[s.location.y] = append(
			exclusionZone.rows[s.location.y],
			segment,
		)

	}
	return &exclusionZone
}

func countExcludedCells(ez *ExclusionZone, s []Sensor, row int) int {
	beacons := make(map[Beacon]bool)
	count := 0
	segments := ez.rows[row]
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].start.x < segments[j].start.x
	})
	nonOverlappingSegments := []Segment{}
	currentSegment := segments[0]
	for i := 1; i < len(segments); i++ {
		if currentSegment.end.x < segments[i].start.x {
			nonOverlappingSegments = append(nonOverlappingSegments, currentSegment)
			currentSegment = segments[i]
		} else if currentSegment.end.x >= segments[i].start.x && currentSegment.end.x <= segments[i].end.x {
			currentSegment.end = segments[i].end
		}
		if i == len(segments)-1 {
			nonOverlappingSegments = append(nonOverlappingSegments, currentSegment)
		}
	}
	for _, segment := range nonOverlappingSegments {
		count += (segment.end.x - segment.start.x) + 1
	}
	for _, sensor := range s {
		if _, ok := beacons[sensor.closestBeacon]; !ok {
			if sensor.closestBeacon.location.y == row {
				count--
			}
			beacons[sensor.closestBeacon] = true
		}
	}

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

func parseInput(input []string) []Sensor {
	var offset int
	sensors := []Sensor{}
	for _, line := range input {
		sensor := Sensor{}
		beacon := Beacon{}
		offset = 12
		sensor.location.x = collectNumber(&offset, line)
		offset += 4
		sensor.location.y = collectNumber(&offset, line)
		offset += 25
		beacon.location.x = collectNumber(&offset, line)
		offset += 4
		beacon.location.y = collectNumber(&offset, line)
		sensor.closestBeacon = beacon
		sensors = append(sensors, sensor)

	}
	return sensors
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
	inputFile := os.Args[1]
	row, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("need to provide row number")
	}
	input := loadInput(inputFile)
	sensors := parseInput(input)
	exclusionZone := calculateExclusionZone(sensors)
	excludedCells := countExcludedCells(exclusionZone, sensors, row)
	fmt.Println(excludedCells)
}
