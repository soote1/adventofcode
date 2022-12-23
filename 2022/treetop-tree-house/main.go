package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MaxNumbersDict map[int]map[int]int

func calcMaxNumbers(grid [][]int) []*MaxNumbersDict {
    maxNumsUp := make(MaxNumbersDict)
    maxNumsLeft := make(MaxNumbersDict)
    maxNumsRight := make(MaxNumbersDict)
    maxNumsDown := make(MaxNumbersDict)

    maxNums := []*MaxNumbersDict{
        &maxNumsLeft,
        &maxNumsUp,
        &maxNumsRight,
        &maxNumsDown,
    }

    for i, x := 0, len(grid)-1; i < len(grid); i, x = i+1, x-1 {
        for j, y := 0, len(grid[0])-1; j < len(grid[i]); j, y = j+1, y-1 {
            if _, ok := maxNumsUp[j]; !ok {
                maxNumsUp[j] = make(map[int]int)
            }
            if _, ok := maxNumsUp[j][i]; !ok {
                maxNumsUp[j][i] = 0
            }

            if _, ok := maxNumsLeft[i]; !ok {
                maxNumsLeft[i] = make(map[int]int)
            }
            if _, ok := maxNumsLeft[i][j]; !ok {
                maxNumsLeft[i][j] = 0
            }

            if _, ok := maxNumsRight[x]; !ok {
                maxNumsRight[x] = make(map[int]int)
            }
            if _, ok := maxNumsRight[x][y]; !ok {
                maxNumsRight[x][y] = 0
            }

            if _, ok := maxNumsDown[y]; !ok {
                maxNumsDown[y] = make(map[int]int)
            }
            if _, ok := maxNumsDown[y][x]; !ok {
                maxNumsDown[y][x] = 0
            }

            if i > 0 { // check up
                if grid[i][j] >= maxNumsUp[j][i-1] {
                    maxNumsUp[j][i] = grid[i][j]
                } else {
                    maxNumsUp[j][i] = maxNumsUp[j][i-1]
                }
            } else {
                maxNumsUp[j][i] = grid[i][j]
            }

            if j > 0 { // check left
                if grid[i][j] >= maxNumsLeft[i][j-1] {
                    maxNumsLeft[i][j] = grid[i][j]
                } else {
                    maxNumsLeft[i][j] = maxNumsLeft[i][j-1]
                }
            } else {
                maxNumsLeft[i][j] = grid[i][j]
            }

            if x < len(grid)-1 { // check down
                if grid[x][y] >= maxNumsDown[y][x+1] {
                    maxNumsDown[y][x] = grid[x][y]
                } else {
                    maxNumsDown[y][x] = maxNumsDown[y][x+1]
                }
            } else {
                maxNumsDown[y][x] = grid[x][y]
            }

            if y < len(grid[0])-1 { // check right
                if grid[x][y] >= maxNumsRight[x][y+1] {
                    maxNumsRight[x][y] = grid[x][y]
                } else {
                    maxNumsRight[x][y] = maxNumsRight[x][y+1]
                }
            } else {
                maxNumsRight[x][y] = grid[x][y]
            }
        }
    }

    return maxNums
}

func visible(grid [][]int, maxNumbers []*MaxNumbersDict) int {
    visible := 0

    maxNumbersLeft := *maxNumbers[0]
    maxNumbersUp := *maxNumbers[1]
    maxNumbersRight := *maxNumbers[2]
    maxNumbersDown := *maxNumbers[3]

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1 {
                visible++
            } else {
                if grid[i][j] > maxNumbersLeft[i][j-1] ||
                    grid[i][j] > maxNumbersUp[j][i-1] ||
                    grid[i][j] > maxNumbersRight[i][j+1] ||
                    grid[i][j] > maxNumbersDown[j][i+1] {
                    visible++
                }
            }
        }
    }

    return visible
}

func maxScenicScore(grid [][]int, maxNumbers []*MaxNumbersDict) int {
    maxScenicScore := 0
    leftVwDs := 0
    upVwDs := 0
    rightVwDs := 0
    downVwDs := 0
    maxNumbersLeft := *maxNumbers[0]
    maxNumbersUp := *maxNumbers[1]
    maxNumbersRight := *maxNumbers[2]
    maxNumbersDown := *maxNumbers[3]

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0])-1; j++ {
            if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1 {
                continue
            } else {
                // calculate viewing distance for each direction
                if grid[i][j] > maxNumbersLeft[i][j-1] {
                    leftVwDs = j
                } else {
                    for k := j-1; k >= 0; k-- {
                        if grid[i][k] >= grid[i][j] {
                            leftVwDs = j-k
                            break
                        }
                    }
                }
                if grid[i][j] > maxNumbersUp[j][i-1] {
                    upVwDs = i
                } else {
                    for k := i-1; k >= 0; k-- {
                        if grid[k][j] >= grid[i][j] {
                            upVwDs = i-k
                            break
                        }
                    }
                }
                if grid[i][j] > maxNumbersRight[i][j+1] {
                    rightVwDs = len(grid[0])-1-j
                } else {
                    for k := j+1; k < len(grid[0]); k++ {
                        if grid[i][k] >= grid[i][j] {
                            rightVwDs = k-j
                            break
                        }
                    }
                }
                if grid[i][j] > maxNumbersDown[j][i+1] {
                    downVwDs = len(grid)-1-i
                } else {
                    for k := i+1; k < len(grid); k++ {
                        if grid[k][j] >= grid[i][j] {
                            downVwDs = k-i
                            break
                        }
                    }
                }
                scenicScore := leftVwDs*upVwDs*rightVwDs*downVwDs
                if scenicScore > maxScenicScore {
                    maxScenicScore = scenicScore
                }
            }
        }
    }

    return maxScenicScore
}

func generateGrid(lines []string) [][]int {
    grid := make([][]int, len(lines)-1)
    for i, line := range lines {
        if line == "" {
            continue
        }
        cols := make([]int, len(line))
        for j := 0; j < len(line); j++ {
            cols[j], _ = strconv.Atoi(string(line[j]))
        }
        grid[i] = cols
    }
    return grid
}

func loadFile(fileName string) []string {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Can't open input file: %v", err)
	}

	return strings.Split(string(bytes), "\n")
}

func main() {
	inputFileName := os.Args[1]
    content := loadFile(inputFileName)
    grid := generateGrid(content)
    maxNumbers := calcMaxNumbers(grid)
    visibleTreesCount := visible(grid, maxNumbers)
    maxScenicScore := maxScenicScore(grid, maxNumbers)
    fmt.Printf("%v\n", visibleTreesCount)
    fmt.Printf("%v\n", maxScenicScore)
}
