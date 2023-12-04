package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findPartNumbers(row []string, rowNo int, adjacentRows [][]string) []int {
	var validPartNumbers []int
	// var notValidPartNumbers []int
	var foundSequences [][]int

	r, _ := regexp.Compile("([0-9]){3}")
	foundSequences = r.FindAllStringIndex(strings.Join(row, ""), -1)

	for _, sequence := range foundSequences {
		var hasSymbolBefore bool
		var hasSymbolAfter bool

		if sequence[0] != 0 {
			hasSymbolBefore = row[sequence[0]-1] != "."
		}

		if len(row)-1 >= sequence[1] {
			hasSymbolAfter = row[sequence[1]] != "."
		}

		if hasSymbolBefore || hasSymbolAfter {
			startIndex := row[sequence[0]:sequence[1]]
			fmt.Printf("Valid PartNo: %v at row %d, starting at idx: %d\n", getValueAtIndex(startIndex), rowNo, sequence)
			validPartNumbers = append(validPartNumbers, getValueAtIndex(startIndex))
		}

		// now check if sequence happens to have symbol in adjacent slot
		if !hasSymbolAfter && !hasSymbolBefore {
			startIndex := row[sequence[0]:sequence[1]]

			var sliceStart, sliceEnd int
			slots := make([][]string, 0)
			if sequence[0] == 0 {
				sliceStart = 0
			} else {
				sliceStart = sequence[0] - 1
			}

			if sequence[1] >= len(row)-1 {
				sliceEnd = len(row) - 1
			} else {
				sliceEnd = sequence[1]
			}

			switch rowNo {
			case 0:
				slots = append(slots, adjacentRows[1][sliceStart:sliceEnd])
			case 139:
				slots = append(slots, adjacentRows[0][sliceStart:sliceEnd])
			default:
				slots = append(slots,
					adjacentRows[0][sliceStart:sliceEnd],
					adjacentRows[2][sliceStart:sliceEnd],
				)

			}

			for _, slot := range slots {
				println(strings.Join(slot, ""))
				foundSymbol := strings.ContainsAny(strings.Join(slot, ""), "@-*=-%&/")
				if foundSymbol {
					fmt.Printf("Valid PartNo: %v at row %d, starting at idx: %d\n", getValueAtIndex(startIndex), rowNo, sequence[0])
					validPartNumbers = append(validPartNumbers, getValueAtIndex(startIndex))
				} else {
					fmt.Printf("Not Valid PartNo: %v at row %d, starting at idx: %d\n", getValueAtIndex(startIndex), rowNo, sequence[0])
				}
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
	}

	for rowNo, row := range engineSchematic {
		switch rowNo {
		case 0:
			{
				adjacentRows := engineSchematic[0:2]
				sumOfPartNumbers = reduce(
					findPartNumbers(row, 0, adjacentRows),
					func(acc, current int) int {
						return acc + current
					}, sumOfPartNumbers)
			}
		case len(engineSchematic) - 1:
			{
				adjacentRows := engineSchematic[len(engineSchematic)-2:]
				sumOfPartNumbers = reduce(
					findPartNumbers(row, rowNo, adjacentRows),
					func(acc, current int) int {
						return acc + current
					}, sumOfPartNumbers)
			}
		default:
			{
				adjacentRows := engineSchematic[rowNo-1 : rowNo+1]
				sumOfPartNumbers = reduce(
					findPartNumbers(row, rowNo, adjacentRows),
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