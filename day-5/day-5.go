package main

import (
	"fmt"
	"os"
	"strings"
)

func aocInput() []string {
	file, _ := os.ReadFile("input.txt")
	return strings.Split(string(file), "\n")
}

func aocOutput(r []int) {
	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprintln(r))
}

func main() {

}
