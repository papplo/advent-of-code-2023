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

func main() {

	input := aocInput()
	fmt.Println(input)

}
