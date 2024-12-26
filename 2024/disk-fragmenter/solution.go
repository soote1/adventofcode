package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Space struct {
	start  int
	length int
	id     int
}

func buildSpaces(diskImage string) ([]Space, []Space) {
	empty := []Space{}
	full := []Space{}
	readMode := "full"
	id := 0
	p := 0
	for i := range diskImage {
		length := int(diskImage[i] - '0')
		switch readMode {
		case "full":
			full = append(full, Space{start: p, length: length, id: id})
			readMode = "empty"
			id += 1
		case "empty":
			empty = append(empty, Space{start: p, length: length})
			readMode = "full"
		}
		p += length
	}
	return empty, full
}

func moveBlocks(empty, full []Space) []Space {
	i := 0
	j := len(full) - 1
	moved := []Space{}
	for len(empty) > 0 && len(full) > 0 && empty[i].start < full[j].start {
		if empty[i].length >= full[j].length {
			space := Space{start: empty[i].start, length: full[j].length, id: full[j].id}
			moved = append(moved, space)
			empty[i].length -= full[j].length
			empty[i].start += full[j].length
			full[j].length = 0
		} else {
			space := Space{start: empty[i].start, length: empty[i].length, id: full[j].id}
			moved = append(moved, space)
			full[j].length -= empty[i].length
			empty[i].length = 0
		}
		if empty[i].length == 0 {
			empty = empty[1:]
		}
		if full[j].length == 0 {
			full = full[:j]
			j = len(full) - 1
		}
	}
	return moved
}

func moveFiles(empty, full []Space) []Space {
	moved := []Space{}
	for i := len(full) - 1; i >= 0; i-- {
		for j := range empty {
			if empty[j].start > full[i].start {
				break
			}
			if full[i].length <= empty[j].length {
				space := Space{start: empty[j].start, length: full[i].length, id: full[i].id}
				moved = append(moved, space)
				empty[j].length -= full[i].length
				empty[j].start += full[i].length
				full[i].length = 0
				break
			}
		}
	}
	return moved
}

func calculateChecksum(full, moved []Space) int {
	i, j := 0, 0
	checksum := 0
	current := Space{}
	for i < len(full) {
		if i < len(full) && j < len(moved) {
			if full[i].start < moved[j].start {
				current = full[i]
				i += 1
			} else {
				current = moved[j]
				j += 1
			}
		} else {
			current = full[i]
			i += 1
		}
		if current.length == 0 {
			continue
		}
		k := current.start
		sum := 0
		for k < (current.start + current.length) {
			sum += k
			k += 1
		}
		checksum += (sum * current.id)

	}
	return checksum
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskImage := scanner.Text()
	empty, full := buildSpaces(diskImage)
	moved := moveFiles(empty, full)
	checksum := calculateChecksum(full, moved)
	fmt.Println(checksum)
}
