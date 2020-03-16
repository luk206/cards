package main

func main() {
	//cards := newDeck()
	////cards.print()
	//
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("cards.txt")

	//cards := newDeckFromFile("cards.txt")
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

}