package main

import (
	"fmt"
	"os"
	"strings"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) shuffle() {

}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}
