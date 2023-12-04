package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func findPartNumbers(rows [][]string) []int {
	var validPartNumbers []int

	// Scan horizontally
	for _, row := range rows {
		var _ []int

		r, _ := regexp.Compile("([0-9]+){3}")

		fmt.Println(r.FindAllStringIndex(strings.Join(row, ""), -1))

	}

	validPartNumbers = append(validPartNumbers, 467, 114)
	return validPartNumbers
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
				findPartNumbers(firstRow)
			}
		case len(lines) - 1:
			{
				lastRow := engineSchematic[139:140]
				findPartNumbers(lastRow)
			}
		default:
			{
				anyRow := engineSchematic[row-1 : row+1]
				findPartNumbers(anyRow)
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
