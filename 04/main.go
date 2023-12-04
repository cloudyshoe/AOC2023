package main

import (
    "os"
    "fmt"
    "slices"
    "strings"
)

func IntPow(base int, exp int) int {
    if exp == 0 { return 1 }
    result := 1
    for i := 1; i <= exp; i++ {
        result *= base
    }
    return result
}


func PartOne(arr *[]string) int  {
    input := *arr
    result := 0

    for _, line := range input {
        gm := strings.Split(line, ": ")
        groups := strings.Split(gm[1], " | ")
        winningNums := strings.Fields(groups[0])
        cardNums := strings.Fields(groups[1])
        exp := -1
        for _, num := range cardNums {
            if slices.Contains(winningNums, num) { exp++ }
        }
        if exp > -1 { result += IntPow(2,exp) }
    }

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    inputLen := len(input)
    result := 0
    copies := make([]int, inputLen)


    for cardIndex, line := range input {
        gm := strings.Split(line, ": ")
        groups := strings.Split(gm[1], " | ")
        winningNums := strings.Fields(groups[0])
        cardNums := strings.Fields(groups[1])
        winners := 0
        copies[cardIndex]++
        for _, num := range cardNums {
            if slices.Contains(winningNums, num) { winners++ }
        }
        for i := cardIndex+1; i < winners + cardIndex + 1; i++ {
            copies[i] += copies[cardIndex]
        }
    }

    for _, num := range copies {
        result += num
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
