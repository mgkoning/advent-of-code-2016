package main

import "fmt"

type step interface {
	execute(input []string) []string
	reverse(input []string) []string
}

type positionSwap struct {
	x int
	y int
}

func (p positionSwap) execute(input []string) []string {
	input[p.x], input[p.y] = input[p.y], input[p.x]
	return input
}

func (p positionSwap) reverse(input []string) []string {
	return p.execute(input)
}

type letterSwap struct {
	x string
	y string
}

func (l letterSwap) execute(input []string) []string {
	return positionSwap{getLetterIndex(input, l.x), getLetterIndex(input, l.y)}.execute(input)
}
func (l letterSwap) reverse(input []string) []string {
	return l.execute(input)
}

type rotate struct {
	x int
}

func (r rotate) execute(input []string) []string {
	rotation := r.x % len(input)
	if rotation == 0 {
		return input
	}
	if rotation < 0 {
		rotation = len(input) + rotation
	}
	splitIndex := len(input) - rotation
	result := make([]string, 0)
	result = append(result, input[splitIndex:]...)
	result = append(result, input[:splitIndex]...)
	return result
}

func (r rotate) reverse(input []string) []string {
	return rotate{-1 * r.x}.execute(input)
}

type letterRotate struct {
	x string
}

func (l letterRotate) execute(input []string) []string {
	letterIndex := getLetterIndex(input, l.x)
	rotation := 1 + letterIndex
	if 3 < letterIndex {
		rotation++
	}
	return rotate{rotation}.execute(input)
}
func (l letterRotate) reverse(input []string) []string {
	letterIndex := getLetterIndex(input, l.x)
	var rotation int
	if letterIndex < 1 {
		letterIndex += 8
	}
	if letterIndex%2 == 1 {
		rotation = -1 * ((letterIndex + 1) / 2)
	} else {
		rotation = -5 - (letterIndex / 2)
	}
	return rotate{rotation}.execute(input)
}

type reverse struct {
	x int
	y int
}

func (r reverse) execute(input []string) []string {
	result := make([]string, 0)
	result = append(result, input[:r.x]...)
	result = append(result, reverseArray(input[r.x:r.y+1])...)
	result = append(result, input[r.y+1:]...)
	return result
}

func (r reverse) reverse(input []string) []string {
	return r.execute(input)
}

type move struct {
	x int
	y int
}

func (m move) execute(input []string) []string {
	temp := make([]string, 0)
	for i, letter := range input {
		if i == m.x {
			continue
		}
		temp = append(temp, letter)
	}
	result := make([]string, 0)
	result = append(result, temp[:m.y]...)
	result = append(result, input[m.x])
	if m.y < len(temp) {
		result = append(result, temp[m.y:]...)
	}
	return result
}

func (m move) reverse(input []string) []string {
	return move{m.y, m.x}.execute(input)
}

func getLetterIndex(input []string, letter string) int {
	for i, l := range input {
		if l == letter {
			return i
		}
	}
	panic(fmt.Sprint("Letter", letter, "not found"))
}

func reverseArray(letters []string) []string {
	result := make([]string, len(letters))
	copy(result, letters)
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result
}
