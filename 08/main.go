package main

import (
    "os"
    "fmt"
    "strings"
)

func GCD(a,b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, list []int) int {
	result := (a*b)/GCD(a,b)
	if len(list) > 0 { result = LCM(result, list[0],list[1:]) }
	return result
}

func PartOne(arr *[]string) int  {
    input := *arr
    result := 0

	instructions := strings.Split(input[0],"")
	mapData := make(map[string][2]string, len(input[1:]))

	lines := strings.Split(input[1],"\n")
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		elements := strings.Split(parts[1],",")
		mapData[parts[0]] = [2]string{elements[0][1:],elements[1][1:4]}
	}

	location := "AAA"
	steps := 0

	for location != "ZZZ" {
		for _, instr := range instructions {
			steps++
			if instr == "L" { location = mapData[location][0] }
			if instr == "R" { location = mapData[location][1] }
		}
	}
	
	result = steps
    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    result := 0

	instructions := strings.Split(input[0],"")
	mapData := make(map[string][2]string, len(input[1:]))
	var nodes []string

	lines := strings.Split(input[1],"\n")
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		elements := strings.Split(parts[1],",")
		key := parts[0]
		node := [2]string{elements[0][1:],elements[1][1:4]}

		if key[2] == 'A' { nodes = append(nodes, key) }
		mapData[key] = node
	}

	stepsList := make([]int, len(nodes))

	for i, location := range nodes {
		steps := 0
		for location[2] != 'Z' {
			for _, instr := range instructions {
				steps++
				if instr == "L" { location = mapData[location][0] }
				if instr == "R" { location = mapData[location][1] }
			}
		}
		stepsList[i] = steps
	}

	a := stepsList[0]
	b := stepsList[1]
	result = LCM(a,b,stepsList[2:])
	
    return result
}

func main () {
    inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-1]), "\n\n")

    partOneResult := PartOne(&input)
    fmt.Printf("Part One Result: %d\n", partOneResult)

    partTwoResult := PartTwo(&input)
    fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
