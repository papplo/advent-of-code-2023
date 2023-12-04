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

// determine the fewest number of cubes of each color
// needed in bag to make game possible
type CubeSet struct {
	red   int
	green int
	blue  int
}

func fewestAmountOfCubes(gameData string) CubeSet {
	gameDataClean := strings.SplitAfter(gameData, "Game")
	gameRounds := strings.Split(gameDataClean[1], ":")[1]
	var red, green, blue int

	for _, round := range strings.Split(gameRounds, ";") {
		for _, cubes := range strings.Split(round, ", ") {
			token := strings.Fields(cubes)
			amount, err := strconv.Atoi(token[0])
			if err == nil {
				switch token[1] {
				case "red":
					if amount > red || red == 0 {
						red = amount
					}
				case "green":
					if amount > green || green == 0 {
						green = amount
					}
				case "blue":
					if amount > blue || blue == 0 {
						blue = amount
					}
				}
			}
		}
	}

	return CubeSet{red: red, green: green, blue: blue}
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
			// println(id, sum)
			if error != nil {
				panic("whoea")
			}
		}
	}

	var power int
	for _, gameData := range strings.Split(string(file), "\n") {
		cubes := fewestAmountOfCubes(gameData)
		powerOfCubes := cubes.blue * cubes.red * cubes.green
		power += powerOfCubes
		fmt.Println(cubes, powerOfCubes, power)

	}

	output, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer output.Close()
	output.WriteString(fmt.Sprint(sum))
	output.WriteString(fmt.Sprint("\n", power))
	fmt.Printf("AoC: Day 2, part 2: %d", power)

}
