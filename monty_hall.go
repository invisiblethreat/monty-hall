package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Round contains all the information about a round of Let's Make a Deal
type Round struct {
	Car    [3]bool // Car randomly assigned to index 0->2
	Pick   int
	Reveal int
	Switch int
}

// NewRound returns a fresh round for the game
func NewRound() Round {
	var init Round
	var used [3]bool

	for i := range init.Car {
		init.Car[i] = false
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	init.Car[random.Intn(3)] = true

	init.Pick = random.Intn(3)
	used[init.Pick] = true

	for {
		rev := random.Intn(3)
		if rev != init.Pick && !init.Car[rev] {
			init.Reveal = rev
			used[rev] = true
			break
		}

	}

	for i, use := range used {
		if !use {
			init.Switch = i
		}
	}

	return init
}

// StratStay never switches doors
func StratStay(tries int) int {
	wins := 0

	for i := 0; i < tries; i++ {
		game := NewRound()

		if game.Car[game.Pick] {
			wins++
		}
	}
	return wins
}

// StratSwitch always switches doors
func StratSwitch(tries int) int {
	wins := 0

	for i := 0; i < tries; i++ {
		game := NewRound()

		if game.Car[game.Switch] {
			wins++
		}
	}
	return wins
}

func main() {
	start := time.Now()
	tries := 1000000
	stay := StratStay(tries)
	fmt.Printf("  Stay Won %d of %d\n", stay, tries)

	sw := StratSwitch(tries)
	fmt.Printf("Switch Won %d of %d\n", sw, tries)
	fmt.Printf("Rutime: %s\n", time.Since(start).String())
}
