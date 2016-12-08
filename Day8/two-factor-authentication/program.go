package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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
		if instruction == nil {
			continue
		}
		fmt.Println(instruction)
		display = instruction.execute(display)
		showDisplay(display)
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

var rectRegexp = regexp.MustCompile(`rect (\d+)x(\d+)`)
var rotateColumnRegexp = regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)
var rotateRowRegexp = regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)

func parseInstruction(line string) instruction {
	if rectValues := rectRegexp.FindStringSubmatch(line); rectValues != nil {
		return rect{columns: mustParseInt(rectValues[1]), rows: mustParseInt(rectValues[2])}
	}
	if rotateColumnValues := rotateColumnRegexp.FindStringSubmatch(line); rotateColumnValues != nil {
		return rotateColumn{
			columnIndex: mustParseInt(rotateColumnValues[1]),
			rotations:   mustParseInt(rotateColumnValues[2]),
		}
	}
	if rotateRowValues := rotateRowRegexp.FindStringSubmatch(line); rotateRowValues != nil {
		return rotateRow{
			rowIndex:  mustParseInt(rotateRowValues[1]),
			rotations: mustParseInt(rotateRowValues[2]),
		}
	}
	panic(fmt.Sprintf("Cannot understand '%v'", line))
}

func mustParseInt(s string) int {
	val, err := strconv.ParseInt(s, 10, 32)
	check(err)
	return int(val)
}

func readInput() []string {
	contents, err := ioutil.ReadFile("input.txt")
	check(err)
	return strings.Split(string(contents), "\n")
}
