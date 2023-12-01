package main

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	exampleBytes, _ := os.ReadFile("example_part_one.txt")
	example := strings.Split(string(exampleBytes), "\n")
	want := 142

	got := PartOne(&example)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}

}


func TestPartTwo(t *testing.T) {
	exampleBytes, _ := os.ReadFile("example_part_two.txt")
	example := strings.Split(string(exampleBytes), "\n")
	want := 281

	got := PartTwo(&example)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}


func BenchmarkPartOne(b *testing.B) {
	inputBytes, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputBytes), "\n")

    for i := 0; i < b.N; i++ {
        PartOne(&input)
    }
}
