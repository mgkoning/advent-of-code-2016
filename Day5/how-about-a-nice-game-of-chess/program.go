package main

import (
	"crypto/md5"
	"fmt"
	"io"
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
	id := "reyedfim"
	var passwordSoFar = make([]rune, 8)
	var counter int64
	var found int
	for ; found < 8; counter++ {
		hash := hash(id, counter)
		if !strings.HasPrefix(hash, "00000") {
			continue
		}
		fmt.Printf("Found hash '%v' (index %v)\n", hash, counter)
		indexAndRune := getRuneSlice(hash, 5, 7)
		index, err := strconv.ParseInt(string(indexAndRune[0]), 16, 64)
		check(err)
		if 7 < index {
			fmt.Println("  ... but invalid position")
			continue
		}
		if passwordSoFar[index] != 0 {
			fmt.Println("  ... but already filled")
			continue
		}
		passwordSoFar[index] = indexAndRune[1]
		found++
	}

	fmt.Printf("%v", string(passwordSoFar))
}

func hash(id string, salt int64) string {
	hasher := md5.New()
	io.WriteString(hasher, fmt.Sprintf("%v%v", id, salt))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func getRuneSlice(s string, from int, to int) []rune {
	var runes []rune
	for i, r := range s {
		if i < from {
			continue
		}
		if to <= i {
			return runes
		}
		runes = append(runes, r)
	}
	panic(fmt.Sprintf("slice %v - %v not in string '%v'", from, to, s))
}
