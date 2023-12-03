package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// determine possible games if
// bag has constraints 12 red, 13 green, 14 blue
func isGamePossible(gameData string) (string, bool) {
	gameRules := make(map[string]int)

	gameRules["red"] = 12
	gameRules["green"] = 13
	gameRules["blue"] = 14

	gameDataClean := strings.SplitAfter(gameData, "Game")
	gameId := strings.Split(gameDataClean[1], ":")[0]
	gameHands := strings.Split(gameDataClean[1], ":")[1]

	var isPossible bool

	for _, hand := range strings.Split(gameHands, ";") {
		for _, cubes := range strings.Split(hand, ",") {
			token := strings.Fields(cubes)
			amount, err := strconv.Atoi(token[0])
			if err == nil {

				if amount > gameRules[token[1]] {
					isPossible = false
					return strings.TrimSpace(gameId), isPossible
				}

				isPossible = amount <= gameRules[token[1]]
			} else {
				println(err)
			}
		}
	}

	return strings.TrimSpace(gameId), isPossible
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	var sum int
	for _, gameData := range strings.Split(string(file), "\n") {
		gameId, isPossible := isGamePossible(gameData)
		if isPossible {
			id, error := strconv.Atoi(gameId)
			sum += id
			println(id, sum)
			if error != nil {
				panic("whoea")
			}
		}
	}

	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprint(sum))
	fmt.Printf("AoC: Day 2, part 1: %d", sum)

}
