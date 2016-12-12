package main

import "testing"
import "fmt"

func TestIt(t *testing.T) {
	s := boop{map[string]int{"s": 4}}
	u := yessirree(s)
	fmt.Println(s)
	fmt.Println(u)
}

type boop struct {
	mab map[string]int
}

func yessirree(boop boop) boop {
	poob := boop
	poob.mab = duplicate(boop.mab)
	poob.mab["b"] = 1
	poob.mab["s"] = 3
	return poob
}
