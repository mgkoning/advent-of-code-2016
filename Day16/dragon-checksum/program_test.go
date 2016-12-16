package main

import (
	"strings"
	"testing"
)

func TestReverseAndFlip(t *testing.T) {
	result := strings.Join(reverseAndFlip(strings.Split("001001010100", "")), "")
	expected := "110101011011"
	if result != expected {
		t.Error("Expected", expected, "but was", result)
	}
}

func TestChecksum(t *testing.T) {
	result := strings.Join(calcChecksum(strings.Split("110010110100", "")), "")
	expected := "110101"
	if result != expected {
		t.Error("Expected", expected, "but was", result)
	}
}
