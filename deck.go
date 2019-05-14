package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (d deck) saveToLocal() {
	joinedString := strings.Join([]string(d), ",")
	ioutil.WriteFile("local.txt", []byte(joinedString), 0666)
}

func loadFromLocal() deck {
	byteSlice, error := ioutil.ReadFile("local.txt")
	if error != nil {
		fmt.Println("Loading data ERROR!")
	}
	splitString := strings.Split(string(byteSlice), ",")
	return deck(splitString)
}
