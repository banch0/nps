package main

import "log"

// ctrl + alt + v
// shift + f6  // menyat vesde

// git commit -m "fix: fixes #1"

// var <name>[<length>]<type>
// _ - don't care

const promotersLowerBound = 9
const detractorsUpperBound = 6

func main() {

	scores := []int{10, 7, 10}

	// problem 3 := magic values

	// nps := (2 - 0) / 3 * 100
	// 2 / 3 * 100 -> 0 * 100 -> 0
	// 2 * 100 / 3 -> 200 / 3 -> 66

	result := nps(scores)
	log.Println(result)
}

// git log --oneline

// HW:1

// nps -> ctrl + shift + t
func nps(scores []int) int {
	promotes := 0
	detractors := 0

	for _, value := range scores {
		if value >= promotersLowerBound {
			promotes = promotes + 1
		}
		if value <= detractorsUpperBound {
			detractors = detractors + 1
		}
	}

	nps := (promotes - detractors) * 100 / len(scores)

	return nps
}
