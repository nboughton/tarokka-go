package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	suit  string
	value int
}

var (
	suits    = []string{"Glyphs", "Swords", "Coins", "Stars"}
	highDeck = []string{"Ghost", "Innocent", "Marionette", "Darklord", "Mists", "Executioner", "Broken One", "Tempter", "Beast", "Donjon", "Raven", "Seer", "Artifact", "Horseman"}
	lowDeck  []card
)

func init() {
	// populate low deck
	for _, suit := range suits {
		for i := 1; i < 11; i++ {
			lowDeck = append(lowDeck, card{suit: suit, value: i})
		}
	}
}

func main() {
	// Set random Seed
	rand.Seed(time.Now().UnixNano())

	// Draw from the low deck
	for i := 1; i < 4; i++ {
		n := rand.Intn(len(lowDeck))
		c := lowDeck[n]
		fmt.Printf("%v: The %v of %v\n", i, c.value, c.suit)

		// remove spent card from deck
		lowDeck = append(lowDeck[:n], lowDeck[n+1:]...)
	}

	// Draw from the high deck
	for i := 4; i < 6; i++ {
		n := rand.Intn(len(highDeck))
		c := highDeck[n]
		fmt.Printf("%v: The %v\n", i, c)

		// remove spent card from deck
		highDeck = append(highDeck[:n], highDeck[n+1:]...)
	}
}
