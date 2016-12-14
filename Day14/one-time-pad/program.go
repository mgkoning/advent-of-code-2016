package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var salt = "ngcjuoqr"
var stretching = 2016

func main() {
	hashesByIndex := make(map[int64]string)
	keyIndexes := make([]int64, 0)
	for next := int64(0); len(keyIndexes) < 64; next++ {
		var hash string
		hash, hashesByIndex = getOrAddHash(hashesByIndex, next)
		triplet := findTriplet(hash)
		if len(triplet) == 0 {
			continue
		}
		for quintetIndex := int64(1); quintetIndex <= 1000; quintetIndex++ {
			hash, hashesByIndex = getOrAddHash(hashesByIndex, next+quintetIndex)
			if hasMatchingQuintet(hash, triplet) {
				keyIndexes = append(keyIndexes, next)
				break
			}
		}
	}
	fmt.Println(keyIndexes)
}

func getOrAddHash(hashesByIndex map[int64]string, index int64) (string, map[int64]string) {
	val, found := hashesByIndex[index]
	if found {
		return val, hashesByIndex
	}
	val = hashIt(salt, index, stretching)
	hashesByIndex[index] = val
	return val, hashesByIndex
}

func findTriplet(s string) string {
	chars := strings.Split(s, "")
	for index := 2; index < len(chars); index++ {
		if chars[index] != chars[index-1] || chars[index] != chars[index-2] {
			continue
		}
		return strings.Repeat(chars[index], 3)
	}
	return ""
}

func hasMatchingQuintet(s string, triplet string) bool {
	return strings.Contains(s, makeQuintet(triplet))
}

func makeQuintet(s string) string {
	return strings.Repeat(strings.Split(s, "")[0], 5)
}

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func hashIt(salt string, index int64, stretching int) string {
	textToHash := fmt.Sprintf("%v%v", salt, index)
	var hash = hashString(textToHash)
	for stretchCount := 0; stretchCount < stretching; stretchCount++ {
		hash = hashString(hash)
	}
	return hash
}

func hashString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
