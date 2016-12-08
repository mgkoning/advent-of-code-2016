package main

import "fmt"

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
