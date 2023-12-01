package main

import (
	"fmt"
    "math"
	"os"
	"strconv"
	"strings"
)

var spelledOut = map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
    "four":  4,
    "five":  5,
    "six":   6,
    "seven": 7,
    "eight": 8,
    "nine":  9,
}

var digits = "123456789"

func DigitsFromString(line string) (int, int, int, int) {
    firstIndex := strings.IndexAny(line,digits)
    if firstIndex < 0 { return -1, -1, -1, -1 }
    first, _ := strconv.Atoi(string(line[firstIndex]))

    lastIndex := strings.LastIndexAny(line,digits)
    last, _ := strconv.Atoi(string(line[lastIndex]))

    return first, firstIndex, last, lastIndex
}

func DigitsFromStringWithWords(line string) int {
    lineLen := len(line)
    firstFound := false
    lastFound := false

    first, firstIndex, last, lastIndex := DigitsFromString(line)
    if firstIndex < 2 && firstIndex >= 0 { firstFound = true }
    if lastIndex > lineLen-3 { lastFound = true }
    if firstIndex == -1 { firstIndex = math.MaxInt }

    for key, value := range spelledOut {
        tmpFi := strings.Index(line, key)
        tmpLi := strings.LastIndex(line, key)
        if !firstFound && tmpFi < firstIndex && tmpFi >= 0 {
               firstIndex = tmpFi
               first = value
           }
        if !lastFound && tmpLi > lastIndex {
               lastIndex = tmpLi
               last = value
           }
    }

    return first*10+last
}

func PartOne(arr *[]string) int {
	input := *arr
	result := 0

	for _, line := range input {
        first, _, last, _ := DigitsFromString(line)
        result += first*10+last
	}

	return result
}

func PartTwo(arr *[]string) int {
	input := *arr
	result := 0

	for _, line := range input {
        result += DigitsFromStringWithWords(line)
	}

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-2]), "\n")

	partOneResult := PartOne(&input)
	fmt.Printf("Part One: %d\n", partOneResult)

	partTwoResult := PartTwo(&input)
	fmt.Printf("Part two: %d\n", partTwoResult)
}
