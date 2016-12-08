package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

/*
For posterity, this is the final output:

####.####.#..#.####..###.####..##...##..###...##..
...#.#....#..#.#....#....#....#..#.#..#.#..#.#..#.
..#..###..####.###..#....###..#..#.#....#..#.#..#.
.#...#....#..#.#.....##..#....#..#.#.##.###..#..#.
#....#....#..#.#.......#.#....#..#.#..#.#....#..#.
####.#....#..#.#....###..#.....##...###.#.....##..
Enabled: 119

*/

func main() {
	instructions := make(chan instruction)
	go parseInput(instructions)
	display := initializeDisplay()
	showDisplay(display)
	for instruction := range instructions {
		fmt.Println(instruction)
		display = instruction.execute(display)
		showDisplay(display)
		time.Sleep(100 * time.Millisecond)
	}
}

func showDisplay(state [][]bool) {
	enabled := 0
	for x, row := range state {
		for y := range row {
			pixel := state[x][y]
			fmt.Print(formatPixel(pixel))
			if pixel {
				enabled++
			}
		}
		fmt.Println()
	}
	fmt.Printf("Enabled: %v\n", enabled)
	fmt.Println()

}

func formatPixel(value bool) string {
	if value {
		return "#"
	}
	return "."
}

func initializeDisplay() [][]bool {
	display := make([][]bool, 6)
	for index := range display {
		display[index] = make([]bool, 50)
	}
	return display
}

func parseInput(instructions chan instruction) {
	lines := readInput()
	for _, line := range lines {
		instructions <- parseInstruction(line)
	}
	close(instructions)
}

func readInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	check(err)
	return strings.Split(string(contents), "\n")
}
