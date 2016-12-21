package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	run("example.txt", "abcde", "")
	run("input.txt", "abcdefgh", "fbgdceah")
}

func run(fileName string, input string, inputToReverse string) {
	steps := parseSteps(fileName)
	start := strings.Split(input, "")
	result := start
	for _, s := range steps {
		result = s.execute(result)
	}
	fmt.Println(strings.Join(result, ""))
	if len(inputToReverse) == 0 {
		return
	}
	reversed := strings.Split(inputToReverse, "")
	for i := len(steps) - 1; 0 <= i; i-- {
		reversed = steps[i].reverse(reversed)
	}
	fmt.Println(strings.Join(reversed, ""))
}

func parseSteps(name string) []step {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\r\n")
	result := make([]step, 0)
lineLoop:
	for _, line := range lines {
		for _, m := range stepMap {
			if strings.HasPrefix(line, m.prefix) {
				result = append(result, m.read(line))
				continue lineLoop
			}
		}
	}
	return result
}

func mustParseInt(s string) int {
	r, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(r)
}

type prefixMap struct {
	prefix string
	read   func(string) step
}

var stepMap = []prefixMap{
	{"swap position", readSwapPosition},
	{"swap letter", readSwapLetter},
	{"rotate based", readLetterRotate},
	{"rotate", readRotate},
	{"reverse", readReverse},
	{"move", readMove},
}

var swapRegexp = regexp.MustCompile(`swap position (\d+) with position (\d+)`)

func readSwapPosition(line string) step {
	results := swapRegexp.FindStringSubmatch(line)
	return positionSwap{mustParseInt(results[1]), mustParseInt(results[2])}
}

var swapLetterRegexp = regexp.MustCompile(`swap letter ([^ ]+) with letter ([^ ]+)`)

func readSwapLetter(line string) step {
	results := swapLetterRegexp.FindStringSubmatch(line)
	return letterSwap{results[1], results[2]}

}

var rotateRegexp = regexp.MustCompile(`rotate (left|right) (\d+) steps?`)

func readRotate(line string) step {
	results := rotateRegexp.FindStringSubmatch(line)
	r := mustParseInt(results[2])
	if results[1] == "left" {
		r *= -1
	}
	return rotate{r}
}

var letterRotateRegexp = regexp.MustCompile(`rotate based on position of letter ([^ ]+)`)

func readLetterRotate(line string) step {
	results := letterRotateRegexp.FindStringSubmatch(line)
	return letterRotate{results[1]}
}

var reverseRegexp = regexp.MustCompile(`reverse positions (\d+) through (\d+)`)

func readReverse(line string) step {
	results := reverseRegexp.FindStringSubmatch(line)
	return reverse{mustParseInt(results[1]), mustParseInt(results[2])}
}

var moveRegexp = regexp.MustCompile(`move position (\d+) to position (\d+)`)

func readMove(line string) step {
	results := moveRegexp.FindStringSubmatch(line)
	return move{mustParseInt(results[1]), mustParseInt(results[2])}
}
