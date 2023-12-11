package main

import (
    "os"
    "fmt"
	"slices"
    "strings"
)

func PartOne(arr *[]string) int  {
    input := *arr
    result := 0
	galaxyMap := make([][]string,0)
	gLocs := make([][2]int,0)
	count := 0

	for _, row := range input {
		galaxyMap = append(galaxyMap, strings.Split(row,""))
		if !strings.Contains(row,"#") {
			galaxyMap = append(galaxyMap, strings.Split(row,""))
		}
	}

	addList := make([]int,0)
	for col := range galaxyMap[0] {
		emptyCount := 0
		for row := range galaxyMap {
			if galaxyMap[row][col] == "." { emptyCount++ }
		}
		if emptyCount == len(galaxyMap) { addList = append(addList, col) }
	}

	for _, col := range addList {
		for row := range galaxyMap {
			galaxyMap[row] = slices.Insert(galaxyMap[row], col+count, ".")
		}
		count++
	}
	for i, row := range galaxyMap {
		for slices.Index(row,"#") != -1 {
			gIndex := slices.Index(row,"#")
			gLocs = append(gLocs,[2]int{i, gIndex})
			row[gIndex] = "@"
		}
	}

	for i := range gLocs {
		for j := i+1; j < len(gLocs); j++ {
			result += max(gLocs[i][0],gLocs[j][0]) - min(gLocs[i][0],gLocs[j][0])
			result += max(gLocs[i][1],gLocs[j][1]) - min(gLocs[i][1],gLocs[j][1])
		}
	}

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    result := 0
	galaxyMap := make([][]string,0)
	gLocs := make([][2]int,0)
	longRows := make([]int,0)

	for i, row := range input {
		galaxyMap = append(galaxyMap, strings.Split(row,""))
		if !strings.Contains(row,"#") {
			longRows = append(longRows, i)
		}
	}

	longCols := make([]int,0)
	for col := range galaxyMap[0] {
		emptyCount := 0
		for row := range galaxyMap {
			if galaxyMap[row][col] == "." { emptyCount++ }
		}
		if emptyCount == len(galaxyMap) { longCols = append(longCols, col) }
	}

	for i, row := range galaxyMap {
		for slices.Index(row,"#") != -1 {
			gIndex := slices.Index(row,"#")
			gLocs = append(gLocs,[2]int{i, gIndex})
			row[gIndex] = "@"
		}
	}

	for i := range gLocs {
		for j := i+1; j < len(gLocs); j++ {
			rowmax := max(gLocs[i][0],gLocs[j][0])
			rowmin := min(gLocs[i][0],gLocs[j][0])
			colmax := max(gLocs[i][1],gLocs[j][1])
			colmin := min(gLocs[i][1],gLocs[j][1])
			rowdiff := rowmax - rowmin
			coldiff := colmax - colmin

			for _, row := range longRows {
				if row > rowmin && row < rowmax {
					rowdiff += 999_999
				}
			}

			for _, col := range longCols {
				if col > colmin && col < colmax {
					coldiff += 999_999
				}
			}
			result += rowdiff + coldiff
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
