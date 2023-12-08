package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	Y, X int
}

func IsDigit(b byte) bool {
	return b-'0' >= 0 && b-'0' <= 9
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(file), "\n")

	// lines := aoc.AsLines("2023/Day03/input.txt")

	ygrid := len(lines)
	xgrid := len(lines[0])

	grid := make([][]byte, ygrid)
	for y, l := range lines {
		grid[y] = make([]byte, xgrid)
		for x := 0; x < len(l); x++ {
			grid[y][x] = l[x]
		}
	}

	neighbors := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	gears := make(map[Point][]int)
	sum := 0
	for y := 0; y < ygrid; y++ {
		for x := 0; x < xgrid; x++ {
			num := 0
			hasSymbol := false
			isGear := false
			var gearCoord Point
			for IsDigit(grid[y][x]) {
				num = num*10 + int(grid[y][x]-'0')
				for _, n := range neighbors {
					if y+n[1] >= 0 && y+n[1] < ygrid &&
						x+n[0] >= 0 && x+n[0] < xgrid {
						v := grid[y+n[1]][x+n[0]]
						if !IsDigit(v) && v != '.' {
							if v == '*' {
								isGear = true
								gearCoord = Point{X: x + n[0], Y: y + n[1]}
							}
							hasSymbol = true
						}
					}
				}
				x += 1
				if x >= xgrid {
					break
				}
			}
			if num > 0 && hasSymbol {
				if isGear {
					gears[gearCoord] = append(gears[gearCoord], num)
				}
				sum += num
			}
		}
	}

	fmt.Println("Part 1: ", sum) // 498559

	sum = 0
	for _, v := range gears {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	fmt.Println("Part 2: ", sum) // 72246648
}
