package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processNumber(row int, col int, arr *[]string) (int, bool, int) {
	input := *arr
	ignore := "0123456789."
	var partNum int
	var lastIndex int
	neighbors := false

	for i := col; i < len(input[row]); i++ {
		val, err := strconv.Atoi(string(input[row][i]))
		if err == nil {
			partNum = partNum * 10
			partNum += val
			lastIndex = i
		} else {
			break
		}
	}

	for sRow := row - 1; sRow <= row+1; sRow++ {
		for sCol := col - 1; sCol <= lastIndex+1; sCol++ {
			if sRow >= 0 && sRow < len(input)-1 && sCol >= 0 && sCol < len(input)-1 &&
				!strings.ContainsAny(string(input[sRow][sCol]), ignore) {
				neighbors = true
				break
			}
		}
	}

	return lastIndex + 1, neighbors, partNum

}

func PartOne(arr *[]string) int {
	input := *arr
	result := 0
	digits := "0123456789"
	var partNum int
	var valid bool

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); {
			char := string(input[row][col])
			if strings.ContainsAny(char, digits) {
				col, valid, partNum = processNumber(row, col, arr)
				if valid {
					result += partNum
				}
			} else {
				col++
			}
		}
	}

	return result
}

func processNumberP2(row int, col int, arr *[]string) (int, map[int][]int) {
	input := *arr
	var partNum int
	var lastIndex int
	gears := make(map[int][]int)

	for i := col; i < len(input[row]); i++ {
		val, err := strconv.Atoi(string(input[row][i]))
		if err == nil {
			partNum = partNum * 10
			partNum += val
			lastIndex = i
		} else {
			break
		}
	}

	for sRow := row - 1; sRow <= row+1; sRow++ {
		for sCol := col - 1; sCol <= lastIndex+1; sCol++ {
			if sRow >= 0 && sRow < len(input)-1 && sCol >= 0 && sCol < len(input)-1 &&
				input[sRow][sCol] == '*' {
				index := sRow*len(input[sRow]) + sCol
				_, ok := gears[index]
				if ok {
					gears[index] = append(gears[index], partNum)
				} else {
					gears[index] = []int{partNum}
				}
			}
		}
	}

	return lastIndex + 1, gears

}

func PartTwo(arr *[]string) int {
	input := *arr
	result := 0
	digits := "0123456789"
	//var partNum int
	potentialGears := make(map[int][]int)

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); {
			char := string(input[row][col])
			if strings.ContainsAny(char, digits) {
				var tmp map[int][]int
				col, tmp = processNumberP2(row, col, arr)
				for k, v := range tmp {
					for _, num := range v {
						_, ok := potentialGears[k]
						if ok {
							potentialGears[k] = append(potentialGears[k], num)
						} else {
							potentialGears[k] = []int{num}
						}
					}
				}
			} else {
				col++
			}
		}
	}

	for _, v := range potentialGears {
		if len(v) == 2 {
			result += v[0] * v[1]
		}
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
