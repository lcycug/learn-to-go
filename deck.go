package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Jan", "Feb", "Mar"}
	cardValues := []string{"1st", "2nd", "3rd"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
