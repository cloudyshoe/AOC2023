package main

import (
    "os"
    "fmt"
	"slices"
	"strconv"
    "strings"
)

func Process(history []int) int {
	zedCount := 0

	for _, val := range history {
		if val == 0 { zedCount++ }
	}

	if zedCount == len(history) { return 0 }

	last := len(history)-1
	diffs := make([]int, last)
	for i := 0; i < last; i++ {
		diffs[i] = history[i+1] - history[i]	
	}

	return history[last] + Process(diffs)
}


func PartOne(arr *[]string) int  {
    input := *arr
    result := 0

	for _, line := range input {
		hstStr := strings.Split(line," ")
		history := make([]int, len(hstStr))
		for i, v := range hstStr {
			history[i], _ = strconv.Atoi(v)
		}
		next := Process(history)
		result += next
	}

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    result := 0

	for _, line := range input {
		hstStr := strings.Split(line," ")
		history := make([]int, len(hstStr))
		for i, v := range hstStr {
			history[i], _ = strconv.Atoi(v)
		}
		slices.Reverse(history)
		prev := Process(history)
		result += prev
	}

    return result
}

func main () {
    inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-1]), "\n")

    partOneResult := PartOne(&input)
    fmt.Printf("Part One Result: %d\n", partOneResult)

    partTwoResult := PartTwo(&input)
    fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
