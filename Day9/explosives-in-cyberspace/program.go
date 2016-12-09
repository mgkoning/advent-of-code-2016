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

func main() {
	file, err := ioutil.ReadFile("input.txt")
	check(err)
	decompressed := decompress(string(file))
	fmt.Println(decompressed)
	fmt.Printf("Decompressed length: %v\n", len(decompressed))
}

func decompress(compressedSequence string) string {
	exploded := strings.Split(replaceWhitespace(compressedSequence), "")
	result := make([]string, 0, len(exploded))
	index := 0
	for index < len(exploded) {
		current := exploded[index]
		if current != "(" {
			result = append(result, current)
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
		length, repeat := parseRepeatSpec(strings.Join(repeatSpec, ""))
		toRepeat := exploded[index : index+length]
		index += length
		for n := 0; n < repeat; n++ {
			for _, char := range toRepeat {
				result = append(result, char)
			}
		}
	}
	return strings.Join(result, "")
}

var whitespaceRegexp = regexp.MustCompile(`\s+`)

func replaceWhitespace(s string) string {
	return whitespaceRegexp.ReplaceAllString(s, "")
}

var repeatSpecRegexp = regexp.MustCompile(`(\d+)x(\d+)`)

func parseRepeatSpec(repeatSpec string) (int, int) {
	matches := repeatSpecRegexp.FindStringSubmatch(repeatSpec)
	if matches == nil {
		panic(fmt.Sprintf("Did not understand %v", repeatSpec))
	}
	return parseInt(matches[1]), parseInt(matches[2])
}

func parseInt(value string) int {
	val, err := strconv.ParseInt(value, 10, 32)
	check(err)
	return int(val)
}
