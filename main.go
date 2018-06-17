package main

func main() {
	cards := newDeck()

	cards.saveToFile("myCards")
	// hand, remainingDeck := deal(cards, 10)
	// remainingDeck.print()
	// hand.print()
}
