package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type instruction interface {
	execute(screen [][]bool) [][]bool
}

type rotateRow struct {
	rowIndex  int
	rotations int
}

func (rotateRow rotateRow) execute(screen [][]bool) [][]bool {
	row := screen[rotateRow.rowIndex]
	newRow := append(make([]bool, 0), row...)
	for index := range newRow {
		newRow[index] = row[(index+len(row)-rotateRow.rotations)%len(row)]
	}
	screen[rotateRow.rowIndex] = newRow
	return screen
}

func (rotateRow rotateRow) String() string {
	return fmt.Sprintf("Rotate row %v by %v", rotateRow.rowIndex, rotateRow.rotations)
}

type rotateColumn struct {
	columnIndex int
	rotations   int
}

func (rotateColumn rotateColumn) execute(screen [][]bool) [][]bool {
	newColumn := make([]bool, len(screen))
	for i := range newColumn {
		newColumn[i] = screen[(i+len(newColumn)-rotateColumn.rotations)%len(newColumn)][rotateColumn.columnIndex]
	}
	for i := range newColumn {
		screen[i][rotateColumn.columnIndex] = newColumn[i]
	}
	return screen
}

func (rotateColumn rotateColumn) String() string {
	return fmt.Sprintf("Rotate column %v by %v", rotateColumn.columnIndex, rotateColumn.rotations)
}

type rect struct {
	rows    int
	columns int
}

func (rect rect) String() string {
	return fmt.Sprintf("Rectangle %v by %v", rect.columns, rect.rows)
}

func (rect rect) execute(screen [][]bool) [][]bool {
	for x := 0; x < rect.rows; x++ {
		for y := 0; y < rect.columns; y++ {
			screen[x][y] = true
		}
	}
	return screen
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
