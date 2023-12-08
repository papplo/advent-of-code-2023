package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func aocInput() []string {
	file, _ := os.ReadFile("input.txt")
	return strings.Split(string(file), "\n")
}

func aocOutput(r int) {
	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprint(r))
}

func main() {

	input := aocInput()
	// each card (line) has a: []nums | b: []nums
	// where for each b in a, you get
	// 1 point for first, double the amount for ...n

	r, _ := regexp.Compile("([0-9]){1,3}")
	var totalPoints int

	for i := 0; i < len(input); i++ {
		nums := r.FindAllString(input[i], -1)

		// stuff we know about a line
		var matchesInGame int
		gameNumber := nums[0:1]
		winningNumbers := nums[1:11]
		myNumbers := nums[11:]

		for _, w := range winningNumbers {
			for _, m := range myNumbers {
				if m == w {
					// fmt.Printf("Win: %v, Scratched: %v\n\n", w, m)
					matchesInGame += 1
				}
			}
		}

		gamePoints := min(matchesInGame, 1)
		for j := 1; j < matchesInGame; j++ {
			gamePoints = gamePoints * 2
		}

		if gamePoints > 0 {
			fmt.Printf("Matches in game %v: %v, POINTS: %v\n", gameNumber, matchesInGame, gamePoints)
		}
		totalPoints += gamePoints
	}
	aocOutput(totalPoints)
}
