package main

import (
    "os"
    "fmt"
    "math"
    "strconv"
    "strings"
)

type CategoryRange struct {
    srcStart int
    srcEnd int
    dstStart int
    dstEnd int
    dstOffset int
}

func (c CategoryRange) InRange (x int) int {
    if x >= c.srcStart && x <= c.srcEnd {
        return x + c.dstOffset
    } else {
        return -1
    }
}

func (c CategoryRange) InDstRange (x int) int {
    if x >= c.dstStart && x <= c.dstEnd {
        return x - c.dstOffset
    } else {
        return -1
    }
}

type SeedRange struct {
    start int
    end int
}

func (s SeedRange) InRange (x int) bool {
    if x >= s.start && x <= s.end {
        return true
    }
    return false
}
func WalkTheMap(categories map[string][]CategoryRange, mapIndex string, val int) int {
    for _, category := range categories[mapIndex] {
        mapping := category.InRange(val)
        if mapping >= 0 {
            return mapping
        }
    }
    return val
}

func WalkTheMaps(seed int, categories map[string][]CategoryRange) int {

    soil := WalkTheMap(categories, "seed-to-soil", seed)
    fertilizer := WalkTheMap(categories, "soil-to-fertilizer", soil)
    water := WalkTheMap(categories, "fertilizer-to-water", fertilizer)
    light := WalkTheMap(categories, "water-to-light", water)
    temperature := WalkTheMap(categories, "light-to-temperature", light)
    humidity := WalkTheMap(categories, "temperature-to-humidity", temperature)
    location := WalkTheMap(categories, "humidity-to-location", humidity)

    return location
}

func WalkBackward(categories map[string][]CategoryRange, mapIndex string, val int) int {
    for _, category := range categories[mapIndex] {
        mapping := category.InDstRange(val)
        if mapping >= 0 {
            return mapping
        }
    }
    return val
}

func WalkBackwards(location int, categories map[string][]CategoryRange) int {

    humidity := WalkBackward(categories, "humidity-to-location", location)
    temperature := WalkBackward(categories, "temperature-to-humidity", humidity)
    light := WalkBackward(categories, "light-to-temperature", temperature)
    water := WalkBackward(categories, "water-to-light", light)
    fertilizer := WalkBackward(categories, "fertilizer-to-water", water)
    soil := WalkBackward(categories, "soil-to-fertilizer", fertilizer)
    seed := WalkBackward(categories, "seed-to-soil", soil)

    return seed
}

func PartOne(arr *[]string) int  {
    input := *arr
    result := math.MaxInt
    var categories = make(map[string][]CategoryRange, len(input)-1)
    var seeds []int

    for i, group := range input {
        if i != 0 {
            categoryLines := strings.Split(group, "\n")
            names := strings.Split(strings.Fields(categoryLines[0])[0], "-to-")
            source := names[0]
            destination := names[1]
            for _, line := range categoryLines[1:] {
                fields := strings.Fields(line)
                dst, _ := strconv.Atoi(fields[0])
                src, _ := strconv.Atoi(fields[1])
                length, _ := strconv.Atoi(fields[2])
                mapName := source+"-to-"+destination
                categories[mapName] = append(
                    categories[mapName],
                    CategoryRange{
                        srcStart: src,
                        srcEnd: src + length -1,
                        dstOffset: dst - src,
                })
            }
        } else {
            seedList := strings.Fields(group)[1:]
            for _, seed := range seedList {
                seedNum, _ := strconv.Atoi(seed)
                seeds = append(seeds, seedNum)
            }
        }

    }

    for _, seed := range seeds {
        tmp := WalkTheMaps(seed, categories)
        if tmp < result { result = tmp }
    }

    return result
}

func PartTwo(arr *[]string) int {
    input := *arr
    //result := math.MaxInt
    var categories = make(map[string][]CategoryRange, len(input)-1)
    var seeds []SeedRange
    minSeed := 0
    maxSeed := 0

    for i, group := range input {
        if i != 0 {
            categoryLines := strings.Split(group, "\n")
            names := strings.Split(strings.Fields(categoryLines[0])[0], "-to-")
            source := names[0]
            destination := names[1]
            for _, line := range categoryLines[1:] {
                fields := strings.Fields(line)
                dst, _ := strconv.Atoi(fields[0])
                src, _ := strconv.Atoi(fields[1])
                length, _ := strconv.Atoi(fields[2])
                mapName := source+"-to-"+destination
                categories[mapName] = append(
                    categories[mapName],
                    CategoryRange{
                        srcStart: src,
                        srcEnd: src + length - 1,
                        dstStart: dst,
                        dstEnd: dst + length - 1,
                        dstOffset: dst - src,
                })
            }
        } else {
            seedList := strings.Fields(group)[1:]
            for i := 0; i < len(seedList); i += 2 {
                start, _:= strconv.Atoi(seedList[i])
                length, _ := strconv.Atoi(seedList[i+1])
                end := start + length - 1
                if start < minSeed { minSeed = start }
                if end > maxSeed { maxSeed = end }
                seeds = append(seeds, SeedRange{ start: start, end: end })
            }
        }
    }

    for location := minSeed; location <= maxSeed; location++ {
        seed := WalkBackwards(location, categories)
        for _, rng := range seeds {
            if rng.InRange(seed) { return location }
        }
    }

    return 0
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
