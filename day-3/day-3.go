package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findPartNumbers(rows [][]string, rowNo int) []int {
	var validPartNumbers []int

	// Scan horizontally
	for _, row := range rows {
		var foundSequences [][]int

		r, _ := regexp.Compile("([0-9]){3}")
		foundSequences = r.FindAllStringIndex(strings.Join(row, ""), -1)
		// fmt.Println(foundSequences)

		for _, sequence := range foundSequences {

			var hasSymbolBefore bool
			var hasSymbolAfter bool
			if sequence[0] == 0 {
				hasSymbolBefore = false
			} else {
				hasSymbolBefore = row[sequence[0]-1] != "."
			}

			if len(row)-1 < sequence[1] {
				hasSymbolAfter = false
			} else {
				hasSymbolAfter = row[sequence[1]] != "."
			}

			if hasSymbolBefore || hasSymbolAfter {
				startIndex := row[sequence[0]:sequence[1]]
				fmt.Printf("Valid PartNo: %v at row %d, starting at idx: %d\n", getValueAtIndex(startIndex), rowNo, sequence[0])
				validPartNumbers = append(validPartNumbers, getValueAtIndex(startIndex))
			}
		}

	}

	return validPartNumbers
}

func getValueAtIndex(args []string) int {
	var res string
	for _, v := range args {
		res += v
	}

	resValue, _ := strconv.Atoi(res)
	return resValue
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
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
				firstRow := engineSchematic[:1]
				sumOfPartNumbers = reduce(
					findPartNumbers(firstRow, row),
					func(acc, current int) int {
						return acc + current
					}, sumOfPartNumbers)
			}
		case len(lines) - 1:
			{
				lastRow := engineSchematic[len(lines)-2:]
				sumOfPartNumbers = reduce(
					findPartNumbers(lastRow, row),
					func(acc, current int) int {
						return acc + current
					}, sumOfPartNumbers)
			}
		default:
			{
				anyRow := engineSchematic[row-1 : row+1]
				sumOfPartNumbers = reduce(
					findPartNumbers(anyRow, row),
					func(acc, current int) int {
						return acc + current
					}, sumOfPartNumbers)
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
