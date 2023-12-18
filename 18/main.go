package main

import (
    "os"
    "fmt"
	"strconv"
    "strings"
)

type Instructions struct {
	dir string
	length int
	color string
}

type Coord struct {
	row int
	col int
}

func EndCoords(dir string, length, row, col int) (int,int) {
	switch dir {
	case "R":
		return row, col + length
	case "L":
		return row, col - length
	case "U":
		return row - length, col
	case "D":
		return row + length, col
	default:
		return row, col
	}
}

var FlipDir = map[string]string{
	"RD": "L",
	"RU": "R",
	"LD": "L",
	"LU": "R",
	"DR": "D",
	"DL": "U",
	"UR": "D",
	"UL": "U",
}

func FloodFill(gridArr *map[Coord]string, row, col int) {
	grid := *gridArr
	_, ok := grid[Coord{row: row, col: col}]
	if !ok {
		grid[Coord{ row: row, col: col }] = "#"

		trow := row-1
		tcol := col
		_, ok = grid[Coord{row: trow, col: tcol}]
		if !ok { FloodFill(gridArr,trow,tcol) }

		trow = row+1
		tcol = col
		_, ok = grid[Coord{row: trow, col: tcol}]
		if !ok { FloodFill(gridArr,trow,tcol) }

		trow = row
		tcol = col-1
		_, ok = grid[Coord{row: trow, col: tcol}]
		if !ok { FloodFill(gridArr,trow,tcol) }

		trow = row
		tcol = col+1
		_, ok = grid[Coord{row: trow, col: tcol}]
		if !ok { FloodFill(gridArr,trow,tcol) }
	}
}

var HexDir = map[string]string {
	"0": "R",
	"1": "D",
	"2": "L",
	"3": "U",
}

func PartOne(arr *[]string) int  {
    input := *arr
    result := 0
	insts := make([]Instructions, len(input))
	grid := make(map[Coord]string)
	igrid := make(map[Coord]string)

	for i, line := range input {
		tmp := strings.Fields(line)	
		insts[i].dir = tmp[0]
		insts[i].length, _ = strconv.Atoi(tmp[1])
		insts[i].color = tmp[2][1:len(tmp[2])-1]
	}

	srow := 0
	scol := 0
	prevDir := ""
	for _, inst := range insts {
		erow, ecol := EndCoords(inst.dir, inst.length, srow, scol)
		minRow := min(srow,erow)
		maxRow := max(srow,erow)
		minCol := min(scol,ecol)
		maxCol := max(scol,ecol)
		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				grid[Coord{row: row, col: col}] = "#"
				if prevDir != "" {
					irow, icol := EndCoords(FlipDir[prevDir+inst.dir],1,row,col)
					_, ok := grid[Coord{row: irow, col: icol}]
					if !ok { igrid[Coord{row: irow, col: icol}] = "I" }
				}
			}
		}
		srow = erow
		scol = ecol
		if inst.dir != prevDir { prevDir = inst.dir }
		
	}

	for k,_ := range igrid {
		_, ok := grid[Coord{ row: k.row, col: k.col }]
		if !ok { FloodFill(&grid,k.row,k.col) }
	}

	for range grid {
		result += 1
	}

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    result := 0

	insts := make([]Instructions, len(input))

	for i, line := range input {
		tmp := strings.Fields(line)	
		tmpd := tmp[2][7:8]
		insts[i].dir = HexDir[tmpd]
		tmpl, _ := strconv.ParseInt(tmp[2][2:7],16,64)
		insts[i].length = int(tmpl)
	}

	srow := 0
	scol := 0
	area := 0
	perimeter := 0
	for _, inst := range insts {
		erow, ecol := EndCoords(inst.dir, inst.length, srow, scol)
		if tmp := (srow + erow) * (scol - ecol); tmp < 0 { tmp *= -1 }
		area += ((srow + erow) * (scol - ecol))/2
		for row := min(srow,erow); row <= max(srow,erow); row++ {
			for col := min(scol,ecol); col <= max(scol,ecol); col++ {
				perimeter++
			}
		}

		srow = erow
		scol = ecol
		
	}

	result = area + perimeter/2+1 - len(input)/2
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
