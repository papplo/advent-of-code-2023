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

func scratchCards(line string) (int, int) {
	r, _ := regexp.Compile("([0-9]){1,3}")
	nums := r.FindAllString(line, -1)

	var scratchedCards int
	var matchesInGame int
	gameNumber := nums[0:1]
	winningNumbers := nums[1:11]
	myNumbers := nums[11:]

	for _, w := range winningNumbers {
		for _, m := range myNumbers {
			if m == w {
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
	return gamePoints, scratchedCards
}

func main() {
	input := aocInput()

	var totalPoints int
	var totalCards int

	for i := 0; i < len(input); i++ {
		gamePoints, scratchedCards := scratchCards(input[i])

		totalCards += scratchedCards
		totalPoints += gamePoints
	}

	aocOutput(totalCards)
}
