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

var debugging = false

func main() {
	file, err := ioutil.ReadFile("input.txt")
	check(err)
	contents := string(file)
	fmt.Printf("Decompressed length (v1): %v\n", decompressedLength(contents, false))
	fmt.Printf("Decompressed length (v2): %v\n", decompressedLength(contents, true))
}

func decompressedLength(compressedSequence string, recurse bool) int64 {
	exploded := strings.Split(replaceWhitespace(compressedSequence), "")
	length := int64(0)
	index := int64(0)
	for index < int64(len(exploded)) {
		current := exploded[index]
		if current != "(" {
			length++
			index++
			continue
		}
		repeatSpec := make([]string, 0)
		index++
		next := exploded[index]
		for next != ")" {
			repeatSpec = append(repeatSpec, next)
			index++
			next = exploded[index]
		}
		index++
		repeatLength, repeatTimes := parseRepeatSpec(strings.Join(repeatSpec, ""))
		toRepeat := exploded[index : index+repeatLength]
		index += repeatLength
		addedLength := int64(repeatLength)
		if recurse {
			addedLength = decompressedLength(strings.Join(toRepeat, ""), recurse)
		}
		length += repeatTimes * addedLength
	}
	if debugging {
		fmt.Println("Sequence", compressedSequence, "has length", length)
	}
	return length
}

var whitespaceRegexp = regexp.MustCompile(`\s+`)

func replaceWhitespace(s string) string {
	return whitespaceRegexp.ReplaceAllString(s, "")
}

var repeatSpecRegexp = regexp.MustCompile(`(\d+)x(\d+)`)

func parseRepeatSpec(repeatSpec string) (int64, int64) {
	matches := repeatSpecRegexp.FindStringSubmatch(repeatSpec)
	if matches == nil {
		panic(fmt.Sprintf("Did not understand %v", repeatSpec))
	}
	return parseInt(matches[1]), parseInt(matches[2])
}

func parseInt(value string) int64 {
	val, err := strconv.ParseInt(value, 10, 32)
	check(err)
	return val
}
