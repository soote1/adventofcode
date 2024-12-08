package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func isOutOfBounds(r, c, nc, nr int) bool {
	return c < 0 || c >= nc || r < 0 || r >= nr
}

func searchLine(data []string) int {
	ncol := len(data[0])
	nrows := len(data)
	found := 0
	fmt.Println(data)
	for i := range data {
		for j := range data[i] {
			valid := 0
			if data[i][j] == 'X' {
				target := "MAS"
				// right
				i1, j1 := i, j+1
				// bottom right
				i2, j2 := i+1, j+1
				// bottom
				i3, j3 := i+1, j
				// bottom left
				i4, j4 := i+1, j-1
				// left
				i5, j5 := i, j-1
				// upper left
				i6, j6 := i-1, j-1
				// up
				i7, j7 := i-1, j
				// upper right
				i8, j8 := i-1, j+1
				for k := range target {
					if !isOutOfBounds(i1, j1, ncol, nrows) {
						if data[i1][j1] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							j1 += 1
						} else {
							i1 = -666
							j1 = -666
						}

					}
					if !isOutOfBounds(i2, j2, ncol, nrows) {
						if data[i2][j2] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i2 += 1
							j2 += 1
						} else {
							i2 = -666
							j2 = -666
						}
					}
					if !isOutOfBounds(i3, j3, ncol, nrows) {
						if data[i3][j3] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i3 += 1
						} else {
							i3 = -666
							j3 = -666
						}
					}
					if !isOutOfBounds(i4, j4, ncol, nrows) {
						if data[i4][j4] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i4 += 1
							j4 -= 1
						} else {
							i4 = -666
							j4 = -666
						}
					}
					if !isOutOfBounds(i5, j5, ncol, nrows) {
						if data[i5][j5] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							j5 -= 1
						} else {
							i5 = -666
							j5 = -666
						}
					}
					if !isOutOfBounds(i6, j6, ncol, nrows) {
						if data[i6][j6] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i6 -= 1
							j6 -= 1
						} else {
							i6 = -666
							j6 = -666
						}
					}
					if !isOutOfBounds(i7, j7, ncol, nrows) {
						if data[i7][j7] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i7 -= 1
						} else {
							i7 = -666
							j7 = -666
						}
					}
					if !isOutOfBounds(i8, j8, ncol, nrows) {
						if data[i8][j8] == target[k] {
							if target[k] == 'S' {
								valid += 1
							}
							i8 -= 1
							j8 += 1
						} else {
							i8 = -666
							j8 = -666
						}
					}
				}

			}
			found += valid
		}
	}
	return found
}

func isValidCorner(ui, uj, bi, bj int, data []string) bool {
	return data[ui][uj] == 'S' && data[bi][bj] == 'M' || data[ui][uj] == 'M' && data[bi][bj] == 'S'
}

func searchCross(data []string) int {
	ncol := len(data[0])
	nrows := len(data)
	found := 0
	fmt.Println(data)
	for i := range data {
		for j := range data[i] {
			valid := 0
			if data[i][j] == 'A' {
				if !isOutOfBounds(i-1, j-1, ncol, nrows) && !isOutOfBounds(i+1, j+1, ncol, nrows) {
					if isValidCorner(i-1, j-1, i+1, j+1, data) {
						if !isOutOfBounds(i+1, j-1, ncol, nrows) && !isOutOfBounds(i-1, j+1, ncol, nrows) {
							if isValidCorner(i+1, j-1, i-1, j+1, data) {
								valid += 1
							}
						}
					}
				}
			}
			found += valid
		}
	}
	return found

}

func main() {
	file, e := os.Open("input.txt")
	defer file.Close()
	checkError(e)

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	foundLine := searchLine(data)
	foundCross := searchCross(data)

	fmt.Println(foundLine)
	fmt.Println(foundCross)
}
