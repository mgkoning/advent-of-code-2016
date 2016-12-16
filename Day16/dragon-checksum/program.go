package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var exampleInput = "10000"
	var exampleDiskLength = 20
	calcAnswer(exampleInput, exampleDiskLength)
	// running time: ~0ms
	var input = "10111011111001111"
	var disk1Length = 272
	calcAnswer(input, disk1Length)
	// running time: ~0.5ms
	var disk2Length = 35651584
	calcAnswer(input, disk2Length)
	// running time: ~5900ms
}

func calcAnswer(input string, length int) {
	defer timeSince(time.Now())
	start := strings.Split(input, "")
	result := expand(start)
	for len(result) < length {
		result = expand(result)
	}
	result = result[:length]
	checksum := calcChecksum(result)
	for len(checksum)%2 == 0 {
		checksum = calcChecksum(checksum)
	}
	fmt.Println("Checksum:", strings.Join(checksum, ""))
}

var flipMap = map[string]string{"0": "1", "1": "0"}

func reverseAndFlip(originalChars []string) []string {
	length := len(originalChars)
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[length-1-i] = flipMap[originalChars[i]]
	}
	return result
}

func expand(s []string) []string {
	result := make([]string, len(s), len(s)*2+1)
	copy(result, s)
	result = append(result, "0")
	result = append(result, reverseAndFlip(s)...)
	return result
}

var checksumMap = map[bool]string{true: "1", false: "0"}

func calcChecksum(s []string) []string {
	result := make([]string, 0, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		result = append(result, checksumMap[s[i] == s[i+1]])
	}
	return result
}

func timeSince(from time.Time) {
	fmt.Println("Time elapsed (ms):", float64(time.Now().Sub(from).Nanoseconds())/float64(1000*1000))
}
