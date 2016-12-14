package main

import "testing"
import "strings"

func TestHash(t *testing.T) {
	hash := hashIt("abc", 18, 0)
	if !strings.Contains(hash, "888") {
		t.Error("Expected hash to contain 888 but was ", hash)
	}
	hash = hashIt("abc", 0, 2016)
	if !strings.HasPrefix(hash, "a107ff") {
		t.Error("Expected hash to have prefix a107ff but was", hash)
	}
}

func TestFindTriplets(t *testing.T) {
	triplet := findTriplet("234gfsfsdgddd45ghfg333")
	if triplet != "ddd" {
		t.Error("Expected triplet 0 to be 'ddd' but was", triplet)
	}
}
