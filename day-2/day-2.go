package main

import (
	"fmt"
	"os"
)

// determine possible games if
// bag has constraints 12 red, 13 green, 15 blue

// sum id's of game id's that are possible
func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	println(file)

	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprint("results here"))
	fmt.Printf("AoC: Day 2, part 1: %d", 000)

}
