package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func aocInput() []string {
	file, _ := os.ReadFile("input_test.txt")
	return strings.Split(string(file), "\n\n")
}

func aocOutput(r []int) {
	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprintln(r))
}

var mapper = map[string][]int{}

func findSequences(in string) (string, []string) {
	rname, _ := regexp.Compile("(.+(-to-).[^\t\n\f\r ]+)")
	r, _ := regexp.Compile("([0-9]).+[^\t\n\f\r ]")

	name := rname.FindString(in)
	seq := r.FindAllString(in, -1)

	return name, seq
}

func main() {
	input := aocInput()

	// sequences
	// _, seeds := findSequences(input[0])
	// var bigmap map[string][][]int

	for _, s := range input[1:] {
		maps, seq := findSequences(s)
		fmt.Println(maps, seq)

	}
}
