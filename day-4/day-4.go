package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type Round struct {
	game    int
	matches int
	points  int
	copies  int
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func scratchCards(lines *[]string, index int) Round {
	r, _ := regexp.Compile("([0-9]){1,3}")
	line := (*lines)[index]
	nums := r.FindAllString(line, -1)

	gameNumber, _ := strconv.Atoi(nums[0:1][0])
	round := Round{game: gameNumber - 1, copies: 1}
	winningNumbers := nums[1:11]
	myNumbers := nums[11:]

	for _, w := range winningNumbers {
		for _, m := range myNumbers {
			if m == w {
				round.matches++
			}
		}
	}

	round.points = min(round.matches, 1)
	for j := 1; j < round.matches; j++ {
		round.points = round.points * 2
	}

	// if gamePoints > 0 {
	// 	fmt.Printf("Matches in game %v: %v, POINTS: %v\n", gameNumber, matchesInGame, gamePoints)
	// }
	return round
}

func main() {
	input := aocInput()

	cards := make([]Round, 0, len(input))
	var total int
	var points int

	for i := 0; i < len(input); i++ {
		r := scratchCards(&input, i)
		cards = append(cards, r)
		points += r.points
	}

	// for _, round := range cards {
	// 	fmt.Printf("Game %v: matches(%v), copies(%v)\n", round.game, round.matches, round.copies)
	// }

	total = reduce(cards,
		func(acc int, current Round) int {
			fmt.Printf("Game %v: matches(%v), copies(%v)\n", current.game, current.matches, current.copies)

			// matches for this card
			for i := current.matches; i >= 0; i-- {
				cards[current.game+i].copies += current.copies
			}

			acc = acc + current.matches*current.copies
			println(acc)
			return acc
		}, len(cards))

	aocOutput(append(make([]int, 0), points, total))
	fmt.Printf("Total: %v scratchcards \n", total)
}
