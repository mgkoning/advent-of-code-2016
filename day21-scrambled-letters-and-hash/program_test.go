package main

import "testing"

func TestPositionSwap(t *testing.T) {
	var swap = positionSwap{1, 3}
	assertEqual(swap.execute([]string{"a", "b", "c", "d", "e"}), []string{"a", "d", "c", "b", "e"}, t)
}

func TestLetterSwap(t *testing.T) {
	var swap = letterSwap{"a", "b"}
	assertEqual(swap.execute([]string{"a", "b", "c", "d", "e"}), []string{"b", "a", "c", "d", "e"}, t)
}

func TestRotateRight(t *testing.T) {
	var rotate = rotate{2}
	assertEqual(rotate.execute([]string{"b", "a", "c", "d", "e"}), []string{"d", "e", "b", "a", "c"}, t)
}

func TestRotateLeft(t *testing.T) {
	var rotate = rotate{-2}
	assertEqual(rotate.execute([]string{"b", "a", "c", "d", "e"}), []string{"c", "d", "e", "b", "a"}, t)
}

func TestLetterRotate(t *testing.T) {
	var rotate = letterRotate{"e"}
	assertEqual(rotate.execute([]string{"a", "d", "c", "b", "e"}), []string{"e", "a", "d", "c", "b"}, t)
}

func TestReverse(t *testing.T) {
	var reverse = reverse{1, 2}
	assertEqual(reverse.execute([]string{"a", "d", "c", "b", "e"}), []string{"a", "c", "d", "b", "e"}, t)
}

func TestMoveBack(t *testing.T) {
	var move = move{3, 1}
	assertEqual(move.execute([]string{"e", "a", "d", "c", "b"}), []string{"e", "c", "a", "d", "b"}, t)
}

func TestMoveForward(t *testing.T) {
	var move = move{1, 3}
	assertEqual(move.execute([]string{"e", "a", "c", "d", "b"}), []string{"e", "c", "d", "a", "b"}, t)
}

func assertEqual(actual []string, expected []string, t *testing.T) {
	if len(actual) != len(expected) {
		t.Error("Expected", expected, "but was", actual)
		return
	}
	for i := range actual {
		if actual[i] == expected[i] {
			continue
		}
		t.Error("Expected", expected, "but was", actual)
		return
	}
}
