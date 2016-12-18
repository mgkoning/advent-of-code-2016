package main

import (
	"fmt"
	"strings"
)

var puzzleRowZero = `^^^^......^...^..^....^^^.^^^.^.^^^^^^..^...^^...^^^.^^....^..^^^.^.^^...^.^...^^.^^^.^^^^.^^.^..^.^`
var example1RowZero = `..^^.`
var example2RowZero = `.^^.^.^^^^`

func main() {
	fmt.Println("Example 1:")
	determineAnswer(example1RowZero, 5, true)
	fmt.Println("Example 2:")
	determineAnswer(example2RowZero, 10, true)
	fmt.Println("Puzzle part 1:")
	determineAnswer(puzzleRowZero, 40, true)
	fmt.Println("Puzzle part 2:")
	determineAnswer(puzzleRowZero, 400000, false)
}

func determineAnswer(rowZero string, rowCount int, print bool) {
	rows := make([][]bool, 1)
	rows[0] = parseRow(rowZero)
	if print {
		printRow(rows[0])
	}
	for len(rows) < rowCount {
		rows = append(rows, determineNext(rows[len(rows)-1], print))
	}
	untrapped := determineUntrapped(rows) - 2*rowCount
	fmt.Println("Untrapped tiles:", untrapped)
}

func parseRow(row string) []bool {
	chars := strings.Split(row, "")
	result := make([]bool, len(chars)+2)
	for i := 0; i < len(chars); i++ {
		result[i+1] = chars[i] == "^"
	}
	return result
}

func determineNext(previousRow []bool, print bool) []bool {
	result := make([]bool, len(previousRow))
	for i := 1; i < len(previousRow)-1; i++ {
		result[i] = previousRow[i-1] != previousRow[i+1]
	}
	if print {
		printRow(result)
	}
	return result
}

func determineUntrapped(rows [][]bool) int {
	sum := 0
	for _, row := range rows {
		for _, cell := range row {
			if cell {
				continue
			}
			sum++
		}
	}
	return sum
}

func printRow(row []bool) {
	for i := 1; i < len(row)-1; i++ {
		if row[i] {
			fmt.Print("^")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
