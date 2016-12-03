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
	var triangles []triangle
	for index := 0; index < len(lines); index++ {
		triangles = append(triangles, parseTriangle(lines[index]))
	}
	return triangles
}

func parseTriangle(line string) triangle {
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
	return triangle{sides[0], sides[1], sides[2]}
}

func (triangle triangle) isPossible() bool {
	return triangle.A + triangle.B > triangle.C &&
		triangle.A + triangle.C > triangle.B &&
		triangle.B + triangle.C > triangle.A
}
