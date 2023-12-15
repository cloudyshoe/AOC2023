package main

import (
    "os"
    "fmt"
	"slices"
	"strconv"
    "strings"
)

type Box struct {
	labels []string
	lenses []int
}

func PartOne(arr *[]string) int {
    input := *arr
    result := 0

	init := strings.Split(input[0],",")

	for _, seq := range init {
		tmp := 0
		for _, r := range seq {
			tmp += int(r)
			tmp *= 17
			tmp %= 256
		}
		result += tmp
	}

    return result
}

func PartTwo(arr *[]string) int  {
    input := *arr
    result := 0

	boxes := make([]Box, 256)

	init := strings.Split(input[0],",")

	for _, seq := range init {

		label := []string{}
		hasDash := strings.Contains(seq, "-")

		if hasDash {
			label = strings.Split(seq, "-")
		} else {
			label = strings.Split(seq, "=")
		}

		hash := 0
		for _, r := range label[0] {
			hash += int(r)
			hash *= 17
			hash %= 256
		}

		lIndex := slices.Index(boxes[hash].labels, label[0])

		if lIndex > -1 {
			if hasDash {
				boxes[hash].labels = slices.Delete(boxes[hash].labels, lIndex, lIndex+1)
				boxes[hash].lenses = slices.Delete(boxes[hash].lenses, lIndex, lIndex+1)
			} else {
				length, _ := strconv.Atoi(label[1])
				boxes[hash].lenses[lIndex] = length
			}
		} else {
			if !hasDash {
				boxes[hash].labels = append(boxes[hash].labels, label[0])
				length, _ := strconv.Atoi(label[1])
				boxes[hash].lenses = append(boxes[hash].lenses, length)
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box.lenses {
			result += (i+1) * (j+1) * lens
		}
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
