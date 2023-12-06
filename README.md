# Advent-of-code-2023
Advent of Code, year of the elf 2023
Pablo Anttila

https://github.com/papplo/advent-of-code-2023


## Day 1:
- pt 1. A little troublesome to start programming under time-pressure, particularly when there is not much free time at all while at parental leave.

Third attempt at calculating the input yields `That's the right answer! You are one gold star closer to restoring snow operations.`

- pt 2.
I knew there was something fishy while inspecting the input, the numerals were going to end up in the calculation somehow, at a later point.
Luckily, I solved the problem with an iteration of the full slice, while stepping forward and checking the index of a numeral was present at index == 0.

## Day 2:
- I started a bit later and im still lagging as saturday was an off day due to other private tasks.
However im starting to get a feel of the slice data structure, the libraries that are available to parse integers and split strings. Nice problems overall.
- pt 2. ———> i did'nt quite catch the spec of "least amount of cubes" so calculated for the lowest instead of highest amount of cubes in the bag, per color.

## Day 3: Gear Ratios
- pt 1. This was quite hard, but finding the patterns for partnumbers taking into consideration edge cases such as first/last index on slices was tricky and quite error prone. The real bug hunt was solving a regex that listed a sequence of {1,3} `one, two or three`digits, unknowingly i only searched for 2/3 long sequences. Finally i gave up and started prowling the input file line by line, until i spotted a single digit which gave up the solution.

-pt 2.
