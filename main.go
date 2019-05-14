package main

func main() {
	//cards := newDeck()
	cards := loadFromLocal()
	cards.print()
	//cards.saveToLocal()
}
