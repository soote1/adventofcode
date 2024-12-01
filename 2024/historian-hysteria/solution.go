package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	Id int
}

type ById []Location

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getSimilarityScore(left, right []Location) int {
	occurrences := make(map[int]int)
	for i := range right {
		if _, ok := occurrences[right[i].Id]; ok {
			occurrences[right[i].Id] += 1
		} else {
			occurrences[right[i].Id] = 1
		}
	}
	score := 0
	for i := range left {
		multiplier, _ := occurrences[left[i].Id]
		score += left[i].Id * multiplier
	}
	return score

}

func main() {
	var a []Location
	var b []Location
	file, err := os.Open("input.txt")
	defer file.Close()
	checkError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		id, err := strconv.Atoi(ids[0])
		checkError(err)
		a = append(a, Location{Id: id})
		id, err = strconv.Atoi(ids[1])
		checkError(err)
		b = append(b, Location{Id: id})

	}
	checkError(scanner.Err())
	sort.Sort(ById(a))
	sort.Sort(ById(b))
	sum := 0
	for i := range a {
		sum += int(math.Abs(float64(a[i].Id) - float64(b[i].Id)))
	}
	fmt.Println(sum)
	similarityScore := getSimilarityScore(a, b)
	fmt.Println(similarityScore)

}
