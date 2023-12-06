package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(arr *[]string) int {
	input := *arr
	result := 1

	var times []int
	var distances []int

	for _, time := range strings.Fields(input[0])[1:] {
		newTime, _ := strconv.Atoi(time)
		times = append(times, newTime)
	}
	for _, distance := range strings.Fields(input[1])[1:] {
		newDistance, _ := strconv.Atoi(distance)
		distances = append(distances, newDistance)
	}
	for i := 0; i < len(times); i++ {
		wins := 0
		for j := 1; j < times[i]; j++ {
			distance := j * (times[i] - j)
			if distance > distances[i] { wins++ }
		}
		result *= wins
	}

	return result
}

func PartTwo(arr *[]string) int {
	input := *arr
	result := 0

	raceTime, _ := strconv.Atoi(strings.Replace(strings.Join(strings.Fields(input[0])[1:], "")," ", "", 0))
	raceDistance, _ := strconv.Atoi(strings.Replace(strings.Join(strings.Fields(input[1])[1:], "")," ", "", 0))

	for i := 1; i < raceTime; i++ {
		distance := i * (raceTime - i)
		if distance > raceDistance { result++ }
	}
	

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	inputLen := len(inputFile)
	input := strings.Split(string(inputFile[:inputLen-1]), "\n")

	partOneResult := PartOne(&input)
	fmt.Printf("Part One Result: %d\n", partOneResult)

	partTwoResult := PartTwo(&input)
	fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
