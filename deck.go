package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type Deck []string

// Creates a new deck of cards
func newDeck() Deck {
	// Creates a deck slice
	cards := Deck{}
	// Creates cardSuits which is a slice of string
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	// Creates cardValues which is a slice of string
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve"}

	// Use of a for loop to create the cards and append them to the cards deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	// returns a new deck
	return cards
}

// Prints with a receiver function
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function to read a file, if it is null exits the program with os.Exit(1)
// strings.Split returns a new slice with a separator of a ","
func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}
