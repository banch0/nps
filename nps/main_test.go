package main

import "testing"

func Test_npc(t *testing.T) {
	scores := []int{10, 7, 10}

	want := 66
	got := nps(scores)

	if got != want {
		t.Error("Error ")
	}
}
