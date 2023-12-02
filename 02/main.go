package main

import (
	"os"
	"fmt"
	"strings"
)

func PartOne(arr *[]string) int {
	input := *arr
	result := 0
	bagContents := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	for _, line := range input {
		var game_id int
		possible := true
		
		game := strings.Split(line,": ")
		fmt.Sscanf(game[0],"Game %d", &game_id)

		rounds := strings.Split(game[1], "; ")
		for _, round := range rounds  {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				var count int
				var color string
				fmt.Sscanf(cube, "%d %s", &count, &color)
				if bagContents[color] < count {
					possible = false
				}
                if !possible { break }
			}
		}
		
		if possible { result += game_id }
		
	}

	return result
}

func PartTwo(arr *[]string) int {
	input := *arr
	result := 0
	bagContents := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	
	for _, line := range input {
		bagContents["red"] = 0
		bagContents["green"] = 0
		bagContents["blue"] = 0

		game := strings.Split(line,": ")
		rounds := strings.Split(game[1], "; ")
		for _, round := range rounds  {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				var count int
				var color string
				fmt.Sscanf(cube, "%d %s", &count, &color)
                bagContents[color] = max(bagContents[color], count)
			}
		}
		
		result += bagContents["red"] * bagContents["green"] * bagContents["blue"]

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
