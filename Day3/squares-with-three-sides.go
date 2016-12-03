package main

import (
	"fmt"
	"strconv"
	"strings"
)

type triangle struct {
	A int64
	B int64
	C int64
}

func main() {
	possible, _ := filterByPossible(getTrianglesFromInput())
	fmt.Printf("Possible: %v", len(possible))
}

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func filterByPossible(triangles []triangle) (possible []triangle, impossible []triangle) {
	for index := 0; index < len(triangles); index++ {
		triangle := triangles[index]
		if triangle.isPossible() {
			possible = append(possible, triangle)
		} else {
			impossible = append(impossible, triangle)
		}
	}
	return possible, impossible
}

func getTrianglesFromInput() []triangle {
	input := getInput()
	lines := strings.Split(input, "\n")
    var column1, column2, column3 []int64
    for index := 0; index < len(lines); index++ {
        var one, two, three = parseSides(lines[index])
        column1 = append(column1, one)
        column2 = append(column2, two)
        column3 = append(column3, three)
	}
    sides := append(column1, append(column2, column3...)...)
	var triangles []triangle
    for index := 0; index < len(sides); index += 3 {
        triangle := triangle{sides[index], sides[index+1], sides[index+2]}
        triangles = append(triangles, triangle)
    }
	return triangles
}

func parseSides(line string) (int64, int64, int64) {
	splitted := strings.Split(line, " ")
	var sides []int64
	for index := 0; index < len(splitted); index++ {
		if len(splitted[index]) < 1 {
			continue
		}
		side, err := strconv.ParseInt(splitted[index], 10, 64)
		check(err)
		sides = append(sides, side)
	}
	if len(sides) != 3 {
		panic(fmt.Sprintf("'%v' could not be parsed", line))
	}
	return sides[0], sides[1], sides[2]
}

func (triangle triangle) isPossible() bool {
	return triangle.A + triangle.B > triangle.C &&
		triangle.A + triangle.C > triangle.B &&
		triangle.B + triangle.C > triangle.A
}
