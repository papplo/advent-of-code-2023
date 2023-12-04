package main

import (
	"fmt"
	"os"
	"strings"
)

func findPartNumbers(rows [][]string) []int {
	var partNumbers []int
	partNumbers = append(partNumbers, 467, 114)
	return partNumbers
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	/* -- Day 3 -- pt.1 */
	var sumOfPartNumbers int

	/* create matrix data structure */
	type LinesOfText [][]string

	lines := strings.Split(string(file), "\n")
	engineSchematic := make(LinesOfText, len(lines[0]))

	for row, line := range lines {
		engineSchematic[row] = strings.Split(line, "")

		switch row {
		case 0:
			{
				firstRow := engineSchematic[0:1]
				res := findPartNumbers(firstRow)
				println(res[0] + res[1])
			}
		case len(lines) - 1:
			{
				// last row
			}
		default:
			{
				// all other rows
			}
		}
	}

	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprint(sumOfPartNumbers))
	fmt.Printf("AoC: Day 3, part 1: %d", sumOfPartNumbers)

}
