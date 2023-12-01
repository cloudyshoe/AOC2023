package main

import (
	"fmt"
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

func DigitFromString(line string, reverse bool) int {
    value := 0
    lineLen := len(line)

    for i, char := range line {
        strChar := ""
        if reverse {
            strChar = string(line[lineLen-i-1])
        } else {
            strChar = string(char)
        }
        value, err := strconv.Atoi(strChar)
        if err == nil { return value }
    }

    return value
}

func DigitFromStringWithWords(line string, reverse bool) int {
    result := 0
    lineLen := len(line)

    for i, char := range line {
        strChar := ""
        if reverse {
            strChar = string(line[lineLen-i-1])
        } else {
            strChar = string(char)
        }

        if value := DigitFromString(strChar, reverse); value != 0 {
            return value
        }

        word := ""
        if reverse {
            word = line[lineLen-i-1:]
        } else {
            word = line[:i]
        }

        for key, value := range spelledOut {
            if strings.Contains(word, key) {
                return value
            }
        }
    }
    return result
}


func PartOne(arr *[]string) int {
	input := *arr
	result := 0

	for _, line := range input {
        result += DigitFromString(line, false) * 10
        result += DigitFromString(line, true)
	}

	return result
}

func PartTwo(arr *[]string) int {
	input := *arr
	result := 0

	for _, line := range input {
        result += DigitFromStringWithWords(line, false) * 10
        result += DigitFromStringWithWords(line, true)
	}

	return result }

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(&input)
	fmt.Printf("Part One: %d\n", partOneResult)

	partTwoResult := PartTwo(&input)
	fmt.Printf("Part two: %d\n", partTwoResult)
}
