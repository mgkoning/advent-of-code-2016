package main

import (
	"fmt"
	"testing"
)

func TestGetMd5Prefix(t *testing.T) {
	assertGetMd5Prefix("hijkl", "", "ced9", t)
	assertGetMd5Prefix("hijkl", "D", "f2bc", t)
	assertGetMd5Prefix("hijkl", "DR", "5745", t)
	assertGetMd5Prefix("hijkl", "DU", "528e", t)
	assertGetMd5Prefix("hijkl", "DUR", "818a", t)
}

func assertGetMd5Prefix(passcode, path, expected string, t *testing.T) {
	var result = fmt.Sprintf("%x", getMd5Prefix(passcode, path))
	if result != expected {
		t.Error("Expected", expected, "for passcode", passcode, "and path", path, "but was", result)
	}
}

func TestGetAllowedDirections(t *testing.T) {
	assertGetAllowedDirections("hijkl", "D", []string{"U", "L", "R"}, t)
	assertGetAllowedDirections("hijkl", "DR", []string{}, t)
	assertGetAllowedDirections("hijkl", "DU", []string{"R"}, t)
	assertGetAllowedDirections("hijkl", "DUR", []string{}, t)
}

func assertGetAllowedDirections(passcode, path string, expected []string, t *testing.T) {
	var result = getAllowedDirections(getMd5Prefix(passcode, path))
	if areEqual(result, expected) {
		return
	}
	t.Error("Expected", expected, "but was", result)

}

func areEqual(one []string, other []string) bool {
	if len(one) != len(other) {
		return false
	}
	for i, v := range one {
		if v != other[i] {
			return false
		}
	}
	return true
}
